// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestInformationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Bit"}`
	
	factory := units.InformationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.InformationBit {
		t.Errorf("expected unit %v, got %v", units.InformationBit, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Bit"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestInformationDto_ToJSON(t *testing.T) {
	dto := units.InformationDto{
		Value: 45,
		Unit:  units.InformationBit,
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
	if result["unit"].(string) != string(units.InformationBit) {
		t.Errorf("expected unit %s, got %v", units.InformationBit, result["unit"])
	}
}

func TestNewInformation_InvalidValue(t *testing.T) {
	factory := units.InformationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateInformation(math.NaN(), units.InformationBit)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateInformation(math.Inf(1), units.InformationBit)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestInformationConversions(t *testing.T) {
	factory := units.InformationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateInformation(180, units.InformationBit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Bytes.
		// No expected conversion value provided for Bytes, verifying result is not NaN.
		result := a.Bytes()
		cacheResult := a.Bytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Bytes returned NaN")
		}
	}
	{
		// Test conversion to Bits.
		// No expected conversion value provided for Bits, verifying result is not NaN.
		result := a.Bits()
		cacheResult := a.Bits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Bits returned NaN")
		}
	}
	{
		// Test conversion to Kilobytes.
		// No expected conversion value provided for Kilobytes, verifying result is not NaN.
		result := a.Kilobytes()
		cacheResult := a.Kilobytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilobytes returned NaN")
		}
	}
	{
		// Test conversion to Megabytes.
		// No expected conversion value provided for Megabytes, verifying result is not NaN.
		result := a.Megabytes()
		cacheResult := a.Megabytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megabytes returned NaN")
		}
	}
	{
		// Test conversion to Gigabytes.
		// No expected conversion value provided for Gigabytes, verifying result is not NaN.
		result := a.Gigabytes()
		cacheResult := a.Gigabytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigabytes returned NaN")
		}
	}
	{
		// Test conversion to Terabytes.
		// No expected conversion value provided for Terabytes, verifying result is not NaN.
		result := a.Terabytes()
		cacheResult := a.Terabytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Terabytes returned NaN")
		}
	}
	{
		// Test conversion to Petabytes.
		// No expected conversion value provided for Petabytes, verifying result is not NaN.
		result := a.Petabytes()
		cacheResult := a.Petabytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Petabytes returned NaN")
		}
	}
	{
		// Test conversion to Exabytes.
		// No expected conversion value provided for Exabytes, verifying result is not NaN.
		result := a.Exabytes()
		cacheResult := a.Exabytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Exabytes returned NaN")
		}
	}
	{
		// Test conversion to Kibibytes.
		// No expected conversion value provided for Kibibytes, verifying result is not NaN.
		result := a.Kibibytes()
		cacheResult := a.Kibibytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kibibytes returned NaN")
		}
	}
	{
		// Test conversion to Mebibytes.
		// No expected conversion value provided for Mebibytes, verifying result is not NaN.
		result := a.Mebibytes()
		cacheResult := a.Mebibytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Mebibytes returned NaN")
		}
	}
	{
		// Test conversion to Gibibytes.
		// No expected conversion value provided for Gibibytes, verifying result is not NaN.
		result := a.Gibibytes()
		cacheResult := a.Gibibytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gibibytes returned NaN")
		}
	}
	{
		// Test conversion to Tebibytes.
		// No expected conversion value provided for Tebibytes, verifying result is not NaN.
		result := a.Tebibytes()
		cacheResult := a.Tebibytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Tebibytes returned NaN")
		}
	}
	{
		// Test conversion to Pebibytes.
		// No expected conversion value provided for Pebibytes, verifying result is not NaN.
		result := a.Pebibytes()
		cacheResult := a.Pebibytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Pebibytes returned NaN")
		}
	}
	{
		// Test conversion to Exbibytes.
		// No expected conversion value provided for Exbibytes, verifying result is not NaN.
		result := a.Exbibytes()
		cacheResult := a.Exbibytes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Exbibytes returned NaN")
		}
	}
	{
		// Test conversion to Kilobits.
		// No expected conversion value provided for Kilobits, verifying result is not NaN.
		result := a.Kilobits()
		cacheResult := a.Kilobits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilobits returned NaN")
		}
	}
	{
		// Test conversion to Megabits.
		// No expected conversion value provided for Megabits, verifying result is not NaN.
		result := a.Megabits()
		cacheResult := a.Megabits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megabits returned NaN")
		}
	}
	{
		// Test conversion to Gigabits.
		// No expected conversion value provided for Gigabits, verifying result is not NaN.
		result := a.Gigabits()
		cacheResult := a.Gigabits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigabits returned NaN")
		}
	}
	{
		// Test conversion to Terabits.
		// No expected conversion value provided for Terabits, verifying result is not NaN.
		result := a.Terabits()
		cacheResult := a.Terabits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Terabits returned NaN")
		}
	}
	{
		// Test conversion to Petabits.
		// No expected conversion value provided for Petabits, verifying result is not NaN.
		result := a.Petabits()
		cacheResult := a.Petabits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Petabits returned NaN")
		}
	}
	{
		// Test conversion to Exabits.
		// No expected conversion value provided for Exabits, verifying result is not NaN.
		result := a.Exabits()
		cacheResult := a.Exabits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Exabits returned NaN")
		}
	}
	{
		// Test conversion to Kibibits.
		// No expected conversion value provided for Kibibits, verifying result is not NaN.
		result := a.Kibibits()
		cacheResult := a.Kibibits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kibibits returned NaN")
		}
	}
	{
		// Test conversion to Mebibits.
		// No expected conversion value provided for Mebibits, verifying result is not NaN.
		result := a.Mebibits()
		cacheResult := a.Mebibits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Mebibits returned NaN")
		}
	}
	{
		// Test conversion to Gibibits.
		// No expected conversion value provided for Gibibits, verifying result is not NaN.
		result := a.Gibibits()
		cacheResult := a.Gibibits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gibibits returned NaN")
		}
	}
	{
		// Test conversion to Tebibits.
		// No expected conversion value provided for Tebibits, verifying result is not NaN.
		result := a.Tebibits()
		cacheResult := a.Tebibits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Tebibits returned NaN")
		}
	}
	{
		// Test conversion to Pebibits.
		// No expected conversion value provided for Pebibits, verifying result is not NaN.
		result := a.Pebibits()
		cacheResult := a.Pebibits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Pebibits returned NaN")
		}
	}
	{
		// Test conversion to Exbibits.
		// No expected conversion value provided for Exbibits, verifying result is not NaN.
		result := a.Exbibits()
		cacheResult := a.Exbibits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Exbibits returned NaN")
		}
	}
}

func TestInformation_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.InformationFactory{}
	a, err := factory.CreateInformation(90, units.InformationBit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.InformationBit {
		t.Errorf("expected default unit Bit, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.InformationByte
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.InformationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.InformationBit {
		t.Errorf("expected unit Bit, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestInformationFactory_FromDto(t *testing.T) {
    factory := units.InformationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationBit,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.InformationDto{
        Value: math.NaN(),
        Unit:  units.InformationBit,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Byte conversion
    bytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationByte,
    }
    
    var bytesResult *units.Information
    bytesResult, err = factory.FromDto(bytesDto)
    if err != nil {
        t.Errorf("FromDto() with Byte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bytesResult.Convert(units.InformationByte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Byte = %v, want %v", converted, 100)
    }
    // Test Bit conversion
    bitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationBit,
    }
    
    var bitsResult *units.Information
    bitsResult, err = factory.FromDto(bitsDto)
    if err != nil {
        t.Errorf("FromDto() with Bit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bitsResult.Convert(units.InformationBit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Bit = %v, want %v", converted, 100)
    }
    // Test Kilobyte conversion
    kilobytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationKilobyte,
    }
    
    var kilobytesResult *units.Information
    kilobytesResult, err = factory.FromDto(kilobytesDto)
    if err != nil {
        t.Errorf("FromDto() with Kilobyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobytesResult.Convert(units.InformationKilobyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilobyte = %v, want %v", converted, 100)
    }
    // Test Megabyte conversion
    megabytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationMegabyte,
    }
    
    var megabytesResult *units.Information
    megabytesResult, err = factory.FromDto(megabytesDto)
    if err != nil {
        t.Errorf("FromDto() with Megabyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabytesResult.Convert(units.InformationMegabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megabyte = %v, want %v", converted, 100)
    }
    // Test Gigabyte conversion
    gigabytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationGigabyte,
    }
    
    var gigabytesResult *units.Information
    gigabytesResult, err = factory.FromDto(gigabytesDto)
    if err != nil {
        t.Errorf("FromDto() with Gigabyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabytesResult.Convert(units.InformationGigabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigabyte = %v, want %v", converted, 100)
    }
    // Test Terabyte conversion
    terabytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationTerabyte,
    }
    
    var terabytesResult *units.Information
    terabytesResult, err = factory.FromDto(terabytesDto)
    if err != nil {
        t.Errorf("FromDto() with Terabyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabytesResult.Convert(units.InformationTerabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terabyte = %v, want %v", converted, 100)
    }
    // Test Petabyte conversion
    petabytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationPetabyte,
    }
    
    var petabytesResult *units.Information
    petabytesResult, err = factory.FromDto(petabytesDto)
    if err != nil {
        t.Errorf("FromDto() with Petabyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabytesResult.Convert(units.InformationPetabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petabyte = %v, want %v", converted, 100)
    }
    // Test Exabyte conversion
    exabytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationExabyte,
    }
    
    var exabytesResult *units.Information
    exabytesResult, err = factory.FromDto(exabytesDto)
    if err != nil {
        t.Errorf("FromDto() with Exabyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabytesResult.Convert(units.InformationExabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exabyte = %v, want %v", converted, 100)
    }
    // Test Kibibyte conversion
    kibibytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationKibibyte,
    }
    
    var kibibytesResult *units.Information
    kibibytesResult, err = factory.FromDto(kibibytesDto)
    if err != nil {
        t.Errorf("FromDto() with Kibibyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibibytesResult.Convert(units.InformationKibibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kibibyte = %v, want %v", converted, 100)
    }
    // Test Mebibyte conversion
    mebibytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationMebibyte,
    }
    
    var mebibytesResult *units.Information
    mebibytesResult, err = factory.FromDto(mebibytesDto)
    if err != nil {
        t.Errorf("FromDto() with Mebibyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebibytesResult.Convert(units.InformationMebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mebibyte = %v, want %v", converted, 100)
    }
    // Test Gibibyte conversion
    gibibytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationGibibyte,
    }
    
    var gibibytesResult *units.Information
    gibibytesResult, err = factory.FromDto(gibibytesDto)
    if err != nil {
        t.Errorf("FromDto() with Gibibyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibibytesResult.Convert(units.InformationGibibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gibibyte = %v, want %v", converted, 100)
    }
    // Test Tebibyte conversion
    tebibytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationTebibyte,
    }
    
    var tebibytesResult *units.Information
    tebibytesResult, err = factory.FromDto(tebibytesDto)
    if err != nil {
        t.Errorf("FromDto() with Tebibyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebibytesResult.Convert(units.InformationTebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tebibyte = %v, want %v", converted, 100)
    }
    // Test Pebibyte conversion
    pebibytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationPebibyte,
    }
    
    var pebibytesResult *units.Information
    pebibytesResult, err = factory.FromDto(pebibytesDto)
    if err != nil {
        t.Errorf("FromDto() with Pebibyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebibytesResult.Convert(units.InformationPebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Pebibyte = %v, want %v", converted, 100)
    }
    // Test Exbibyte conversion
    exbibytesDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationExbibyte,
    }
    
    var exbibytesResult *units.Information
    exbibytesResult, err = factory.FromDto(exbibytesDto)
    if err != nil {
        t.Errorf("FromDto() with Exbibyte returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbibytesResult.Convert(units.InformationExbibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exbibyte = %v, want %v", converted, 100)
    }
    // Test Kilobit conversion
    kilobitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationKilobit,
    }
    
    var kilobitsResult *units.Information
    kilobitsResult, err = factory.FromDto(kilobitsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilobit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobitsResult.Convert(units.InformationKilobit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilobit = %v, want %v", converted, 100)
    }
    // Test Megabit conversion
    megabitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationMegabit,
    }
    
    var megabitsResult *units.Information
    megabitsResult, err = factory.FromDto(megabitsDto)
    if err != nil {
        t.Errorf("FromDto() with Megabit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabitsResult.Convert(units.InformationMegabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megabit = %v, want %v", converted, 100)
    }
    // Test Gigabit conversion
    gigabitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationGigabit,
    }
    
    var gigabitsResult *units.Information
    gigabitsResult, err = factory.FromDto(gigabitsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigabit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabitsResult.Convert(units.InformationGigabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigabit = %v, want %v", converted, 100)
    }
    // Test Terabit conversion
    terabitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationTerabit,
    }
    
    var terabitsResult *units.Information
    terabitsResult, err = factory.FromDto(terabitsDto)
    if err != nil {
        t.Errorf("FromDto() with Terabit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabitsResult.Convert(units.InformationTerabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terabit = %v, want %v", converted, 100)
    }
    // Test Petabit conversion
    petabitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationPetabit,
    }
    
    var petabitsResult *units.Information
    petabitsResult, err = factory.FromDto(petabitsDto)
    if err != nil {
        t.Errorf("FromDto() with Petabit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabitsResult.Convert(units.InformationPetabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petabit = %v, want %v", converted, 100)
    }
    // Test Exabit conversion
    exabitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationExabit,
    }
    
    var exabitsResult *units.Information
    exabitsResult, err = factory.FromDto(exabitsDto)
    if err != nil {
        t.Errorf("FromDto() with Exabit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabitsResult.Convert(units.InformationExabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exabit = %v, want %v", converted, 100)
    }
    // Test Kibibit conversion
    kibibitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationKibibit,
    }
    
    var kibibitsResult *units.Information
    kibibitsResult, err = factory.FromDto(kibibitsDto)
    if err != nil {
        t.Errorf("FromDto() with Kibibit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibibitsResult.Convert(units.InformationKibibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kibibit = %v, want %v", converted, 100)
    }
    // Test Mebibit conversion
    mebibitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationMebibit,
    }
    
    var mebibitsResult *units.Information
    mebibitsResult, err = factory.FromDto(mebibitsDto)
    if err != nil {
        t.Errorf("FromDto() with Mebibit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebibitsResult.Convert(units.InformationMebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mebibit = %v, want %v", converted, 100)
    }
    // Test Gibibit conversion
    gibibitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationGibibit,
    }
    
    var gibibitsResult *units.Information
    gibibitsResult, err = factory.FromDto(gibibitsDto)
    if err != nil {
        t.Errorf("FromDto() with Gibibit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibibitsResult.Convert(units.InformationGibibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gibibit = %v, want %v", converted, 100)
    }
    // Test Tebibit conversion
    tebibitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationTebibit,
    }
    
    var tebibitsResult *units.Information
    tebibitsResult, err = factory.FromDto(tebibitsDto)
    if err != nil {
        t.Errorf("FromDto() with Tebibit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebibitsResult.Convert(units.InformationTebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tebibit = %v, want %v", converted, 100)
    }
    // Test Pebibit conversion
    pebibitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationPebibit,
    }
    
    var pebibitsResult *units.Information
    pebibitsResult, err = factory.FromDto(pebibitsDto)
    if err != nil {
        t.Errorf("FromDto() with Pebibit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebibitsResult.Convert(units.InformationPebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Pebibit = %v, want %v", converted, 100)
    }
    // Test Exbibit conversion
    exbibitsDto := units.InformationDto{
        Value: 100,
        Unit:  units.InformationExbibit,
    }
    
    var exbibitsResult *units.Information
    exbibitsResult, err = factory.FromDto(exbibitsDto)
    if err != nil {
        t.Errorf("FromDto() with Exbibit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbibitsResult.Convert(units.InformationExbibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exbibit = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.InformationDto{
        Value: 0,
        Unit:  units.InformationBit,
    }
    
    var zeroResult *units.Information
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestInformationFactory_FromDtoJSON(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Bit"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Bit"}`)
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
    nanJSON, _ := json.Marshal(units.InformationDto{
        Value: nanValue,
        Unit:  units.InformationBit,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Byte unit
    bytesJSON := []byte(`{"value": 100, "unit": "Byte"}`)
    bytesResult, err := factory.FromDtoJSON(bytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Byte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bytesResult.Convert(units.InformationByte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Byte = %v, want %v", converted, 100)
    }
    // Test JSON with Bit unit
    bitsJSON := []byte(`{"value": 100, "unit": "Bit"}`)
    bitsResult, err := factory.FromDtoJSON(bitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Bit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bitsResult.Convert(units.InformationBit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Bit = %v, want %v", converted, 100)
    }
    // Test JSON with Kilobyte unit
    kilobytesJSON := []byte(`{"value": 100, "unit": "Kilobyte"}`)
    kilobytesResult, err := factory.FromDtoJSON(kilobytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilobyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobytesResult.Convert(units.InformationKilobyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilobyte = %v, want %v", converted, 100)
    }
    // Test JSON with Megabyte unit
    megabytesJSON := []byte(`{"value": 100, "unit": "Megabyte"}`)
    megabytesResult, err := factory.FromDtoJSON(megabytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megabyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabytesResult.Convert(units.InformationMegabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megabyte = %v, want %v", converted, 100)
    }
    // Test JSON with Gigabyte unit
    gigabytesJSON := []byte(`{"value": 100, "unit": "Gigabyte"}`)
    gigabytesResult, err := factory.FromDtoJSON(gigabytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigabyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabytesResult.Convert(units.InformationGigabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigabyte = %v, want %v", converted, 100)
    }
    // Test JSON with Terabyte unit
    terabytesJSON := []byte(`{"value": 100, "unit": "Terabyte"}`)
    terabytesResult, err := factory.FromDtoJSON(terabytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terabyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabytesResult.Convert(units.InformationTerabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terabyte = %v, want %v", converted, 100)
    }
    // Test JSON with Petabyte unit
    petabytesJSON := []byte(`{"value": 100, "unit": "Petabyte"}`)
    petabytesResult, err := factory.FromDtoJSON(petabytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Petabyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabytesResult.Convert(units.InformationPetabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petabyte = %v, want %v", converted, 100)
    }
    // Test JSON with Exabyte unit
    exabytesJSON := []byte(`{"value": 100, "unit": "Exabyte"}`)
    exabytesResult, err := factory.FromDtoJSON(exabytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Exabyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabytesResult.Convert(units.InformationExabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exabyte = %v, want %v", converted, 100)
    }
    // Test JSON with Kibibyte unit
    kibibytesJSON := []byte(`{"value": 100, "unit": "Kibibyte"}`)
    kibibytesResult, err := factory.FromDtoJSON(kibibytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kibibyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibibytesResult.Convert(units.InformationKibibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kibibyte = %v, want %v", converted, 100)
    }
    // Test JSON with Mebibyte unit
    mebibytesJSON := []byte(`{"value": 100, "unit": "Mebibyte"}`)
    mebibytesResult, err := factory.FromDtoJSON(mebibytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Mebibyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebibytesResult.Convert(units.InformationMebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mebibyte = %v, want %v", converted, 100)
    }
    // Test JSON with Gibibyte unit
    gibibytesJSON := []byte(`{"value": 100, "unit": "Gibibyte"}`)
    gibibytesResult, err := factory.FromDtoJSON(gibibytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gibibyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibibytesResult.Convert(units.InformationGibibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gibibyte = %v, want %v", converted, 100)
    }
    // Test JSON with Tebibyte unit
    tebibytesJSON := []byte(`{"value": 100, "unit": "Tebibyte"}`)
    tebibytesResult, err := factory.FromDtoJSON(tebibytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Tebibyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebibytesResult.Convert(units.InformationTebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tebibyte = %v, want %v", converted, 100)
    }
    // Test JSON with Pebibyte unit
    pebibytesJSON := []byte(`{"value": 100, "unit": "Pebibyte"}`)
    pebibytesResult, err := factory.FromDtoJSON(pebibytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Pebibyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebibytesResult.Convert(units.InformationPebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Pebibyte = %v, want %v", converted, 100)
    }
    // Test JSON with Exbibyte unit
    exbibytesJSON := []byte(`{"value": 100, "unit": "Exbibyte"}`)
    exbibytesResult, err := factory.FromDtoJSON(exbibytesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Exbibyte unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbibytesResult.Convert(units.InformationExbibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exbibyte = %v, want %v", converted, 100)
    }
    // Test JSON with Kilobit unit
    kilobitsJSON := []byte(`{"value": 100, "unit": "Kilobit"}`)
    kilobitsResult, err := factory.FromDtoJSON(kilobitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilobit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobitsResult.Convert(units.InformationKilobit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilobit = %v, want %v", converted, 100)
    }
    // Test JSON with Megabit unit
    megabitsJSON := []byte(`{"value": 100, "unit": "Megabit"}`)
    megabitsResult, err := factory.FromDtoJSON(megabitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megabit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabitsResult.Convert(units.InformationMegabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megabit = %v, want %v", converted, 100)
    }
    // Test JSON with Gigabit unit
    gigabitsJSON := []byte(`{"value": 100, "unit": "Gigabit"}`)
    gigabitsResult, err := factory.FromDtoJSON(gigabitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigabit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabitsResult.Convert(units.InformationGigabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigabit = %v, want %v", converted, 100)
    }
    // Test JSON with Terabit unit
    terabitsJSON := []byte(`{"value": 100, "unit": "Terabit"}`)
    terabitsResult, err := factory.FromDtoJSON(terabitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terabit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabitsResult.Convert(units.InformationTerabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terabit = %v, want %v", converted, 100)
    }
    // Test JSON with Petabit unit
    petabitsJSON := []byte(`{"value": 100, "unit": "Petabit"}`)
    petabitsResult, err := factory.FromDtoJSON(petabitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Petabit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabitsResult.Convert(units.InformationPetabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petabit = %v, want %v", converted, 100)
    }
    // Test JSON with Exabit unit
    exabitsJSON := []byte(`{"value": 100, "unit": "Exabit"}`)
    exabitsResult, err := factory.FromDtoJSON(exabitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Exabit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabitsResult.Convert(units.InformationExabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exabit = %v, want %v", converted, 100)
    }
    // Test JSON with Kibibit unit
    kibibitsJSON := []byte(`{"value": 100, "unit": "Kibibit"}`)
    kibibitsResult, err := factory.FromDtoJSON(kibibitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kibibit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibibitsResult.Convert(units.InformationKibibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kibibit = %v, want %v", converted, 100)
    }
    // Test JSON with Mebibit unit
    mebibitsJSON := []byte(`{"value": 100, "unit": "Mebibit"}`)
    mebibitsResult, err := factory.FromDtoJSON(mebibitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Mebibit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebibitsResult.Convert(units.InformationMebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mebibit = %v, want %v", converted, 100)
    }
    // Test JSON with Gibibit unit
    gibibitsJSON := []byte(`{"value": 100, "unit": "Gibibit"}`)
    gibibitsResult, err := factory.FromDtoJSON(gibibitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gibibit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibibitsResult.Convert(units.InformationGibibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gibibit = %v, want %v", converted, 100)
    }
    // Test JSON with Tebibit unit
    tebibitsJSON := []byte(`{"value": 100, "unit": "Tebibit"}`)
    tebibitsResult, err := factory.FromDtoJSON(tebibitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Tebibit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebibitsResult.Convert(units.InformationTebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tebibit = %v, want %v", converted, 100)
    }
    // Test JSON with Pebibit unit
    pebibitsJSON := []byte(`{"value": 100, "unit": "Pebibit"}`)
    pebibitsResult, err := factory.FromDtoJSON(pebibitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Pebibit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebibitsResult.Convert(units.InformationPebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Pebibit = %v, want %v", converted, 100)
    }
    // Test JSON with Exbibit unit
    exbibitsJSON := []byte(`{"value": 100, "unit": "Exbibit"}`)
    exbibitsResult, err := factory.FromDtoJSON(exbibitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Exbibit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbibitsResult.Convert(units.InformationExbibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exbibit = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Bit"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromBytes function
func TestInformationFactory_FromBytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBytes(100)
    if err != nil {
        t.Errorf("FromBytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationByte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBytes(math.NaN())
    if err == nil {
        t.Error("FromBytes() with NaN value should return error")
    }

    _, err = factory.FromBytes(math.Inf(1))
    if err == nil {
        t.Error("FromBytes() with +Inf value should return error")
    }

    _, err = factory.FromBytes(math.Inf(-1))
    if err == nil {
        t.Error("FromBytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBytes(0)
    if err != nil {
        t.Errorf("FromBytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationByte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBytes() with zero value = %v, want 0", converted)
    }
}
// Test FromBits function
func TestInformationFactory_FromBits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBits(100)
    if err != nil {
        t.Errorf("FromBits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationBit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBits(math.NaN())
    if err == nil {
        t.Error("FromBits() with NaN value should return error")
    }

    _, err = factory.FromBits(math.Inf(1))
    if err == nil {
        t.Error("FromBits() with +Inf value should return error")
    }

    _, err = factory.FromBits(math.Inf(-1))
    if err == nil {
        t.Error("FromBits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBits(0)
    if err != nil {
        t.Errorf("FromBits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationBit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBits() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobytes function
func TestInformationFactory_FromKilobytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobytes(100)
    if err != nil {
        t.Errorf("FromKilobytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationKilobyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobytes(math.NaN())
    if err == nil {
        t.Error("FromKilobytes() with NaN value should return error")
    }

    _, err = factory.FromKilobytes(math.Inf(1))
    if err == nil {
        t.Error("FromKilobytes() with +Inf value should return error")
    }

    _, err = factory.FromKilobytes(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobytes(0)
    if err != nil {
        t.Errorf("FromKilobytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationKilobyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobytes() with zero value = %v, want 0", converted)
    }
}
// Test FromMegabytes function
func TestInformationFactory_FromMegabytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegabytes(100)
    if err != nil {
        t.Errorf("FromMegabytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationMegabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegabytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegabytes(math.NaN())
    if err == nil {
        t.Error("FromMegabytes() with NaN value should return error")
    }

    _, err = factory.FromMegabytes(math.Inf(1))
    if err == nil {
        t.Error("FromMegabytes() with +Inf value should return error")
    }

    _, err = factory.FromMegabytes(math.Inf(-1))
    if err == nil {
        t.Error("FromMegabytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegabytes(0)
    if err != nil {
        t.Errorf("FromMegabytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationMegabyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegabytes() with zero value = %v, want 0", converted)
    }
}
// Test FromGigabytes function
func TestInformationFactory_FromGigabytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigabytes(100)
    if err != nil {
        t.Errorf("FromGigabytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationGigabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigabytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigabytes(math.NaN())
    if err == nil {
        t.Error("FromGigabytes() with NaN value should return error")
    }

    _, err = factory.FromGigabytes(math.Inf(1))
    if err == nil {
        t.Error("FromGigabytes() with +Inf value should return error")
    }

    _, err = factory.FromGigabytes(math.Inf(-1))
    if err == nil {
        t.Error("FromGigabytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigabytes(0)
    if err != nil {
        t.Errorf("FromGigabytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationGigabyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigabytes() with zero value = %v, want 0", converted)
    }
}
// Test FromTerabytes function
func TestInformationFactory_FromTerabytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerabytes(100)
    if err != nil {
        t.Errorf("FromTerabytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationTerabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerabytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerabytes(math.NaN())
    if err == nil {
        t.Error("FromTerabytes() with NaN value should return error")
    }

    _, err = factory.FromTerabytes(math.Inf(1))
    if err == nil {
        t.Error("FromTerabytes() with +Inf value should return error")
    }

    _, err = factory.FromTerabytes(math.Inf(-1))
    if err == nil {
        t.Error("FromTerabytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerabytes(0)
    if err != nil {
        t.Errorf("FromTerabytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationTerabyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerabytes() with zero value = %v, want 0", converted)
    }
}
// Test FromPetabytes function
func TestInformationFactory_FromPetabytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetabytes(100)
    if err != nil {
        t.Errorf("FromPetabytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationPetabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetabytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetabytes(math.NaN())
    if err == nil {
        t.Error("FromPetabytes() with NaN value should return error")
    }

    _, err = factory.FromPetabytes(math.Inf(1))
    if err == nil {
        t.Error("FromPetabytes() with +Inf value should return error")
    }

    _, err = factory.FromPetabytes(math.Inf(-1))
    if err == nil {
        t.Error("FromPetabytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetabytes(0)
    if err != nil {
        t.Errorf("FromPetabytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationPetabyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetabytes() with zero value = %v, want 0", converted)
    }
}
// Test FromExabytes function
func TestInformationFactory_FromExabytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExabytes(100)
    if err != nil {
        t.Errorf("FromExabytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationExabyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExabytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExabytes(math.NaN())
    if err == nil {
        t.Error("FromExabytes() with NaN value should return error")
    }

    _, err = factory.FromExabytes(math.Inf(1))
    if err == nil {
        t.Error("FromExabytes() with +Inf value should return error")
    }

    _, err = factory.FromExabytes(math.Inf(-1))
    if err == nil {
        t.Error("FromExabytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExabytes(0)
    if err != nil {
        t.Errorf("FromExabytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationExabyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExabytes() with zero value = %v, want 0", converted)
    }
}
// Test FromKibibytes function
func TestInformationFactory_FromKibibytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKibibytes(100)
    if err != nil {
        t.Errorf("FromKibibytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationKibibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKibibytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKibibytes(math.NaN())
    if err == nil {
        t.Error("FromKibibytes() with NaN value should return error")
    }

    _, err = factory.FromKibibytes(math.Inf(1))
    if err == nil {
        t.Error("FromKibibytes() with +Inf value should return error")
    }

    _, err = factory.FromKibibytes(math.Inf(-1))
    if err == nil {
        t.Error("FromKibibytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKibibytes(0)
    if err != nil {
        t.Errorf("FromKibibytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationKibibyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKibibytes() with zero value = %v, want 0", converted)
    }
}
// Test FromMebibytes function
func TestInformationFactory_FromMebibytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMebibytes(100)
    if err != nil {
        t.Errorf("FromMebibytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationMebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMebibytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMebibytes(math.NaN())
    if err == nil {
        t.Error("FromMebibytes() with NaN value should return error")
    }

    _, err = factory.FromMebibytes(math.Inf(1))
    if err == nil {
        t.Error("FromMebibytes() with +Inf value should return error")
    }

    _, err = factory.FromMebibytes(math.Inf(-1))
    if err == nil {
        t.Error("FromMebibytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMebibytes(0)
    if err != nil {
        t.Errorf("FromMebibytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationMebibyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMebibytes() with zero value = %v, want 0", converted)
    }
}
// Test FromGibibytes function
func TestInformationFactory_FromGibibytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGibibytes(100)
    if err != nil {
        t.Errorf("FromGibibytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationGibibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGibibytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGibibytes(math.NaN())
    if err == nil {
        t.Error("FromGibibytes() with NaN value should return error")
    }

    _, err = factory.FromGibibytes(math.Inf(1))
    if err == nil {
        t.Error("FromGibibytes() with +Inf value should return error")
    }

    _, err = factory.FromGibibytes(math.Inf(-1))
    if err == nil {
        t.Error("FromGibibytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGibibytes(0)
    if err != nil {
        t.Errorf("FromGibibytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationGibibyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGibibytes() with zero value = %v, want 0", converted)
    }
}
// Test FromTebibytes function
func TestInformationFactory_FromTebibytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTebibytes(100)
    if err != nil {
        t.Errorf("FromTebibytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationTebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTebibytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTebibytes(math.NaN())
    if err == nil {
        t.Error("FromTebibytes() with NaN value should return error")
    }

    _, err = factory.FromTebibytes(math.Inf(1))
    if err == nil {
        t.Error("FromTebibytes() with +Inf value should return error")
    }

    _, err = factory.FromTebibytes(math.Inf(-1))
    if err == nil {
        t.Error("FromTebibytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTebibytes(0)
    if err != nil {
        t.Errorf("FromTebibytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationTebibyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTebibytes() with zero value = %v, want 0", converted)
    }
}
// Test FromPebibytes function
func TestInformationFactory_FromPebibytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPebibytes(100)
    if err != nil {
        t.Errorf("FromPebibytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationPebibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPebibytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPebibytes(math.NaN())
    if err == nil {
        t.Error("FromPebibytes() with NaN value should return error")
    }

    _, err = factory.FromPebibytes(math.Inf(1))
    if err == nil {
        t.Error("FromPebibytes() with +Inf value should return error")
    }

    _, err = factory.FromPebibytes(math.Inf(-1))
    if err == nil {
        t.Error("FromPebibytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPebibytes(0)
    if err != nil {
        t.Errorf("FromPebibytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationPebibyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPebibytes() with zero value = %v, want 0", converted)
    }
}
// Test FromExbibytes function
func TestInformationFactory_FromExbibytes(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExbibytes(100)
    if err != nil {
        t.Errorf("FromExbibytes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationExbibyte)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExbibytes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExbibytes(math.NaN())
    if err == nil {
        t.Error("FromExbibytes() with NaN value should return error")
    }

    _, err = factory.FromExbibytes(math.Inf(1))
    if err == nil {
        t.Error("FromExbibytes() with +Inf value should return error")
    }

    _, err = factory.FromExbibytes(math.Inf(-1))
    if err == nil {
        t.Error("FromExbibytes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExbibytes(0)
    if err != nil {
        t.Errorf("FromExbibytes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationExbibyte)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExbibytes() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobits function
func TestInformationFactory_FromKilobits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobits(100)
    if err != nil {
        t.Errorf("FromKilobits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationKilobit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobits(math.NaN())
    if err == nil {
        t.Error("FromKilobits() with NaN value should return error")
    }

    _, err = factory.FromKilobits(math.Inf(1))
    if err == nil {
        t.Error("FromKilobits() with +Inf value should return error")
    }

    _, err = factory.FromKilobits(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobits(0)
    if err != nil {
        t.Errorf("FromKilobits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationKilobit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobits() with zero value = %v, want 0", converted)
    }
}
// Test FromMegabits function
func TestInformationFactory_FromMegabits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegabits(100)
    if err != nil {
        t.Errorf("FromMegabits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationMegabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegabits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegabits(math.NaN())
    if err == nil {
        t.Error("FromMegabits() with NaN value should return error")
    }

    _, err = factory.FromMegabits(math.Inf(1))
    if err == nil {
        t.Error("FromMegabits() with +Inf value should return error")
    }

    _, err = factory.FromMegabits(math.Inf(-1))
    if err == nil {
        t.Error("FromMegabits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegabits(0)
    if err != nil {
        t.Errorf("FromMegabits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationMegabit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegabits() with zero value = %v, want 0", converted)
    }
}
// Test FromGigabits function
func TestInformationFactory_FromGigabits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigabits(100)
    if err != nil {
        t.Errorf("FromGigabits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationGigabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigabits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigabits(math.NaN())
    if err == nil {
        t.Error("FromGigabits() with NaN value should return error")
    }

    _, err = factory.FromGigabits(math.Inf(1))
    if err == nil {
        t.Error("FromGigabits() with +Inf value should return error")
    }

    _, err = factory.FromGigabits(math.Inf(-1))
    if err == nil {
        t.Error("FromGigabits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigabits(0)
    if err != nil {
        t.Errorf("FromGigabits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationGigabit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigabits() with zero value = %v, want 0", converted)
    }
}
// Test FromTerabits function
func TestInformationFactory_FromTerabits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerabits(100)
    if err != nil {
        t.Errorf("FromTerabits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationTerabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerabits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerabits(math.NaN())
    if err == nil {
        t.Error("FromTerabits() with NaN value should return error")
    }

    _, err = factory.FromTerabits(math.Inf(1))
    if err == nil {
        t.Error("FromTerabits() with +Inf value should return error")
    }

    _, err = factory.FromTerabits(math.Inf(-1))
    if err == nil {
        t.Error("FromTerabits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerabits(0)
    if err != nil {
        t.Errorf("FromTerabits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationTerabit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerabits() with zero value = %v, want 0", converted)
    }
}
// Test FromPetabits function
func TestInformationFactory_FromPetabits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetabits(100)
    if err != nil {
        t.Errorf("FromPetabits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationPetabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetabits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetabits(math.NaN())
    if err == nil {
        t.Error("FromPetabits() with NaN value should return error")
    }

    _, err = factory.FromPetabits(math.Inf(1))
    if err == nil {
        t.Error("FromPetabits() with +Inf value should return error")
    }

    _, err = factory.FromPetabits(math.Inf(-1))
    if err == nil {
        t.Error("FromPetabits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetabits(0)
    if err != nil {
        t.Errorf("FromPetabits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationPetabit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetabits() with zero value = %v, want 0", converted)
    }
}
// Test FromExabits function
func TestInformationFactory_FromExabits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExabits(100)
    if err != nil {
        t.Errorf("FromExabits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationExabit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExabits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExabits(math.NaN())
    if err == nil {
        t.Error("FromExabits() with NaN value should return error")
    }

    _, err = factory.FromExabits(math.Inf(1))
    if err == nil {
        t.Error("FromExabits() with +Inf value should return error")
    }

    _, err = factory.FromExabits(math.Inf(-1))
    if err == nil {
        t.Error("FromExabits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExabits(0)
    if err != nil {
        t.Errorf("FromExabits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationExabit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExabits() with zero value = %v, want 0", converted)
    }
}
// Test FromKibibits function
func TestInformationFactory_FromKibibits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKibibits(100)
    if err != nil {
        t.Errorf("FromKibibits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationKibibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKibibits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKibibits(math.NaN())
    if err == nil {
        t.Error("FromKibibits() with NaN value should return error")
    }

    _, err = factory.FromKibibits(math.Inf(1))
    if err == nil {
        t.Error("FromKibibits() with +Inf value should return error")
    }

    _, err = factory.FromKibibits(math.Inf(-1))
    if err == nil {
        t.Error("FromKibibits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKibibits(0)
    if err != nil {
        t.Errorf("FromKibibits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationKibibit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKibibits() with zero value = %v, want 0", converted)
    }
}
// Test FromMebibits function
func TestInformationFactory_FromMebibits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMebibits(100)
    if err != nil {
        t.Errorf("FromMebibits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationMebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMebibits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMebibits(math.NaN())
    if err == nil {
        t.Error("FromMebibits() with NaN value should return error")
    }

    _, err = factory.FromMebibits(math.Inf(1))
    if err == nil {
        t.Error("FromMebibits() with +Inf value should return error")
    }

    _, err = factory.FromMebibits(math.Inf(-1))
    if err == nil {
        t.Error("FromMebibits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMebibits(0)
    if err != nil {
        t.Errorf("FromMebibits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationMebibit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMebibits() with zero value = %v, want 0", converted)
    }
}
// Test FromGibibits function
func TestInformationFactory_FromGibibits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGibibits(100)
    if err != nil {
        t.Errorf("FromGibibits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationGibibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGibibits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGibibits(math.NaN())
    if err == nil {
        t.Error("FromGibibits() with NaN value should return error")
    }

    _, err = factory.FromGibibits(math.Inf(1))
    if err == nil {
        t.Error("FromGibibits() with +Inf value should return error")
    }

    _, err = factory.FromGibibits(math.Inf(-1))
    if err == nil {
        t.Error("FromGibibits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGibibits(0)
    if err != nil {
        t.Errorf("FromGibibits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationGibibit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGibibits() with zero value = %v, want 0", converted)
    }
}
// Test FromTebibits function
func TestInformationFactory_FromTebibits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTebibits(100)
    if err != nil {
        t.Errorf("FromTebibits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationTebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTebibits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTebibits(math.NaN())
    if err == nil {
        t.Error("FromTebibits() with NaN value should return error")
    }

    _, err = factory.FromTebibits(math.Inf(1))
    if err == nil {
        t.Error("FromTebibits() with +Inf value should return error")
    }

    _, err = factory.FromTebibits(math.Inf(-1))
    if err == nil {
        t.Error("FromTebibits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTebibits(0)
    if err != nil {
        t.Errorf("FromTebibits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationTebibit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTebibits() with zero value = %v, want 0", converted)
    }
}
// Test FromPebibits function
func TestInformationFactory_FromPebibits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPebibits(100)
    if err != nil {
        t.Errorf("FromPebibits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationPebibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPebibits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPebibits(math.NaN())
    if err == nil {
        t.Error("FromPebibits() with NaN value should return error")
    }

    _, err = factory.FromPebibits(math.Inf(1))
    if err == nil {
        t.Error("FromPebibits() with +Inf value should return error")
    }

    _, err = factory.FromPebibits(math.Inf(-1))
    if err == nil {
        t.Error("FromPebibits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPebibits(0)
    if err != nil {
        t.Errorf("FromPebibits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationPebibit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPebibits() with zero value = %v, want 0", converted)
    }
}
// Test FromExbibits function
func TestInformationFactory_FromExbibits(t *testing.T) {
    factory := units.InformationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExbibits(100)
    if err != nil {
        t.Errorf("FromExbibits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.InformationExbibit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExbibits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExbibits(math.NaN())
    if err == nil {
        t.Error("FromExbibits() with NaN value should return error")
    }

    _, err = factory.FromExbibits(math.Inf(1))
    if err == nil {
        t.Error("FromExbibits() with +Inf value should return error")
    }

    _, err = factory.FromExbibits(math.Inf(-1))
    if err == nil {
        t.Error("FromExbibits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExbibits(0)
    if err != nil {
        t.Errorf("FromExbibits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.InformationExbibit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExbibits() with zero value = %v, want 0", converted)
    }
}

func TestInformationToString(t *testing.T) {
	factory := units.InformationFactory{}
	a, err := factory.CreateInformation(45, units.InformationBit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.InformationBit, 2)
	expected := "45.00 " + units.GetInformationAbbreviation(units.InformationBit)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.InformationBit, -1)
	expected = "45 " + units.GetInformationAbbreviation(units.InformationBit)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestInformation_EqualityAndComparison(t *testing.T) {
	factory := units.InformationFactory{}
	a1, _ := factory.CreateInformation(60, units.InformationBit)
	a2, _ := factory.CreateInformation(60, units.InformationBit)
	a3, _ := factory.CreateInformation(120, units.InformationBit)

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

func TestInformation_Arithmetic(t *testing.T) {
	factory := units.InformationFactory{}
	a1, _ := factory.CreateInformation(30, units.InformationBit)
	a2, _ := factory.CreateInformation(45, units.InformationBit)

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


func TestGetInformationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.InformationUnits
        want string
    }{
        {
            name: "Byte abbreviation",
            unit: units.InformationByte,
            want: "B",
        },
        {
            name: "Bit abbreviation",
            unit: units.InformationBit,
            want: "b",
        },
        {
            name: "Kilobyte abbreviation",
            unit: units.InformationKilobyte,
            want: "kB",
        },
        {
            name: "Megabyte abbreviation",
            unit: units.InformationMegabyte,
            want: "MB",
        },
        {
            name: "Gigabyte abbreviation",
            unit: units.InformationGigabyte,
            want: "GB",
        },
        {
            name: "Terabyte abbreviation",
            unit: units.InformationTerabyte,
            want: "TB",
        },
        {
            name: "Petabyte abbreviation",
            unit: units.InformationPetabyte,
            want: "PB",
        },
        {
            name: "Exabyte abbreviation",
            unit: units.InformationExabyte,
            want: "EB",
        },
        {
            name: "Kibibyte abbreviation",
            unit: units.InformationKibibyte,
            want: "KiBB",
        },
        {
            name: "Mebibyte abbreviation",
            unit: units.InformationMebibyte,
            want: "MiBB",
        },
        {
            name: "Gibibyte abbreviation",
            unit: units.InformationGibibyte,
            want: "GiBB",
        },
        {
            name: "Tebibyte abbreviation",
            unit: units.InformationTebibyte,
            want: "TiBB",
        },
        {
            name: "Pebibyte abbreviation",
            unit: units.InformationPebibyte,
            want: "PiBB",
        },
        {
            name: "Exbibyte abbreviation",
            unit: units.InformationExbibyte,
            want: "EiBB",
        },
        {
            name: "Kilobit abbreviation",
            unit: units.InformationKilobit,
            want: "kb",
        },
        {
            name: "Megabit abbreviation",
            unit: units.InformationMegabit,
            want: "Mb",
        },
        {
            name: "Gigabit abbreviation",
            unit: units.InformationGigabit,
            want: "Gb",
        },
        {
            name: "Terabit abbreviation",
            unit: units.InformationTerabit,
            want: "Tb",
        },
        {
            name: "Petabit abbreviation",
            unit: units.InformationPetabit,
            want: "Pb",
        },
        {
            name: "Exabit abbreviation",
            unit: units.InformationExabit,
            want: "Eb",
        },
        {
            name: "Kibibit abbreviation",
            unit: units.InformationKibibit,
            want: "KiBb",
        },
        {
            name: "Mebibit abbreviation",
            unit: units.InformationMebibit,
            want: "MiBb",
        },
        {
            name: "Gibibit abbreviation",
            unit: units.InformationGibibit,
            want: "GiBb",
        },
        {
            name: "Tebibit abbreviation",
            unit: units.InformationTebibit,
            want: "TiBb",
        },
        {
            name: "Pebibit abbreviation",
            unit: units.InformationPebibit,
            want: "PiBb",
        },
        {
            name: "Exbibit abbreviation",
            unit: units.InformationExbibit,
            want: "EiBb",
        },
        {
            name: "invalid unit",
            unit: units.InformationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetInformationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetInformationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestInformation_String(t *testing.T) {
    factory := units.InformationFactory{}
    
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
            unit, err := factory.CreateInformation(tt.value, units.InformationBit)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Information.String() = %v, want %v", got, tt.want)
            }
        })
    }
}