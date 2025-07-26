// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestBitRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "BitPerSecond"}`
	
	factory := units.BitRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.BitRateBitPerSecond {
		t.Errorf("expected unit %v, got %v", units.BitRateBitPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "BitPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestBitRateDto_ToJSON(t *testing.T) {
	dto := units.BitRateDto{
		Value: 45,
		Unit:  units.BitRateBitPerSecond,
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
	if result["unit"].(string) != string(units.BitRateBitPerSecond) {
		t.Errorf("expected unit %s, got %v", units.BitRateBitPerSecond, result["unit"])
	}
}

func TestNewBitRate_InvalidValue(t *testing.T) {
	factory := units.BitRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateBitRate(math.NaN(), units.BitRateBitPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateBitRate(math.Inf(1), units.BitRateBitPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestBitRateConversions(t *testing.T) {
	factory := units.BitRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateBitRate(180, units.BitRateBitPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to BitsPerSecond.
		// No expected conversion value provided for BitsPerSecond, verifying result is not NaN.
		result := a.BitsPerSecond()
		cacheResult := a.BitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to BytesPerSecond.
		// No expected conversion value provided for BytesPerSecond, verifying result is not NaN.
		result := a.BytesPerSecond()
		cacheResult := a.BytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to OctetsPerSecond.
		// No expected conversion value provided for OctetsPerSecond, verifying result is not NaN.
		result := a.OctetsPerSecond()
		cacheResult := a.OctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilobitsPerSecond.
		// No expected conversion value provided for KilobitsPerSecond, verifying result is not NaN.
		result := a.KilobitsPerSecond()
		cacheResult := a.KilobitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilobitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegabitsPerSecond.
		// No expected conversion value provided for MegabitsPerSecond, verifying result is not NaN.
		result := a.MegabitsPerSecond()
		cacheResult := a.MegabitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GigabitsPerSecond.
		// No expected conversion value provided for GigabitsPerSecond, verifying result is not NaN.
		result := a.GigabitsPerSecond()
		cacheResult := a.GigabitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TerabitsPerSecond.
		// No expected conversion value provided for TerabitsPerSecond, verifying result is not NaN.
		result := a.TerabitsPerSecond()
		cacheResult := a.TerabitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PetabitsPerSecond.
		// No expected conversion value provided for PetabitsPerSecond, verifying result is not NaN.
		result := a.PetabitsPerSecond()
		cacheResult := a.PetabitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PetabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExabitsPerSecond.
		// No expected conversion value provided for ExabitsPerSecond, verifying result is not NaN.
		result := a.ExabitsPerSecond()
		cacheResult := a.ExabitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ExabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KibibitsPerSecond.
		// No expected conversion value provided for KibibitsPerSecond, verifying result is not NaN.
		result := a.KibibitsPerSecond()
		cacheResult := a.KibibitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KibibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MebibitsPerSecond.
		// No expected conversion value provided for MebibitsPerSecond, verifying result is not NaN.
		result := a.MebibitsPerSecond()
		cacheResult := a.MebibitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MebibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GibibitsPerSecond.
		// No expected conversion value provided for GibibitsPerSecond, verifying result is not NaN.
		result := a.GibibitsPerSecond()
		cacheResult := a.GibibitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GibibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TebibitsPerSecond.
		// No expected conversion value provided for TebibitsPerSecond, verifying result is not NaN.
		result := a.TebibitsPerSecond()
		cacheResult := a.TebibitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TebibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PebibitsPerSecond.
		// No expected conversion value provided for PebibitsPerSecond, verifying result is not NaN.
		result := a.PebibitsPerSecond()
		cacheResult := a.PebibitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PebibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExbibitsPerSecond.
		// No expected conversion value provided for ExbibitsPerSecond, verifying result is not NaN.
		result := a.ExbibitsPerSecond()
		cacheResult := a.ExbibitsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ExbibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilobytesPerSecond.
		// No expected conversion value provided for KilobytesPerSecond, verifying result is not NaN.
		result := a.KilobytesPerSecond()
		cacheResult := a.KilobytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilobytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegabytesPerSecond.
		// No expected conversion value provided for MegabytesPerSecond, verifying result is not NaN.
		result := a.MegabytesPerSecond()
		cacheResult := a.MegabytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GigabytesPerSecond.
		// No expected conversion value provided for GigabytesPerSecond, verifying result is not NaN.
		result := a.GigabytesPerSecond()
		cacheResult := a.GigabytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TerabytesPerSecond.
		// No expected conversion value provided for TerabytesPerSecond, verifying result is not NaN.
		result := a.TerabytesPerSecond()
		cacheResult := a.TerabytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PetabytesPerSecond.
		// No expected conversion value provided for PetabytesPerSecond, verifying result is not NaN.
		result := a.PetabytesPerSecond()
		cacheResult := a.PetabytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PetabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExabytesPerSecond.
		// No expected conversion value provided for ExabytesPerSecond, verifying result is not NaN.
		result := a.ExabytesPerSecond()
		cacheResult := a.ExabytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ExabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KibibytesPerSecond.
		// No expected conversion value provided for KibibytesPerSecond, verifying result is not NaN.
		result := a.KibibytesPerSecond()
		cacheResult := a.KibibytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KibibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MebibytesPerSecond.
		// No expected conversion value provided for MebibytesPerSecond, verifying result is not NaN.
		result := a.MebibytesPerSecond()
		cacheResult := a.MebibytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MebibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GibibytesPerSecond.
		// No expected conversion value provided for GibibytesPerSecond, verifying result is not NaN.
		result := a.GibibytesPerSecond()
		cacheResult := a.GibibytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GibibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TebibytesPerSecond.
		// No expected conversion value provided for TebibytesPerSecond, verifying result is not NaN.
		result := a.TebibytesPerSecond()
		cacheResult := a.TebibytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TebibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PebibytesPerSecond.
		// No expected conversion value provided for PebibytesPerSecond, verifying result is not NaN.
		result := a.PebibytesPerSecond()
		cacheResult := a.PebibytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PebibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExbibytesPerSecond.
		// No expected conversion value provided for ExbibytesPerSecond, verifying result is not NaN.
		result := a.ExbibytesPerSecond()
		cacheResult := a.ExbibytesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ExbibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilooctetsPerSecond.
		// No expected conversion value provided for KilooctetsPerSecond, verifying result is not NaN.
		result := a.KilooctetsPerSecond()
		cacheResult := a.KilooctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilooctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegaoctetsPerSecond.
		// No expected conversion value provided for MegaoctetsPerSecond, verifying result is not NaN.
		result := a.MegaoctetsPerSecond()
		cacheResult := a.MegaoctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegaoctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GigaoctetsPerSecond.
		// No expected conversion value provided for GigaoctetsPerSecond, verifying result is not NaN.
		result := a.GigaoctetsPerSecond()
		cacheResult := a.GigaoctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigaoctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TeraoctetsPerSecond.
		// No expected conversion value provided for TeraoctetsPerSecond, verifying result is not NaN.
		result := a.TeraoctetsPerSecond()
		cacheResult := a.TeraoctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TeraoctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PetaoctetsPerSecond.
		// No expected conversion value provided for PetaoctetsPerSecond, verifying result is not NaN.
		result := a.PetaoctetsPerSecond()
		cacheResult := a.PetaoctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PetaoctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExaoctetsPerSecond.
		// No expected conversion value provided for ExaoctetsPerSecond, verifying result is not NaN.
		result := a.ExaoctetsPerSecond()
		cacheResult := a.ExaoctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ExaoctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KibioctetsPerSecond.
		// No expected conversion value provided for KibioctetsPerSecond, verifying result is not NaN.
		result := a.KibioctetsPerSecond()
		cacheResult := a.KibioctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KibioctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MebioctetsPerSecond.
		// No expected conversion value provided for MebioctetsPerSecond, verifying result is not NaN.
		result := a.MebioctetsPerSecond()
		cacheResult := a.MebioctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MebioctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GibioctetsPerSecond.
		// No expected conversion value provided for GibioctetsPerSecond, verifying result is not NaN.
		result := a.GibioctetsPerSecond()
		cacheResult := a.GibioctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GibioctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TebioctetsPerSecond.
		// No expected conversion value provided for TebioctetsPerSecond, verifying result is not NaN.
		result := a.TebioctetsPerSecond()
		cacheResult := a.TebioctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TebioctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PebioctetsPerSecond.
		// No expected conversion value provided for PebioctetsPerSecond, verifying result is not NaN.
		result := a.PebioctetsPerSecond()
		cacheResult := a.PebioctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PebioctetsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExbioctetsPerSecond.
		// No expected conversion value provided for ExbioctetsPerSecond, verifying result is not NaN.
		result := a.ExbioctetsPerSecond()
		cacheResult := a.ExbioctetsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ExbioctetsPerSecond returned NaN")
		}
	}
}

func TestBitRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.BitRateFactory{}
	a, err := factory.CreateBitRate(90, units.BitRateBitPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.BitRateBitPerSecond {
		t.Errorf("expected default unit BitPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.BitRateBitPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.BitRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.BitRateBitPerSecond {
		t.Errorf("expected unit BitPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestBitRateFactory_FromDto(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateBitPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.BitRateDto{
        Value: math.NaN(),
        Unit:  units.BitRateBitPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test BitPerSecond conversion
    bits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateBitPerSecond,
    }
    
    var bits_per_secondResult *units.BitRate
    bits_per_secondResult, err = factory.FromDto(bits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with BitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bits_per_secondResult.Convert(units.BitRateBitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BitPerSecond = %v, want %v", converted, 100)
    }
    // Test BytePerSecond conversion
    bytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateBytePerSecond,
    }
    
    var bytes_per_secondResult *units.BitRate
    bytes_per_secondResult, err = factory.FromDto(bytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with BytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bytes_per_secondResult.Convert(units.BitRateBytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BytePerSecond = %v, want %v", converted, 100)
    }
    // Test OctetPerSecond conversion
    octets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateOctetPerSecond,
    }
    
    var octets_per_secondResult *units.BitRate
    octets_per_secondResult, err = factory.FromDto(octets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with OctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = octets_per_secondResult.Convert(units.BitRateOctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OctetPerSecond = %v, want %v", converted, 100)
    }
    // Test KilobitPerSecond conversion
    kilobits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateKilobitPerSecond,
    }
    
    var kilobits_per_secondResult *units.BitRate
    kilobits_per_secondResult, err = factory.FromDto(kilobits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilobitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobits_per_secondResult.Convert(units.BitRateKilobitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobitPerSecond = %v, want %v", converted, 100)
    }
    // Test MegabitPerSecond conversion
    megabits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateMegabitPerSecond,
    }
    
    var megabits_per_secondResult *units.BitRate
    megabits_per_secondResult, err = factory.FromDto(megabits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegabitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabits_per_secondResult.Convert(units.BitRateMegabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegabitPerSecond = %v, want %v", converted, 100)
    }
    // Test GigabitPerSecond conversion
    gigabits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateGigabitPerSecond,
    }
    
    var gigabits_per_secondResult *units.BitRate
    gigabits_per_secondResult, err = factory.FromDto(gigabits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with GigabitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabits_per_secondResult.Convert(units.BitRateGigabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigabitPerSecond = %v, want %v", converted, 100)
    }
    // Test TerabitPerSecond conversion
    terabits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateTerabitPerSecond,
    }
    
    var terabits_per_secondResult *units.BitRate
    terabits_per_secondResult, err = factory.FromDto(terabits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with TerabitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabits_per_secondResult.Convert(units.BitRateTerabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerabitPerSecond = %v, want %v", converted, 100)
    }
    // Test PetabitPerSecond conversion
    petabits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRatePetabitPerSecond,
    }
    
    var petabits_per_secondResult *units.BitRate
    petabits_per_secondResult, err = factory.FromDto(petabits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PetabitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabits_per_secondResult.Convert(units.BitRatePetabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetabitPerSecond = %v, want %v", converted, 100)
    }
    // Test ExabitPerSecond conversion
    exabits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateExabitPerSecond,
    }
    
    var exabits_per_secondResult *units.BitRate
    exabits_per_secondResult, err = factory.FromDto(exabits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with ExabitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabits_per_secondResult.Convert(units.BitRateExabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExabitPerSecond = %v, want %v", converted, 100)
    }
    // Test KibibitPerSecond conversion
    kibibits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateKibibitPerSecond,
    }
    
    var kibibits_per_secondResult *units.BitRate
    kibibits_per_secondResult, err = factory.FromDto(kibibits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KibibitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibibits_per_secondResult.Convert(units.BitRateKibibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KibibitPerSecond = %v, want %v", converted, 100)
    }
    // Test MebibitPerSecond conversion
    mebibits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateMebibitPerSecond,
    }
    
    var mebibits_per_secondResult *units.BitRate
    mebibits_per_secondResult, err = factory.FromDto(mebibits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MebibitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebibits_per_secondResult.Convert(units.BitRateMebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MebibitPerSecond = %v, want %v", converted, 100)
    }
    // Test GibibitPerSecond conversion
    gibibits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateGibibitPerSecond,
    }
    
    var gibibits_per_secondResult *units.BitRate
    gibibits_per_secondResult, err = factory.FromDto(gibibits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with GibibitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibibits_per_secondResult.Convert(units.BitRateGibibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GibibitPerSecond = %v, want %v", converted, 100)
    }
    // Test TebibitPerSecond conversion
    tebibits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateTebibitPerSecond,
    }
    
    var tebibits_per_secondResult *units.BitRate
    tebibits_per_secondResult, err = factory.FromDto(tebibits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with TebibitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebibits_per_secondResult.Convert(units.BitRateTebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TebibitPerSecond = %v, want %v", converted, 100)
    }
    // Test PebibitPerSecond conversion
    pebibits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRatePebibitPerSecond,
    }
    
    var pebibits_per_secondResult *units.BitRate
    pebibits_per_secondResult, err = factory.FromDto(pebibits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PebibitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebibits_per_secondResult.Convert(units.BitRatePebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PebibitPerSecond = %v, want %v", converted, 100)
    }
    // Test ExbibitPerSecond conversion
    exbibits_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateExbibitPerSecond,
    }
    
    var exbibits_per_secondResult *units.BitRate
    exbibits_per_secondResult, err = factory.FromDto(exbibits_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with ExbibitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbibits_per_secondResult.Convert(units.BitRateExbibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExbibitPerSecond = %v, want %v", converted, 100)
    }
    // Test KilobytePerSecond conversion
    kilobytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateKilobytePerSecond,
    }
    
    var kilobytes_per_secondResult *units.BitRate
    kilobytes_per_secondResult, err = factory.FromDto(kilobytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilobytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobytes_per_secondResult.Convert(units.BitRateKilobytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobytePerSecond = %v, want %v", converted, 100)
    }
    // Test MegabytePerSecond conversion
    megabytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateMegabytePerSecond,
    }
    
    var megabytes_per_secondResult *units.BitRate
    megabytes_per_secondResult, err = factory.FromDto(megabytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegabytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabytes_per_secondResult.Convert(units.BitRateMegabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegabytePerSecond = %v, want %v", converted, 100)
    }
    // Test GigabytePerSecond conversion
    gigabytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateGigabytePerSecond,
    }
    
    var gigabytes_per_secondResult *units.BitRate
    gigabytes_per_secondResult, err = factory.FromDto(gigabytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with GigabytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabytes_per_secondResult.Convert(units.BitRateGigabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigabytePerSecond = %v, want %v", converted, 100)
    }
    // Test TerabytePerSecond conversion
    terabytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateTerabytePerSecond,
    }
    
    var terabytes_per_secondResult *units.BitRate
    terabytes_per_secondResult, err = factory.FromDto(terabytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with TerabytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabytes_per_secondResult.Convert(units.BitRateTerabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerabytePerSecond = %v, want %v", converted, 100)
    }
    // Test PetabytePerSecond conversion
    petabytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRatePetabytePerSecond,
    }
    
    var petabytes_per_secondResult *units.BitRate
    petabytes_per_secondResult, err = factory.FromDto(petabytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PetabytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabytes_per_secondResult.Convert(units.BitRatePetabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetabytePerSecond = %v, want %v", converted, 100)
    }
    // Test ExabytePerSecond conversion
    exabytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateExabytePerSecond,
    }
    
    var exabytes_per_secondResult *units.BitRate
    exabytes_per_secondResult, err = factory.FromDto(exabytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with ExabytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabytes_per_secondResult.Convert(units.BitRateExabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExabytePerSecond = %v, want %v", converted, 100)
    }
    // Test KibibytePerSecond conversion
    kibibytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateKibibytePerSecond,
    }
    
    var kibibytes_per_secondResult *units.BitRate
    kibibytes_per_secondResult, err = factory.FromDto(kibibytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KibibytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibibytes_per_secondResult.Convert(units.BitRateKibibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KibibytePerSecond = %v, want %v", converted, 100)
    }
    // Test MebibytePerSecond conversion
    mebibytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateMebibytePerSecond,
    }
    
    var mebibytes_per_secondResult *units.BitRate
    mebibytes_per_secondResult, err = factory.FromDto(mebibytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MebibytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebibytes_per_secondResult.Convert(units.BitRateMebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MebibytePerSecond = %v, want %v", converted, 100)
    }
    // Test GibibytePerSecond conversion
    gibibytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateGibibytePerSecond,
    }
    
    var gibibytes_per_secondResult *units.BitRate
    gibibytes_per_secondResult, err = factory.FromDto(gibibytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with GibibytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibibytes_per_secondResult.Convert(units.BitRateGibibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GibibytePerSecond = %v, want %v", converted, 100)
    }
    // Test TebibytePerSecond conversion
    tebibytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateTebibytePerSecond,
    }
    
    var tebibytes_per_secondResult *units.BitRate
    tebibytes_per_secondResult, err = factory.FromDto(tebibytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with TebibytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebibytes_per_secondResult.Convert(units.BitRateTebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TebibytePerSecond = %v, want %v", converted, 100)
    }
    // Test PebibytePerSecond conversion
    pebibytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRatePebibytePerSecond,
    }
    
    var pebibytes_per_secondResult *units.BitRate
    pebibytes_per_secondResult, err = factory.FromDto(pebibytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PebibytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebibytes_per_secondResult.Convert(units.BitRatePebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PebibytePerSecond = %v, want %v", converted, 100)
    }
    // Test ExbibytePerSecond conversion
    exbibytes_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateExbibytePerSecond,
    }
    
    var exbibytes_per_secondResult *units.BitRate
    exbibytes_per_secondResult, err = factory.FromDto(exbibytes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with ExbibytePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbibytes_per_secondResult.Convert(units.BitRateExbibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExbibytePerSecond = %v, want %v", converted, 100)
    }
    // Test KilooctetPerSecond conversion
    kilooctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateKilooctetPerSecond,
    }
    
    var kilooctets_per_secondResult *units.BitRate
    kilooctets_per_secondResult, err = factory.FromDto(kilooctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilooctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilooctets_per_secondResult.Convert(units.BitRateKilooctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilooctetPerSecond = %v, want %v", converted, 100)
    }
    // Test MegaoctetPerSecond conversion
    megaoctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateMegaoctetPerSecond,
    }
    
    var megaoctets_per_secondResult *units.BitRate
    megaoctets_per_secondResult, err = factory.FromDto(megaoctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegaoctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaoctets_per_secondResult.Convert(units.BitRateMegaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test GigaoctetPerSecond conversion
    gigaoctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateGigaoctetPerSecond,
    }
    
    var gigaoctets_per_secondResult *units.BitRate
    gigaoctets_per_secondResult, err = factory.FromDto(gigaoctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with GigaoctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigaoctets_per_secondResult.Convert(units.BitRateGigaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigaoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test TeraoctetPerSecond conversion
    teraoctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateTeraoctetPerSecond,
    }
    
    var teraoctets_per_secondResult *units.BitRate
    teraoctets_per_secondResult, err = factory.FromDto(teraoctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with TeraoctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teraoctets_per_secondResult.Convert(units.BitRateTeraoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TeraoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test PetaoctetPerSecond conversion
    petaoctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRatePetaoctetPerSecond,
    }
    
    var petaoctets_per_secondResult *units.BitRate
    petaoctets_per_secondResult, err = factory.FromDto(petaoctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PetaoctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petaoctets_per_secondResult.Convert(units.BitRatePetaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetaoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test ExaoctetPerSecond conversion
    exaoctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateExaoctetPerSecond,
    }
    
    var exaoctets_per_secondResult *units.BitRate
    exaoctets_per_secondResult, err = factory.FromDto(exaoctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with ExaoctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exaoctets_per_secondResult.Convert(units.BitRateExaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExaoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test KibioctetPerSecond conversion
    kibioctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateKibioctetPerSecond,
    }
    
    var kibioctets_per_secondResult *units.BitRate
    kibioctets_per_secondResult, err = factory.FromDto(kibioctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KibioctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibioctets_per_secondResult.Convert(units.BitRateKibioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KibioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test MebioctetPerSecond conversion
    mebioctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateMebioctetPerSecond,
    }
    
    var mebioctets_per_secondResult *units.BitRate
    mebioctets_per_secondResult, err = factory.FromDto(mebioctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MebioctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebioctets_per_secondResult.Convert(units.BitRateMebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MebioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test GibioctetPerSecond conversion
    gibioctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateGibioctetPerSecond,
    }
    
    var gibioctets_per_secondResult *units.BitRate
    gibioctets_per_secondResult, err = factory.FromDto(gibioctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with GibioctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibioctets_per_secondResult.Convert(units.BitRateGibioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GibioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test TebioctetPerSecond conversion
    tebioctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateTebioctetPerSecond,
    }
    
    var tebioctets_per_secondResult *units.BitRate
    tebioctets_per_secondResult, err = factory.FromDto(tebioctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with TebioctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebioctets_per_secondResult.Convert(units.BitRateTebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TebioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test PebioctetPerSecond conversion
    pebioctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRatePebioctetPerSecond,
    }
    
    var pebioctets_per_secondResult *units.BitRate
    pebioctets_per_secondResult, err = factory.FromDto(pebioctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PebioctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebioctets_per_secondResult.Convert(units.BitRatePebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PebioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test ExbioctetPerSecond conversion
    exbioctets_per_secondDto := units.BitRateDto{
        Value: 100,
        Unit:  units.BitRateExbioctetPerSecond,
    }
    
    var exbioctets_per_secondResult *units.BitRate
    exbioctets_per_secondResult, err = factory.FromDto(exbioctets_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with ExbioctetPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbioctets_per_secondResult.Convert(units.BitRateExbioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExbioctetPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.BitRateDto{
        Value: 0,
        Unit:  units.BitRateBitPerSecond,
    }
    
    var zeroResult *units.BitRate
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestBitRateFactory_FromDtoJSON(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "BitPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "BitPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.BitRateDto{
        Value: nanValue,
        Unit:  units.BitRateBitPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with BitPerSecond unit
    bits_per_secondJSON := []byte(`{"value": 100, "unit": "BitPerSecond"}`)
    bits_per_secondResult, err := factory.FromDtoJSON(bits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bits_per_secondResult.Convert(units.BitRateBitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with BytePerSecond unit
    bytes_per_secondJSON := []byte(`{"value": 100, "unit": "BytePerSecond"}`)
    bytes_per_secondResult, err := factory.FromDtoJSON(bytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bytes_per_secondResult.Convert(units.BitRateBytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with OctetPerSecond unit
    octets_per_secondJSON := []byte(`{"value": 100, "unit": "OctetPerSecond"}`)
    octets_per_secondResult, err := factory.FromDtoJSON(octets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = octets_per_secondResult.Convert(units.BitRateOctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilobitPerSecond unit
    kilobits_per_secondJSON := []byte(`{"value": 100, "unit": "KilobitPerSecond"}`)
    kilobits_per_secondResult, err := factory.FromDtoJSON(kilobits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilobitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobits_per_secondResult.Convert(units.BitRateKilobitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegabitPerSecond unit
    megabits_per_secondJSON := []byte(`{"value": 100, "unit": "MegabitPerSecond"}`)
    megabits_per_secondResult, err := factory.FromDtoJSON(megabits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegabitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabits_per_secondResult.Convert(units.BitRateMegabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegabitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with GigabitPerSecond unit
    gigabits_per_secondJSON := []byte(`{"value": 100, "unit": "GigabitPerSecond"}`)
    gigabits_per_secondResult, err := factory.FromDtoJSON(gigabits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigabitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabits_per_secondResult.Convert(units.BitRateGigabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigabitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with TerabitPerSecond unit
    terabits_per_secondJSON := []byte(`{"value": 100, "unit": "TerabitPerSecond"}`)
    terabits_per_secondResult, err := factory.FromDtoJSON(terabits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerabitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabits_per_secondResult.Convert(units.BitRateTerabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerabitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PetabitPerSecond unit
    petabits_per_secondJSON := []byte(`{"value": 100, "unit": "PetabitPerSecond"}`)
    petabits_per_secondResult, err := factory.FromDtoJSON(petabits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PetabitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabits_per_secondResult.Convert(units.BitRatePetabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetabitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with ExabitPerSecond unit
    exabits_per_secondJSON := []byte(`{"value": 100, "unit": "ExabitPerSecond"}`)
    exabits_per_secondResult, err := factory.FromDtoJSON(exabits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ExabitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabits_per_secondResult.Convert(units.BitRateExabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExabitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KibibitPerSecond unit
    kibibits_per_secondJSON := []byte(`{"value": 100, "unit": "KibibitPerSecond"}`)
    kibibits_per_secondResult, err := factory.FromDtoJSON(kibibits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KibibitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibibits_per_secondResult.Convert(units.BitRateKibibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KibibitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MebibitPerSecond unit
    mebibits_per_secondJSON := []byte(`{"value": 100, "unit": "MebibitPerSecond"}`)
    mebibits_per_secondResult, err := factory.FromDtoJSON(mebibits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MebibitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebibits_per_secondResult.Convert(units.BitRateMebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MebibitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with GibibitPerSecond unit
    gibibits_per_secondJSON := []byte(`{"value": 100, "unit": "GibibitPerSecond"}`)
    gibibits_per_secondResult, err := factory.FromDtoJSON(gibibits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GibibitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibibits_per_secondResult.Convert(units.BitRateGibibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GibibitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with TebibitPerSecond unit
    tebibits_per_secondJSON := []byte(`{"value": 100, "unit": "TebibitPerSecond"}`)
    tebibits_per_secondResult, err := factory.FromDtoJSON(tebibits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TebibitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebibits_per_secondResult.Convert(units.BitRateTebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TebibitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PebibitPerSecond unit
    pebibits_per_secondJSON := []byte(`{"value": 100, "unit": "PebibitPerSecond"}`)
    pebibits_per_secondResult, err := factory.FromDtoJSON(pebibits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PebibitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebibits_per_secondResult.Convert(units.BitRatePebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PebibitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with ExbibitPerSecond unit
    exbibits_per_secondJSON := []byte(`{"value": 100, "unit": "ExbibitPerSecond"}`)
    exbibits_per_secondResult, err := factory.FromDtoJSON(exbibits_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ExbibitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbibits_per_secondResult.Convert(units.BitRateExbibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExbibitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilobytePerSecond unit
    kilobytes_per_secondJSON := []byte(`{"value": 100, "unit": "KilobytePerSecond"}`)
    kilobytes_per_secondResult, err := factory.FromDtoJSON(kilobytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilobytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobytes_per_secondResult.Convert(units.BitRateKilobytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegabytePerSecond unit
    megabytes_per_secondJSON := []byte(`{"value": 100, "unit": "MegabytePerSecond"}`)
    megabytes_per_secondResult, err := factory.FromDtoJSON(megabytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegabytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabytes_per_secondResult.Convert(units.BitRateMegabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegabytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with GigabytePerSecond unit
    gigabytes_per_secondJSON := []byte(`{"value": 100, "unit": "GigabytePerSecond"}`)
    gigabytes_per_secondResult, err := factory.FromDtoJSON(gigabytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigabytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabytes_per_secondResult.Convert(units.BitRateGigabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigabytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with TerabytePerSecond unit
    terabytes_per_secondJSON := []byte(`{"value": 100, "unit": "TerabytePerSecond"}`)
    terabytes_per_secondResult, err := factory.FromDtoJSON(terabytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerabytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabytes_per_secondResult.Convert(units.BitRateTerabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerabytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PetabytePerSecond unit
    petabytes_per_secondJSON := []byte(`{"value": 100, "unit": "PetabytePerSecond"}`)
    petabytes_per_secondResult, err := factory.FromDtoJSON(petabytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PetabytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabytes_per_secondResult.Convert(units.BitRatePetabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetabytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with ExabytePerSecond unit
    exabytes_per_secondJSON := []byte(`{"value": 100, "unit": "ExabytePerSecond"}`)
    exabytes_per_secondResult, err := factory.FromDtoJSON(exabytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ExabytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabytes_per_secondResult.Convert(units.BitRateExabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExabytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KibibytePerSecond unit
    kibibytes_per_secondJSON := []byte(`{"value": 100, "unit": "KibibytePerSecond"}`)
    kibibytes_per_secondResult, err := factory.FromDtoJSON(kibibytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KibibytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibibytes_per_secondResult.Convert(units.BitRateKibibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KibibytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MebibytePerSecond unit
    mebibytes_per_secondJSON := []byte(`{"value": 100, "unit": "MebibytePerSecond"}`)
    mebibytes_per_secondResult, err := factory.FromDtoJSON(mebibytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MebibytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebibytes_per_secondResult.Convert(units.BitRateMebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MebibytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with GibibytePerSecond unit
    gibibytes_per_secondJSON := []byte(`{"value": 100, "unit": "GibibytePerSecond"}`)
    gibibytes_per_secondResult, err := factory.FromDtoJSON(gibibytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GibibytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibibytes_per_secondResult.Convert(units.BitRateGibibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GibibytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with TebibytePerSecond unit
    tebibytes_per_secondJSON := []byte(`{"value": 100, "unit": "TebibytePerSecond"}`)
    tebibytes_per_secondResult, err := factory.FromDtoJSON(tebibytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TebibytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebibytes_per_secondResult.Convert(units.BitRateTebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TebibytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PebibytePerSecond unit
    pebibytes_per_secondJSON := []byte(`{"value": 100, "unit": "PebibytePerSecond"}`)
    pebibytes_per_secondResult, err := factory.FromDtoJSON(pebibytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PebibytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebibytes_per_secondResult.Convert(units.BitRatePebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PebibytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with ExbibytePerSecond unit
    exbibytes_per_secondJSON := []byte(`{"value": 100, "unit": "ExbibytePerSecond"}`)
    exbibytes_per_secondResult, err := factory.FromDtoJSON(exbibytes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ExbibytePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbibytes_per_secondResult.Convert(units.BitRateExbibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExbibytePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilooctetPerSecond unit
    kilooctets_per_secondJSON := []byte(`{"value": 100, "unit": "KilooctetPerSecond"}`)
    kilooctets_per_secondResult, err := factory.FromDtoJSON(kilooctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilooctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilooctets_per_secondResult.Convert(units.BitRateKilooctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilooctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegaoctetPerSecond unit
    megaoctets_per_secondJSON := []byte(`{"value": 100, "unit": "MegaoctetPerSecond"}`)
    megaoctets_per_secondResult, err := factory.FromDtoJSON(megaoctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaoctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaoctets_per_secondResult.Convert(units.BitRateMegaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with GigaoctetPerSecond unit
    gigaoctets_per_secondJSON := []byte(`{"value": 100, "unit": "GigaoctetPerSecond"}`)
    gigaoctets_per_secondResult, err := factory.FromDtoJSON(gigaoctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigaoctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigaoctets_per_secondResult.Convert(units.BitRateGigaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigaoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with TeraoctetPerSecond unit
    teraoctets_per_secondJSON := []byte(`{"value": 100, "unit": "TeraoctetPerSecond"}`)
    teraoctets_per_secondResult, err := factory.FromDtoJSON(teraoctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TeraoctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teraoctets_per_secondResult.Convert(units.BitRateTeraoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TeraoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PetaoctetPerSecond unit
    petaoctets_per_secondJSON := []byte(`{"value": 100, "unit": "PetaoctetPerSecond"}`)
    petaoctets_per_secondResult, err := factory.FromDtoJSON(petaoctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PetaoctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petaoctets_per_secondResult.Convert(units.BitRatePetaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetaoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with ExaoctetPerSecond unit
    exaoctets_per_secondJSON := []byte(`{"value": 100, "unit": "ExaoctetPerSecond"}`)
    exaoctets_per_secondResult, err := factory.FromDtoJSON(exaoctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ExaoctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exaoctets_per_secondResult.Convert(units.BitRateExaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExaoctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KibioctetPerSecond unit
    kibioctets_per_secondJSON := []byte(`{"value": 100, "unit": "KibioctetPerSecond"}`)
    kibioctets_per_secondResult, err := factory.FromDtoJSON(kibioctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KibioctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kibioctets_per_secondResult.Convert(units.BitRateKibioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KibioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MebioctetPerSecond unit
    mebioctets_per_secondJSON := []byte(`{"value": 100, "unit": "MebioctetPerSecond"}`)
    mebioctets_per_secondResult, err := factory.FromDtoJSON(mebioctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MebioctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mebioctets_per_secondResult.Convert(units.BitRateMebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MebioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with GibioctetPerSecond unit
    gibioctets_per_secondJSON := []byte(`{"value": 100, "unit": "GibioctetPerSecond"}`)
    gibioctets_per_secondResult, err := factory.FromDtoJSON(gibioctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GibioctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gibioctets_per_secondResult.Convert(units.BitRateGibioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GibioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with TebioctetPerSecond unit
    tebioctets_per_secondJSON := []byte(`{"value": 100, "unit": "TebioctetPerSecond"}`)
    tebioctets_per_secondResult, err := factory.FromDtoJSON(tebioctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TebioctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tebioctets_per_secondResult.Convert(units.BitRateTebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TebioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PebioctetPerSecond unit
    pebioctets_per_secondJSON := []byte(`{"value": 100, "unit": "PebioctetPerSecond"}`)
    pebioctets_per_secondResult, err := factory.FromDtoJSON(pebioctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PebioctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pebioctets_per_secondResult.Convert(units.BitRatePebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PebioctetPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with ExbioctetPerSecond unit
    exbioctets_per_secondJSON := []byte(`{"value": 100, "unit": "ExbioctetPerSecond"}`)
    exbioctets_per_secondResult, err := factory.FromDtoJSON(exbioctets_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ExbioctetPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exbioctets_per_secondResult.Convert(units.BitRateExbioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ExbioctetPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "BitPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromBitsPerSecond function
func TestBitRateFactory_FromBitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBitsPerSecond(100)
    if err != nil {
        t.Errorf("FromBitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateBitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromBitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromBitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromBitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromBitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromBitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBitsPerSecond(0)
    if err != nil {
        t.Errorf("FromBitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateBitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromBytesPerSecond function
func TestBitRateFactory_FromBytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBytesPerSecond(100)
    if err != nil {
        t.Errorf("FromBytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateBytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromBytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromBytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromBytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromBytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromBytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBytesPerSecond(0)
    if err != nil {
        t.Errorf("FromBytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateBytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromOctetsPerSecond function
func TestBitRateFactory_FromOctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromOctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateOctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromOctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromOctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromOctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromOctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromOctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromOctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateOctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobitsPerSecond function
func TestBitRateFactory_FromKilobitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobitsPerSecond(100)
    if err != nil {
        t.Errorf("FromKilobitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateKilobitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilobitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilobitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilobitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilobitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobitsPerSecond(0)
    if err != nil {
        t.Errorf("FromKilobitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateKilobitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMegabitsPerSecond function
func TestBitRateFactory_FromMegabitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegabitsPerSecond(100)
    if err != nil {
        t.Errorf("FromMegabitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateMegabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegabitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegabitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegabitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegabitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegabitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegabitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegabitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegabitsPerSecond(0)
    if err != nil {
        t.Errorf("FromMegabitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateMegabitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegabitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromGigabitsPerSecond function
func TestBitRateFactory_FromGigabitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigabitsPerSecond(100)
    if err != nil {
        t.Errorf("FromGigabitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateGigabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigabitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigabitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromGigabitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromGigabitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromGigabitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromGigabitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromGigabitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigabitsPerSecond(0)
    if err != nil {
        t.Errorf("FromGigabitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateGigabitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigabitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromTerabitsPerSecond function
func TestBitRateFactory_FromTerabitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerabitsPerSecond(100)
    if err != nil {
        t.Errorf("FromTerabitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateTerabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerabitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerabitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromTerabitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromTerabitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromTerabitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromTerabitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromTerabitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerabitsPerSecond(0)
    if err != nil {
        t.Errorf("FromTerabitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateTerabitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerabitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPetabitsPerSecond function
func TestBitRateFactory_FromPetabitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetabitsPerSecond(100)
    if err != nil {
        t.Errorf("FromPetabitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRatePetabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetabitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetabitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPetabitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPetabitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPetabitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPetabitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPetabitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetabitsPerSecond(0)
    if err != nil {
        t.Errorf("FromPetabitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRatePetabitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetabitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromExabitsPerSecond function
func TestBitRateFactory_FromExabitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExabitsPerSecond(100)
    if err != nil {
        t.Errorf("FromExabitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateExabitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExabitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExabitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromExabitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromExabitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromExabitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromExabitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromExabitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExabitsPerSecond(0)
    if err != nil {
        t.Errorf("FromExabitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateExabitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExabitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKibibitsPerSecond function
func TestBitRateFactory_FromKibibitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKibibitsPerSecond(100)
    if err != nil {
        t.Errorf("FromKibibitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateKibibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKibibitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKibibitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKibibitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKibibitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKibibitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKibibitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKibibitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKibibitsPerSecond(0)
    if err != nil {
        t.Errorf("FromKibibitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateKibibitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKibibitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMebibitsPerSecond function
func TestBitRateFactory_FromMebibitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMebibitsPerSecond(100)
    if err != nil {
        t.Errorf("FromMebibitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateMebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMebibitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMebibitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMebibitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMebibitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMebibitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMebibitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMebibitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMebibitsPerSecond(0)
    if err != nil {
        t.Errorf("FromMebibitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateMebibitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMebibitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromGibibitsPerSecond function
func TestBitRateFactory_FromGibibitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGibibitsPerSecond(100)
    if err != nil {
        t.Errorf("FromGibibitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateGibibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGibibitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGibibitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromGibibitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromGibibitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromGibibitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromGibibitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromGibibitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGibibitsPerSecond(0)
    if err != nil {
        t.Errorf("FromGibibitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateGibibitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGibibitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromTebibitsPerSecond function
func TestBitRateFactory_FromTebibitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTebibitsPerSecond(100)
    if err != nil {
        t.Errorf("FromTebibitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateTebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTebibitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTebibitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromTebibitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromTebibitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromTebibitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromTebibitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromTebibitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTebibitsPerSecond(0)
    if err != nil {
        t.Errorf("FromTebibitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateTebibitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTebibitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPebibitsPerSecond function
func TestBitRateFactory_FromPebibitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPebibitsPerSecond(100)
    if err != nil {
        t.Errorf("FromPebibitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRatePebibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPebibitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPebibitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPebibitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPebibitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPebibitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPebibitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPebibitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPebibitsPerSecond(0)
    if err != nil {
        t.Errorf("FromPebibitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRatePebibitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPebibitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromExbibitsPerSecond function
func TestBitRateFactory_FromExbibitsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExbibitsPerSecond(100)
    if err != nil {
        t.Errorf("FromExbibitsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateExbibitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExbibitsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExbibitsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromExbibitsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromExbibitsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromExbibitsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromExbibitsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromExbibitsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExbibitsPerSecond(0)
    if err != nil {
        t.Errorf("FromExbibitsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateExbibitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExbibitsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobytesPerSecond function
func TestBitRateFactory_FromKilobytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobytesPerSecond(100)
    if err != nil {
        t.Errorf("FromKilobytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateKilobytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilobytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilobytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilobytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilobytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobytesPerSecond(0)
    if err != nil {
        t.Errorf("FromKilobytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateKilobytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMegabytesPerSecond function
func TestBitRateFactory_FromMegabytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegabytesPerSecond(100)
    if err != nil {
        t.Errorf("FromMegabytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateMegabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegabytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegabytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegabytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegabytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegabytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegabytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegabytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegabytesPerSecond(0)
    if err != nil {
        t.Errorf("FromMegabytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateMegabytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegabytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromGigabytesPerSecond function
func TestBitRateFactory_FromGigabytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigabytesPerSecond(100)
    if err != nil {
        t.Errorf("FromGigabytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateGigabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigabytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigabytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromGigabytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromGigabytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromGigabytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromGigabytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromGigabytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigabytesPerSecond(0)
    if err != nil {
        t.Errorf("FromGigabytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateGigabytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigabytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromTerabytesPerSecond function
func TestBitRateFactory_FromTerabytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerabytesPerSecond(100)
    if err != nil {
        t.Errorf("FromTerabytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateTerabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerabytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerabytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromTerabytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromTerabytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromTerabytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromTerabytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromTerabytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerabytesPerSecond(0)
    if err != nil {
        t.Errorf("FromTerabytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateTerabytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerabytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPetabytesPerSecond function
func TestBitRateFactory_FromPetabytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetabytesPerSecond(100)
    if err != nil {
        t.Errorf("FromPetabytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRatePetabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetabytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetabytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPetabytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPetabytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPetabytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPetabytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPetabytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetabytesPerSecond(0)
    if err != nil {
        t.Errorf("FromPetabytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRatePetabytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetabytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromExabytesPerSecond function
func TestBitRateFactory_FromExabytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExabytesPerSecond(100)
    if err != nil {
        t.Errorf("FromExabytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateExabytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExabytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExabytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromExabytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromExabytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromExabytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromExabytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromExabytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExabytesPerSecond(0)
    if err != nil {
        t.Errorf("FromExabytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateExabytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExabytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKibibytesPerSecond function
func TestBitRateFactory_FromKibibytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKibibytesPerSecond(100)
    if err != nil {
        t.Errorf("FromKibibytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateKibibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKibibytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKibibytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKibibytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKibibytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKibibytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKibibytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKibibytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKibibytesPerSecond(0)
    if err != nil {
        t.Errorf("FromKibibytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateKibibytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKibibytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMebibytesPerSecond function
func TestBitRateFactory_FromMebibytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMebibytesPerSecond(100)
    if err != nil {
        t.Errorf("FromMebibytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateMebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMebibytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMebibytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMebibytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMebibytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMebibytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMebibytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMebibytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMebibytesPerSecond(0)
    if err != nil {
        t.Errorf("FromMebibytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateMebibytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMebibytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromGibibytesPerSecond function
func TestBitRateFactory_FromGibibytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGibibytesPerSecond(100)
    if err != nil {
        t.Errorf("FromGibibytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateGibibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGibibytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGibibytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromGibibytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromGibibytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromGibibytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromGibibytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromGibibytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGibibytesPerSecond(0)
    if err != nil {
        t.Errorf("FromGibibytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateGibibytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGibibytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromTebibytesPerSecond function
func TestBitRateFactory_FromTebibytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTebibytesPerSecond(100)
    if err != nil {
        t.Errorf("FromTebibytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateTebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTebibytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTebibytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromTebibytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromTebibytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromTebibytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromTebibytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromTebibytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTebibytesPerSecond(0)
    if err != nil {
        t.Errorf("FromTebibytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateTebibytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTebibytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPebibytesPerSecond function
func TestBitRateFactory_FromPebibytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPebibytesPerSecond(100)
    if err != nil {
        t.Errorf("FromPebibytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRatePebibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPebibytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPebibytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPebibytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPebibytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPebibytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPebibytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPebibytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPebibytesPerSecond(0)
    if err != nil {
        t.Errorf("FromPebibytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRatePebibytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPebibytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromExbibytesPerSecond function
func TestBitRateFactory_FromExbibytesPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExbibytesPerSecond(100)
    if err != nil {
        t.Errorf("FromExbibytesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateExbibytePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExbibytesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExbibytesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromExbibytesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromExbibytesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromExbibytesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromExbibytesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromExbibytesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExbibytesPerSecond(0)
    if err != nil {
        t.Errorf("FromExbibytesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateExbibytePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExbibytesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilooctetsPerSecond function
func TestBitRateFactory_FromKilooctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilooctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromKilooctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateKilooctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilooctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilooctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilooctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilooctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilooctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilooctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilooctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilooctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromKilooctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateKilooctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilooctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaoctetsPerSecond function
func TestBitRateFactory_FromMegaoctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaoctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromMegaoctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateMegaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaoctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaoctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegaoctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegaoctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegaoctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegaoctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaoctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaoctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromMegaoctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateMegaoctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaoctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromGigaoctetsPerSecond function
func TestBitRateFactory_FromGigaoctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigaoctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromGigaoctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateGigaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigaoctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigaoctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromGigaoctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromGigaoctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromGigaoctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromGigaoctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromGigaoctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigaoctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromGigaoctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateGigaoctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigaoctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromTeraoctetsPerSecond function
func TestBitRateFactory_FromTeraoctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeraoctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromTeraoctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateTeraoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTeraoctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTeraoctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromTeraoctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromTeraoctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromTeraoctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromTeraoctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromTeraoctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTeraoctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromTeraoctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateTeraoctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeraoctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPetaoctetsPerSecond function
func TestBitRateFactory_FromPetaoctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetaoctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromPetaoctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRatePetaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetaoctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetaoctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPetaoctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPetaoctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPetaoctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPetaoctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPetaoctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetaoctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromPetaoctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRatePetaoctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetaoctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromExaoctetsPerSecond function
func TestBitRateFactory_FromExaoctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExaoctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromExaoctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateExaoctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExaoctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExaoctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromExaoctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromExaoctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromExaoctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromExaoctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromExaoctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExaoctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromExaoctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateExaoctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExaoctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKibioctetsPerSecond function
func TestBitRateFactory_FromKibioctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKibioctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromKibioctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateKibioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKibioctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKibioctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKibioctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKibioctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKibioctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKibioctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKibioctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKibioctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromKibioctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateKibioctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKibioctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMebioctetsPerSecond function
func TestBitRateFactory_FromMebioctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMebioctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromMebioctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateMebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMebioctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMebioctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMebioctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMebioctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMebioctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMebioctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMebioctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMebioctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromMebioctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateMebioctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMebioctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromGibioctetsPerSecond function
func TestBitRateFactory_FromGibioctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGibioctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromGibioctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateGibioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGibioctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGibioctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromGibioctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromGibioctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromGibioctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromGibioctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromGibioctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGibioctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromGibioctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateGibioctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGibioctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromTebioctetsPerSecond function
func TestBitRateFactory_FromTebioctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTebioctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromTebioctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateTebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTebioctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTebioctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromTebioctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromTebioctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromTebioctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromTebioctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromTebioctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTebioctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromTebioctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateTebioctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTebioctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPebioctetsPerSecond function
func TestBitRateFactory_FromPebioctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPebioctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromPebioctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRatePebioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPebioctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPebioctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPebioctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPebioctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPebioctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPebioctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPebioctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPebioctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromPebioctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRatePebioctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPebioctetsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromExbioctetsPerSecond function
func TestBitRateFactory_FromExbioctetsPerSecond(t *testing.T) {
    factory := units.BitRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExbioctetsPerSecond(100)
    if err != nil {
        t.Errorf("FromExbioctetsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BitRateExbioctetPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExbioctetsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExbioctetsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromExbioctetsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromExbioctetsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromExbioctetsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromExbioctetsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromExbioctetsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExbioctetsPerSecond(0)
    if err != nil {
        t.Errorf("FromExbioctetsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BitRateExbioctetPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExbioctetsPerSecond() with zero value = %v, want 0", converted)
    }
}

func TestBitRateToString(t *testing.T) {
	factory := units.BitRateFactory{}
	a, err := factory.CreateBitRate(45, units.BitRateBitPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.BitRateBitPerSecond, 2)
	expected := "45.00 " + units.GetBitRateAbbreviation(units.BitRateBitPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.BitRateBitPerSecond, -1)
	expected = "45 " + units.GetBitRateAbbreviation(units.BitRateBitPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestBitRate_EqualityAndComparison(t *testing.T) {
	factory := units.BitRateFactory{}
	a1, _ := factory.CreateBitRate(60, units.BitRateBitPerSecond)
	a2, _ := factory.CreateBitRate(60, units.BitRateBitPerSecond)
	a3, _ := factory.CreateBitRate(120, units.BitRateBitPerSecond)

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

func TestBitRate_Arithmetic(t *testing.T) {
	factory := units.BitRateFactory{}
	a1, _ := factory.CreateBitRate(30, units.BitRateBitPerSecond)
	a2, _ := factory.CreateBitRate(45, units.BitRateBitPerSecond)

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


func TestGetBitRateAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.BitRateUnits
        want string
    }{
        {
            name: "BitPerSecond abbreviation",
            unit: units.BitRateBitPerSecond,
            want: "bit/s",
        },
        {
            name: "BytePerSecond abbreviation",
            unit: units.BitRateBytePerSecond,
            want: "B/s",
        },
        {
            name: "OctetPerSecond abbreviation",
            unit: units.BitRateOctetPerSecond,
            want: "o/s",
        },
        {
            name: "KilobitPerSecond abbreviation",
            unit: units.BitRateKilobitPerSecond,
            want: "kbit/s",
        },
        {
            name: "MegabitPerSecond abbreviation",
            unit: units.BitRateMegabitPerSecond,
            want: "Mbit/s",
        },
        {
            name: "GigabitPerSecond abbreviation",
            unit: units.BitRateGigabitPerSecond,
            want: "Gbit/s",
        },
        {
            name: "TerabitPerSecond abbreviation",
            unit: units.BitRateTerabitPerSecond,
            want: "Tbit/s",
        },
        {
            name: "PetabitPerSecond abbreviation",
            unit: units.BitRatePetabitPerSecond,
            want: "Pbit/s",
        },
        {
            name: "ExabitPerSecond abbreviation",
            unit: units.BitRateExabitPerSecond,
            want: "Ebit/s",
        },
        {
            name: "KibibitPerSecond abbreviation",
            unit: units.BitRateKibibitPerSecond,
            want: "KiBbit/s",
        },
        {
            name: "MebibitPerSecond abbreviation",
            unit: units.BitRateMebibitPerSecond,
            want: "MiBbit/s",
        },
        {
            name: "GibibitPerSecond abbreviation",
            unit: units.BitRateGibibitPerSecond,
            want: "GiBbit/s",
        },
        {
            name: "TebibitPerSecond abbreviation",
            unit: units.BitRateTebibitPerSecond,
            want: "TiBbit/s",
        },
        {
            name: "PebibitPerSecond abbreviation",
            unit: units.BitRatePebibitPerSecond,
            want: "PiBbit/s",
        },
        {
            name: "ExbibitPerSecond abbreviation",
            unit: units.BitRateExbibitPerSecond,
            want: "EiBbit/s",
        },
        {
            name: "KilobytePerSecond abbreviation",
            unit: units.BitRateKilobytePerSecond,
            want: "kB/s",
        },
        {
            name: "MegabytePerSecond abbreviation",
            unit: units.BitRateMegabytePerSecond,
            want: "MB/s",
        },
        {
            name: "GigabytePerSecond abbreviation",
            unit: units.BitRateGigabytePerSecond,
            want: "GB/s",
        },
        {
            name: "TerabytePerSecond abbreviation",
            unit: units.BitRateTerabytePerSecond,
            want: "TB/s",
        },
        {
            name: "PetabytePerSecond abbreviation",
            unit: units.BitRatePetabytePerSecond,
            want: "PB/s",
        },
        {
            name: "ExabytePerSecond abbreviation",
            unit: units.BitRateExabytePerSecond,
            want: "EB/s",
        },
        {
            name: "KibibytePerSecond abbreviation",
            unit: units.BitRateKibibytePerSecond,
            want: "KiBB/s",
        },
        {
            name: "MebibytePerSecond abbreviation",
            unit: units.BitRateMebibytePerSecond,
            want: "MiBB/s",
        },
        {
            name: "GibibytePerSecond abbreviation",
            unit: units.BitRateGibibytePerSecond,
            want: "GiBB/s",
        },
        {
            name: "TebibytePerSecond abbreviation",
            unit: units.BitRateTebibytePerSecond,
            want: "TiBB/s",
        },
        {
            name: "PebibytePerSecond abbreviation",
            unit: units.BitRatePebibytePerSecond,
            want: "PiBB/s",
        },
        {
            name: "ExbibytePerSecond abbreviation",
            unit: units.BitRateExbibytePerSecond,
            want: "EiBB/s",
        },
        {
            name: "KilooctetPerSecond abbreviation",
            unit: units.BitRateKilooctetPerSecond,
            want: "ko/s",
        },
        {
            name: "MegaoctetPerSecond abbreviation",
            unit: units.BitRateMegaoctetPerSecond,
            want: "Mo/s",
        },
        {
            name: "GigaoctetPerSecond abbreviation",
            unit: units.BitRateGigaoctetPerSecond,
            want: "Go/s",
        },
        {
            name: "TeraoctetPerSecond abbreviation",
            unit: units.BitRateTeraoctetPerSecond,
            want: "To/s",
        },
        {
            name: "PetaoctetPerSecond abbreviation",
            unit: units.BitRatePetaoctetPerSecond,
            want: "Po/s",
        },
        {
            name: "ExaoctetPerSecond abbreviation",
            unit: units.BitRateExaoctetPerSecond,
            want: "Eo/s",
        },
        {
            name: "KibioctetPerSecond abbreviation",
            unit: units.BitRateKibioctetPerSecond,
            want: "KiBo/s",
        },
        {
            name: "MebioctetPerSecond abbreviation",
            unit: units.BitRateMebioctetPerSecond,
            want: "MiBo/s",
        },
        {
            name: "GibioctetPerSecond abbreviation",
            unit: units.BitRateGibioctetPerSecond,
            want: "GiBo/s",
        },
        {
            name: "TebioctetPerSecond abbreviation",
            unit: units.BitRateTebioctetPerSecond,
            want: "TiBo/s",
        },
        {
            name: "PebioctetPerSecond abbreviation",
            unit: units.BitRatePebioctetPerSecond,
            want: "PiBo/s",
        },
        {
            name: "ExbioctetPerSecond abbreviation",
            unit: units.BitRateExbioctetPerSecond,
            want: "EiBo/s",
        },
        {
            name: "invalid unit",
            unit: units.BitRateUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetBitRateAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetBitRateAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestBitRate_String(t *testing.T) {
    factory := units.BitRateFactory{}
    
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
            unit, err := factory.CreateBitRate(tt.value, units.BitRateBitPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("BitRate.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestBitRate_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.BitRateFactory{}

	_, err := uf.CreateBitRate(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}