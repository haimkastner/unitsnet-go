// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Kilogram"}`
	
	factory := units.MassDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassKilogram {
		t.Errorf("expected unit %v, got %v", units.MassKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Kilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassDto_ToJSON(t *testing.T) {
	dto := units.MassDto{
		Value: 45,
		Unit:  units.MassKilogram,
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
	if result["unit"].(string) != string(units.MassKilogram) {
		t.Errorf("expected unit %s, got %v", units.MassKilogram, result["unit"])
	}
}

func TestNewMass_InvalidValue(t *testing.T) {
	factory := units.MassFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMass(math.NaN(), units.MassKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMass(math.Inf(1), units.MassKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassConversions(t *testing.T) {
	factory := units.MassFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMass(180, units.MassKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Grams.
		// No expected conversion value provided for Grams, verifying result is not NaN.
		result := a.Grams()
		cacheResult := a.Grams()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Grams returned NaN")
		}
	}
	{
		// Test conversion to Tonnes.
		// No expected conversion value provided for Tonnes, verifying result is not NaN.
		result := a.Tonnes()
		cacheResult := a.Tonnes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Tonnes returned NaN")
		}
	}
	{
		// Test conversion to ShortTons.
		// No expected conversion value provided for ShortTons, verifying result is not NaN.
		result := a.ShortTons()
		cacheResult := a.ShortTons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ShortTons returned NaN")
		}
	}
	{
		// Test conversion to LongTons.
		// No expected conversion value provided for LongTons, verifying result is not NaN.
		result := a.LongTons()
		cacheResult := a.LongTons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LongTons returned NaN")
		}
	}
	{
		// Test conversion to Pounds.
		// No expected conversion value provided for Pounds, verifying result is not NaN.
		result := a.Pounds()
		cacheResult := a.Pounds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Pounds returned NaN")
		}
	}
	{
		// Test conversion to Ounces.
		// No expected conversion value provided for Ounces, verifying result is not NaN.
		result := a.Ounces()
		cacheResult := a.Ounces()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Ounces returned NaN")
		}
	}
	{
		// Test conversion to Slugs.
		// No expected conversion value provided for Slugs, verifying result is not NaN.
		result := a.Slugs()
		cacheResult := a.Slugs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Slugs returned NaN")
		}
	}
	{
		// Test conversion to Stone.
		// No expected conversion value provided for Stone, verifying result is not NaN.
		result := a.Stone()
		cacheResult := a.Stone()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Stone returned NaN")
		}
	}
	{
		// Test conversion to ShortHundredweight.
		// No expected conversion value provided for ShortHundredweight, verifying result is not NaN.
		result := a.ShortHundredweight()
		cacheResult := a.ShortHundredweight()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ShortHundredweight returned NaN")
		}
	}
	{
		// Test conversion to LongHundredweight.
		// No expected conversion value provided for LongHundredweight, verifying result is not NaN.
		result := a.LongHundredweight()
		cacheResult := a.LongHundredweight()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LongHundredweight returned NaN")
		}
	}
	{
		// Test conversion to Grains.
		// No expected conversion value provided for Grains, verifying result is not NaN.
		result := a.Grains()
		cacheResult := a.Grains()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Grains returned NaN")
		}
	}
	{
		// Test conversion to SolarMasses.
		// No expected conversion value provided for SolarMasses, verifying result is not NaN.
		result := a.SolarMasses()
		cacheResult := a.SolarMasses()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SolarMasses returned NaN")
		}
	}
	{
		// Test conversion to EarthMasses.
		// No expected conversion value provided for EarthMasses, verifying result is not NaN.
		result := a.EarthMasses()
		cacheResult := a.EarthMasses()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to EarthMasses returned NaN")
		}
	}
	{
		// Test conversion to Femtograms.
		// No expected conversion value provided for Femtograms, verifying result is not NaN.
		result := a.Femtograms()
		cacheResult := a.Femtograms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Femtograms returned NaN")
		}
	}
	{
		// Test conversion to Picograms.
		// No expected conversion value provided for Picograms, verifying result is not NaN.
		result := a.Picograms()
		cacheResult := a.Picograms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picograms returned NaN")
		}
	}
	{
		// Test conversion to Nanograms.
		// No expected conversion value provided for Nanograms, verifying result is not NaN.
		result := a.Nanograms()
		cacheResult := a.Nanograms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanograms returned NaN")
		}
	}
	{
		// Test conversion to Micrograms.
		// No expected conversion value provided for Micrograms, verifying result is not NaN.
		result := a.Micrograms()
		cacheResult := a.Micrograms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Micrograms returned NaN")
		}
	}
	{
		// Test conversion to Milligrams.
		// No expected conversion value provided for Milligrams, verifying result is not NaN.
		result := a.Milligrams()
		cacheResult := a.Milligrams()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milligrams returned NaN")
		}
	}
	{
		// Test conversion to Centigrams.
		// No expected conversion value provided for Centigrams, verifying result is not NaN.
		result := a.Centigrams()
		cacheResult := a.Centigrams()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Centigrams returned NaN")
		}
	}
	{
		// Test conversion to Decigrams.
		// No expected conversion value provided for Decigrams, verifying result is not NaN.
		result := a.Decigrams()
		cacheResult := a.Decigrams()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decigrams returned NaN")
		}
	}
	{
		// Test conversion to Decagrams.
		// No expected conversion value provided for Decagrams, verifying result is not NaN.
		result := a.Decagrams()
		cacheResult := a.Decagrams()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decagrams returned NaN")
		}
	}
	{
		// Test conversion to Hectograms.
		// No expected conversion value provided for Hectograms, verifying result is not NaN.
		result := a.Hectograms()
		cacheResult := a.Hectograms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Hectograms returned NaN")
		}
	}
	{
		// Test conversion to Kilograms.
		// No expected conversion value provided for Kilograms, verifying result is not NaN.
		result := a.Kilograms()
		cacheResult := a.Kilograms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilograms returned NaN")
		}
	}
	{
		// Test conversion to Kilotonnes.
		// No expected conversion value provided for Kilotonnes, verifying result is not NaN.
		result := a.Kilotonnes()
		cacheResult := a.Kilotonnes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilotonnes returned NaN")
		}
	}
	{
		// Test conversion to Megatonnes.
		// No expected conversion value provided for Megatonnes, verifying result is not NaN.
		result := a.Megatonnes()
		cacheResult := a.Megatonnes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megatonnes returned NaN")
		}
	}
	{
		// Test conversion to Kilopounds.
		// No expected conversion value provided for Kilopounds, verifying result is not NaN.
		result := a.Kilopounds()
		cacheResult := a.Kilopounds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilopounds returned NaN")
		}
	}
	{
		// Test conversion to Megapounds.
		// No expected conversion value provided for Megapounds, verifying result is not NaN.
		result := a.Megapounds()
		cacheResult := a.Megapounds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megapounds returned NaN")
		}
	}
}

func TestMass_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassFactory{}
	a, err := factory.CreateMass(90, units.MassKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassKilogram {
		t.Errorf("expected default unit Kilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassGram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassKilogram {
		t.Errorf("expected unit Kilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassFactory_FromDto(t *testing.T) {
    factory := units.MassFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MassDto{
        Value: 100,
        Unit:  units.MassKilogram,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MassDto{
        Value: math.NaN(),
        Unit:  units.MassKilogram,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Gram conversion
    gramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassGram,
    }
    
    var gramsResult *units.Mass
    gramsResult, err = factory.FromDto(gramsDto)
    if err != nil {
        t.Errorf("FromDto() with Gram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gramsResult.Convert(units.MassGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gram = %v, want %v", converted, 100)
    }
    // Test Tonne conversion
    tonnesDto := units.MassDto{
        Value: 100,
        Unit:  units.MassTonne,
    }
    
    var tonnesResult *units.Mass
    tonnesResult, err = factory.FromDto(tonnesDto)
    if err != nil {
        t.Errorf("FromDto() with Tonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnesResult.Convert(units.MassTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tonne = %v, want %v", converted, 100)
    }
    // Test ShortTon conversion
    short_tonsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassShortTon,
    }
    
    var short_tonsResult *units.Mass
    short_tonsResult, err = factory.FromDto(short_tonsDto)
    if err != nil {
        t.Errorf("FromDto() with ShortTon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = short_tonsResult.Convert(units.MassShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ShortTon = %v, want %v", converted, 100)
    }
    // Test LongTon conversion
    long_tonsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassLongTon,
    }
    
    var long_tonsResult *units.Mass
    long_tonsResult, err = factory.FromDto(long_tonsDto)
    if err != nil {
        t.Errorf("FromDto() with LongTon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = long_tonsResult.Convert(units.MassLongTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LongTon = %v, want %v", converted, 100)
    }
    // Test Pound conversion
    poundsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassPound,
    }
    
    var poundsResult *units.Mass
    poundsResult, err = factory.FromDto(poundsDto)
    if err != nil {
        t.Errorf("FromDto() with Pound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = poundsResult.Convert(units.MassPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Pound = %v, want %v", converted, 100)
    }
    // Test Ounce conversion
    ouncesDto := units.MassDto{
        Value: 100,
        Unit:  units.MassOunce,
    }
    
    var ouncesResult *units.Mass
    ouncesResult, err = factory.FromDto(ouncesDto)
    if err != nil {
        t.Errorf("FromDto() with Ounce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ouncesResult.Convert(units.MassOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Ounce = %v, want %v", converted, 100)
    }
    // Test Slug conversion
    slugsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassSlug,
    }
    
    var slugsResult *units.Mass
    slugsResult, err = factory.FromDto(slugsDto)
    if err != nil {
        t.Errorf("FromDto() with Slug returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugsResult.Convert(units.MassSlug)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Slug = %v, want %v", converted, 100)
    }
    // Test Stone conversion
    stoneDto := units.MassDto{
        Value: 100,
        Unit:  units.MassStone,
    }
    
    var stoneResult *units.Mass
    stoneResult, err = factory.FromDto(stoneDto)
    if err != nil {
        t.Errorf("FromDto() with Stone returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = stoneResult.Convert(units.MassStone)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Stone = %v, want %v", converted, 100)
    }
    // Test ShortHundredweight conversion
    short_hundredweightDto := units.MassDto{
        Value: 100,
        Unit:  units.MassShortHundredweight,
    }
    
    var short_hundredweightResult *units.Mass
    short_hundredweightResult, err = factory.FromDto(short_hundredweightDto)
    if err != nil {
        t.Errorf("FromDto() with ShortHundredweight returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = short_hundredweightResult.Convert(units.MassShortHundredweight)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ShortHundredweight = %v, want %v", converted, 100)
    }
    // Test LongHundredweight conversion
    long_hundredweightDto := units.MassDto{
        Value: 100,
        Unit:  units.MassLongHundredweight,
    }
    
    var long_hundredweightResult *units.Mass
    long_hundredweightResult, err = factory.FromDto(long_hundredweightDto)
    if err != nil {
        t.Errorf("FromDto() with LongHundredweight returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = long_hundredweightResult.Convert(units.MassLongHundredweight)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LongHundredweight = %v, want %v", converted, 100)
    }
    // Test Grain conversion
    grainsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassGrain,
    }
    
    var grainsResult *units.Mass
    grainsResult, err = factory.FromDto(grainsDto)
    if err != nil {
        t.Errorf("FromDto() with Grain returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grainsResult.Convert(units.MassGrain)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Grain = %v, want %v", converted, 100)
    }
    // Test SolarMass conversion
    solar_massesDto := units.MassDto{
        Value: 100,
        Unit:  units.MassSolarMass,
    }
    
    var solar_massesResult *units.Mass
    solar_massesResult, err = factory.FromDto(solar_massesDto)
    if err != nil {
        t.Errorf("FromDto() with SolarMass returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solar_massesResult.Convert(units.MassSolarMass)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SolarMass = %v, want %v", converted, 100)
    }
    // Test EarthMass conversion
    earth_massesDto := units.MassDto{
        Value: 100,
        Unit:  units.MassEarthMass,
    }
    
    var earth_massesResult *units.Mass
    earth_massesResult, err = factory.FromDto(earth_massesDto)
    if err != nil {
        t.Errorf("FromDto() with EarthMass returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = earth_massesResult.Convert(units.MassEarthMass)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for EarthMass = %v, want %v", converted, 100)
    }
    // Test Femtogram conversion
    femtogramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassFemtogram,
    }
    
    var femtogramsResult *units.Mass
    femtogramsResult, err = factory.FromDto(femtogramsDto)
    if err != nil {
        t.Errorf("FromDto() with Femtogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtogramsResult.Convert(units.MassFemtogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtogram = %v, want %v", converted, 100)
    }
    // Test Picogram conversion
    picogramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassPicogram,
    }
    
    var picogramsResult *units.Mass
    picogramsResult, err = factory.FromDto(picogramsDto)
    if err != nil {
        t.Errorf("FromDto() with Picogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picogramsResult.Convert(units.MassPicogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picogram = %v, want %v", converted, 100)
    }
    // Test Nanogram conversion
    nanogramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassNanogram,
    }
    
    var nanogramsResult *units.Mass
    nanogramsResult, err = factory.FromDto(nanogramsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanogramsResult.Convert(units.MassNanogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanogram = %v, want %v", converted, 100)
    }
    // Test Microgram conversion
    microgramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassMicrogram,
    }
    
    var microgramsResult *units.Mass
    microgramsResult, err = factory.FromDto(microgramsDto)
    if err != nil {
        t.Errorf("FromDto() with Microgram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgramsResult.Convert(units.MassMicrogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microgram = %v, want %v", converted, 100)
    }
    // Test Milligram conversion
    milligramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassMilligram,
    }
    
    var milligramsResult *units.Mass
    milligramsResult, err = factory.FromDto(milligramsDto)
    if err != nil {
        t.Errorf("FromDto() with Milligram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligramsResult.Convert(units.MassMilligram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milligram = %v, want %v", converted, 100)
    }
    // Test Centigram conversion
    centigramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassCentigram,
    }
    
    var centigramsResult *units.Mass
    centigramsResult, err = factory.FromDto(centigramsDto)
    if err != nil {
        t.Errorf("FromDto() with Centigram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigramsResult.Convert(units.MassCentigram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centigram = %v, want %v", converted, 100)
    }
    // Test Decigram conversion
    decigramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassDecigram,
    }
    
    var decigramsResult *units.Mass
    decigramsResult, err = factory.FromDto(decigramsDto)
    if err != nil {
        t.Errorf("FromDto() with Decigram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigramsResult.Convert(units.MassDecigram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decigram = %v, want %v", converted, 100)
    }
    // Test Decagram conversion
    decagramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassDecagram,
    }
    
    var decagramsResult *units.Mass
    decagramsResult, err = factory.FromDto(decagramsDto)
    if err != nil {
        t.Errorf("FromDto() with Decagram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagramsResult.Convert(units.MassDecagram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decagram = %v, want %v", converted, 100)
    }
    // Test Hectogram conversion
    hectogramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassHectogram,
    }
    
    var hectogramsResult *units.Mass
    hectogramsResult, err = factory.FromDto(hectogramsDto)
    if err != nil {
        t.Errorf("FromDto() with Hectogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectogramsResult.Convert(units.MassHectogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectogram = %v, want %v", converted, 100)
    }
    // Test Kilogram conversion
    kilogramsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassKilogram,
    }
    
    var kilogramsResult *units.Mass
    kilogramsResult, err = factory.FromDto(kilogramsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogramsResult.Convert(units.MassKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilogram = %v, want %v", converted, 100)
    }
    // Test Kilotonne conversion
    kilotonnesDto := units.MassDto{
        Value: 100,
        Unit:  units.MassKilotonne,
    }
    
    var kilotonnesResult *units.Mass
    kilotonnesResult, err = factory.FromDto(kilotonnesDto)
    if err != nil {
        t.Errorf("FromDto() with Kilotonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonnesResult.Convert(units.MassKilotonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilotonne = %v, want %v", converted, 100)
    }
    // Test Megatonne conversion
    megatonnesDto := units.MassDto{
        Value: 100,
        Unit:  units.MassMegatonne,
    }
    
    var megatonnesResult *units.Mass
    megatonnesResult, err = factory.FromDto(megatonnesDto)
    if err != nil {
        t.Errorf("FromDto() with Megatonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonnesResult.Convert(units.MassMegatonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megatonne = %v, want %v", converted, 100)
    }
    // Test Kilopound conversion
    kilopoundsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassKilopound,
    }
    
    var kilopoundsResult *units.Mass
    kilopoundsResult, err = factory.FromDto(kilopoundsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilopound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopoundsResult.Convert(units.MassKilopound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilopound = %v, want %v", converted, 100)
    }
    // Test Megapound conversion
    megapoundsDto := units.MassDto{
        Value: 100,
        Unit:  units.MassMegapound,
    }
    
    var megapoundsResult *units.Mass
    megapoundsResult, err = factory.FromDto(megapoundsDto)
    if err != nil {
        t.Errorf("FromDto() with Megapound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapoundsResult.Convert(units.MassMegapound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megapound = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MassDto{
        Value: 0,
        Unit:  units.MassKilogram,
    }
    
    var zeroResult *units.Mass
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMassFactory_FromDtoJSON(t *testing.T) {
    factory := units.MassFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Kilogram"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Kilogram"}`)
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
    nanJSON, _ := json.Marshal(units.MassDto{
        Value: nanValue,
        Unit:  units.MassKilogram,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Gram unit
    gramsJSON := []byte(`{"value": 100, "unit": "Gram"}`)
    gramsResult, err := factory.FromDtoJSON(gramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gramsResult.Convert(units.MassGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gram = %v, want %v", converted, 100)
    }
    // Test JSON with Tonne unit
    tonnesJSON := []byte(`{"value": 100, "unit": "Tonne"}`)
    tonnesResult, err := factory.FromDtoJSON(tonnesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Tonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnesResult.Convert(units.MassTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tonne = %v, want %v", converted, 100)
    }
    // Test JSON with ShortTon unit
    short_tonsJSON := []byte(`{"value": 100, "unit": "ShortTon"}`)
    short_tonsResult, err := factory.FromDtoJSON(short_tonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ShortTon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = short_tonsResult.Convert(units.MassShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ShortTon = %v, want %v", converted, 100)
    }
    // Test JSON with LongTon unit
    long_tonsJSON := []byte(`{"value": 100, "unit": "LongTon"}`)
    long_tonsResult, err := factory.FromDtoJSON(long_tonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LongTon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = long_tonsResult.Convert(units.MassLongTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LongTon = %v, want %v", converted, 100)
    }
    // Test JSON with Pound unit
    poundsJSON := []byte(`{"value": 100, "unit": "Pound"}`)
    poundsResult, err := factory.FromDtoJSON(poundsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Pound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = poundsResult.Convert(units.MassPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Pound = %v, want %v", converted, 100)
    }
    // Test JSON with Ounce unit
    ouncesJSON := []byte(`{"value": 100, "unit": "Ounce"}`)
    ouncesResult, err := factory.FromDtoJSON(ouncesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Ounce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ouncesResult.Convert(units.MassOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Ounce = %v, want %v", converted, 100)
    }
    // Test JSON with Slug unit
    slugsJSON := []byte(`{"value": 100, "unit": "Slug"}`)
    slugsResult, err := factory.FromDtoJSON(slugsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Slug unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugsResult.Convert(units.MassSlug)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Slug = %v, want %v", converted, 100)
    }
    // Test JSON with Stone unit
    stoneJSON := []byte(`{"value": 100, "unit": "Stone"}`)
    stoneResult, err := factory.FromDtoJSON(stoneJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Stone unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = stoneResult.Convert(units.MassStone)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Stone = %v, want %v", converted, 100)
    }
    // Test JSON with ShortHundredweight unit
    short_hundredweightJSON := []byte(`{"value": 100, "unit": "ShortHundredweight"}`)
    short_hundredweightResult, err := factory.FromDtoJSON(short_hundredweightJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ShortHundredweight unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = short_hundredweightResult.Convert(units.MassShortHundredweight)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ShortHundredweight = %v, want %v", converted, 100)
    }
    // Test JSON with LongHundredweight unit
    long_hundredweightJSON := []byte(`{"value": 100, "unit": "LongHundredweight"}`)
    long_hundredweightResult, err := factory.FromDtoJSON(long_hundredweightJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LongHundredweight unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = long_hundredweightResult.Convert(units.MassLongHundredweight)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LongHundredweight = %v, want %v", converted, 100)
    }
    // Test JSON with Grain unit
    grainsJSON := []byte(`{"value": 100, "unit": "Grain"}`)
    grainsResult, err := factory.FromDtoJSON(grainsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Grain unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grainsResult.Convert(units.MassGrain)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Grain = %v, want %v", converted, 100)
    }
    // Test JSON with SolarMass unit
    solar_massesJSON := []byte(`{"value": 100, "unit": "SolarMass"}`)
    solar_massesResult, err := factory.FromDtoJSON(solar_massesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SolarMass unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solar_massesResult.Convert(units.MassSolarMass)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SolarMass = %v, want %v", converted, 100)
    }
    // Test JSON with EarthMass unit
    earth_massesJSON := []byte(`{"value": 100, "unit": "EarthMass"}`)
    earth_massesResult, err := factory.FromDtoJSON(earth_massesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with EarthMass unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = earth_massesResult.Convert(units.MassEarthMass)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for EarthMass = %v, want %v", converted, 100)
    }
    // Test JSON with Femtogram unit
    femtogramsJSON := []byte(`{"value": 100, "unit": "Femtogram"}`)
    femtogramsResult, err := factory.FromDtoJSON(femtogramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Femtogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtogramsResult.Convert(units.MassFemtogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtogram = %v, want %v", converted, 100)
    }
    // Test JSON with Picogram unit
    picogramsJSON := []byte(`{"value": 100, "unit": "Picogram"}`)
    picogramsResult, err := factory.FromDtoJSON(picogramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picogramsResult.Convert(units.MassPicogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picogram = %v, want %v", converted, 100)
    }
    // Test JSON with Nanogram unit
    nanogramsJSON := []byte(`{"value": 100, "unit": "Nanogram"}`)
    nanogramsResult, err := factory.FromDtoJSON(nanogramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanogramsResult.Convert(units.MassNanogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanogram = %v, want %v", converted, 100)
    }
    // Test JSON with Microgram unit
    microgramsJSON := []byte(`{"value": 100, "unit": "Microgram"}`)
    microgramsResult, err := factory.FromDtoJSON(microgramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microgram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgramsResult.Convert(units.MassMicrogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microgram = %v, want %v", converted, 100)
    }
    // Test JSON with Milligram unit
    milligramsJSON := []byte(`{"value": 100, "unit": "Milligram"}`)
    milligramsResult, err := factory.FromDtoJSON(milligramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milligram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligramsResult.Convert(units.MassMilligram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milligram = %v, want %v", converted, 100)
    }
    // Test JSON with Centigram unit
    centigramsJSON := []byte(`{"value": 100, "unit": "Centigram"}`)
    centigramsResult, err := factory.FromDtoJSON(centigramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centigram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigramsResult.Convert(units.MassCentigram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centigram = %v, want %v", converted, 100)
    }
    // Test JSON with Decigram unit
    decigramsJSON := []byte(`{"value": 100, "unit": "Decigram"}`)
    decigramsResult, err := factory.FromDtoJSON(decigramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decigram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigramsResult.Convert(units.MassDecigram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decigram = %v, want %v", converted, 100)
    }
    // Test JSON with Decagram unit
    decagramsJSON := []byte(`{"value": 100, "unit": "Decagram"}`)
    decagramsResult, err := factory.FromDtoJSON(decagramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decagram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagramsResult.Convert(units.MassDecagram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decagram = %v, want %v", converted, 100)
    }
    // Test JSON with Hectogram unit
    hectogramsJSON := []byte(`{"value": 100, "unit": "Hectogram"}`)
    hectogramsResult, err := factory.FromDtoJSON(hectogramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Hectogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectogramsResult.Convert(units.MassHectogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectogram = %v, want %v", converted, 100)
    }
    // Test JSON with Kilogram unit
    kilogramsJSON := []byte(`{"value": 100, "unit": "Kilogram"}`)
    kilogramsResult, err := factory.FromDtoJSON(kilogramsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogramsResult.Convert(units.MassKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilogram = %v, want %v", converted, 100)
    }
    // Test JSON with Kilotonne unit
    kilotonnesJSON := []byte(`{"value": 100, "unit": "Kilotonne"}`)
    kilotonnesResult, err := factory.FromDtoJSON(kilotonnesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilotonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonnesResult.Convert(units.MassKilotonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilotonne = %v, want %v", converted, 100)
    }
    // Test JSON with Megatonne unit
    megatonnesJSON := []byte(`{"value": 100, "unit": "Megatonne"}`)
    megatonnesResult, err := factory.FromDtoJSON(megatonnesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megatonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonnesResult.Convert(units.MassMegatonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megatonne = %v, want %v", converted, 100)
    }
    // Test JSON with Kilopound unit
    kilopoundsJSON := []byte(`{"value": 100, "unit": "Kilopound"}`)
    kilopoundsResult, err := factory.FromDtoJSON(kilopoundsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilopound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopoundsResult.Convert(units.MassKilopound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilopound = %v, want %v", converted, 100)
    }
    // Test JSON with Megapound unit
    megapoundsJSON := []byte(`{"value": 100, "unit": "Megapound"}`)
    megapoundsResult, err := factory.FromDtoJSON(megapoundsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megapound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapoundsResult.Convert(units.MassMegapound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megapound = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Kilogram"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGrams function
func TestMassFactory_FromGrams(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGrams(100)
    if err != nil {
        t.Errorf("FromGrams() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGrams() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGrams(math.NaN())
    if err == nil {
        t.Error("FromGrams() with NaN value should return error")
    }

    _, err = factory.FromGrams(math.Inf(1))
    if err == nil {
        t.Error("FromGrams() with +Inf value should return error")
    }

    _, err = factory.FromGrams(math.Inf(-1))
    if err == nil {
        t.Error("FromGrams() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGrams(0)
    if err != nil {
        t.Errorf("FromGrams() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGrams() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnes function
func TestMassFactory_FromTonnes(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnes(100)
    if err != nil {
        t.Errorf("FromTonnes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnes(math.NaN())
    if err == nil {
        t.Error("FromTonnes() with NaN value should return error")
    }

    _, err = factory.FromTonnes(math.Inf(1))
    if err == nil {
        t.Error("FromTonnes() with +Inf value should return error")
    }

    _, err = factory.FromTonnes(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnes(0)
    if err != nil {
        t.Errorf("FromTonnes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassTonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnes() with zero value = %v, want 0", converted)
    }
}
// Test FromShortTons function
func TestMassFactory_FromShortTons(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromShortTons(100)
    if err != nil {
        t.Errorf("FromShortTons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromShortTons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromShortTons(math.NaN())
    if err == nil {
        t.Error("FromShortTons() with NaN value should return error")
    }

    _, err = factory.FromShortTons(math.Inf(1))
    if err == nil {
        t.Error("FromShortTons() with +Inf value should return error")
    }

    _, err = factory.FromShortTons(math.Inf(-1))
    if err == nil {
        t.Error("FromShortTons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromShortTons(0)
    if err != nil {
        t.Errorf("FromShortTons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassShortTon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromShortTons() with zero value = %v, want 0", converted)
    }
}
// Test FromLongTons function
func TestMassFactory_FromLongTons(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLongTons(100)
    if err != nil {
        t.Errorf("FromLongTons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassLongTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLongTons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLongTons(math.NaN())
    if err == nil {
        t.Error("FromLongTons() with NaN value should return error")
    }

    _, err = factory.FromLongTons(math.Inf(1))
    if err == nil {
        t.Error("FromLongTons() with +Inf value should return error")
    }

    _, err = factory.FromLongTons(math.Inf(-1))
    if err == nil {
        t.Error("FromLongTons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLongTons(0)
    if err != nil {
        t.Errorf("FromLongTons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassLongTon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLongTons() with zero value = %v, want 0", converted)
    }
}
// Test FromPounds function
func TestMassFactory_FromPounds(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPounds(100)
    if err != nil {
        t.Errorf("FromPounds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPounds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPounds(math.NaN())
    if err == nil {
        t.Error("FromPounds() with NaN value should return error")
    }

    _, err = factory.FromPounds(math.Inf(1))
    if err == nil {
        t.Error("FromPounds() with +Inf value should return error")
    }

    _, err = factory.FromPounds(math.Inf(-1))
    if err == nil {
        t.Error("FromPounds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPounds(0)
    if err != nil {
        t.Errorf("FromPounds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassPound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPounds() with zero value = %v, want 0", converted)
    }
}
// Test FromOunces function
func TestMassFactory_FromOunces(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOunces(100)
    if err != nil {
        t.Errorf("FromOunces() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOunces() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOunces(math.NaN())
    if err == nil {
        t.Error("FromOunces() with NaN value should return error")
    }

    _, err = factory.FromOunces(math.Inf(1))
    if err == nil {
        t.Error("FromOunces() with +Inf value should return error")
    }

    _, err = factory.FromOunces(math.Inf(-1))
    if err == nil {
        t.Error("FromOunces() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOunces(0)
    if err != nil {
        t.Errorf("FromOunces() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassOunce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOunces() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugs function
func TestMassFactory_FromSlugs(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugs(100)
    if err != nil {
        t.Errorf("FromSlugs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassSlug)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugs(math.NaN())
    if err == nil {
        t.Error("FromSlugs() with NaN value should return error")
    }

    _, err = factory.FromSlugs(math.Inf(1))
    if err == nil {
        t.Error("FromSlugs() with +Inf value should return error")
    }

    _, err = factory.FromSlugs(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugs(0)
    if err != nil {
        t.Errorf("FromSlugs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassSlug)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugs() with zero value = %v, want 0", converted)
    }
}
// Test FromStone function
func TestMassFactory_FromStone(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStone(100)
    if err != nil {
        t.Errorf("FromStone() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassStone)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStone() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStone(math.NaN())
    if err == nil {
        t.Error("FromStone() with NaN value should return error")
    }

    _, err = factory.FromStone(math.Inf(1))
    if err == nil {
        t.Error("FromStone() with +Inf value should return error")
    }

    _, err = factory.FromStone(math.Inf(-1))
    if err == nil {
        t.Error("FromStone() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStone(0)
    if err != nil {
        t.Errorf("FromStone() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassStone)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStone() with zero value = %v, want 0", converted)
    }
}
// Test FromShortHundredweight function
func TestMassFactory_FromShortHundredweight(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromShortHundredweight(100)
    if err != nil {
        t.Errorf("FromShortHundredweight() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassShortHundredweight)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromShortHundredweight() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromShortHundredweight(math.NaN())
    if err == nil {
        t.Error("FromShortHundredweight() with NaN value should return error")
    }

    _, err = factory.FromShortHundredweight(math.Inf(1))
    if err == nil {
        t.Error("FromShortHundredweight() with +Inf value should return error")
    }

    _, err = factory.FromShortHundredweight(math.Inf(-1))
    if err == nil {
        t.Error("FromShortHundredweight() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromShortHundredweight(0)
    if err != nil {
        t.Errorf("FromShortHundredweight() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassShortHundredweight)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromShortHundredweight() with zero value = %v, want 0", converted)
    }
}
// Test FromLongHundredweight function
func TestMassFactory_FromLongHundredweight(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLongHundredweight(100)
    if err != nil {
        t.Errorf("FromLongHundredweight() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassLongHundredweight)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLongHundredweight() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLongHundredweight(math.NaN())
    if err == nil {
        t.Error("FromLongHundredweight() with NaN value should return error")
    }

    _, err = factory.FromLongHundredweight(math.Inf(1))
    if err == nil {
        t.Error("FromLongHundredweight() with +Inf value should return error")
    }

    _, err = factory.FromLongHundredweight(math.Inf(-1))
    if err == nil {
        t.Error("FromLongHundredweight() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLongHundredweight(0)
    if err != nil {
        t.Errorf("FromLongHundredweight() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassLongHundredweight)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLongHundredweight() with zero value = %v, want 0", converted)
    }
}
// Test FromGrains function
func TestMassFactory_FromGrains(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGrains(100)
    if err != nil {
        t.Errorf("FromGrains() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassGrain)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGrains() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGrains(math.NaN())
    if err == nil {
        t.Error("FromGrains() with NaN value should return error")
    }

    _, err = factory.FromGrains(math.Inf(1))
    if err == nil {
        t.Error("FromGrains() with +Inf value should return error")
    }

    _, err = factory.FromGrains(math.Inf(-1))
    if err == nil {
        t.Error("FromGrains() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGrains(0)
    if err != nil {
        t.Errorf("FromGrains() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassGrain)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGrains() with zero value = %v, want 0", converted)
    }
}
// Test FromSolarMasses function
func TestMassFactory_FromSolarMasses(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSolarMasses(100)
    if err != nil {
        t.Errorf("FromSolarMasses() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassSolarMass)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSolarMasses() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSolarMasses(math.NaN())
    if err == nil {
        t.Error("FromSolarMasses() with NaN value should return error")
    }

    _, err = factory.FromSolarMasses(math.Inf(1))
    if err == nil {
        t.Error("FromSolarMasses() with +Inf value should return error")
    }

    _, err = factory.FromSolarMasses(math.Inf(-1))
    if err == nil {
        t.Error("FromSolarMasses() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSolarMasses(0)
    if err != nil {
        t.Errorf("FromSolarMasses() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassSolarMass)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSolarMasses() with zero value = %v, want 0", converted)
    }
}
// Test FromEarthMasses function
func TestMassFactory_FromEarthMasses(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromEarthMasses(100)
    if err != nil {
        t.Errorf("FromEarthMasses() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassEarthMass)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromEarthMasses() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromEarthMasses(math.NaN())
    if err == nil {
        t.Error("FromEarthMasses() with NaN value should return error")
    }

    _, err = factory.FromEarthMasses(math.Inf(1))
    if err == nil {
        t.Error("FromEarthMasses() with +Inf value should return error")
    }

    _, err = factory.FromEarthMasses(math.Inf(-1))
    if err == nil {
        t.Error("FromEarthMasses() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromEarthMasses(0)
    if err != nil {
        t.Errorf("FromEarthMasses() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassEarthMass)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromEarthMasses() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtograms function
func TestMassFactory_FromFemtograms(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtograms(100)
    if err != nil {
        t.Errorf("FromFemtograms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFemtogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtograms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtograms(math.NaN())
    if err == nil {
        t.Error("FromFemtograms() with NaN value should return error")
    }

    _, err = factory.FromFemtograms(math.Inf(1))
    if err == nil {
        t.Error("FromFemtograms() with +Inf value should return error")
    }

    _, err = factory.FromFemtograms(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtograms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtograms(0)
    if err != nil {
        t.Errorf("FromFemtograms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFemtogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtograms() with zero value = %v, want 0", converted)
    }
}
// Test FromPicograms function
func TestMassFactory_FromPicograms(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicograms(100)
    if err != nil {
        t.Errorf("FromPicograms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassPicogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicograms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicograms(math.NaN())
    if err == nil {
        t.Error("FromPicograms() with NaN value should return error")
    }

    _, err = factory.FromPicograms(math.Inf(1))
    if err == nil {
        t.Error("FromPicograms() with +Inf value should return error")
    }

    _, err = factory.FromPicograms(math.Inf(-1))
    if err == nil {
        t.Error("FromPicograms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicograms(0)
    if err != nil {
        t.Errorf("FromPicograms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassPicogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicograms() with zero value = %v, want 0", converted)
    }
}
// Test FromNanograms function
func TestMassFactory_FromNanograms(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanograms(100)
    if err != nil {
        t.Errorf("FromNanograms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassNanogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanograms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanograms(math.NaN())
    if err == nil {
        t.Error("FromNanograms() with NaN value should return error")
    }

    _, err = factory.FromNanograms(math.Inf(1))
    if err == nil {
        t.Error("FromNanograms() with +Inf value should return error")
    }

    _, err = factory.FromNanograms(math.Inf(-1))
    if err == nil {
        t.Error("FromNanograms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanograms(0)
    if err != nil {
        t.Errorf("FromNanograms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassNanogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanograms() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrograms function
func TestMassFactory_FromMicrograms(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrograms(100)
    if err != nil {
        t.Errorf("FromMicrograms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMicrogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrograms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrograms(math.NaN())
    if err == nil {
        t.Error("FromMicrograms() with NaN value should return error")
    }

    _, err = factory.FromMicrograms(math.Inf(1))
    if err == nil {
        t.Error("FromMicrograms() with +Inf value should return error")
    }

    _, err = factory.FromMicrograms(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrograms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrograms(0)
    if err != nil {
        t.Errorf("FromMicrograms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMicrogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrograms() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligrams function
func TestMassFactory_FromMilligrams(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligrams(100)
    if err != nil {
        t.Errorf("FromMilligrams() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMilligram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligrams() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligrams(math.NaN())
    if err == nil {
        t.Error("FromMilligrams() with NaN value should return error")
    }

    _, err = factory.FromMilligrams(math.Inf(1))
    if err == nil {
        t.Error("FromMilligrams() with +Inf value should return error")
    }

    _, err = factory.FromMilligrams(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligrams() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligrams(0)
    if err != nil {
        t.Errorf("FromMilligrams() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMilligram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligrams() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigrams function
func TestMassFactory_FromCentigrams(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigrams(100)
    if err != nil {
        t.Errorf("FromCentigrams() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassCentigram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigrams() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigrams(math.NaN())
    if err == nil {
        t.Error("FromCentigrams() with NaN value should return error")
    }

    _, err = factory.FromCentigrams(math.Inf(1))
    if err == nil {
        t.Error("FromCentigrams() with +Inf value should return error")
    }

    _, err = factory.FromCentigrams(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigrams() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigrams(0)
    if err != nil {
        t.Errorf("FromCentigrams() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassCentigram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigrams() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigrams function
func TestMassFactory_FromDecigrams(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigrams(100)
    if err != nil {
        t.Errorf("FromDecigrams() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassDecigram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigrams() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigrams(math.NaN())
    if err == nil {
        t.Error("FromDecigrams() with NaN value should return error")
    }

    _, err = factory.FromDecigrams(math.Inf(1))
    if err == nil {
        t.Error("FromDecigrams() with +Inf value should return error")
    }

    _, err = factory.FromDecigrams(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigrams() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigrams(0)
    if err != nil {
        t.Errorf("FromDecigrams() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassDecigram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigrams() with zero value = %v, want 0", converted)
    }
}
// Test FromDecagrams function
func TestMassFactory_FromDecagrams(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecagrams(100)
    if err != nil {
        t.Errorf("FromDecagrams() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassDecagram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecagrams() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecagrams(math.NaN())
    if err == nil {
        t.Error("FromDecagrams() with NaN value should return error")
    }

    _, err = factory.FromDecagrams(math.Inf(1))
    if err == nil {
        t.Error("FromDecagrams() with +Inf value should return error")
    }

    _, err = factory.FromDecagrams(math.Inf(-1))
    if err == nil {
        t.Error("FromDecagrams() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecagrams(0)
    if err != nil {
        t.Errorf("FromDecagrams() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassDecagram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecagrams() with zero value = %v, want 0", converted)
    }
}
// Test FromHectograms function
func TestMassFactory_FromHectograms(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectograms(100)
    if err != nil {
        t.Errorf("FromHectograms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassHectogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectograms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectograms(math.NaN())
    if err == nil {
        t.Error("FromHectograms() with NaN value should return error")
    }

    _, err = factory.FromHectograms(math.Inf(1))
    if err == nil {
        t.Error("FromHectograms() with +Inf value should return error")
    }

    _, err = factory.FromHectograms(math.Inf(-1))
    if err == nil {
        t.Error("FromHectograms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectograms(0)
    if err != nil {
        t.Errorf("FromHectograms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassHectogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectograms() with zero value = %v, want 0", converted)
    }
}
// Test FromKilograms function
func TestMassFactory_FromKilograms(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilograms(100)
    if err != nil {
        t.Errorf("FromKilograms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilograms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilograms(math.NaN())
    if err == nil {
        t.Error("FromKilograms() with NaN value should return error")
    }

    _, err = factory.FromKilograms(math.Inf(1))
    if err == nil {
        t.Error("FromKilograms() with +Inf value should return error")
    }

    _, err = factory.FromKilograms(math.Inf(-1))
    if err == nil {
        t.Error("FromKilograms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilograms(0)
    if err != nil {
        t.Errorf("FromKilograms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilograms() with zero value = %v, want 0", converted)
    }
}
// Test FromKilotonnes function
func TestMassFactory_FromKilotonnes(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilotonnes(100)
    if err != nil {
        t.Errorf("FromKilotonnes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassKilotonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilotonnes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilotonnes(math.NaN())
    if err == nil {
        t.Error("FromKilotonnes() with NaN value should return error")
    }

    _, err = factory.FromKilotonnes(math.Inf(1))
    if err == nil {
        t.Error("FromKilotonnes() with +Inf value should return error")
    }

    _, err = factory.FromKilotonnes(math.Inf(-1))
    if err == nil {
        t.Error("FromKilotonnes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilotonnes(0)
    if err != nil {
        t.Errorf("FromKilotonnes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassKilotonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilotonnes() with zero value = %v, want 0", converted)
    }
}
// Test FromMegatonnes function
func TestMassFactory_FromMegatonnes(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegatonnes(100)
    if err != nil {
        t.Errorf("FromMegatonnes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMegatonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegatonnes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegatonnes(math.NaN())
    if err == nil {
        t.Error("FromMegatonnes() with NaN value should return error")
    }

    _, err = factory.FromMegatonnes(math.Inf(1))
    if err == nil {
        t.Error("FromMegatonnes() with +Inf value should return error")
    }

    _, err = factory.FromMegatonnes(math.Inf(-1))
    if err == nil {
        t.Error("FromMegatonnes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegatonnes(0)
    if err != nil {
        t.Errorf("FromMegatonnes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMegatonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegatonnes() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopounds function
func TestMassFactory_FromKilopounds(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopounds(100)
    if err != nil {
        t.Errorf("FromKilopounds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassKilopound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopounds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopounds(math.NaN())
    if err == nil {
        t.Error("FromKilopounds() with NaN value should return error")
    }

    _, err = factory.FromKilopounds(math.Inf(1))
    if err == nil {
        t.Error("FromKilopounds() with +Inf value should return error")
    }

    _, err = factory.FromKilopounds(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopounds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopounds(0)
    if err != nil {
        t.Errorf("FromKilopounds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassKilopound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopounds() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapounds function
func TestMassFactory_FromMegapounds(t *testing.T) {
    factory := units.MassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapounds(100)
    if err != nil {
        t.Errorf("FromMegapounds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMegapound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapounds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapounds(math.NaN())
    if err == nil {
        t.Error("FromMegapounds() with NaN value should return error")
    }

    _, err = factory.FromMegapounds(math.Inf(1))
    if err == nil {
        t.Error("FromMegapounds() with +Inf value should return error")
    }

    _, err = factory.FromMegapounds(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapounds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapounds(0)
    if err != nil {
        t.Errorf("FromMegapounds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMegapound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapounds() with zero value = %v, want 0", converted)
    }
}

func TestMassToString(t *testing.T) {
	factory := units.MassFactory{}
	a, err := factory.CreateMass(45, units.MassKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassKilogram, 2)
	expected := "45.00 " + units.GetMassAbbreviation(units.MassKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassKilogram, -1)
	expected = "45 " + units.GetMassAbbreviation(units.MassKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMass_EqualityAndComparison(t *testing.T) {
	factory := units.MassFactory{}
	a1, _ := factory.CreateMass(60, units.MassKilogram)
	a2, _ := factory.CreateMass(60, units.MassKilogram)
	a3, _ := factory.CreateMass(120, units.MassKilogram)

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

func TestMass_Arithmetic(t *testing.T) {
	factory := units.MassFactory{}
	a1, _ := factory.CreateMass(30, units.MassKilogram)
	a2, _ := factory.CreateMass(45, units.MassKilogram)

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


func TestGetMassAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MassUnits
        want string
    }{
        {
            name: "Gram abbreviation",
            unit: units.MassGram,
            want: "g",
        },
        {
            name: "Tonne abbreviation",
            unit: units.MassTonne,
            want: "t",
        },
        {
            name: "ShortTon abbreviation",
            unit: units.MassShortTon,
            want: "t (short)",
        },
        {
            name: "LongTon abbreviation",
            unit: units.MassLongTon,
            want: "long tn",
        },
        {
            name: "Pound abbreviation",
            unit: units.MassPound,
            want: "lb",
        },
        {
            name: "Ounce abbreviation",
            unit: units.MassOunce,
            want: "oz",
        },
        {
            name: "Slug abbreviation",
            unit: units.MassSlug,
            want: "slug",
        },
        {
            name: "Stone abbreviation",
            unit: units.MassStone,
            want: "st",
        },
        {
            name: "ShortHundredweight abbreviation",
            unit: units.MassShortHundredweight,
            want: "cwt",
        },
        {
            name: "LongHundredweight abbreviation",
            unit: units.MassLongHundredweight,
            want: "cwt",
        },
        {
            name: "Grain abbreviation",
            unit: units.MassGrain,
            want: "gr",
        },
        {
            name: "SolarMass abbreviation",
            unit: units.MassSolarMass,
            want: "M☉",
        },
        {
            name: "EarthMass abbreviation",
            unit: units.MassEarthMass,
            want: "em",
        },
        {
            name: "Femtogram abbreviation",
            unit: units.MassFemtogram,
            want: "fg",
        },
        {
            name: "Picogram abbreviation",
            unit: units.MassPicogram,
            want: "pg",
        },
        {
            name: "Nanogram abbreviation",
            unit: units.MassNanogram,
            want: "ng",
        },
        {
            name: "Microgram abbreviation",
            unit: units.MassMicrogram,
            want: "μg",
        },
        {
            name: "Milligram abbreviation",
            unit: units.MassMilligram,
            want: "mg",
        },
        {
            name: "Centigram abbreviation",
            unit: units.MassCentigram,
            want: "cg",
        },
        {
            name: "Decigram abbreviation",
            unit: units.MassDecigram,
            want: "dg",
        },
        {
            name: "Decagram abbreviation",
            unit: units.MassDecagram,
            want: "dag",
        },
        {
            name: "Hectogram abbreviation",
            unit: units.MassHectogram,
            want: "hg",
        },
        {
            name: "Kilogram abbreviation",
            unit: units.MassKilogram,
            want: "kg",
        },
        {
            name: "Kilotonne abbreviation",
            unit: units.MassKilotonne,
            want: "kt",
        },
        {
            name: "Megatonne abbreviation",
            unit: units.MassMegatonne,
            want: "Mt",
        },
        {
            name: "Kilopound abbreviation",
            unit: units.MassKilopound,
            want: "klb",
        },
        {
            name: "Megapound abbreviation",
            unit: units.MassMegapound,
            want: "Mlb",
        },
        {
            name: "invalid unit",
            unit: units.MassUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMassAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMassAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMass_String(t *testing.T) {
    factory := units.MassFactory{}
    
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
            unit, err := factory.CreateMass(tt.value, units.MassKilogram)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Mass.String() = %v, want %v", got, tt.want)
            }
        })
    }
}