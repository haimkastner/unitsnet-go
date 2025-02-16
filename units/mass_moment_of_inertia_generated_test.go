// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassMomentOfInertiaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramSquareMeter"}`
	
	factory := units.MassMomentOfInertiaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassMomentOfInertiaKilogramSquareMeter {
		t.Errorf("expected unit %v, got %v", units.MassMomentOfInertiaKilogramSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassMomentOfInertiaDto_ToJSON(t *testing.T) {
	dto := units.MassMomentOfInertiaDto{
		Value: 45,
		Unit:  units.MassMomentOfInertiaKilogramSquareMeter,
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
	if result["unit"].(string) != string(units.MassMomentOfInertiaKilogramSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.MassMomentOfInertiaKilogramSquareMeter, result["unit"])
	}
}

func TestNewMassMomentOfInertia_InvalidValue(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMassMomentOfInertia(math.NaN(), units.MassMomentOfInertiaKilogramSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMassMomentOfInertia(math.Inf(1), units.MassMomentOfInertiaKilogramSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassMomentOfInertiaConversions(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMassMomentOfInertia(180, units.MassMomentOfInertiaKilogramSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramSquareMeters.
		// No expected conversion value provided for GramSquareMeters, verifying result is not NaN.
		result := a.GramSquareMeters()
		cacheResult := a.GramSquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to GramSquareDecimeters.
		// No expected conversion value provided for GramSquareDecimeters, verifying result is not NaN.
		result := a.GramSquareDecimeters()
		cacheResult := a.GramSquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to GramSquareCentimeters.
		// No expected conversion value provided for GramSquareCentimeters, verifying result is not NaN.
		result := a.GramSquareCentimeters()
		cacheResult := a.GramSquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to GramSquareMillimeters.
		// No expected conversion value provided for GramSquareMillimeters, verifying result is not NaN.
		result := a.GramSquareMillimeters()
		cacheResult := a.GramSquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramSquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneSquareMeters.
		// No expected conversion value provided for TonneSquareMeters, verifying result is not NaN.
		result := a.TonneSquareMeters()
		cacheResult := a.TonneSquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonneSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to TonneSquareDecimeters.
		// No expected conversion value provided for TonneSquareDecimeters, verifying result is not NaN.
		result := a.TonneSquareDecimeters()
		cacheResult := a.TonneSquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonneSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneSquareCentimeters.
		// No expected conversion value provided for TonneSquareCentimeters, verifying result is not NaN.
		result := a.TonneSquareCentimeters()
		cacheResult := a.TonneSquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonneSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneSquareMilimeters.
		// No expected conversion value provided for TonneSquareMilimeters, verifying result is not NaN.
		result := a.TonneSquareMilimeters()
		cacheResult := a.TonneSquareMilimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonneSquareMilimeters returned NaN")
		}
	}
	{
		// Test conversion to PoundSquareFeet.
		// No expected conversion value provided for PoundSquareFeet, verifying result is not NaN.
		result := a.PoundSquareFeet()
		cacheResult := a.PoundSquareFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundSquareFeet returned NaN")
		}
	}
	{
		// Test conversion to PoundSquareInches.
		// No expected conversion value provided for PoundSquareInches, verifying result is not NaN.
		result := a.PoundSquareInches()
		cacheResult := a.PoundSquareInches()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundSquareInches returned NaN")
		}
	}
	{
		// Test conversion to SlugSquareFeet.
		// No expected conversion value provided for SlugSquareFeet, verifying result is not NaN.
		result := a.SlugSquareFeet()
		cacheResult := a.SlugSquareFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SlugSquareFeet returned NaN")
		}
	}
	{
		// Test conversion to SlugSquareInches.
		// No expected conversion value provided for SlugSquareInches, verifying result is not NaN.
		result := a.SlugSquareInches()
		cacheResult := a.SlugSquareInches()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SlugSquareInches returned NaN")
		}
	}
	{
		// Test conversion to MilligramSquareMeters.
		// No expected conversion value provided for MilligramSquareMeters, verifying result is not NaN.
		result := a.MilligramSquareMeters()
		cacheResult := a.MilligramSquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramSquareMeters.
		// No expected conversion value provided for KilogramSquareMeters, verifying result is not NaN.
		result := a.KilogramSquareMeters()
		cacheResult := a.KilogramSquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to MilligramSquareDecimeters.
		// No expected conversion value provided for MilligramSquareDecimeters, verifying result is not NaN.
		result := a.MilligramSquareDecimeters()
		cacheResult := a.MilligramSquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramSquareDecimeters.
		// No expected conversion value provided for KilogramSquareDecimeters, verifying result is not NaN.
		result := a.KilogramSquareDecimeters()
		cacheResult := a.KilogramSquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligramSquareCentimeters.
		// No expected conversion value provided for MilligramSquareCentimeters, verifying result is not NaN.
		result := a.MilligramSquareCentimeters()
		cacheResult := a.MilligramSquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramSquareCentimeters.
		// No expected conversion value provided for KilogramSquareCentimeters, verifying result is not NaN.
		result := a.KilogramSquareCentimeters()
		cacheResult := a.KilogramSquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligramSquareMillimeters.
		// No expected conversion value provided for MilligramSquareMillimeters, verifying result is not NaN.
		result := a.MilligramSquareMillimeters()
		cacheResult := a.MilligramSquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramSquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramSquareMillimeters.
		// No expected conversion value provided for KilogramSquareMillimeters, verifying result is not NaN.
		result := a.KilogramSquareMillimeters()
		cacheResult := a.KilogramSquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramSquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to KilotonneSquareMeters.
		// No expected conversion value provided for KilotonneSquareMeters, verifying result is not NaN.
		result := a.KilotonneSquareMeters()
		cacheResult := a.KilotonneSquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilotonneSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to MegatonneSquareMeters.
		// No expected conversion value provided for MegatonneSquareMeters, verifying result is not NaN.
		result := a.MegatonneSquareMeters()
		cacheResult := a.MegatonneSquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegatonneSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to KilotonneSquareDecimeters.
		// No expected conversion value provided for KilotonneSquareDecimeters, verifying result is not NaN.
		result := a.KilotonneSquareDecimeters()
		cacheResult := a.KilotonneSquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilotonneSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to MegatonneSquareDecimeters.
		// No expected conversion value provided for MegatonneSquareDecimeters, verifying result is not NaN.
		result := a.MegatonneSquareDecimeters()
		cacheResult := a.MegatonneSquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegatonneSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to KilotonneSquareCentimeters.
		// No expected conversion value provided for KilotonneSquareCentimeters, verifying result is not NaN.
		result := a.KilotonneSquareCentimeters()
		cacheResult := a.KilotonneSquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilotonneSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MegatonneSquareCentimeters.
		// No expected conversion value provided for MegatonneSquareCentimeters, verifying result is not NaN.
		result := a.MegatonneSquareCentimeters()
		cacheResult := a.MegatonneSquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegatonneSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to KilotonneSquareMilimeters.
		// No expected conversion value provided for KilotonneSquareMilimeters, verifying result is not NaN.
		result := a.KilotonneSquareMilimeters()
		cacheResult := a.KilotonneSquareMilimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilotonneSquareMilimeters returned NaN")
		}
	}
	{
		// Test conversion to MegatonneSquareMilimeters.
		// No expected conversion value provided for MegatonneSquareMilimeters, verifying result is not NaN.
		result := a.MegatonneSquareMilimeters()
		cacheResult := a.MegatonneSquareMilimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegatonneSquareMilimeters returned NaN")
		}
	}
}

func TestMassMomentOfInertia_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	a, err := factory.CreateMassMomentOfInertia(90, units.MassMomentOfInertiaKilogramSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassMomentOfInertiaKilogramSquareMeter {
		t.Errorf("expected default unit KilogramSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassMomentOfInertiaGramSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassMomentOfInertiaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassMomentOfInertiaKilogramSquareMeter {
		t.Errorf("expected unit KilogramSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassMomentOfInertiaFactory_FromDto(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilogramSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MassMomentOfInertiaDto{
        Value: math.NaN(),
        Unit:  units.MassMomentOfInertiaKilogramSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GramSquareMeter conversion
    gram_square_metersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaGramSquareMeter,
    }
    
    var gram_square_metersResult *units.MassMomentOfInertia
    gram_square_metersResult, err = factory.FromDto(gram_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with GramSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_square_metersResult.Convert(units.MassMomentOfInertiaGramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramSquareMeter = %v, want %v", converted, 100)
    }
    // Test GramSquareDecimeter conversion
    gram_square_decimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaGramSquareDecimeter,
    }
    
    var gram_square_decimetersResult *units.MassMomentOfInertia
    gram_square_decimetersResult, err = factory.FromDto(gram_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with GramSquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_square_decimetersResult.Convert(units.MassMomentOfInertiaGramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test GramSquareCentimeter conversion
    gram_square_centimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaGramSquareCentimeter,
    }
    
    var gram_square_centimetersResult *units.MassMomentOfInertia
    gram_square_centimetersResult, err = factory.FromDto(gram_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with GramSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_square_centimetersResult.Convert(units.MassMomentOfInertiaGramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test GramSquareMillimeter conversion
    gram_square_millimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaGramSquareMillimeter,
    }
    
    var gram_square_millimetersResult *units.MassMomentOfInertia
    gram_square_millimetersResult, err = factory.FromDto(gram_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with GramSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_square_millimetersResult.Convert(units.MassMomentOfInertiaGramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test TonneSquareMeter conversion
    tonne_square_metersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaTonneSquareMeter,
    }
    
    var tonne_square_metersResult *units.MassMomentOfInertia
    tonne_square_metersResult, err = factory.FromDto(tonne_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with TonneSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_square_metersResult.Convert(units.MassMomentOfInertiaTonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneSquareMeter = %v, want %v", converted, 100)
    }
    // Test TonneSquareDecimeter conversion
    tonne_square_decimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaTonneSquareDecimeter,
    }
    
    var tonne_square_decimetersResult *units.MassMomentOfInertia
    tonne_square_decimetersResult, err = factory.FromDto(tonne_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with TonneSquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_square_decimetersResult.Convert(units.MassMomentOfInertiaTonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test TonneSquareCentimeter conversion
    tonne_square_centimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaTonneSquareCentimeter,
    }
    
    var tonne_square_centimetersResult *units.MassMomentOfInertia
    tonne_square_centimetersResult, err = factory.FromDto(tonne_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with TonneSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_square_centimetersResult.Convert(units.MassMomentOfInertiaTonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test TonneSquareMilimeter conversion
    tonne_square_milimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaTonneSquareMilimeter,
    }
    
    var tonne_square_milimetersResult *units.MassMomentOfInertia
    tonne_square_milimetersResult, err = factory.FromDto(tonne_square_milimetersDto)
    if err != nil {
        t.Errorf("FromDto() with TonneSquareMilimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_square_milimetersResult.Convert(units.MassMomentOfInertiaTonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneSquareMilimeter = %v, want %v", converted, 100)
    }
    // Test PoundSquareFoot conversion
    pound_square_feetDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaPoundSquareFoot,
    }
    
    var pound_square_feetResult *units.MassMomentOfInertia
    pound_square_feetResult, err = factory.FromDto(pound_square_feetDto)
    if err != nil {
        t.Errorf("FromDto() with PoundSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_square_feetResult.Convert(units.MassMomentOfInertiaPoundSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundSquareFoot = %v, want %v", converted, 100)
    }
    // Test PoundSquareInch conversion
    pound_square_inchesDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaPoundSquareInch,
    }
    
    var pound_square_inchesResult *units.MassMomentOfInertia
    pound_square_inchesResult, err = factory.FromDto(pound_square_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with PoundSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_square_inchesResult.Convert(units.MassMomentOfInertiaPoundSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundSquareInch = %v, want %v", converted, 100)
    }
    // Test SlugSquareFoot conversion
    slug_square_feetDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaSlugSquareFoot,
    }
    
    var slug_square_feetResult *units.MassMomentOfInertia
    slug_square_feetResult, err = factory.FromDto(slug_square_feetDto)
    if err != nil {
        t.Errorf("FromDto() with SlugSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slug_square_feetResult.Convert(units.MassMomentOfInertiaSlugSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugSquareFoot = %v, want %v", converted, 100)
    }
    // Test SlugSquareInch conversion
    slug_square_inchesDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaSlugSquareInch,
    }
    
    var slug_square_inchesResult *units.MassMomentOfInertia
    slug_square_inchesResult, err = factory.FromDto(slug_square_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with SlugSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slug_square_inchesResult.Convert(units.MassMomentOfInertiaSlugSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugSquareInch = %v, want %v", converted, 100)
    }
    // Test MilligramSquareMeter conversion
    milligram_square_metersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaMilligramSquareMeter,
    }
    
    var milligram_square_metersResult *units.MassMomentOfInertia
    milligram_square_metersResult, err = factory.FromDto(milligram_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligram_square_metersResult.Convert(units.MassMomentOfInertiaMilligramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilogramSquareMeter conversion
    kilogram_square_metersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilogramSquareMeter,
    }
    
    var kilogram_square_metersResult *units.MassMomentOfInertia
    kilogram_square_metersResult, err = factory.FromDto(kilogram_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_square_metersResult.Convert(units.MassMomentOfInertiaKilogramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramSquareMeter = %v, want %v", converted, 100)
    }
    // Test MilligramSquareDecimeter conversion
    milligram_square_decimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaMilligramSquareDecimeter,
    }
    
    var milligram_square_decimetersResult *units.MassMomentOfInertia
    milligram_square_decimetersResult, err = factory.FromDto(milligram_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramSquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligram_square_decimetersResult.Convert(units.MassMomentOfInertiaMilligramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test KilogramSquareDecimeter conversion
    kilogram_square_decimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilogramSquareDecimeter,
    }
    
    var kilogram_square_decimetersResult *units.MassMomentOfInertia
    kilogram_square_decimetersResult, err = factory.FromDto(kilogram_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramSquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_square_decimetersResult.Convert(units.MassMomentOfInertiaKilogramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test MilligramSquareCentimeter conversion
    milligram_square_centimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaMilligramSquareCentimeter,
    }
    
    var milligram_square_centimetersResult *units.MassMomentOfInertia
    milligram_square_centimetersResult, err = factory.FromDto(milligram_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligram_square_centimetersResult.Convert(units.MassMomentOfInertiaMilligramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramSquareCentimeter conversion
    kilogram_square_centimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilogramSquareCentimeter,
    }
    
    var kilogram_square_centimetersResult *units.MassMomentOfInertia
    kilogram_square_centimetersResult, err = factory.FromDto(kilogram_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_square_centimetersResult.Convert(units.MassMomentOfInertiaKilogramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test MilligramSquareMillimeter conversion
    milligram_square_millimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaMilligramSquareMillimeter,
    }
    
    var milligram_square_millimetersResult *units.MassMomentOfInertia
    milligram_square_millimetersResult, err = factory.FromDto(milligram_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligram_square_millimetersResult.Convert(units.MassMomentOfInertiaMilligramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramSquareMillimeter conversion
    kilogram_square_millimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilogramSquareMillimeter,
    }
    
    var kilogram_square_millimetersResult *units.MassMomentOfInertia
    kilogram_square_millimetersResult, err = factory.FromDto(kilogram_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_square_millimetersResult.Convert(units.MassMomentOfInertiaKilogramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test KilotonneSquareMeter conversion
    kilotonne_square_metersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilotonneSquareMeter,
    }
    
    var kilotonne_square_metersResult *units.MassMomentOfInertia
    kilotonne_square_metersResult, err = factory.FromDto(kilotonne_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with KilotonneSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonne_square_metersResult.Convert(units.MassMomentOfInertiaKilotonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilotonneSquareMeter = %v, want %v", converted, 100)
    }
    // Test MegatonneSquareMeter conversion
    megatonne_square_metersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaMegatonneSquareMeter,
    }
    
    var megatonne_square_metersResult *units.MassMomentOfInertia
    megatonne_square_metersResult, err = factory.FromDto(megatonne_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with MegatonneSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonne_square_metersResult.Convert(units.MassMomentOfInertiaMegatonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegatonneSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilotonneSquareDecimeter conversion
    kilotonne_square_decimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilotonneSquareDecimeter,
    }
    
    var kilotonne_square_decimetersResult *units.MassMomentOfInertia
    kilotonne_square_decimetersResult, err = factory.FromDto(kilotonne_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilotonneSquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonne_square_decimetersResult.Convert(units.MassMomentOfInertiaKilotonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilotonneSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test MegatonneSquareDecimeter conversion
    megatonne_square_decimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaMegatonneSquareDecimeter,
    }
    
    var megatonne_square_decimetersResult *units.MassMomentOfInertia
    megatonne_square_decimetersResult, err = factory.FromDto(megatonne_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MegatonneSquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonne_square_decimetersResult.Convert(units.MassMomentOfInertiaMegatonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegatonneSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test KilotonneSquareCentimeter conversion
    kilotonne_square_centimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilotonneSquareCentimeter,
    }
    
    var kilotonne_square_centimetersResult *units.MassMomentOfInertia
    kilotonne_square_centimetersResult, err = factory.FromDto(kilotonne_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilotonneSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonne_square_centimetersResult.Convert(units.MassMomentOfInertiaKilotonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilotonneSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test MegatonneSquareCentimeter conversion
    megatonne_square_centimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaMegatonneSquareCentimeter,
    }
    
    var megatonne_square_centimetersResult *units.MassMomentOfInertia
    megatonne_square_centimetersResult, err = factory.FromDto(megatonne_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MegatonneSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonne_square_centimetersResult.Convert(units.MassMomentOfInertiaMegatonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegatonneSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilotonneSquareMilimeter conversion
    kilotonne_square_milimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaKilotonneSquareMilimeter,
    }
    
    var kilotonne_square_milimetersResult *units.MassMomentOfInertia
    kilotonne_square_milimetersResult, err = factory.FromDto(kilotonne_square_milimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilotonneSquareMilimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonne_square_milimetersResult.Convert(units.MassMomentOfInertiaKilotonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilotonneSquareMilimeter = %v, want %v", converted, 100)
    }
    // Test MegatonneSquareMilimeter conversion
    megatonne_square_milimetersDto := units.MassMomentOfInertiaDto{
        Value: 100,
        Unit:  units.MassMomentOfInertiaMegatonneSquareMilimeter,
    }
    
    var megatonne_square_milimetersResult *units.MassMomentOfInertia
    megatonne_square_milimetersResult, err = factory.FromDto(megatonne_square_milimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MegatonneSquareMilimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonne_square_milimetersResult.Convert(units.MassMomentOfInertiaMegatonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegatonneSquareMilimeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MassMomentOfInertiaDto{
        Value: 0,
        Unit:  units.MassMomentOfInertiaKilogramSquareMeter,
    }
    
    var zeroResult *units.MassMomentOfInertia
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMassMomentOfInertiaFactory_FromDtoJSON(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KilogramSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KilogramSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.MassMomentOfInertiaDto{
        Value: nanValue,
        Unit:  units.MassMomentOfInertiaKilogramSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with GramSquareMeter unit
    gram_square_metersJSON := []byte(`{"value": 100, "unit": "GramSquareMeter"}`)
    gram_square_metersResult, err := factory.FromDtoJSON(gram_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_square_metersResult.Convert(units.MassMomentOfInertiaGramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramSquareDecimeter unit
    gram_square_decimetersJSON := []byte(`{"value": 100, "unit": "GramSquareDecimeter"}`)
    gram_square_decimetersResult, err := factory.FromDtoJSON(gram_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramSquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_square_decimetersResult.Convert(units.MassMomentOfInertiaGramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramSquareCentimeter unit
    gram_square_centimetersJSON := []byte(`{"value": 100, "unit": "GramSquareCentimeter"}`)
    gram_square_centimetersResult, err := factory.FromDtoJSON(gram_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_square_centimetersResult.Convert(units.MassMomentOfInertiaGramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramSquareMillimeter unit
    gram_square_millimetersJSON := []byte(`{"value": 100, "unit": "GramSquareMillimeter"}`)
    gram_square_millimetersResult, err := factory.FromDtoJSON(gram_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_square_millimetersResult.Convert(units.MassMomentOfInertiaGramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneSquareMeter unit
    tonne_square_metersJSON := []byte(`{"value": 100, "unit": "TonneSquareMeter"}`)
    tonne_square_metersResult, err := factory.FromDtoJSON(tonne_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_square_metersResult.Convert(units.MassMomentOfInertiaTonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneSquareDecimeter unit
    tonne_square_decimetersJSON := []byte(`{"value": 100, "unit": "TonneSquareDecimeter"}`)
    tonne_square_decimetersResult, err := factory.FromDtoJSON(tonne_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneSquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_square_decimetersResult.Convert(units.MassMomentOfInertiaTonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneSquareCentimeter unit
    tonne_square_centimetersJSON := []byte(`{"value": 100, "unit": "TonneSquareCentimeter"}`)
    tonne_square_centimetersResult, err := factory.FromDtoJSON(tonne_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_square_centimetersResult.Convert(units.MassMomentOfInertiaTonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneSquareMilimeter unit
    tonne_square_milimetersJSON := []byte(`{"value": 100, "unit": "TonneSquareMilimeter"}`)
    tonne_square_milimetersResult, err := factory.FromDtoJSON(tonne_square_milimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneSquareMilimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_square_milimetersResult.Convert(units.MassMomentOfInertiaTonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneSquareMilimeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundSquareFoot unit
    pound_square_feetJSON := []byte(`{"value": 100, "unit": "PoundSquareFoot"}`)
    pound_square_feetResult, err := factory.FromDtoJSON(pound_square_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_square_feetResult.Convert(units.MassMomentOfInertiaPoundSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with PoundSquareInch unit
    pound_square_inchesJSON := []byte(`{"value": 100, "unit": "PoundSquareInch"}`)
    pound_square_inchesResult, err := factory.FromDtoJSON(pound_square_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_square_inchesResult.Convert(units.MassMomentOfInertiaPoundSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with SlugSquareFoot unit
    slug_square_feetJSON := []byte(`{"value": 100, "unit": "SlugSquareFoot"}`)
    slug_square_feetResult, err := factory.FromDtoJSON(slug_square_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slug_square_feetResult.Convert(units.MassMomentOfInertiaSlugSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with SlugSquareInch unit
    slug_square_inchesJSON := []byte(`{"value": 100, "unit": "SlugSquareInch"}`)
    slug_square_inchesResult, err := factory.FromDtoJSON(slug_square_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slug_square_inchesResult.Convert(units.MassMomentOfInertiaSlugSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramSquareMeter unit
    milligram_square_metersJSON := []byte(`{"value": 100, "unit": "MilligramSquareMeter"}`)
    milligram_square_metersResult, err := factory.FromDtoJSON(milligram_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligram_square_metersResult.Convert(units.MassMomentOfInertiaMilligramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramSquareMeter unit
    kilogram_square_metersJSON := []byte(`{"value": 100, "unit": "KilogramSquareMeter"}`)
    kilogram_square_metersResult, err := factory.FromDtoJSON(kilogram_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_square_metersResult.Convert(units.MassMomentOfInertiaKilogramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramSquareDecimeter unit
    milligram_square_decimetersJSON := []byte(`{"value": 100, "unit": "MilligramSquareDecimeter"}`)
    milligram_square_decimetersResult, err := factory.FromDtoJSON(milligram_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramSquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligram_square_decimetersResult.Convert(units.MassMomentOfInertiaMilligramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramSquareDecimeter unit
    kilogram_square_decimetersJSON := []byte(`{"value": 100, "unit": "KilogramSquareDecimeter"}`)
    kilogram_square_decimetersResult, err := factory.FromDtoJSON(kilogram_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramSquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_square_decimetersResult.Convert(units.MassMomentOfInertiaKilogramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramSquareCentimeter unit
    milligram_square_centimetersJSON := []byte(`{"value": 100, "unit": "MilligramSquareCentimeter"}`)
    milligram_square_centimetersResult, err := factory.FromDtoJSON(milligram_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligram_square_centimetersResult.Convert(units.MassMomentOfInertiaMilligramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramSquareCentimeter unit
    kilogram_square_centimetersJSON := []byte(`{"value": 100, "unit": "KilogramSquareCentimeter"}`)
    kilogram_square_centimetersResult, err := factory.FromDtoJSON(kilogram_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_square_centimetersResult.Convert(units.MassMomentOfInertiaKilogramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramSquareMillimeter unit
    milligram_square_millimetersJSON := []byte(`{"value": 100, "unit": "MilligramSquareMillimeter"}`)
    milligram_square_millimetersResult, err := factory.FromDtoJSON(milligram_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligram_square_millimetersResult.Convert(units.MassMomentOfInertiaMilligramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramSquareMillimeter unit
    kilogram_square_millimetersJSON := []byte(`{"value": 100, "unit": "KilogramSquareMillimeter"}`)
    kilogram_square_millimetersResult, err := factory.FromDtoJSON(kilogram_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_square_millimetersResult.Convert(units.MassMomentOfInertiaKilogramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilotonneSquareMeter unit
    kilotonne_square_metersJSON := []byte(`{"value": 100, "unit": "KilotonneSquareMeter"}`)
    kilotonne_square_metersResult, err := factory.FromDtoJSON(kilotonne_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilotonneSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonne_square_metersResult.Convert(units.MassMomentOfInertiaKilotonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilotonneSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegatonneSquareMeter unit
    megatonne_square_metersJSON := []byte(`{"value": 100, "unit": "MegatonneSquareMeter"}`)
    megatonne_square_metersResult, err := factory.FromDtoJSON(megatonne_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegatonneSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonne_square_metersResult.Convert(units.MassMomentOfInertiaMegatonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegatonneSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilotonneSquareDecimeter unit
    kilotonne_square_decimetersJSON := []byte(`{"value": 100, "unit": "KilotonneSquareDecimeter"}`)
    kilotonne_square_decimetersResult, err := factory.FromDtoJSON(kilotonne_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilotonneSquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonne_square_decimetersResult.Convert(units.MassMomentOfInertiaKilotonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilotonneSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegatonneSquareDecimeter unit
    megatonne_square_decimetersJSON := []byte(`{"value": 100, "unit": "MegatonneSquareDecimeter"}`)
    megatonne_square_decimetersResult, err := factory.FromDtoJSON(megatonne_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegatonneSquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonne_square_decimetersResult.Convert(units.MassMomentOfInertiaMegatonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegatonneSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilotonneSquareCentimeter unit
    kilotonne_square_centimetersJSON := []byte(`{"value": 100, "unit": "KilotonneSquareCentimeter"}`)
    kilotonne_square_centimetersResult, err := factory.FromDtoJSON(kilotonne_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilotonneSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonne_square_centimetersResult.Convert(units.MassMomentOfInertiaKilotonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilotonneSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegatonneSquareCentimeter unit
    megatonne_square_centimetersJSON := []byte(`{"value": 100, "unit": "MegatonneSquareCentimeter"}`)
    megatonne_square_centimetersResult, err := factory.FromDtoJSON(megatonne_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegatonneSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonne_square_centimetersResult.Convert(units.MassMomentOfInertiaMegatonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegatonneSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilotonneSquareMilimeter unit
    kilotonne_square_milimetersJSON := []byte(`{"value": 100, "unit": "KilotonneSquareMilimeter"}`)
    kilotonne_square_milimetersResult, err := factory.FromDtoJSON(kilotonne_square_milimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilotonneSquareMilimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilotonne_square_milimetersResult.Convert(units.MassMomentOfInertiaKilotonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilotonneSquareMilimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegatonneSquareMilimeter unit
    megatonne_square_milimetersJSON := []byte(`{"value": 100, "unit": "MegatonneSquareMilimeter"}`)
    megatonne_square_milimetersResult, err := factory.FromDtoJSON(megatonne_square_milimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegatonneSquareMilimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megatonne_square_milimetersResult.Convert(units.MassMomentOfInertiaMegatonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegatonneSquareMilimeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KilogramSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGramSquareMeters function
func TestMassMomentOfInertiaFactory_FromGramSquareMeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramSquareMeters(100)
    if err != nil {
        t.Errorf("FromGramSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaGramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromGramSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromGramSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromGramSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromGramSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGramSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramSquareMeters(0)
    if err != nil {
        t.Errorf("FromGramSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaGramSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGramSquareDecimeters function
func TestMassMomentOfInertiaFactory_FromGramSquareDecimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramSquareDecimeters(100)
    if err != nil {
        t.Errorf("FromGramSquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaGramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramSquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramSquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromGramSquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromGramSquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromGramSquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromGramSquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGramSquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramSquareDecimeters(0)
    if err != nil {
        t.Errorf("FromGramSquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaGramSquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramSquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGramSquareCentimeters function
func TestMassMomentOfInertiaFactory_FromGramSquareCentimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromGramSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaGramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromGramSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromGramSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromGramSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromGramSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGramSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromGramSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaGramSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGramSquareMillimeters function
func TestMassMomentOfInertiaFactory_FromGramSquareMillimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramSquareMillimeters(100)
    if err != nil {
        t.Errorf("FromGramSquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaGramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramSquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramSquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromGramSquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromGramSquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromGramSquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromGramSquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGramSquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramSquareMillimeters(0)
    if err != nil {
        t.Errorf("FromGramSquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaGramSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramSquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneSquareMeters function
func TestMassMomentOfInertiaFactory_FromTonneSquareMeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneSquareMeters(100)
    if err != nil {
        t.Errorf("FromTonneSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaTonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromTonneSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromTonneSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromTonneSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromTonneSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneSquareMeters(0)
    if err != nil {
        t.Errorf("FromTonneSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaTonneSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneSquareDecimeters function
func TestMassMomentOfInertiaFactory_FromTonneSquareDecimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneSquareDecimeters(100)
    if err != nil {
        t.Errorf("FromTonneSquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaTonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneSquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneSquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromTonneSquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromTonneSquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromTonneSquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromTonneSquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneSquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneSquareDecimeters(0)
    if err != nil {
        t.Errorf("FromTonneSquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaTonneSquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneSquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneSquareCentimeters function
func TestMassMomentOfInertiaFactory_FromTonneSquareCentimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromTonneSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaTonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromTonneSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromTonneSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromTonneSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromTonneSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromTonneSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaTonneSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneSquareMilimeters function
func TestMassMomentOfInertiaFactory_FromTonneSquareMilimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneSquareMilimeters(100)
    if err != nil {
        t.Errorf("FromTonneSquareMilimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaTonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneSquareMilimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneSquareMilimeters(math.NaN())
    if err == nil {
        t.Error("FromTonneSquareMilimeters() with NaN value should return error")
    }

    _, err = factory.FromTonneSquareMilimeters(math.Inf(1))
    if err == nil {
        t.Error("FromTonneSquareMilimeters() with +Inf value should return error")
    }

    _, err = factory.FromTonneSquareMilimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneSquareMilimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneSquareMilimeters(0)
    if err != nil {
        t.Errorf("FromTonneSquareMilimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaTonneSquareMilimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneSquareMilimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundSquareFeet function
func TestMassMomentOfInertiaFactory_FromPoundSquareFeet(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundSquareFeet(100)
    if err != nil {
        t.Errorf("FromPoundSquareFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaPoundSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundSquareFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundSquareFeet(math.NaN())
    if err == nil {
        t.Error("FromPoundSquareFeet() with NaN value should return error")
    }

    _, err = factory.FromPoundSquareFeet(math.Inf(1))
    if err == nil {
        t.Error("FromPoundSquareFeet() with +Inf value should return error")
    }

    _, err = factory.FromPoundSquareFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundSquareFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundSquareFeet(0)
    if err != nil {
        t.Errorf("FromPoundSquareFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaPoundSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundSquareFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundSquareInches function
func TestMassMomentOfInertiaFactory_FromPoundSquareInches(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundSquareInches(100)
    if err != nil {
        t.Errorf("FromPoundSquareInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaPoundSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundSquareInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundSquareInches(math.NaN())
    if err == nil {
        t.Error("FromPoundSquareInches() with NaN value should return error")
    }

    _, err = factory.FromPoundSquareInches(math.Inf(1))
    if err == nil {
        t.Error("FromPoundSquareInches() with +Inf value should return error")
    }

    _, err = factory.FromPoundSquareInches(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundSquareInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundSquareInches(0)
    if err != nil {
        t.Errorf("FromPoundSquareInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaPoundSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundSquareInches() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugSquareFeet function
func TestMassMomentOfInertiaFactory_FromSlugSquareFeet(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugSquareFeet(100)
    if err != nil {
        t.Errorf("FromSlugSquareFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaSlugSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugSquareFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugSquareFeet(math.NaN())
    if err == nil {
        t.Error("FromSlugSquareFeet() with NaN value should return error")
    }

    _, err = factory.FromSlugSquareFeet(math.Inf(1))
    if err == nil {
        t.Error("FromSlugSquareFeet() with +Inf value should return error")
    }

    _, err = factory.FromSlugSquareFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugSquareFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugSquareFeet(0)
    if err != nil {
        t.Errorf("FromSlugSquareFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaSlugSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugSquareFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugSquareInches function
func TestMassMomentOfInertiaFactory_FromSlugSquareInches(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugSquareInches(100)
    if err != nil {
        t.Errorf("FromSlugSquareInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaSlugSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugSquareInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugSquareInches(math.NaN())
    if err == nil {
        t.Error("FromSlugSquareInches() with NaN value should return error")
    }

    _, err = factory.FromSlugSquareInches(math.Inf(1))
    if err == nil {
        t.Error("FromSlugSquareInches() with +Inf value should return error")
    }

    _, err = factory.FromSlugSquareInches(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugSquareInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugSquareInches(0)
    if err != nil {
        t.Errorf("FromSlugSquareInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaSlugSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugSquareInches() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramSquareMeters function
func TestMassMomentOfInertiaFactory_FromMilligramSquareMeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramSquareMeters(100)
    if err != nil {
        t.Errorf("FromMilligramSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaMilligramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromMilligramSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromMilligramSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromMilligramSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramSquareMeters(0)
    if err != nil {
        t.Errorf("FromMilligramSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaMilligramSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramSquareMeters function
func TestMassMomentOfInertiaFactory_FromKilogramSquareMeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramSquareMeters(100)
    if err != nil {
        t.Errorf("FromKilogramSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaKilogramSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromKilogramSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromKilogramSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromKilogramSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramSquareMeters(0)
    if err != nil {
        t.Errorf("FromKilogramSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaKilogramSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramSquareDecimeters function
func TestMassMomentOfInertiaFactory_FromMilligramSquareDecimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramSquareDecimeters(100)
    if err != nil {
        t.Errorf("FromMilligramSquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaMilligramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramSquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramSquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromMilligramSquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromMilligramSquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramSquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromMilligramSquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramSquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramSquareDecimeters(0)
    if err != nil {
        t.Errorf("FromMilligramSquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaMilligramSquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramSquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramSquareDecimeters function
func TestMassMomentOfInertiaFactory_FromKilogramSquareDecimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramSquareDecimeters(100)
    if err != nil {
        t.Errorf("FromKilogramSquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaKilogramSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramSquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramSquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromKilogramSquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromKilogramSquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramSquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilogramSquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramSquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramSquareDecimeters(0)
    if err != nil {
        t.Errorf("FromKilogramSquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaKilogramSquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramSquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramSquareCentimeters function
func TestMassMomentOfInertiaFactory_FromMilligramSquareCentimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromMilligramSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaMilligramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromMilligramSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromMilligramSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromMilligramSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromMilligramSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaMilligramSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramSquareCentimeters function
func TestMassMomentOfInertiaFactory_FromKilogramSquareCentimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromKilogramSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaKilogramSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromKilogramSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromKilogramSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilogramSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromKilogramSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaKilogramSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramSquareMillimeters function
func TestMassMomentOfInertiaFactory_FromMilligramSquareMillimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramSquareMillimeters(100)
    if err != nil {
        t.Errorf("FromMilligramSquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaMilligramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramSquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramSquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromMilligramSquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromMilligramSquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramSquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromMilligramSquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramSquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramSquareMillimeters(0)
    if err != nil {
        t.Errorf("FromMilligramSquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaMilligramSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramSquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramSquareMillimeters function
func TestMassMomentOfInertiaFactory_FromKilogramSquareMillimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramSquareMillimeters(100)
    if err != nil {
        t.Errorf("FromKilogramSquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaKilogramSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramSquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramSquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromKilogramSquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromKilogramSquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramSquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilogramSquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramSquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramSquareMillimeters(0)
    if err != nil {
        t.Errorf("FromKilogramSquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaKilogramSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramSquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilotonneSquareMeters function
func TestMassMomentOfInertiaFactory_FromKilotonneSquareMeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilotonneSquareMeters(100)
    if err != nil {
        t.Errorf("FromKilotonneSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaKilotonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilotonneSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilotonneSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromKilotonneSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromKilotonneSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilotonneSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromKilotonneSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilotonneSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilotonneSquareMeters(0)
    if err != nil {
        t.Errorf("FromKilotonneSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaKilotonneSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilotonneSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMegatonneSquareMeters function
func TestMassMomentOfInertiaFactory_FromMegatonneSquareMeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegatonneSquareMeters(100)
    if err != nil {
        t.Errorf("FromMegatonneSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaMegatonneSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegatonneSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegatonneSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromMegatonneSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromMegatonneSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMegatonneSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromMegatonneSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMegatonneSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegatonneSquareMeters(0)
    if err != nil {
        t.Errorf("FromMegatonneSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaMegatonneSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegatonneSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilotonneSquareDecimeters function
func TestMassMomentOfInertiaFactory_FromKilotonneSquareDecimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilotonneSquareDecimeters(100)
    if err != nil {
        t.Errorf("FromKilotonneSquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaKilotonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilotonneSquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilotonneSquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromKilotonneSquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromKilotonneSquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilotonneSquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilotonneSquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilotonneSquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilotonneSquareDecimeters(0)
    if err != nil {
        t.Errorf("FromKilotonneSquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaKilotonneSquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilotonneSquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMegatonneSquareDecimeters function
func TestMassMomentOfInertiaFactory_FromMegatonneSquareDecimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegatonneSquareDecimeters(100)
    if err != nil {
        t.Errorf("FromMegatonneSquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaMegatonneSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegatonneSquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegatonneSquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromMegatonneSquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromMegatonneSquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMegatonneSquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromMegatonneSquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMegatonneSquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegatonneSquareDecimeters(0)
    if err != nil {
        t.Errorf("FromMegatonneSquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaMegatonneSquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegatonneSquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilotonneSquareCentimeters function
func TestMassMomentOfInertiaFactory_FromKilotonneSquareCentimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilotonneSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromKilotonneSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaKilotonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilotonneSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilotonneSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromKilotonneSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromKilotonneSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilotonneSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilotonneSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilotonneSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilotonneSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromKilotonneSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaKilotonneSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilotonneSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMegatonneSquareCentimeters function
func TestMassMomentOfInertiaFactory_FromMegatonneSquareCentimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegatonneSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromMegatonneSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaMegatonneSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegatonneSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegatonneSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromMegatonneSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromMegatonneSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMegatonneSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromMegatonneSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMegatonneSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegatonneSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromMegatonneSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaMegatonneSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegatonneSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilotonneSquareMilimeters function
func TestMassMomentOfInertiaFactory_FromKilotonneSquareMilimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilotonneSquareMilimeters(100)
    if err != nil {
        t.Errorf("FromKilotonneSquareMilimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaKilotonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilotonneSquareMilimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilotonneSquareMilimeters(math.NaN())
    if err == nil {
        t.Error("FromKilotonneSquareMilimeters() with NaN value should return error")
    }

    _, err = factory.FromKilotonneSquareMilimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilotonneSquareMilimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilotonneSquareMilimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilotonneSquareMilimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilotonneSquareMilimeters(0)
    if err != nil {
        t.Errorf("FromKilotonneSquareMilimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaKilotonneSquareMilimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilotonneSquareMilimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMegatonneSquareMilimeters function
func TestMassMomentOfInertiaFactory_FromMegatonneSquareMilimeters(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegatonneSquareMilimeters(100)
    if err != nil {
        t.Errorf("FromMegatonneSquareMilimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassMomentOfInertiaMegatonneSquareMilimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegatonneSquareMilimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegatonneSquareMilimeters(math.NaN())
    if err == nil {
        t.Error("FromMegatonneSquareMilimeters() with NaN value should return error")
    }

    _, err = factory.FromMegatonneSquareMilimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMegatonneSquareMilimeters() with +Inf value should return error")
    }

    _, err = factory.FromMegatonneSquareMilimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMegatonneSquareMilimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegatonneSquareMilimeters(0)
    if err != nil {
        t.Errorf("FromMegatonneSquareMilimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassMomentOfInertiaMegatonneSquareMilimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegatonneSquareMilimeters() with zero value = %v, want 0", converted)
    }
}

func TestMassMomentOfInertiaToString(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	a, err := factory.CreateMassMomentOfInertia(45, units.MassMomentOfInertiaKilogramSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassMomentOfInertiaKilogramSquareMeter, 2)
	expected := "45.00 " + units.GetMassMomentOfInertiaAbbreviation(units.MassMomentOfInertiaKilogramSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassMomentOfInertiaKilogramSquareMeter, -1)
	expected = "45 " + units.GetMassMomentOfInertiaAbbreviation(units.MassMomentOfInertiaKilogramSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMassMomentOfInertia_EqualityAndComparison(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	a1, _ := factory.CreateMassMomentOfInertia(60, units.MassMomentOfInertiaKilogramSquareMeter)
	a2, _ := factory.CreateMassMomentOfInertia(60, units.MassMomentOfInertiaKilogramSquareMeter)
	a3, _ := factory.CreateMassMomentOfInertia(120, units.MassMomentOfInertiaKilogramSquareMeter)

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

func TestMassMomentOfInertia_Arithmetic(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	a1, _ := factory.CreateMassMomentOfInertia(30, units.MassMomentOfInertiaKilogramSquareMeter)
	a2, _ := factory.CreateMassMomentOfInertia(45, units.MassMomentOfInertiaKilogramSquareMeter)

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


func TestGetMassMomentOfInertiaAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MassMomentOfInertiaUnits
        want string
    }{
        {
            name: "GramSquareMeter abbreviation",
            unit: units.MassMomentOfInertiaGramSquareMeter,
            want: "g·m²",
        },
        {
            name: "GramSquareDecimeter abbreviation",
            unit: units.MassMomentOfInertiaGramSquareDecimeter,
            want: "g·dm²",
        },
        {
            name: "GramSquareCentimeter abbreviation",
            unit: units.MassMomentOfInertiaGramSquareCentimeter,
            want: "g·cm²",
        },
        {
            name: "GramSquareMillimeter abbreviation",
            unit: units.MassMomentOfInertiaGramSquareMillimeter,
            want: "g·mm²",
        },
        {
            name: "TonneSquareMeter abbreviation",
            unit: units.MassMomentOfInertiaTonneSquareMeter,
            want: "t·m²",
        },
        {
            name: "TonneSquareDecimeter abbreviation",
            unit: units.MassMomentOfInertiaTonneSquareDecimeter,
            want: "t·dm²",
        },
        {
            name: "TonneSquareCentimeter abbreviation",
            unit: units.MassMomentOfInertiaTonneSquareCentimeter,
            want: "t·cm²",
        },
        {
            name: "TonneSquareMilimeter abbreviation",
            unit: units.MassMomentOfInertiaTonneSquareMilimeter,
            want: "t·mm²",
        },
        {
            name: "PoundSquareFoot abbreviation",
            unit: units.MassMomentOfInertiaPoundSquareFoot,
            want: "lb·ft²",
        },
        {
            name: "PoundSquareInch abbreviation",
            unit: units.MassMomentOfInertiaPoundSquareInch,
            want: "lb·in²",
        },
        {
            name: "SlugSquareFoot abbreviation",
            unit: units.MassMomentOfInertiaSlugSquareFoot,
            want: "slug·ft²",
        },
        {
            name: "SlugSquareInch abbreviation",
            unit: units.MassMomentOfInertiaSlugSquareInch,
            want: "slug·in²",
        },
        {
            name: "MilligramSquareMeter abbreviation",
            unit: units.MassMomentOfInertiaMilligramSquareMeter,
            want: "mg·m²",
        },
        {
            name: "KilogramSquareMeter abbreviation",
            unit: units.MassMomentOfInertiaKilogramSquareMeter,
            want: "kg·m²",
        },
        {
            name: "MilligramSquareDecimeter abbreviation",
            unit: units.MassMomentOfInertiaMilligramSquareDecimeter,
            want: "mg·dm²",
        },
        {
            name: "KilogramSquareDecimeter abbreviation",
            unit: units.MassMomentOfInertiaKilogramSquareDecimeter,
            want: "kg·dm²",
        },
        {
            name: "MilligramSquareCentimeter abbreviation",
            unit: units.MassMomentOfInertiaMilligramSquareCentimeter,
            want: "mg·cm²",
        },
        {
            name: "KilogramSquareCentimeter abbreviation",
            unit: units.MassMomentOfInertiaKilogramSquareCentimeter,
            want: "kg·cm²",
        },
        {
            name: "MilligramSquareMillimeter abbreviation",
            unit: units.MassMomentOfInertiaMilligramSquareMillimeter,
            want: "mg·mm²",
        },
        {
            name: "KilogramSquareMillimeter abbreviation",
            unit: units.MassMomentOfInertiaKilogramSquareMillimeter,
            want: "kg·mm²",
        },
        {
            name: "KilotonneSquareMeter abbreviation",
            unit: units.MassMomentOfInertiaKilotonneSquareMeter,
            want: "kt·m²",
        },
        {
            name: "MegatonneSquareMeter abbreviation",
            unit: units.MassMomentOfInertiaMegatonneSquareMeter,
            want: "Mt·m²",
        },
        {
            name: "KilotonneSquareDecimeter abbreviation",
            unit: units.MassMomentOfInertiaKilotonneSquareDecimeter,
            want: "kt·dm²",
        },
        {
            name: "MegatonneSquareDecimeter abbreviation",
            unit: units.MassMomentOfInertiaMegatonneSquareDecimeter,
            want: "Mt·dm²",
        },
        {
            name: "KilotonneSquareCentimeter abbreviation",
            unit: units.MassMomentOfInertiaKilotonneSquareCentimeter,
            want: "kt·cm²",
        },
        {
            name: "MegatonneSquareCentimeter abbreviation",
            unit: units.MassMomentOfInertiaMegatonneSquareCentimeter,
            want: "Mt·cm²",
        },
        {
            name: "KilotonneSquareMilimeter abbreviation",
            unit: units.MassMomentOfInertiaKilotonneSquareMilimeter,
            want: "kt·mm²",
        },
        {
            name: "MegatonneSquareMilimeter abbreviation",
            unit: units.MassMomentOfInertiaMegatonneSquareMilimeter,
            want: "Mt·mm²",
        },
        {
            name: "invalid unit",
            unit: units.MassMomentOfInertiaUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMassMomentOfInertiaAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMassMomentOfInertiaAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMassMomentOfInertia_String(t *testing.T) {
    factory := units.MassMomentOfInertiaFactory{}
    
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
            unit, err := factory.CreateMassMomentOfInertia(tt.value, units.MassMomentOfInertiaKilogramSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("MassMomentOfInertia.String() = %v, want %v", got, tt.want)
            }
        })
    }
}