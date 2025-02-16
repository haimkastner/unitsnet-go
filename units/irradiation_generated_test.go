// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestIrradiationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerSquareMeter"}`
	
	factory := units.IrradiationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.IrradiationJoulePerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.IrradiationJoulePerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestIrradiationDto_ToJSON(t *testing.T) {
	dto := units.IrradiationDto{
		Value: 45,
		Unit:  units.IrradiationJoulePerSquareMeter,
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
	if result["unit"].(string) != string(units.IrradiationJoulePerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.IrradiationJoulePerSquareMeter, result["unit"])
	}
}

func TestNewIrradiation_InvalidValue(t *testing.T) {
	factory := units.IrradiationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateIrradiation(math.NaN(), units.IrradiationJoulePerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateIrradiation(math.Inf(1), units.IrradiationJoulePerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestIrradiationConversions(t *testing.T) {
	factory := units.IrradiationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateIrradiation(180, units.IrradiationJoulePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerSquareMeter.
		// No expected conversion value provided for JoulesPerSquareMeter, verifying result is not NaN.
		result := a.JoulesPerSquareMeter()
		cacheResult := a.JoulesPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerSquareCentimeter.
		// No expected conversion value provided for JoulesPerSquareCentimeter, verifying result is not NaN.
		result := a.JoulesPerSquareCentimeter()
		cacheResult := a.JoulesPerSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerSquareMillimeter.
		// No expected conversion value provided for JoulesPerSquareMillimeter, verifying result is not NaN.
		result := a.JoulesPerSquareMillimeter()
		cacheResult := a.JoulesPerSquareMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to WattHoursPerSquareMeter.
		// No expected conversion value provided for WattHoursPerSquareMeter, verifying result is not NaN.
		result := a.WattHoursPerSquareMeter()
		cacheResult := a.WattHoursPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattHoursPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to BtusPerSquareFoot.
		// No expected conversion value provided for BtusPerSquareFoot, verifying result is not NaN.
		result := a.BtusPerSquareFoot()
		cacheResult := a.BtusPerSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BtusPerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerSquareMeter.
		// No expected conversion value provided for KilojoulesPerSquareMeter, verifying result is not NaN.
		result := a.KilojoulesPerSquareMeter()
		cacheResult := a.KilojoulesPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MillijoulesPerSquareCentimeter.
		// No expected conversion value provided for MillijoulesPerSquareCentimeter, verifying result is not NaN.
		result := a.MillijoulesPerSquareCentimeter()
		cacheResult := a.MillijoulesPerSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillijoulesPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattHoursPerSquareMeter.
		// No expected conversion value provided for KilowattHoursPerSquareMeter, verifying result is not NaN.
		result := a.KilowattHoursPerSquareMeter()
		cacheResult := a.KilowattHoursPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattHoursPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilobtusPerSquareFoot.
		// No expected conversion value provided for KilobtusPerSquareFoot, verifying result is not NaN.
		result := a.KilobtusPerSquareFoot()
		cacheResult := a.KilobtusPerSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilobtusPerSquareFoot returned NaN")
		}
	}
}

func TestIrradiation_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.IrradiationFactory{}
	a, err := factory.CreateIrradiation(90, units.IrradiationJoulePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.IrradiationJoulePerSquareMeter {
		t.Errorf("expected default unit JoulePerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.IrradiationJoulePerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.IrradiationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.IrradiationJoulePerSquareMeter {
		t.Errorf("expected unit JoulePerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestIrradiationFactory_FromDto(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationJoulePerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.IrradiationDto{
        Value: math.NaN(),
        Unit:  units.IrradiationJoulePerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test JoulePerSquareMeter conversion
    joules_per_square_meterDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationJoulePerSquareMeter,
    }
    
    var joules_per_square_meterResult *units.Irradiation
    joules_per_square_meterResult, err = factory.FromDto(joules_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_square_meterResult.Convert(units.IrradiationJoulePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JoulePerSquareCentimeter conversion
    joules_per_square_centimeterDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationJoulePerSquareCentimeter,
    }
    
    var joules_per_square_centimeterResult *units.Irradiation
    joules_per_square_centimeterResult, err = factory.FromDto(joules_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_square_centimeterResult.Convert(units.IrradiationJoulePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JoulePerSquareMillimeter conversion
    joules_per_square_millimeterDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationJoulePerSquareMillimeter,
    }
    
    var joules_per_square_millimeterResult *units.Irradiation
    joules_per_square_millimeterResult, err = factory.FromDto(joules_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_square_millimeterResult.Convert(units.IrradiationJoulePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test WattHourPerSquareMeter conversion
    watt_hours_per_square_meterDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationWattHourPerSquareMeter,
    }
    
    var watt_hours_per_square_meterResult *units.Irradiation
    watt_hours_per_square_meterResult, err = factory.FromDto(watt_hours_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with WattHourPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hours_per_square_meterResult.Convert(units.IrradiationWattHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHourPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test BtuPerSquareFoot conversion
    btus_per_square_footDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationBtuPerSquareFoot,
    }
    
    var btus_per_square_footResult *units.Irradiation
    btus_per_square_footResult, err = factory.FromDto(btus_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_square_footResult.Convert(units.IrradiationBtuPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerSquareFoot = %v, want %v", converted, 100)
    }
    // Test KilojoulePerSquareMeter conversion
    kilojoules_per_square_meterDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationKilojoulePerSquareMeter,
    }
    
    var kilojoules_per_square_meterResult *units.Irradiation
    kilojoules_per_square_meterResult, err = factory.FromDto(kilojoules_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_square_meterResult.Convert(units.IrradiationKilojoulePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MillijoulePerSquareCentimeter conversion
    millijoules_per_square_centimeterDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationMillijoulePerSquareCentimeter,
    }
    
    var millijoules_per_square_centimeterResult *units.Irradiation
    millijoules_per_square_centimeterResult, err = factory.FromDto(millijoules_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MillijoulePerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millijoules_per_square_centimeterResult.Convert(units.IrradiationMillijoulePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillijoulePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilowattHourPerSquareMeter conversion
    kilowatt_hours_per_square_meterDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationKilowattHourPerSquareMeter,
    }
    
    var kilowatt_hours_per_square_meterResult *units.Irradiation
    kilowatt_hours_per_square_meterResult, err = factory.FromDto(kilowatt_hours_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattHourPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hours_per_square_meterResult.Convert(units.IrradiationKilowattHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHourPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilobtuPerSquareFoot conversion
    kilobtus_per_square_footDto := units.IrradiationDto{
        Value: 100,
        Unit:  units.IrradiationKilobtuPerSquareFoot,
    }
    
    var kilobtus_per_square_footResult *units.Irradiation
    kilobtus_per_square_footResult, err = factory.FromDto(kilobtus_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilobtuPerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobtus_per_square_footResult.Convert(units.IrradiationKilobtuPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobtuPerSquareFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.IrradiationDto{
        Value: 0,
        Unit:  units.IrradiationJoulePerSquareMeter,
    }
    
    var zeroResult *units.Irradiation
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestIrradiationFactory_FromDtoJSON(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "JoulePerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "JoulePerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.IrradiationDto{
        Value: nanValue,
        Unit:  units.IrradiationJoulePerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with JoulePerSquareMeter unit
    joules_per_square_meterJSON := []byte(`{"value": 100, "unit": "JoulePerSquareMeter"}`)
    joules_per_square_meterResult, err := factory.FromDtoJSON(joules_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_square_meterResult.Convert(units.IrradiationJoulePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with JoulePerSquareCentimeter unit
    joules_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "JoulePerSquareCentimeter"}`)
    joules_per_square_centimeterResult, err := factory.FromDtoJSON(joules_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_square_centimeterResult.Convert(units.IrradiationJoulePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with JoulePerSquareMillimeter unit
    joules_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "JoulePerSquareMillimeter"}`)
    joules_per_square_millimeterResult, err := factory.FromDtoJSON(joules_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_square_millimeterResult.Convert(units.IrradiationJoulePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with WattHourPerSquareMeter unit
    watt_hours_per_square_meterJSON := []byte(`{"value": 100, "unit": "WattHourPerSquareMeter"}`)
    watt_hours_per_square_meterResult, err := factory.FromDtoJSON(watt_hours_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattHourPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hours_per_square_meterResult.Convert(units.IrradiationWattHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHourPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerSquareFoot unit
    btus_per_square_footJSON := []byte(`{"value": 100, "unit": "BtuPerSquareFoot"}`)
    btus_per_square_footResult, err := factory.FromDtoJSON(btus_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_square_footResult.Convert(units.IrradiationBtuPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerSquareMeter unit
    kilojoules_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilojoulePerSquareMeter"}`)
    kilojoules_per_square_meterResult, err := factory.FromDtoJSON(kilojoules_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_square_meterResult.Convert(units.IrradiationKilojoulePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillijoulePerSquareCentimeter unit
    millijoules_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "MillijoulePerSquareCentimeter"}`)
    millijoules_per_square_centimeterResult, err := factory.FromDtoJSON(millijoules_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillijoulePerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millijoules_per_square_centimeterResult.Convert(units.IrradiationMillijoulePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillijoulePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattHourPerSquareMeter unit
    kilowatt_hours_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilowattHourPerSquareMeter"}`)
    kilowatt_hours_per_square_meterResult, err := factory.FromDtoJSON(kilowatt_hours_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattHourPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hours_per_square_meterResult.Convert(units.IrradiationKilowattHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHourPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilobtuPerSquareFoot unit
    kilobtus_per_square_footJSON := []byte(`{"value": 100, "unit": "KilobtuPerSquareFoot"}`)
    kilobtus_per_square_footResult, err := factory.FromDtoJSON(kilobtus_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilobtuPerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobtus_per_square_footResult.Convert(units.IrradiationKilobtuPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobtuPerSquareFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "JoulePerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoulesPerSquareMeter function
func TestIrradiationFactory_FromJoulesPerSquareMeter(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromJoulesPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationJoulePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromJoulesPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationJoulePerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromJoulesPerSquareCentimeter function
func TestIrradiationFactory_FromJoulesPerSquareCentimeter(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromJoulesPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationJoulePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromJoulesPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationJoulePerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromJoulesPerSquareMillimeter function
func TestIrradiationFactory_FromJoulesPerSquareMillimeter(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromJoulesPerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationJoulePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromJoulesPerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationJoulePerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWattHoursPerSquareMeter function
func TestIrradiationFactory_FromWattHoursPerSquareMeter(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattHoursPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromWattHoursPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationWattHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattHoursPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattHoursPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromWattHoursPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromWattHoursPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattHoursPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromWattHoursPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattHoursPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattHoursPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromWattHoursPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationWattHourPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattHoursPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerSquareFoot function
func TestIrradiationFactory_FromBtusPerSquareFoot(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerSquareFoot(100)
    if err != nil {
        t.Errorf("FromBtusPerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationBtuPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromBtusPerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromBtusPerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerSquareFoot(0)
    if err != nil {
        t.Errorf("FromBtusPerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationBtuPerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerSquareMeter function
func TestIrradiationFactory_FromKilojoulesPerSquareMeter(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationKilojoulePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationKilojoulePerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillijoulesPerSquareCentimeter function
func TestIrradiationFactory_FromMillijoulesPerSquareCentimeter(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillijoulesPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromMillijoulesPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationMillijoulePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillijoulesPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillijoulesPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMillijoulesPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMillijoulesPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillijoulesPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMillijoulesPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillijoulesPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillijoulesPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromMillijoulesPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationMillijoulePerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillijoulesPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattHoursPerSquareMeter function
func TestIrradiationFactory_FromKilowattHoursPerSquareMeter(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattHoursPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilowattHoursPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationKilowattHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattHoursPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattHoursPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattHoursPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattHoursPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattHoursPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattHoursPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattHoursPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattHoursPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilowattHoursPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationKilowattHourPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattHoursPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobtusPerSquareFoot function
func TestIrradiationFactory_FromKilobtusPerSquareFoot(t *testing.T) {
    factory := units.IrradiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobtusPerSquareFoot(100)
    if err != nil {
        t.Errorf("FromKilobtusPerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiationKilobtuPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobtusPerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobtusPerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromKilobtusPerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromKilobtusPerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilobtusPerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilobtusPerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobtusPerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobtusPerSquareFoot(0)
    if err != nil {
        t.Errorf("FromKilobtusPerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiationKilobtuPerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobtusPerSquareFoot() with zero value = %v, want 0", converted)
    }
}

func TestIrradiationToString(t *testing.T) {
	factory := units.IrradiationFactory{}
	a, err := factory.CreateIrradiation(45, units.IrradiationJoulePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.IrradiationJoulePerSquareMeter, 2)
	expected := "45.00 " + units.GetIrradiationAbbreviation(units.IrradiationJoulePerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.IrradiationJoulePerSquareMeter, -1)
	expected = "45 " + units.GetIrradiationAbbreviation(units.IrradiationJoulePerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestIrradiation_EqualityAndComparison(t *testing.T) {
	factory := units.IrradiationFactory{}
	a1, _ := factory.CreateIrradiation(60, units.IrradiationJoulePerSquareMeter)
	a2, _ := factory.CreateIrradiation(60, units.IrradiationJoulePerSquareMeter)
	a3, _ := factory.CreateIrradiation(120, units.IrradiationJoulePerSquareMeter)

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

func TestIrradiation_Arithmetic(t *testing.T) {
	factory := units.IrradiationFactory{}
	a1, _ := factory.CreateIrradiation(30, units.IrradiationJoulePerSquareMeter)
	a2, _ := factory.CreateIrradiation(45, units.IrradiationJoulePerSquareMeter)

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


func TestGetIrradiationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.IrradiationUnits
        want string
    }{
        {
            name: "JoulePerSquareMeter abbreviation",
            unit: units.IrradiationJoulePerSquareMeter,
            want: "J/m²",
        },
        {
            name: "JoulePerSquareCentimeter abbreviation",
            unit: units.IrradiationJoulePerSquareCentimeter,
            want: "J/cm²",
        },
        {
            name: "JoulePerSquareMillimeter abbreviation",
            unit: units.IrradiationJoulePerSquareMillimeter,
            want: "J/mm²",
        },
        {
            name: "WattHourPerSquareMeter abbreviation",
            unit: units.IrradiationWattHourPerSquareMeter,
            want: "Wh/m²",
        },
        {
            name: "BtuPerSquareFoot abbreviation",
            unit: units.IrradiationBtuPerSquareFoot,
            want: "Btu/ft²",
        },
        {
            name: "KilojoulePerSquareMeter abbreviation",
            unit: units.IrradiationKilojoulePerSquareMeter,
            want: "kJ/m²",
        },
        {
            name: "MillijoulePerSquareCentimeter abbreviation",
            unit: units.IrradiationMillijoulePerSquareCentimeter,
            want: "mJ/cm²",
        },
        {
            name: "KilowattHourPerSquareMeter abbreviation",
            unit: units.IrradiationKilowattHourPerSquareMeter,
            want: "kWh/m²",
        },
        {
            name: "KilobtuPerSquareFoot abbreviation",
            unit: units.IrradiationKilobtuPerSquareFoot,
            want: "kBtu/ft²",
        },
        {
            name: "invalid unit",
            unit: units.IrradiationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetIrradiationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetIrradiationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestIrradiation_String(t *testing.T) {
    factory := units.IrradiationFactory{}
    
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
            unit, err := factory.CreateIrradiation(tt.value, units.IrradiationJoulePerSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Irradiation.String() = %v, want %v", got, tt.want)
            }
        })
    }
}