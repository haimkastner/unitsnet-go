// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricResistivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "OhmMeter"}`
	
	factory := units.ElectricResistivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricResistivityOhmMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricResistivityOhmMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "OhmMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricResistivityDto_ToJSON(t *testing.T) {
	dto := units.ElectricResistivityDto{
		Value: 45,
		Unit:  units.ElectricResistivityOhmMeter,
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
	if result["unit"].(string) != string(units.ElectricResistivityOhmMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricResistivityOhmMeter, result["unit"])
	}
}

func TestNewElectricResistivity_InvalidValue(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricResistivity(math.NaN(), units.ElectricResistivityOhmMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricResistivity(math.Inf(1), units.ElectricResistivityOhmMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricResistivityConversions(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricResistivity(180, units.ElectricResistivityOhmMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to OhmMeters.
		// No expected conversion value provided for OhmMeters, verifying result is not NaN.
		result := a.OhmMeters()
		cacheResult := a.OhmMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OhmMeters returned NaN")
		}
	}
	{
		// Test conversion to OhmsCentimeter.
		// No expected conversion value provided for OhmsCentimeter, verifying result is not NaN.
		result := a.OhmsCentimeter()
		cacheResult := a.OhmsCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OhmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PicoohmMeters.
		// No expected conversion value provided for PicoohmMeters, verifying result is not NaN.
		result := a.PicoohmMeters()
		cacheResult := a.PicoohmMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicoohmMeters returned NaN")
		}
	}
	{
		// Test conversion to NanoohmMeters.
		// No expected conversion value provided for NanoohmMeters, verifying result is not NaN.
		result := a.NanoohmMeters()
		cacheResult := a.NanoohmMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanoohmMeters returned NaN")
		}
	}
	{
		// Test conversion to MicroohmMeters.
		// No expected conversion value provided for MicroohmMeters, verifying result is not NaN.
		result := a.MicroohmMeters()
		cacheResult := a.MicroohmMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicroohmMeters returned NaN")
		}
	}
	{
		// Test conversion to MilliohmMeters.
		// No expected conversion value provided for MilliohmMeters, verifying result is not NaN.
		result := a.MilliohmMeters()
		cacheResult := a.MilliohmMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliohmMeters returned NaN")
		}
	}
	{
		// Test conversion to KiloohmMeters.
		// No expected conversion value provided for KiloohmMeters, verifying result is not NaN.
		result := a.KiloohmMeters()
		cacheResult := a.KiloohmMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KiloohmMeters returned NaN")
		}
	}
	{
		// Test conversion to MegaohmMeters.
		// No expected conversion value provided for MegaohmMeters, verifying result is not NaN.
		result := a.MegaohmMeters()
		cacheResult := a.MegaohmMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegaohmMeters returned NaN")
		}
	}
	{
		// Test conversion to PicoohmsCentimeter.
		// No expected conversion value provided for PicoohmsCentimeter, verifying result is not NaN.
		result := a.PicoohmsCentimeter()
		cacheResult := a.PicoohmsCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicoohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NanoohmsCentimeter.
		// No expected conversion value provided for NanoohmsCentimeter, verifying result is not NaN.
		result := a.NanoohmsCentimeter()
		cacheResult := a.NanoohmsCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanoohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicroohmsCentimeter.
		// No expected conversion value provided for MicroohmsCentimeter, verifying result is not NaN.
		result := a.MicroohmsCentimeter()
		cacheResult := a.MicroohmsCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicroohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MilliohmsCentimeter.
		// No expected conversion value provided for MilliohmsCentimeter, verifying result is not NaN.
		result := a.MilliohmsCentimeter()
		cacheResult := a.MilliohmsCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KiloohmsCentimeter.
		// No expected conversion value provided for KiloohmsCentimeter, verifying result is not NaN.
		result := a.KiloohmsCentimeter()
		cacheResult := a.KiloohmsCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KiloohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MegaohmsCentimeter.
		// No expected conversion value provided for MegaohmsCentimeter, verifying result is not NaN.
		result := a.MegaohmsCentimeter()
		cacheResult := a.MegaohmsCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegaohmsCentimeter returned NaN")
		}
	}
}

func TestElectricResistivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	a, err := factory.CreateElectricResistivity(90, units.ElectricResistivityOhmMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricResistivityOhmMeter {
		t.Errorf("expected default unit OhmMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricResistivityOhmMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricResistivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricResistivityOhmMeter {
		t.Errorf("expected unit OhmMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricResistivityFactory_FromDto(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityOhmMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricResistivityDto{
        Value: math.NaN(),
        Unit:  units.ElectricResistivityOhmMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test OhmMeter conversion
    ohm_metersDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityOhmMeter,
    }
    
    var ohm_metersResult *units.ElectricResistivity
    ohm_metersResult, err = factory.FromDto(ohm_metersDto)
    if err != nil {
        t.Errorf("FromDto() with OhmMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ohm_metersResult.Convert(units.ElectricResistivityOhmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OhmMeter = %v, want %v", converted, 100)
    }
    // Test OhmCentimeter conversion
    ohms_centimeterDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityOhmCentimeter,
    }
    
    var ohms_centimeterResult *units.ElectricResistivity
    ohms_centimeterResult, err = factory.FromDto(ohms_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with OhmCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ohms_centimeterResult.Convert(units.ElectricResistivityOhmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OhmCentimeter = %v, want %v", converted, 100)
    }
    // Test PicoohmMeter conversion
    picoohm_metersDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityPicoohmMeter,
    }
    
    var picoohm_metersResult *units.ElectricResistivity
    picoohm_metersResult, err = factory.FromDto(picoohm_metersDto)
    if err != nil {
        t.Errorf("FromDto() with PicoohmMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoohm_metersResult.Convert(units.ElectricResistivityPicoohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicoohmMeter = %v, want %v", converted, 100)
    }
    // Test NanoohmMeter conversion
    nanoohm_metersDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityNanoohmMeter,
    }
    
    var nanoohm_metersResult *units.ElectricResistivity
    nanoohm_metersResult, err = factory.FromDto(nanoohm_metersDto)
    if err != nil {
        t.Errorf("FromDto() with NanoohmMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoohm_metersResult.Convert(units.ElectricResistivityNanoohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoohmMeter = %v, want %v", converted, 100)
    }
    // Test MicroohmMeter conversion
    microohm_metersDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityMicroohmMeter,
    }
    
    var microohm_metersResult *units.ElectricResistivity
    microohm_metersResult, err = factory.FromDto(microohm_metersDto)
    if err != nil {
        t.Errorf("FromDto() with MicroohmMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microohm_metersResult.Convert(units.ElectricResistivityMicroohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroohmMeter = %v, want %v", converted, 100)
    }
    // Test MilliohmMeter conversion
    milliohm_metersDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityMilliohmMeter,
    }
    
    var milliohm_metersResult *units.ElectricResistivity
    milliohm_metersResult, err = factory.FromDto(milliohm_metersDto)
    if err != nil {
        t.Errorf("FromDto() with MilliohmMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliohm_metersResult.Convert(units.ElectricResistivityMilliohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliohmMeter = %v, want %v", converted, 100)
    }
    // Test KiloohmMeter conversion
    kiloohm_metersDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityKiloohmMeter,
    }
    
    var kiloohm_metersResult *units.ElectricResistivity
    kiloohm_metersResult, err = factory.FromDto(kiloohm_metersDto)
    if err != nil {
        t.Errorf("FromDto() with KiloohmMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloohm_metersResult.Convert(units.ElectricResistivityKiloohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloohmMeter = %v, want %v", converted, 100)
    }
    // Test MegaohmMeter conversion
    megaohm_metersDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityMegaohmMeter,
    }
    
    var megaohm_metersResult *units.ElectricResistivity
    megaohm_metersResult, err = factory.FromDto(megaohm_metersDto)
    if err != nil {
        t.Errorf("FromDto() with MegaohmMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaohm_metersResult.Convert(units.ElectricResistivityMegaohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaohmMeter = %v, want %v", converted, 100)
    }
    // Test PicoohmCentimeter conversion
    picoohms_centimeterDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityPicoohmCentimeter,
    }
    
    var picoohms_centimeterResult *units.ElectricResistivity
    picoohms_centimeterResult, err = factory.FromDto(picoohms_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with PicoohmCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoohms_centimeterResult.Convert(units.ElectricResistivityPicoohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicoohmCentimeter = %v, want %v", converted, 100)
    }
    // Test NanoohmCentimeter conversion
    nanoohms_centimeterDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityNanoohmCentimeter,
    }
    
    var nanoohms_centimeterResult *units.ElectricResistivity
    nanoohms_centimeterResult, err = factory.FromDto(nanoohms_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NanoohmCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoohms_centimeterResult.Convert(units.ElectricResistivityNanoohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoohmCentimeter = %v, want %v", converted, 100)
    }
    // Test MicroohmCentimeter conversion
    microohms_centimeterDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityMicroohmCentimeter,
    }
    
    var microohms_centimeterResult *units.ElectricResistivity
    microohms_centimeterResult, err = factory.FromDto(microohms_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MicroohmCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microohms_centimeterResult.Convert(units.ElectricResistivityMicroohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroohmCentimeter = %v, want %v", converted, 100)
    }
    // Test MilliohmCentimeter conversion
    milliohms_centimeterDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityMilliohmCentimeter,
    }
    
    var milliohms_centimeterResult *units.ElectricResistivity
    milliohms_centimeterResult, err = factory.FromDto(milliohms_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliohmCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliohms_centimeterResult.Convert(units.ElectricResistivityMilliohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliohmCentimeter = %v, want %v", converted, 100)
    }
    // Test KiloohmCentimeter conversion
    kiloohms_centimeterDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityKiloohmCentimeter,
    }
    
    var kiloohms_centimeterResult *units.ElectricResistivity
    kiloohms_centimeterResult, err = factory.FromDto(kiloohms_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KiloohmCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloohms_centimeterResult.Convert(units.ElectricResistivityKiloohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloohmCentimeter = %v, want %v", converted, 100)
    }
    // Test MegaohmCentimeter conversion
    megaohms_centimeterDto := units.ElectricResistivityDto{
        Value: 100,
        Unit:  units.ElectricResistivityMegaohmCentimeter,
    }
    
    var megaohms_centimeterResult *units.ElectricResistivity
    megaohms_centimeterResult, err = factory.FromDto(megaohms_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MegaohmCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaohms_centimeterResult.Convert(units.ElectricResistivityMegaohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaohmCentimeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricResistivityDto{
        Value: 0,
        Unit:  units.ElectricResistivityOhmMeter,
    }
    
    var zeroResult *units.ElectricResistivity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricResistivityFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "OhmMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "OhmMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricResistivityDto{
        Value: nanValue,
        Unit:  units.ElectricResistivityOhmMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with OhmMeter unit
    ohm_metersJSON := []byte(`{"value": 100, "unit": "OhmMeter"}`)
    ohm_metersResult, err := factory.FromDtoJSON(ohm_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OhmMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ohm_metersResult.Convert(units.ElectricResistivityOhmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OhmMeter = %v, want %v", converted, 100)
    }
    // Test JSON with OhmCentimeter unit
    ohms_centimeterJSON := []byte(`{"value": 100, "unit": "OhmCentimeter"}`)
    ohms_centimeterResult, err := factory.FromDtoJSON(ohms_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OhmCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ohms_centimeterResult.Convert(units.ElectricResistivityOhmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OhmCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with PicoohmMeter unit
    picoohm_metersJSON := []byte(`{"value": 100, "unit": "PicoohmMeter"}`)
    picoohm_metersResult, err := factory.FromDtoJSON(picoohm_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicoohmMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoohm_metersResult.Convert(units.ElectricResistivityPicoohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicoohmMeter = %v, want %v", converted, 100)
    }
    // Test JSON with NanoohmMeter unit
    nanoohm_metersJSON := []byte(`{"value": 100, "unit": "NanoohmMeter"}`)
    nanoohm_metersResult, err := factory.FromDtoJSON(nanoohm_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoohmMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoohm_metersResult.Convert(units.ElectricResistivityNanoohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoohmMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicroohmMeter unit
    microohm_metersJSON := []byte(`{"value": 100, "unit": "MicroohmMeter"}`)
    microohm_metersResult, err := factory.FromDtoJSON(microohm_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroohmMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microohm_metersResult.Convert(units.ElectricResistivityMicroohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroohmMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliohmMeter unit
    milliohm_metersJSON := []byte(`{"value": 100, "unit": "MilliohmMeter"}`)
    milliohm_metersResult, err := factory.FromDtoJSON(milliohm_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliohmMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliohm_metersResult.Convert(units.ElectricResistivityMilliohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliohmMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KiloohmMeter unit
    kiloohm_metersJSON := []byte(`{"value": 100, "unit": "KiloohmMeter"}`)
    kiloohm_metersResult, err := factory.FromDtoJSON(kiloohm_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloohmMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloohm_metersResult.Convert(units.ElectricResistivityKiloohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloohmMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegaohmMeter unit
    megaohm_metersJSON := []byte(`{"value": 100, "unit": "MegaohmMeter"}`)
    megaohm_metersResult, err := factory.FromDtoJSON(megaohm_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaohmMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaohm_metersResult.Convert(units.ElectricResistivityMegaohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaohmMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PicoohmCentimeter unit
    picoohms_centimeterJSON := []byte(`{"value": 100, "unit": "PicoohmCentimeter"}`)
    picoohms_centimeterResult, err := factory.FromDtoJSON(picoohms_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicoohmCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoohms_centimeterResult.Convert(units.ElectricResistivityPicoohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicoohmCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NanoohmCentimeter unit
    nanoohms_centimeterJSON := []byte(`{"value": 100, "unit": "NanoohmCentimeter"}`)
    nanoohms_centimeterResult, err := factory.FromDtoJSON(nanoohms_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoohmCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoohms_centimeterResult.Convert(units.ElectricResistivityNanoohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoohmCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicroohmCentimeter unit
    microohms_centimeterJSON := []byte(`{"value": 100, "unit": "MicroohmCentimeter"}`)
    microohms_centimeterResult, err := factory.FromDtoJSON(microohms_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroohmCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microohms_centimeterResult.Convert(units.ElectricResistivityMicroohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroohmCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliohmCentimeter unit
    milliohms_centimeterJSON := []byte(`{"value": 100, "unit": "MilliohmCentimeter"}`)
    milliohms_centimeterResult, err := factory.FromDtoJSON(milliohms_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliohmCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliohms_centimeterResult.Convert(units.ElectricResistivityMilliohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliohmCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KiloohmCentimeter unit
    kiloohms_centimeterJSON := []byte(`{"value": 100, "unit": "KiloohmCentimeter"}`)
    kiloohms_centimeterResult, err := factory.FromDtoJSON(kiloohms_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloohmCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloohms_centimeterResult.Convert(units.ElectricResistivityKiloohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloohmCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegaohmCentimeter unit
    megaohms_centimeterJSON := []byte(`{"value": 100, "unit": "MegaohmCentimeter"}`)
    megaohms_centimeterResult, err := factory.FromDtoJSON(megaohms_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaohmCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaohms_centimeterResult.Convert(units.ElectricResistivityMegaohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaohmCentimeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "OhmMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromOhmMeters function
func TestElectricResistivityFactory_FromOhmMeters(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOhmMeters(100)
    if err != nil {
        t.Errorf("FromOhmMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityOhmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOhmMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOhmMeters(math.NaN())
    if err == nil {
        t.Error("FromOhmMeters() with NaN value should return error")
    }

    _, err = factory.FromOhmMeters(math.Inf(1))
    if err == nil {
        t.Error("FromOhmMeters() with +Inf value should return error")
    }

    _, err = factory.FromOhmMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromOhmMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOhmMeters(0)
    if err != nil {
        t.Errorf("FromOhmMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityOhmMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOhmMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromOhmsCentimeter function
func TestElectricResistivityFactory_FromOhmsCentimeter(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOhmsCentimeter(100)
    if err != nil {
        t.Errorf("FromOhmsCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityOhmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOhmsCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOhmsCentimeter(math.NaN())
    if err == nil {
        t.Error("FromOhmsCentimeter() with NaN value should return error")
    }

    _, err = factory.FromOhmsCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromOhmsCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromOhmsCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromOhmsCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOhmsCentimeter(0)
    if err != nil {
        t.Errorf("FromOhmsCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityOhmCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOhmsCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicoohmMeters function
func TestElectricResistivityFactory_FromPicoohmMeters(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicoohmMeters(100)
    if err != nil {
        t.Errorf("FromPicoohmMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityPicoohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicoohmMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicoohmMeters(math.NaN())
    if err == nil {
        t.Error("FromPicoohmMeters() with NaN value should return error")
    }

    _, err = factory.FromPicoohmMeters(math.Inf(1))
    if err == nil {
        t.Error("FromPicoohmMeters() with +Inf value should return error")
    }

    _, err = factory.FromPicoohmMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromPicoohmMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicoohmMeters(0)
    if err != nil {
        t.Errorf("FromPicoohmMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityPicoohmMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicoohmMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoohmMeters function
func TestElectricResistivityFactory_FromNanoohmMeters(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoohmMeters(100)
    if err != nil {
        t.Errorf("FromNanoohmMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityNanoohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoohmMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoohmMeters(math.NaN())
    if err == nil {
        t.Error("FromNanoohmMeters() with NaN value should return error")
    }

    _, err = factory.FromNanoohmMeters(math.Inf(1))
    if err == nil {
        t.Error("FromNanoohmMeters() with +Inf value should return error")
    }

    _, err = factory.FromNanoohmMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoohmMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoohmMeters(0)
    if err != nil {
        t.Errorf("FromNanoohmMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityNanoohmMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoohmMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroohmMeters function
func TestElectricResistivityFactory_FromMicroohmMeters(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroohmMeters(100)
    if err != nil {
        t.Errorf("FromMicroohmMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityMicroohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroohmMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroohmMeters(math.NaN())
    if err == nil {
        t.Error("FromMicroohmMeters() with NaN value should return error")
    }

    _, err = factory.FromMicroohmMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMicroohmMeters() with +Inf value should return error")
    }

    _, err = factory.FromMicroohmMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroohmMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroohmMeters(0)
    if err != nil {
        t.Errorf("FromMicroohmMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityMicroohmMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroohmMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliohmMeters function
func TestElectricResistivityFactory_FromMilliohmMeters(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliohmMeters(100)
    if err != nil {
        t.Errorf("FromMilliohmMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityMilliohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliohmMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliohmMeters(math.NaN())
    if err == nil {
        t.Error("FromMilliohmMeters() with NaN value should return error")
    }

    _, err = factory.FromMilliohmMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilliohmMeters() with +Inf value should return error")
    }

    _, err = factory.FromMilliohmMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliohmMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliohmMeters(0)
    if err != nil {
        t.Errorf("FromMilliohmMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityMilliohmMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliohmMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloohmMeters function
func TestElectricResistivityFactory_FromKiloohmMeters(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloohmMeters(100)
    if err != nil {
        t.Errorf("FromKiloohmMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityKiloohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloohmMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloohmMeters(math.NaN())
    if err == nil {
        t.Error("FromKiloohmMeters() with NaN value should return error")
    }

    _, err = factory.FromKiloohmMeters(math.Inf(1))
    if err == nil {
        t.Error("FromKiloohmMeters() with +Inf value should return error")
    }

    _, err = factory.FromKiloohmMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloohmMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloohmMeters(0)
    if err != nil {
        t.Errorf("FromKiloohmMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityKiloohmMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloohmMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaohmMeters function
func TestElectricResistivityFactory_FromMegaohmMeters(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaohmMeters(100)
    if err != nil {
        t.Errorf("FromMegaohmMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityMegaohmMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaohmMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaohmMeters(math.NaN())
    if err == nil {
        t.Error("FromMegaohmMeters() with NaN value should return error")
    }

    _, err = factory.FromMegaohmMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMegaohmMeters() with +Inf value should return error")
    }

    _, err = factory.FromMegaohmMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaohmMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaohmMeters(0)
    if err != nil {
        t.Errorf("FromMegaohmMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityMegaohmMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaohmMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromPicoohmsCentimeter function
func TestElectricResistivityFactory_FromPicoohmsCentimeter(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicoohmsCentimeter(100)
    if err != nil {
        t.Errorf("FromPicoohmsCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityPicoohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicoohmsCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicoohmsCentimeter(math.NaN())
    if err == nil {
        t.Error("FromPicoohmsCentimeter() with NaN value should return error")
    }

    _, err = factory.FromPicoohmsCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromPicoohmsCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromPicoohmsCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicoohmsCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicoohmsCentimeter(0)
    if err != nil {
        t.Errorf("FromPicoohmsCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityPicoohmCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicoohmsCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoohmsCentimeter function
func TestElectricResistivityFactory_FromNanoohmsCentimeter(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoohmsCentimeter(100)
    if err != nil {
        t.Errorf("FromNanoohmsCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityNanoohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoohmsCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoohmsCentimeter(math.NaN())
    if err == nil {
        t.Error("FromNanoohmsCentimeter() with NaN value should return error")
    }

    _, err = factory.FromNanoohmsCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanoohmsCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromNanoohmsCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoohmsCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoohmsCentimeter(0)
    if err != nil {
        t.Errorf("FromNanoohmsCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityNanoohmCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoohmsCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroohmsCentimeter function
func TestElectricResistivityFactory_FromMicroohmsCentimeter(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroohmsCentimeter(100)
    if err != nil {
        t.Errorf("FromMicroohmsCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityMicroohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroohmsCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroohmsCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMicroohmsCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMicroohmsCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicroohmsCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMicroohmsCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroohmsCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroohmsCentimeter(0)
    if err != nil {
        t.Errorf("FromMicroohmsCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityMicroohmCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroohmsCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliohmsCentimeter function
func TestElectricResistivityFactory_FromMilliohmsCentimeter(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliohmsCentimeter(100)
    if err != nil {
        t.Errorf("FromMilliohmsCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityMilliohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliohmsCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliohmsCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMilliohmsCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMilliohmsCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliohmsCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMilliohmsCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliohmsCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliohmsCentimeter(0)
    if err != nil {
        t.Errorf("FromMilliohmsCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityMilliohmCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliohmsCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloohmsCentimeter function
func TestElectricResistivityFactory_FromKiloohmsCentimeter(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloohmsCentimeter(100)
    if err != nil {
        t.Errorf("FromKiloohmsCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityKiloohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloohmsCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloohmsCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKiloohmsCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKiloohmsCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKiloohmsCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKiloohmsCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloohmsCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloohmsCentimeter(0)
    if err != nil {
        t.Errorf("FromKiloohmsCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityKiloohmCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloohmsCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaohmsCentimeter function
func TestElectricResistivityFactory_FromMegaohmsCentimeter(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaohmsCentimeter(100)
    if err != nil {
        t.Errorf("FromMegaohmsCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistivityMegaohmCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaohmsCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaohmsCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMegaohmsCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMegaohmsCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegaohmsCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMegaohmsCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaohmsCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaohmsCentimeter(0)
    if err != nil {
        t.Errorf("FromMegaohmsCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistivityMegaohmCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaohmsCentimeter() with zero value = %v, want 0", converted)
    }
}

func TestElectricResistivityToString(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	a, err := factory.CreateElectricResistivity(45, units.ElectricResistivityOhmMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricResistivityOhmMeter, 2)
	expected := "45.00 " + units.GetElectricResistivityAbbreviation(units.ElectricResistivityOhmMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricResistivityOhmMeter, -1)
	expected = "45 " + units.GetElectricResistivityAbbreviation(units.ElectricResistivityOhmMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricResistivity_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	a1, _ := factory.CreateElectricResistivity(60, units.ElectricResistivityOhmMeter)
	a2, _ := factory.CreateElectricResistivity(60, units.ElectricResistivityOhmMeter)
	a3, _ := factory.CreateElectricResistivity(120, units.ElectricResistivityOhmMeter)

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

func TestElectricResistivity_Arithmetic(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	a1, _ := factory.CreateElectricResistivity(30, units.ElectricResistivityOhmMeter)
	a2, _ := factory.CreateElectricResistivity(45, units.ElectricResistivityOhmMeter)

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


func TestGetElectricResistivityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricResistivityUnits
        want string
    }{
        {
            name: "OhmMeter abbreviation",
            unit: units.ElectricResistivityOhmMeter,
            want: "Ω·m",
        },
        {
            name: "OhmCentimeter abbreviation",
            unit: units.ElectricResistivityOhmCentimeter,
            want: "Ω·cm",
        },
        {
            name: "PicoohmMeter abbreviation",
            unit: units.ElectricResistivityPicoohmMeter,
            want: "pΩ·m",
        },
        {
            name: "NanoohmMeter abbreviation",
            unit: units.ElectricResistivityNanoohmMeter,
            want: "nΩ·m",
        },
        {
            name: "MicroohmMeter abbreviation",
            unit: units.ElectricResistivityMicroohmMeter,
            want: "μΩ·m",
        },
        {
            name: "MilliohmMeter abbreviation",
            unit: units.ElectricResistivityMilliohmMeter,
            want: "mΩ·m",
        },
        {
            name: "KiloohmMeter abbreviation",
            unit: units.ElectricResistivityKiloohmMeter,
            want: "kΩ·m",
        },
        {
            name: "MegaohmMeter abbreviation",
            unit: units.ElectricResistivityMegaohmMeter,
            want: "MΩ·m",
        },
        {
            name: "PicoohmCentimeter abbreviation",
            unit: units.ElectricResistivityPicoohmCentimeter,
            want: "pΩ·cm",
        },
        {
            name: "NanoohmCentimeter abbreviation",
            unit: units.ElectricResistivityNanoohmCentimeter,
            want: "nΩ·cm",
        },
        {
            name: "MicroohmCentimeter abbreviation",
            unit: units.ElectricResistivityMicroohmCentimeter,
            want: "μΩ·cm",
        },
        {
            name: "MilliohmCentimeter abbreviation",
            unit: units.ElectricResistivityMilliohmCentimeter,
            want: "mΩ·cm",
        },
        {
            name: "KiloohmCentimeter abbreviation",
            unit: units.ElectricResistivityKiloohmCentimeter,
            want: "kΩ·cm",
        },
        {
            name: "MegaohmCentimeter abbreviation",
            unit: units.ElectricResistivityMegaohmCentimeter,
            want: "MΩ·cm",
        },
        {
            name: "invalid unit",
            unit: units.ElectricResistivityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricResistivityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricResistivityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricResistivity_String(t *testing.T) {
    factory := units.ElectricResistivityFactory{}
    
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
            unit, err := factory.CreateElectricResistivity(tt.value, units.ElectricResistivityOhmMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricResistivity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}