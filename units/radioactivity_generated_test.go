// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRadioactivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Becquerel"}`
	
	factory := units.RadioactivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RadioactivityBecquerel {
		t.Errorf("expected unit %v, got %v", units.RadioactivityBecquerel, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Becquerel"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRadioactivityDto_ToJSON(t *testing.T) {
	dto := units.RadioactivityDto{
		Value: 45,
		Unit:  units.RadioactivityBecquerel,
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
	if result["unit"].(string) != string(units.RadioactivityBecquerel) {
		t.Errorf("expected unit %s, got %v", units.RadioactivityBecquerel, result["unit"])
	}
}

func TestNewRadioactivity_InvalidValue(t *testing.T) {
	factory := units.RadioactivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRadioactivity(math.NaN(), units.RadioactivityBecquerel)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRadioactivity(math.Inf(1), units.RadioactivityBecquerel)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRadioactivityConversions(t *testing.T) {
	factory := units.RadioactivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRadioactivity(180, units.RadioactivityBecquerel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Becquerels.
		// No expected conversion value provided for Becquerels, verifying result is not NaN.
		result := a.Becquerels()
		cacheResult := a.Becquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Becquerels returned NaN")
		}
	}
	{
		// Test conversion to Curies.
		// No expected conversion value provided for Curies, verifying result is not NaN.
		result := a.Curies()
		cacheResult := a.Curies()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Curies returned NaN")
		}
	}
	{
		// Test conversion to Rutherfords.
		// No expected conversion value provided for Rutherfords, verifying result is not NaN.
		result := a.Rutherfords()
		cacheResult := a.Rutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Rutherfords returned NaN")
		}
	}
	{
		// Test conversion to Picobecquerels.
		// No expected conversion value provided for Picobecquerels, verifying result is not NaN.
		result := a.Picobecquerels()
		cacheResult := a.Picobecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picobecquerels returned NaN")
		}
	}
	{
		// Test conversion to Nanobecquerels.
		// No expected conversion value provided for Nanobecquerels, verifying result is not NaN.
		result := a.Nanobecquerels()
		cacheResult := a.Nanobecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanobecquerels returned NaN")
		}
	}
	{
		// Test conversion to Microbecquerels.
		// No expected conversion value provided for Microbecquerels, verifying result is not NaN.
		result := a.Microbecquerels()
		cacheResult := a.Microbecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microbecquerels returned NaN")
		}
	}
	{
		// Test conversion to Millibecquerels.
		// No expected conversion value provided for Millibecquerels, verifying result is not NaN.
		result := a.Millibecquerels()
		cacheResult := a.Millibecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millibecquerels returned NaN")
		}
	}
	{
		// Test conversion to Kilobecquerels.
		// No expected conversion value provided for Kilobecquerels, verifying result is not NaN.
		result := a.Kilobecquerels()
		cacheResult := a.Kilobecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilobecquerels returned NaN")
		}
	}
	{
		// Test conversion to Megabecquerels.
		// No expected conversion value provided for Megabecquerels, verifying result is not NaN.
		result := a.Megabecquerels()
		cacheResult := a.Megabecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Gigabecquerels.
		// No expected conversion value provided for Gigabecquerels, verifying result is not NaN.
		result := a.Gigabecquerels()
		cacheResult := a.Gigabecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Terabecquerels.
		// No expected conversion value provided for Terabecquerels, verifying result is not NaN.
		result := a.Terabecquerels()
		cacheResult := a.Terabecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Terabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Petabecquerels.
		// No expected conversion value provided for Petabecquerels, verifying result is not NaN.
		result := a.Petabecquerels()
		cacheResult := a.Petabecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Petabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Exabecquerels.
		// No expected conversion value provided for Exabecquerels, verifying result is not NaN.
		result := a.Exabecquerels()
		cacheResult := a.Exabecquerels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Exabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Picocuries.
		// No expected conversion value provided for Picocuries, verifying result is not NaN.
		result := a.Picocuries()
		cacheResult := a.Picocuries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picocuries returned NaN")
		}
	}
	{
		// Test conversion to Nanocuries.
		// No expected conversion value provided for Nanocuries, verifying result is not NaN.
		result := a.Nanocuries()
		cacheResult := a.Nanocuries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanocuries returned NaN")
		}
	}
	{
		// Test conversion to Microcuries.
		// No expected conversion value provided for Microcuries, verifying result is not NaN.
		result := a.Microcuries()
		cacheResult := a.Microcuries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microcuries returned NaN")
		}
	}
	{
		// Test conversion to Millicuries.
		// No expected conversion value provided for Millicuries, verifying result is not NaN.
		result := a.Millicuries()
		cacheResult := a.Millicuries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millicuries returned NaN")
		}
	}
	{
		// Test conversion to Kilocuries.
		// No expected conversion value provided for Kilocuries, verifying result is not NaN.
		result := a.Kilocuries()
		cacheResult := a.Kilocuries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilocuries returned NaN")
		}
	}
	{
		// Test conversion to Megacuries.
		// No expected conversion value provided for Megacuries, verifying result is not NaN.
		result := a.Megacuries()
		cacheResult := a.Megacuries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megacuries returned NaN")
		}
	}
	{
		// Test conversion to Gigacuries.
		// No expected conversion value provided for Gigacuries, verifying result is not NaN.
		result := a.Gigacuries()
		cacheResult := a.Gigacuries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigacuries returned NaN")
		}
	}
	{
		// Test conversion to Teracuries.
		// No expected conversion value provided for Teracuries, verifying result is not NaN.
		result := a.Teracuries()
		cacheResult := a.Teracuries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Teracuries returned NaN")
		}
	}
	{
		// Test conversion to Picorutherfords.
		// No expected conversion value provided for Picorutherfords, verifying result is not NaN.
		result := a.Picorutherfords()
		cacheResult := a.Picorutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picorutherfords returned NaN")
		}
	}
	{
		// Test conversion to Nanorutherfords.
		// No expected conversion value provided for Nanorutherfords, verifying result is not NaN.
		result := a.Nanorutherfords()
		cacheResult := a.Nanorutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanorutherfords returned NaN")
		}
	}
	{
		// Test conversion to Microrutherfords.
		// No expected conversion value provided for Microrutherfords, verifying result is not NaN.
		result := a.Microrutherfords()
		cacheResult := a.Microrutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microrutherfords returned NaN")
		}
	}
	{
		// Test conversion to Millirutherfords.
		// No expected conversion value provided for Millirutherfords, verifying result is not NaN.
		result := a.Millirutherfords()
		cacheResult := a.Millirutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millirutherfords returned NaN")
		}
	}
	{
		// Test conversion to Kilorutherfords.
		// No expected conversion value provided for Kilorutherfords, verifying result is not NaN.
		result := a.Kilorutherfords()
		cacheResult := a.Kilorutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilorutherfords returned NaN")
		}
	}
	{
		// Test conversion to Megarutherfords.
		// No expected conversion value provided for Megarutherfords, verifying result is not NaN.
		result := a.Megarutherfords()
		cacheResult := a.Megarutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megarutherfords returned NaN")
		}
	}
	{
		// Test conversion to Gigarutherfords.
		// No expected conversion value provided for Gigarutherfords, verifying result is not NaN.
		result := a.Gigarutherfords()
		cacheResult := a.Gigarutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigarutherfords returned NaN")
		}
	}
	{
		// Test conversion to Terarutherfords.
		// No expected conversion value provided for Terarutherfords, verifying result is not NaN.
		result := a.Terarutherfords()
		cacheResult := a.Terarutherfords()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Terarutherfords returned NaN")
		}
	}
}

func TestRadioactivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RadioactivityFactory{}
	a, err := factory.CreateRadioactivity(90, units.RadioactivityBecquerel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RadioactivityBecquerel {
		t.Errorf("expected default unit Becquerel, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RadioactivityBecquerel
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RadioactivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RadioactivityBecquerel {
		t.Errorf("expected unit Becquerel, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRadioactivityFactory_FromDto(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityBecquerel,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RadioactivityDto{
        Value: math.NaN(),
        Unit:  units.RadioactivityBecquerel,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Becquerel conversion
    becquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityBecquerel,
    }
    
    var becquerelsResult *units.Radioactivity
    becquerelsResult, err = factory.FromDto(becquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Becquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = becquerelsResult.Convert(units.RadioactivityBecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Becquerel = %v, want %v", converted, 100)
    }
    // Test Curie conversion
    curiesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityCurie,
    }
    
    var curiesResult *units.Radioactivity
    curiesResult, err = factory.FromDto(curiesDto)
    if err != nil {
        t.Errorf("FromDto() with Curie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = curiesResult.Convert(units.RadioactivityCurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Curie = %v, want %v", converted, 100)
    }
    // Test Rutherford conversion
    rutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityRutherford,
    }
    
    var rutherfordsResult *units.Radioactivity
    rutherfordsResult, err = factory.FromDto(rutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Rutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = rutherfordsResult.Convert(units.RadioactivityRutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Rutherford = %v, want %v", converted, 100)
    }
    // Test Picobecquerel conversion
    picobecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityPicobecquerel,
    }
    
    var picobecquerelsResult *units.Radioactivity
    picobecquerelsResult, err = factory.FromDto(picobecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Picobecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picobecquerelsResult.Convert(units.RadioactivityPicobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picobecquerel = %v, want %v", converted, 100)
    }
    // Test Nanobecquerel conversion
    nanobecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityNanobecquerel,
    }
    
    var nanobecquerelsResult *units.Radioactivity
    nanobecquerelsResult, err = factory.FromDto(nanobecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanobecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanobecquerelsResult.Convert(units.RadioactivityNanobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanobecquerel = %v, want %v", converted, 100)
    }
    // Test Microbecquerel conversion
    microbecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMicrobecquerel,
    }
    
    var microbecquerelsResult *units.Radioactivity
    microbecquerelsResult, err = factory.FromDto(microbecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Microbecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microbecquerelsResult.Convert(units.RadioactivityMicrobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microbecquerel = %v, want %v", converted, 100)
    }
    // Test Millibecquerel conversion
    millibecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMillibecquerel,
    }
    
    var millibecquerelsResult *units.Radioactivity
    millibecquerelsResult, err = factory.FromDto(millibecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Millibecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibecquerelsResult.Convert(units.RadioactivityMillibecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millibecquerel = %v, want %v", converted, 100)
    }
    // Test Kilobecquerel conversion
    kilobecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityKilobecquerel,
    }
    
    var kilobecquerelsResult *units.Radioactivity
    kilobecquerelsResult, err = factory.FromDto(kilobecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilobecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobecquerelsResult.Convert(units.RadioactivityKilobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilobecquerel = %v, want %v", converted, 100)
    }
    // Test Megabecquerel conversion
    megabecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMegabecquerel,
    }
    
    var megabecquerelsResult *units.Radioactivity
    megabecquerelsResult, err = factory.FromDto(megabecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Megabecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabecquerelsResult.Convert(units.RadioactivityMegabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megabecquerel = %v, want %v", converted, 100)
    }
    // Test Gigabecquerel conversion
    gigabecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityGigabecquerel,
    }
    
    var gigabecquerelsResult *units.Radioactivity
    gigabecquerelsResult, err = factory.FromDto(gigabecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigabecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabecquerelsResult.Convert(units.RadioactivityGigabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigabecquerel = %v, want %v", converted, 100)
    }
    // Test Terabecquerel conversion
    terabecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityTerabecquerel,
    }
    
    var terabecquerelsResult *units.Radioactivity
    terabecquerelsResult, err = factory.FromDto(terabecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Terabecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabecquerelsResult.Convert(units.RadioactivityTerabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terabecquerel = %v, want %v", converted, 100)
    }
    // Test Petabecquerel conversion
    petabecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityPetabecquerel,
    }
    
    var petabecquerelsResult *units.Radioactivity
    petabecquerelsResult, err = factory.FromDto(petabecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Petabecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabecquerelsResult.Convert(units.RadioactivityPetabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petabecquerel = %v, want %v", converted, 100)
    }
    // Test Exabecquerel conversion
    exabecquerelsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityExabecquerel,
    }
    
    var exabecquerelsResult *units.Radioactivity
    exabecquerelsResult, err = factory.FromDto(exabecquerelsDto)
    if err != nil {
        t.Errorf("FromDto() with Exabecquerel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabecquerelsResult.Convert(units.RadioactivityExabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exabecquerel = %v, want %v", converted, 100)
    }
    // Test Picocurie conversion
    picocuriesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityPicocurie,
    }
    
    var picocuriesResult *units.Radioactivity
    picocuriesResult, err = factory.FromDto(picocuriesDto)
    if err != nil {
        t.Errorf("FromDto() with Picocurie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picocuriesResult.Convert(units.RadioactivityPicocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picocurie = %v, want %v", converted, 100)
    }
    // Test Nanocurie conversion
    nanocuriesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityNanocurie,
    }
    
    var nanocuriesResult *units.Radioactivity
    nanocuriesResult, err = factory.FromDto(nanocuriesDto)
    if err != nil {
        t.Errorf("FromDto() with Nanocurie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanocuriesResult.Convert(units.RadioactivityNanocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanocurie = %v, want %v", converted, 100)
    }
    // Test Microcurie conversion
    microcuriesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMicrocurie,
    }
    
    var microcuriesResult *units.Radioactivity
    microcuriesResult, err = factory.FromDto(microcuriesDto)
    if err != nil {
        t.Errorf("FromDto() with Microcurie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microcuriesResult.Convert(units.RadioactivityMicrocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microcurie = %v, want %v", converted, 100)
    }
    // Test Millicurie conversion
    millicuriesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMillicurie,
    }
    
    var millicuriesResult *units.Radioactivity
    millicuriesResult, err = factory.FromDto(millicuriesDto)
    if err != nil {
        t.Errorf("FromDto() with Millicurie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicuriesResult.Convert(units.RadioactivityMillicurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millicurie = %v, want %v", converted, 100)
    }
    // Test Kilocurie conversion
    kilocuriesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityKilocurie,
    }
    
    var kilocuriesResult *units.Radioactivity
    kilocuriesResult, err = factory.FromDto(kilocuriesDto)
    if err != nil {
        t.Errorf("FromDto() with Kilocurie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocuriesResult.Convert(units.RadioactivityKilocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilocurie = %v, want %v", converted, 100)
    }
    // Test Megacurie conversion
    megacuriesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMegacurie,
    }
    
    var megacuriesResult *units.Radioactivity
    megacuriesResult, err = factory.FromDto(megacuriesDto)
    if err != nil {
        t.Errorf("FromDto() with Megacurie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megacuriesResult.Convert(units.RadioactivityMegacurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megacurie = %v, want %v", converted, 100)
    }
    // Test Gigacurie conversion
    gigacuriesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityGigacurie,
    }
    
    var gigacuriesResult *units.Radioactivity
    gigacuriesResult, err = factory.FromDto(gigacuriesDto)
    if err != nil {
        t.Errorf("FromDto() with Gigacurie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigacuriesResult.Convert(units.RadioactivityGigacurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigacurie = %v, want %v", converted, 100)
    }
    // Test Teracurie conversion
    teracuriesDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityTeracurie,
    }
    
    var teracuriesResult *units.Radioactivity
    teracuriesResult, err = factory.FromDto(teracuriesDto)
    if err != nil {
        t.Errorf("FromDto() with Teracurie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teracuriesResult.Convert(units.RadioactivityTeracurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teracurie = %v, want %v", converted, 100)
    }
    // Test Picorutherford conversion
    picorutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityPicorutherford,
    }
    
    var picorutherfordsResult *units.Radioactivity
    picorutherfordsResult, err = factory.FromDto(picorutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Picorutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picorutherfordsResult.Convert(units.RadioactivityPicorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picorutherford = %v, want %v", converted, 100)
    }
    // Test Nanorutherford conversion
    nanorutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityNanorutherford,
    }
    
    var nanorutherfordsResult *units.Radioactivity
    nanorutherfordsResult, err = factory.FromDto(nanorutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanorutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanorutherfordsResult.Convert(units.RadioactivityNanorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanorutherford = %v, want %v", converted, 100)
    }
    // Test Microrutherford conversion
    microrutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMicrorutherford,
    }
    
    var microrutherfordsResult *units.Radioactivity
    microrutherfordsResult, err = factory.FromDto(microrutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Microrutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microrutherfordsResult.Convert(units.RadioactivityMicrorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microrutherford = %v, want %v", converted, 100)
    }
    // Test Millirutherford conversion
    millirutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMillirutherford,
    }
    
    var millirutherfordsResult *units.Radioactivity
    millirutherfordsResult, err = factory.FromDto(millirutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Millirutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millirutherfordsResult.Convert(units.RadioactivityMillirutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millirutherford = %v, want %v", converted, 100)
    }
    // Test Kilorutherford conversion
    kilorutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityKilorutherford,
    }
    
    var kilorutherfordsResult *units.Radioactivity
    kilorutherfordsResult, err = factory.FromDto(kilorutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilorutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilorutherfordsResult.Convert(units.RadioactivityKilorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilorutherford = %v, want %v", converted, 100)
    }
    // Test Megarutherford conversion
    megarutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityMegarutherford,
    }
    
    var megarutherfordsResult *units.Radioactivity
    megarutherfordsResult, err = factory.FromDto(megarutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Megarutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megarutherfordsResult.Convert(units.RadioactivityMegarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megarutherford = %v, want %v", converted, 100)
    }
    // Test Gigarutherford conversion
    gigarutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityGigarutherford,
    }
    
    var gigarutherfordsResult *units.Radioactivity
    gigarutherfordsResult, err = factory.FromDto(gigarutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigarutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigarutherfordsResult.Convert(units.RadioactivityGigarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigarutherford = %v, want %v", converted, 100)
    }
    // Test Terarutherford conversion
    terarutherfordsDto := units.RadioactivityDto{
        Value: 100,
        Unit:  units.RadioactivityTerarutherford,
    }
    
    var terarutherfordsResult *units.Radioactivity
    terarutherfordsResult, err = factory.FromDto(terarutherfordsDto)
    if err != nil {
        t.Errorf("FromDto() with Terarutherford returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terarutherfordsResult.Convert(units.RadioactivityTerarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terarutherford = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RadioactivityDto{
        Value: 0,
        Unit:  units.RadioactivityBecquerel,
    }
    
    var zeroResult *units.Radioactivity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRadioactivityFactory_FromDtoJSON(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Becquerel"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Becquerel"}`)
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
    nanJSON, _ := json.Marshal(units.RadioactivityDto{
        Value: nanValue,
        Unit:  units.RadioactivityBecquerel,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Becquerel unit
    becquerelsJSON := []byte(`{"value": 100, "unit": "Becquerel"}`)
    becquerelsResult, err := factory.FromDtoJSON(becquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Becquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = becquerelsResult.Convert(units.RadioactivityBecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Becquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Curie unit
    curiesJSON := []byte(`{"value": 100, "unit": "Curie"}`)
    curiesResult, err := factory.FromDtoJSON(curiesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Curie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = curiesResult.Convert(units.RadioactivityCurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Curie = %v, want %v", converted, 100)
    }
    // Test JSON with Rutherford unit
    rutherfordsJSON := []byte(`{"value": 100, "unit": "Rutherford"}`)
    rutherfordsResult, err := factory.FromDtoJSON(rutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Rutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = rutherfordsResult.Convert(units.RadioactivityRutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Rutherford = %v, want %v", converted, 100)
    }
    // Test JSON with Picobecquerel unit
    picobecquerelsJSON := []byte(`{"value": 100, "unit": "Picobecquerel"}`)
    picobecquerelsResult, err := factory.FromDtoJSON(picobecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picobecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picobecquerelsResult.Convert(units.RadioactivityPicobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picobecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Nanobecquerel unit
    nanobecquerelsJSON := []byte(`{"value": 100, "unit": "Nanobecquerel"}`)
    nanobecquerelsResult, err := factory.FromDtoJSON(nanobecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanobecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanobecquerelsResult.Convert(units.RadioactivityNanobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanobecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Microbecquerel unit
    microbecquerelsJSON := []byte(`{"value": 100, "unit": "Microbecquerel"}`)
    microbecquerelsResult, err := factory.FromDtoJSON(microbecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microbecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microbecquerelsResult.Convert(units.RadioactivityMicrobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microbecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Millibecquerel unit
    millibecquerelsJSON := []byte(`{"value": 100, "unit": "Millibecquerel"}`)
    millibecquerelsResult, err := factory.FromDtoJSON(millibecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millibecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibecquerelsResult.Convert(units.RadioactivityMillibecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millibecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Kilobecquerel unit
    kilobecquerelsJSON := []byte(`{"value": 100, "unit": "Kilobecquerel"}`)
    kilobecquerelsResult, err := factory.FromDtoJSON(kilobecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilobecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobecquerelsResult.Convert(units.RadioactivityKilobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilobecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Megabecquerel unit
    megabecquerelsJSON := []byte(`{"value": 100, "unit": "Megabecquerel"}`)
    megabecquerelsResult, err := factory.FromDtoJSON(megabecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megabecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabecquerelsResult.Convert(units.RadioactivityMegabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megabecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Gigabecquerel unit
    gigabecquerelsJSON := []byte(`{"value": 100, "unit": "Gigabecquerel"}`)
    gigabecquerelsResult, err := factory.FromDtoJSON(gigabecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigabecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabecquerelsResult.Convert(units.RadioactivityGigabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigabecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Terabecquerel unit
    terabecquerelsJSON := []byte(`{"value": 100, "unit": "Terabecquerel"}`)
    terabecquerelsResult, err := factory.FromDtoJSON(terabecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terabecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terabecquerelsResult.Convert(units.RadioactivityTerabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terabecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Petabecquerel unit
    petabecquerelsJSON := []byte(`{"value": 100, "unit": "Petabecquerel"}`)
    petabecquerelsResult, err := factory.FromDtoJSON(petabecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Petabecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petabecquerelsResult.Convert(units.RadioactivityPetabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petabecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Exabecquerel unit
    exabecquerelsJSON := []byte(`{"value": 100, "unit": "Exabecquerel"}`)
    exabecquerelsResult, err := factory.FromDtoJSON(exabecquerelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Exabecquerel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = exabecquerelsResult.Convert(units.RadioactivityExabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Exabecquerel = %v, want %v", converted, 100)
    }
    // Test JSON with Picocurie unit
    picocuriesJSON := []byte(`{"value": 100, "unit": "Picocurie"}`)
    picocuriesResult, err := factory.FromDtoJSON(picocuriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picocurie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picocuriesResult.Convert(units.RadioactivityPicocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picocurie = %v, want %v", converted, 100)
    }
    // Test JSON with Nanocurie unit
    nanocuriesJSON := []byte(`{"value": 100, "unit": "Nanocurie"}`)
    nanocuriesResult, err := factory.FromDtoJSON(nanocuriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanocurie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanocuriesResult.Convert(units.RadioactivityNanocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanocurie = %v, want %v", converted, 100)
    }
    // Test JSON with Microcurie unit
    microcuriesJSON := []byte(`{"value": 100, "unit": "Microcurie"}`)
    microcuriesResult, err := factory.FromDtoJSON(microcuriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microcurie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microcuriesResult.Convert(units.RadioactivityMicrocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microcurie = %v, want %v", converted, 100)
    }
    // Test JSON with Millicurie unit
    millicuriesJSON := []byte(`{"value": 100, "unit": "Millicurie"}`)
    millicuriesResult, err := factory.FromDtoJSON(millicuriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millicurie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicuriesResult.Convert(units.RadioactivityMillicurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millicurie = %v, want %v", converted, 100)
    }
    // Test JSON with Kilocurie unit
    kilocuriesJSON := []byte(`{"value": 100, "unit": "Kilocurie"}`)
    kilocuriesResult, err := factory.FromDtoJSON(kilocuriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilocurie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocuriesResult.Convert(units.RadioactivityKilocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilocurie = %v, want %v", converted, 100)
    }
    // Test JSON with Megacurie unit
    megacuriesJSON := []byte(`{"value": 100, "unit": "Megacurie"}`)
    megacuriesResult, err := factory.FromDtoJSON(megacuriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megacurie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megacuriesResult.Convert(units.RadioactivityMegacurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megacurie = %v, want %v", converted, 100)
    }
    // Test JSON with Gigacurie unit
    gigacuriesJSON := []byte(`{"value": 100, "unit": "Gigacurie"}`)
    gigacuriesResult, err := factory.FromDtoJSON(gigacuriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigacurie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigacuriesResult.Convert(units.RadioactivityGigacurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigacurie = %v, want %v", converted, 100)
    }
    // Test JSON with Teracurie unit
    teracuriesJSON := []byte(`{"value": 100, "unit": "Teracurie"}`)
    teracuriesResult, err := factory.FromDtoJSON(teracuriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Teracurie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teracuriesResult.Convert(units.RadioactivityTeracurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teracurie = %v, want %v", converted, 100)
    }
    // Test JSON with Picorutherford unit
    picorutherfordsJSON := []byte(`{"value": 100, "unit": "Picorutherford"}`)
    picorutherfordsResult, err := factory.FromDtoJSON(picorutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picorutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picorutherfordsResult.Convert(units.RadioactivityPicorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picorutherford = %v, want %v", converted, 100)
    }
    // Test JSON with Nanorutherford unit
    nanorutherfordsJSON := []byte(`{"value": 100, "unit": "Nanorutherford"}`)
    nanorutherfordsResult, err := factory.FromDtoJSON(nanorutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanorutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanorutherfordsResult.Convert(units.RadioactivityNanorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanorutherford = %v, want %v", converted, 100)
    }
    // Test JSON with Microrutherford unit
    microrutherfordsJSON := []byte(`{"value": 100, "unit": "Microrutherford"}`)
    microrutherfordsResult, err := factory.FromDtoJSON(microrutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microrutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microrutherfordsResult.Convert(units.RadioactivityMicrorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microrutherford = %v, want %v", converted, 100)
    }
    // Test JSON with Millirutherford unit
    millirutherfordsJSON := []byte(`{"value": 100, "unit": "Millirutherford"}`)
    millirutherfordsResult, err := factory.FromDtoJSON(millirutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millirutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millirutherfordsResult.Convert(units.RadioactivityMillirutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millirutherford = %v, want %v", converted, 100)
    }
    // Test JSON with Kilorutherford unit
    kilorutherfordsJSON := []byte(`{"value": 100, "unit": "Kilorutherford"}`)
    kilorutherfordsResult, err := factory.FromDtoJSON(kilorutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilorutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilorutherfordsResult.Convert(units.RadioactivityKilorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilorutherford = %v, want %v", converted, 100)
    }
    // Test JSON with Megarutherford unit
    megarutherfordsJSON := []byte(`{"value": 100, "unit": "Megarutherford"}`)
    megarutherfordsResult, err := factory.FromDtoJSON(megarutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megarutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megarutherfordsResult.Convert(units.RadioactivityMegarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megarutherford = %v, want %v", converted, 100)
    }
    // Test JSON with Gigarutherford unit
    gigarutherfordsJSON := []byte(`{"value": 100, "unit": "Gigarutherford"}`)
    gigarutherfordsResult, err := factory.FromDtoJSON(gigarutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigarutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigarutherfordsResult.Convert(units.RadioactivityGigarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigarutherford = %v, want %v", converted, 100)
    }
    // Test JSON with Terarutherford unit
    terarutherfordsJSON := []byte(`{"value": 100, "unit": "Terarutherford"}`)
    terarutherfordsResult, err := factory.FromDtoJSON(terarutherfordsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terarutherford unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terarutherfordsResult.Convert(units.RadioactivityTerarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terarutherford = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Becquerel"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromBecquerels function
func TestRadioactivityFactory_FromBecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBecquerels(100)
    if err != nil {
        t.Errorf("FromBecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityBecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBecquerels(math.NaN())
    if err == nil {
        t.Error("FromBecquerels() with NaN value should return error")
    }

    _, err = factory.FromBecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromBecquerels() with +Inf value should return error")
    }

    _, err = factory.FromBecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromBecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBecquerels(0)
    if err != nil {
        t.Errorf("FromBecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityBecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromCuries function
func TestRadioactivityFactory_FromCuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCuries(100)
    if err != nil {
        t.Errorf("FromCuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityCurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCuries(math.NaN())
    if err == nil {
        t.Error("FromCuries() with NaN value should return error")
    }

    _, err = factory.FromCuries(math.Inf(1))
    if err == nil {
        t.Error("FromCuries() with +Inf value should return error")
    }

    _, err = factory.FromCuries(math.Inf(-1))
    if err == nil {
        t.Error("FromCuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCuries(0)
    if err != nil {
        t.Errorf("FromCuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityCurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCuries() with zero value = %v, want 0", converted)
    }
}
// Test FromRutherfords function
func TestRadioactivityFactory_FromRutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRutherfords(100)
    if err != nil {
        t.Errorf("FromRutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityRutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRutherfords(math.NaN())
    if err == nil {
        t.Error("FromRutherfords() with NaN value should return error")
    }

    _, err = factory.FromRutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromRutherfords() with +Inf value should return error")
    }

    _, err = factory.FromRutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromRutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRutherfords(0)
    if err != nil {
        t.Errorf("FromRutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityRutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRutherfords() with zero value = %v, want 0", converted)
    }
}
// Test FromPicobecquerels function
func TestRadioactivityFactory_FromPicobecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicobecquerels(100)
    if err != nil {
        t.Errorf("FromPicobecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityPicobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicobecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicobecquerels(math.NaN())
    if err == nil {
        t.Error("FromPicobecquerels() with NaN value should return error")
    }

    _, err = factory.FromPicobecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromPicobecquerels() with +Inf value should return error")
    }

    _, err = factory.FromPicobecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromPicobecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicobecquerels(0)
    if err != nil {
        t.Errorf("FromPicobecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityPicobecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicobecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromNanobecquerels function
func TestRadioactivityFactory_FromNanobecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanobecquerels(100)
    if err != nil {
        t.Errorf("FromNanobecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityNanobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanobecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanobecquerels(math.NaN())
    if err == nil {
        t.Error("FromNanobecquerels() with NaN value should return error")
    }

    _, err = factory.FromNanobecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromNanobecquerels() with +Inf value should return error")
    }

    _, err = factory.FromNanobecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromNanobecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanobecquerels(0)
    if err != nil {
        t.Errorf("FromNanobecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityNanobecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanobecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrobecquerels function
func TestRadioactivityFactory_FromMicrobecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrobecquerels(100)
    if err != nil {
        t.Errorf("FromMicrobecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMicrobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrobecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrobecquerels(math.NaN())
    if err == nil {
        t.Error("FromMicrobecquerels() with NaN value should return error")
    }

    _, err = factory.FromMicrobecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromMicrobecquerels() with +Inf value should return error")
    }

    _, err = factory.FromMicrobecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrobecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrobecquerels(0)
    if err != nil {
        t.Errorf("FromMicrobecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMicrobecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrobecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromMillibecquerels function
func TestRadioactivityFactory_FromMillibecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillibecquerels(100)
    if err != nil {
        t.Errorf("FromMillibecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMillibecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillibecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillibecquerels(math.NaN())
    if err == nil {
        t.Error("FromMillibecquerels() with NaN value should return error")
    }

    _, err = factory.FromMillibecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromMillibecquerels() with +Inf value should return error")
    }

    _, err = factory.FromMillibecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromMillibecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillibecquerels(0)
    if err != nil {
        t.Errorf("FromMillibecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMillibecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillibecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobecquerels function
func TestRadioactivityFactory_FromKilobecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobecquerels(100)
    if err != nil {
        t.Errorf("FromKilobecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityKilobecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobecquerels(math.NaN())
    if err == nil {
        t.Error("FromKilobecquerels() with NaN value should return error")
    }

    _, err = factory.FromKilobecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromKilobecquerels() with +Inf value should return error")
    }

    _, err = factory.FromKilobecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobecquerels(0)
    if err != nil {
        t.Errorf("FromKilobecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityKilobecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromMegabecquerels function
func TestRadioactivityFactory_FromMegabecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegabecquerels(100)
    if err != nil {
        t.Errorf("FromMegabecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMegabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegabecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegabecquerels(math.NaN())
    if err == nil {
        t.Error("FromMegabecquerels() with NaN value should return error")
    }

    _, err = factory.FromMegabecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromMegabecquerels() with +Inf value should return error")
    }

    _, err = factory.FromMegabecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromMegabecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegabecquerels(0)
    if err != nil {
        t.Errorf("FromMegabecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMegabecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegabecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromGigabecquerels function
func TestRadioactivityFactory_FromGigabecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigabecquerels(100)
    if err != nil {
        t.Errorf("FromGigabecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityGigabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigabecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigabecquerels(math.NaN())
    if err == nil {
        t.Error("FromGigabecquerels() with NaN value should return error")
    }

    _, err = factory.FromGigabecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromGigabecquerels() with +Inf value should return error")
    }

    _, err = factory.FromGigabecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromGigabecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigabecquerels(0)
    if err != nil {
        t.Errorf("FromGigabecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityGigabecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigabecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromTerabecquerels function
func TestRadioactivityFactory_FromTerabecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerabecquerels(100)
    if err != nil {
        t.Errorf("FromTerabecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityTerabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerabecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerabecquerels(math.NaN())
    if err == nil {
        t.Error("FromTerabecquerels() with NaN value should return error")
    }

    _, err = factory.FromTerabecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromTerabecquerels() with +Inf value should return error")
    }

    _, err = factory.FromTerabecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromTerabecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerabecquerels(0)
    if err != nil {
        t.Errorf("FromTerabecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityTerabecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerabecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromPetabecquerels function
func TestRadioactivityFactory_FromPetabecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetabecquerels(100)
    if err != nil {
        t.Errorf("FromPetabecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityPetabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetabecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetabecquerels(math.NaN())
    if err == nil {
        t.Error("FromPetabecquerels() with NaN value should return error")
    }

    _, err = factory.FromPetabecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromPetabecquerels() with +Inf value should return error")
    }

    _, err = factory.FromPetabecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromPetabecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetabecquerels(0)
    if err != nil {
        t.Errorf("FromPetabecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityPetabecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetabecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromExabecquerels function
func TestRadioactivityFactory_FromExabecquerels(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromExabecquerels(100)
    if err != nil {
        t.Errorf("FromExabecquerels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityExabecquerel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromExabecquerels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromExabecquerels(math.NaN())
    if err == nil {
        t.Error("FromExabecquerels() with NaN value should return error")
    }

    _, err = factory.FromExabecquerels(math.Inf(1))
    if err == nil {
        t.Error("FromExabecquerels() with +Inf value should return error")
    }

    _, err = factory.FromExabecquerels(math.Inf(-1))
    if err == nil {
        t.Error("FromExabecquerels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromExabecquerels(0)
    if err != nil {
        t.Errorf("FromExabecquerels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityExabecquerel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromExabecquerels() with zero value = %v, want 0", converted)
    }
}
// Test FromPicocuries function
func TestRadioactivityFactory_FromPicocuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicocuries(100)
    if err != nil {
        t.Errorf("FromPicocuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityPicocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicocuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicocuries(math.NaN())
    if err == nil {
        t.Error("FromPicocuries() with NaN value should return error")
    }

    _, err = factory.FromPicocuries(math.Inf(1))
    if err == nil {
        t.Error("FromPicocuries() with +Inf value should return error")
    }

    _, err = factory.FromPicocuries(math.Inf(-1))
    if err == nil {
        t.Error("FromPicocuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicocuries(0)
    if err != nil {
        t.Errorf("FromPicocuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityPicocurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicocuries() with zero value = %v, want 0", converted)
    }
}
// Test FromNanocuries function
func TestRadioactivityFactory_FromNanocuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanocuries(100)
    if err != nil {
        t.Errorf("FromNanocuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityNanocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanocuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanocuries(math.NaN())
    if err == nil {
        t.Error("FromNanocuries() with NaN value should return error")
    }

    _, err = factory.FromNanocuries(math.Inf(1))
    if err == nil {
        t.Error("FromNanocuries() with +Inf value should return error")
    }

    _, err = factory.FromNanocuries(math.Inf(-1))
    if err == nil {
        t.Error("FromNanocuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanocuries(0)
    if err != nil {
        t.Errorf("FromNanocuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityNanocurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanocuries() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrocuries function
func TestRadioactivityFactory_FromMicrocuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrocuries(100)
    if err != nil {
        t.Errorf("FromMicrocuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMicrocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrocuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrocuries(math.NaN())
    if err == nil {
        t.Error("FromMicrocuries() with NaN value should return error")
    }

    _, err = factory.FromMicrocuries(math.Inf(1))
    if err == nil {
        t.Error("FromMicrocuries() with +Inf value should return error")
    }

    _, err = factory.FromMicrocuries(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrocuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrocuries(0)
    if err != nil {
        t.Errorf("FromMicrocuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMicrocurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrocuries() with zero value = %v, want 0", converted)
    }
}
// Test FromMillicuries function
func TestRadioactivityFactory_FromMillicuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillicuries(100)
    if err != nil {
        t.Errorf("FromMillicuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMillicurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillicuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillicuries(math.NaN())
    if err == nil {
        t.Error("FromMillicuries() with NaN value should return error")
    }

    _, err = factory.FromMillicuries(math.Inf(1))
    if err == nil {
        t.Error("FromMillicuries() with +Inf value should return error")
    }

    _, err = factory.FromMillicuries(math.Inf(-1))
    if err == nil {
        t.Error("FromMillicuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillicuries(0)
    if err != nil {
        t.Errorf("FromMillicuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMillicurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillicuries() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocuries function
func TestRadioactivityFactory_FromKilocuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocuries(100)
    if err != nil {
        t.Errorf("FromKilocuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityKilocurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocuries(math.NaN())
    if err == nil {
        t.Error("FromKilocuries() with NaN value should return error")
    }

    _, err = factory.FromKilocuries(math.Inf(1))
    if err == nil {
        t.Error("FromKilocuries() with +Inf value should return error")
    }

    _, err = factory.FromKilocuries(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocuries(0)
    if err != nil {
        t.Errorf("FromKilocuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityKilocurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocuries() with zero value = %v, want 0", converted)
    }
}
// Test FromMegacuries function
func TestRadioactivityFactory_FromMegacuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegacuries(100)
    if err != nil {
        t.Errorf("FromMegacuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMegacurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegacuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegacuries(math.NaN())
    if err == nil {
        t.Error("FromMegacuries() with NaN value should return error")
    }

    _, err = factory.FromMegacuries(math.Inf(1))
    if err == nil {
        t.Error("FromMegacuries() with +Inf value should return error")
    }

    _, err = factory.FromMegacuries(math.Inf(-1))
    if err == nil {
        t.Error("FromMegacuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegacuries(0)
    if err != nil {
        t.Errorf("FromMegacuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMegacurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegacuries() with zero value = %v, want 0", converted)
    }
}
// Test FromGigacuries function
func TestRadioactivityFactory_FromGigacuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigacuries(100)
    if err != nil {
        t.Errorf("FromGigacuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityGigacurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigacuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigacuries(math.NaN())
    if err == nil {
        t.Error("FromGigacuries() with NaN value should return error")
    }

    _, err = factory.FromGigacuries(math.Inf(1))
    if err == nil {
        t.Error("FromGigacuries() with +Inf value should return error")
    }

    _, err = factory.FromGigacuries(math.Inf(-1))
    if err == nil {
        t.Error("FromGigacuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigacuries(0)
    if err != nil {
        t.Errorf("FromGigacuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityGigacurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigacuries() with zero value = %v, want 0", converted)
    }
}
// Test FromTeracuries function
func TestRadioactivityFactory_FromTeracuries(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeracuries(100)
    if err != nil {
        t.Errorf("FromTeracuries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityTeracurie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTeracuries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTeracuries(math.NaN())
    if err == nil {
        t.Error("FromTeracuries() with NaN value should return error")
    }

    _, err = factory.FromTeracuries(math.Inf(1))
    if err == nil {
        t.Error("FromTeracuries() with +Inf value should return error")
    }

    _, err = factory.FromTeracuries(math.Inf(-1))
    if err == nil {
        t.Error("FromTeracuries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTeracuries(0)
    if err != nil {
        t.Errorf("FromTeracuries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityTeracurie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeracuries() with zero value = %v, want 0", converted)
    }
}
// Test FromPicorutherfords function
func TestRadioactivityFactory_FromPicorutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicorutherfords(100)
    if err != nil {
        t.Errorf("FromPicorutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityPicorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicorutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicorutherfords(math.NaN())
    if err == nil {
        t.Error("FromPicorutherfords() with NaN value should return error")
    }

    _, err = factory.FromPicorutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromPicorutherfords() with +Inf value should return error")
    }

    _, err = factory.FromPicorutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromPicorutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicorutherfords(0)
    if err != nil {
        t.Errorf("FromPicorutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityPicorutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicorutherfords() with zero value = %v, want 0", converted)
    }
}
// Test FromNanorutherfords function
func TestRadioactivityFactory_FromNanorutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanorutherfords(100)
    if err != nil {
        t.Errorf("FromNanorutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityNanorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanorutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanorutherfords(math.NaN())
    if err == nil {
        t.Error("FromNanorutherfords() with NaN value should return error")
    }

    _, err = factory.FromNanorutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromNanorutherfords() with +Inf value should return error")
    }

    _, err = factory.FromNanorutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromNanorutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanorutherfords(0)
    if err != nil {
        t.Errorf("FromNanorutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityNanorutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanorutherfords() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrorutherfords function
func TestRadioactivityFactory_FromMicrorutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrorutherfords(100)
    if err != nil {
        t.Errorf("FromMicrorutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMicrorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrorutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrorutherfords(math.NaN())
    if err == nil {
        t.Error("FromMicrorutherfords() with NaN value should return error")
    }

    _, err = factory.FromMicrorutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromMicrorutherfords() with +Inf value should return error")
    }

    _, err = factory.FromMicrorutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrorutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrorutherfords(0)
    if err != nil {
        t.Errorf("FromMicrorutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMicrorutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrorutherfords() with zero value = %v, want 0", converted)
    }
}
// Test FromMillirutherfords function
func TestRadioactivityFactory_FromMillirutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillirutherfords(100)
    if err != nil {
        t.Errorf("FromMillirutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMillirutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillirutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillirutherfords(math.NaN())
    if err == nil {
        t.Error("FromMillirutherfords() with NaN value should return error")
    }

    _, err = factory.FromMillirutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromMillirutherfords() with +Inf value should return error")
    }

    _, err = factory.FromMillirutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromMillirutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillirutherfords(0)
    if err != nil {
        t.Errorf("FromMillirutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMillirutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillirutherfords() with zero value = %v, want 0", converted)
    }
}
// Test FromKilorutherfords function
func TestRadioactivityFactory_FromKilorutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilorutherfords(100)
    if err != nil {
        t.Errorf("FromKilorutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityKilorutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilorutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilorutherfords(math.NaN())
    if err == nil {
        t.Error("FromKilorutherfords() with NaN value should return error")
    }

    _, err = factory.FromKilorutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromKilorutherfords() with +Inf value should return error")
    }

    _, err = factory.FromKilorutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromKilorutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilorutherfords(0)
    if err != nil {
        t.Errorf("FromKilorutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityKilorutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilorutherfords() with zero value = %v, want 0", converted)
    }
}
// Test FromMegarutherfords function
func TestRadioactivityFactory_FromMegarutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegarutherfords(100)
    if err != nil {
        t.Errorf("FromMegarutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityMegarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegarutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegarutherfords(math.NaN())
    if err == nil {
        t.Error("FromMegarutherfords() with NaN value should return error")
    }

    _, err = factory.FromMegarutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromMegarutherfords() with +Inf value should return error")
    }

    _, err = factory.FromMegarutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromMegarutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegarutherfords(0)
    if err != nil {
        t.Errorf("FromMegarutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityMegarutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegarutherfords() with zero value = %v, want 0", converted)
    }
}
// Test FromGigarutherfords function
func TestRadioactivityFactory_FromGigarutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigarutherfords(100)
    if err != nil {
        t.Errorf("FromGigarutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityGigarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigarutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigarutherfords(math.NaN())
    if err == nil {
        t.Error("FromGigarutherfords() with NaN value should return error")
    }

    _, err = factory.FromGigarutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromGigarutherfords() with +Inf value should return error")
    }

    _, err = factory.FromGigarutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromGigarutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigarutherfords(0)
    if err != nil {
        t.Errorf("FromGigarutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityGigarutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigarutherfords() with zero value = %v, want 0", converted)
    }
}
// Test FromTerarutherfords function
func TestRadioactivityFactory_FromTerarutherfords(t *testing.T) {
    factory := units.RadioactivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerarutherfords(100)
    if err != nil {
        t.Errorf("FromTerarutherfords() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadioactivityTerarutherford)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerarutherfords() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerarutherfords(math.NaN())
    if err == nil {
        t.Error("FromTerarutherfords() with NaN value should return error")
    }

    _, err = factory.FromTerarutherfords(math.Inf(1))
    if err == nil {
        t.Error("FromTerarutherfords() with +Inf value should return error")
    }

    _, err = factory.FromTerarutherfords(math.Inf(-1))
    if err == nil {
        t.Error("FromTerarutherfords() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerarutherfords(0)
    if err != nil {
        t.Errorf("FromTerarutherfords() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadioactivityTerarutherford)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerarutherfords() with zero value = %v, want 0", converted)
    }
}

func TestRadioactivityToString(t *testing.T) {
	factory := units.RadioactivityFactory{}
	a, err := factory.CreateRadioactivity(45, units.RadioactivityBecquerel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RadioactivityBecquerel, 2)
	expected := "45.00 " + units.GetRadioactivityAbbreviation(units.RadioactivityBecquerel)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RadioactivityBecquerel, -1)
	expected = "45 " + units.GetRadioactivityAbbreviation(units.RadioactivityBecquerel)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRadioactivity_EqualityAndComparison(t *testing.T) {
	factory := units.RadioactivityFactory{}
	a1, _ := factory.CreateRadioactivity(60, units.RadioactivityBecquerel)
	a2, _ := factory.CreateRadioactivity(60, units.RadioactivityBecquerel)
	a3, _ := factory.CreateRadioactivity(120, units.RadioactivityBecquerel)

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

func TestRadioactivity_Arithmetic(t *testing.T) {
	factory := units.RadioactivityFactory{}
	a1, _ := factory.CreateRadioactivity(30, units.RadioactivityBecquerel)
	a2, _ := factory.CreateRadioactivity(45, units.RadioactivityBecquerel)

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


func TestGetRadioactivityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.RadioactivityUnits
        want string
    }{
        {
            name: "Becquerel abbreviation",
            unit: units.RadioactivityBecquerel,
            want: "Bq",
        },
        {
            name: "Curie abbreviation",
            unit: units.RadioactivityCurie,
            want: "Ci",
        },
        {
            name: "Rutherford abbreviation",
            unit: units.RadioactivityRutherford,
            want: "Rd",
        },
        {
            name: "Picobecquerel abbreviation",
            unit: units.RadioactivityPicobecquerel,
            want: "pBq",
        },
        {
            name: "Nanobecquerel abbreviation",
            unit: units.RadioactivityNanobecquerel,
            want: "nBq",
        },
        {
            name: "Microbecquerel abbreviation",
            unit: units.RadioactivityMicrobecquerel,
            want: "μBq",
        },
        {
            name: "Millibecquerel abbreviation",
            unit: units.RadioactivityMillibecquerel,
            want: "mBq",
        },
        {
            name: "Kilobecquerel abbreviation",
            unit: units.RadioactivityKilobecquerel,
            want: "kBq",
        },
        {
            name: "Megabecquerel abbreviation",
            unit: units.RadioactivityMegabecquerel,
            want: "MBq",
        },
        {
            name: "Gigabecquerel abbreviation",
            unit: units.RadioactivityGigabecquerel,
            want: "GBq",
        },
        {
            name: "Terabecquerel abbreviation",
            unit: units.RadioactivityTerabecquerel,
            want: "TBq",
        },
        {
            name: "Petabecquerel abbreviation",
            unit: units.RadioactivityPetabecquerel,
            want: "PBq",
        },
        {
            name: "Exabecquerel abbreviation",
            unit: units.RadioactivityExabecquerel,
            want: "EBq",
        },
        {
            name: "Picocurie abbreviation",
            unit: units.RadioactivityPicocurie,
            want: "pCi",
        },
        {
            name: "Nanocurie abbreviation",
            unit: units.RadioactivityNanocurie,
            want: "nCi",
        },
        {
            name: "Microcurie abbreviation",
            unit: units.RadioactivityMicrocurie,
            want: "μCi",
        },
        {
            name: "Millicurie abbreviation",
            unit: units.RadioactivityMillicurie,
            want: "mCi",
        },
        {
            name: "Kilocurie abbreviation",
            unit: units.RadioactivityKilocurie,
            want: "kCi",
        },
        {
            name: "Megacurie abbreviation",
            unit: units.RadioactivityMegacurie,
            want: "MCi",
        },
        {
            name: "Gigacurie abbreviation",
            unit: units.RadioactivityGigacurie,
            want: "GCi",
        },
        {
            name: "Teracurie abbreviation",
            unit: units.RadioactivityTeracurie,
            want: "TCi",
        },
        {
            name: "Picorutherford abbreviation",
            unit: units.RadioactivityPicorutherford,
            want: "pRd",
        },
        {
            name: "Nanorutherford abbreviation",
            unit: units.RadioactivityNanorutherford,
            want: "nRd",
        },
        {
            name: "Microrutherford abbreviation",
            unit: units.RadioactivityMicrorutherford,
            want: "μRd",
        },
        {
            name: "Millirutherford abbreviation",
            unit: units.RadioactivityMillirutherford,
            want: "mRd",
        },
        {
            name: "Kilorutherford abbreviation",
            unit: units.RadioactivityKilorutherford,
            want: "kRd",
        },
        {
            name: "Megarutherford abbreviation",
            unit: units.RadioactivityMegarutherford,
            want: "MRd",
        },
        {
            name: "Gigarutherford abbreviation",
            unit: units.RadioactivityGigarutherford,
            want: "GRd",
        },
        {
            name: "Terarutherford abbreviation",
            unit: units.RadioactivityTerarutherford,
            want: "TRd",
        },
        {
            name: "invalid unit",
            unit: units.RadioactivityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetRadioactivityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetRadioactivityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestRadioactivity_String(t *testing.T) {
    factory := units.RadioactivityFactory{}
    
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
            unit, err := factory.CreateRadioactivity(tt.value, units.RadioactivityBecquerel)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Radioactivity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}