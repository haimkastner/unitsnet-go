// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassFractionDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecimalFraction"}`
	
	factory := units.MassFractionDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassFractionDecimalFraction {
		t.Errorf("expected unit %v, got %v", units.MassFractionDecimalFraction, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecimalFraction"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassFractionDto_ToJSON(t *testing.T) {
	dto := units.MassFractionDto{
		Value: 45,
		Unit:  units.MassFractionDecimalFraction,
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
	if result["unit"].(string) != string(units.MassFractionDecimalFraction) {
		t.Errorf("expected unit %s, got %v", units.MassFractionDecimalFraction, result["unit"])
	}
}

func TestNewMassFraction_InvalidValue(t *testing.T) {
	factory := units.MassFractionFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMassFraction(math.NaN(), units.MassFractionDecimalFraction)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMassFraction(math.Inf(1), units.MassFractionDecimalFraction)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassFractionConversions(t *testing.T) {
	factory := units.MassFractionFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMassFraction(180, units.MassFractionDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to DecimalFractions.
		// No expected conversion value provided for DecimalFractions, verifying result is not NaN.
		result := a.DecimalFractions()
		cacheResult := a.DecimalFractions()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecimalFractions returned NaN")
		}
	}
	{
		// Test conversion to GramsPerGram.
		// No expected conversion value provided for GramsPerGram, verifying result is not NaN.
		result := a.GramsPerGram()
		cacheResult := a.GramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to GramsPerKilogram.
		// No expected conversion value provided for GramsPerKilogram, verifying result is not NaN.
		result := a.GramsPerKilogram()
		cacheResult := a.GramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to Percent.
		// No expected conversion value provided for Percent, verifying result is not NaN.
		result := a.Percent()
		cacheResult := a.Percent()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Percent returned NaN")
		}
	}
	{
		// Test conversion to PartsPerThousand.
		// No expected conversion value provided for PartsPerThousand, verifying result is not NaN.
		result := a.PartsPerThousand()
		cacheResult := a.PartsPerThousand()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PartsPerThousand returned NaN")
		}
	}
	{
		// Test conversion to PartsPerMillion.
		// No expected conversion value provided for PartsPerMillion, verifying result is not NaN.
		result := a.PartsPerMillion()
		cacheResult := a.PartsPerMillion()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PartsPerMillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerBillion.
		// No expected conversion value provided for PartsPerBillion, verifying result is not NaN.
		result := a.PartsPerBillion()
		cacheResult := a.PartsPerBillion()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PartsPerBillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerTrillion.
		// No expected conversion value provided for PartsPerTrillion, verifying result is not NaN.
		result := a.PartsPerTrillion()
		cacheResult := a.PartsPerTrillion()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PartsPerTrillion returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerGram.
		// No expected conversion value provided for NanogramsPerGram, verifying result is not NaN.
		result := a.NanogramsPerGram()
		cacheResult := a.NanogramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanogramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerGram.
		// No expected conversion value provided for MicrogramsPerGram, verifying result is not NaN.
		result := a.MicrogramsPerGram()
		cacheResult := a.MicrogramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrogramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerGram.
		// No expected conversion value provided for MilligramsPerGram, verifying result is not NaN.
		result := a.MilligramsPerGram()
		cacheResult := a.MilligramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerGram.
		// No expected conversion value provided for CentigramsPerGram, verifying result is not NaN.
		result := a.CentigramsPerGram()
		cacheResult := a.CentigramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerGram.
		// No expected conversion value provided for DecigramsPerGram, verifying result is not NaN.
		result := a.DecigramsPerGram()
		cacheResult := a.DecigramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to DecagramsPerGram.
		// No expected conversion value provided for DecagramsPerGram, verifying result is not NaN.
		result := a.DecagramsPerGram()
		cacheResult := a.DecagramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecagramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to HectogramsPerGram.
		// No expected conversion value provided for HectogramsPerGram, verifying result is not NaN.
		result := a.HectogramsPerGram()
		cacheResult := a.HectogramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectogramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerGram.
		// No expected conversion value provided for KilogramsPerGram, verifying result is not NaN.
		result := a.KilogramsPerGram()
		cacheResult := a.KilogramsPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerKilogram.
		// No expected conversion value provided for NanogramsPerKilogram, verifying result is not NaN.
		result := a.NanogramsPerKilogram()
		cacheResult := a.NanogramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanogramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerKilogram.
		// No expected conversion value provided for MicrogramsPerKilogram, verifying result is not NaN.
		result := a.MicrogramsPerKilogram()
		cacheResult := a.MicrogramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrogramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerKilogram.
		// No expected conversion value provided for MilligramsPerKilogram, verifying result is not NaN.
		result := a.MilligramsPerKilogram()
		cacheResult := a.MilligramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerKilogram.
		// No expected conversion value provided for CentigramsPerKilogram, verifying result is not NaN.
		result := a.CentigramsPerKilogram()
		cacheResult := a.CentigramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerKilogram.
		// No expected conversion value provided for DecigramsPerKilogram, verifying result is not NaN.
		result := a.DecigramsPerKilogram()
		cacheResult := a.DecigramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to DecagramsPerKilogram.
		// No expected conversion value provided for DecagramsPerKilogram, verifying result is not NaN.
		result := a.DecagramsPerKilogram()
		cacheResult := a.DecagramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecagramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to HectogramsPerKilogram.
		// No expected conversion value provided for HectogramsPerKilogram, verifying result is not NaN.
		result := a.HectogramsPerKilogram()
		cacheResult := a.HectogramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectogramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerKilogram.
		// No expected conversion value provided for KilogramsPerKilogram, verifying result is not NaN.
		result := a.KilogramsPerKilogram()
		cacheResult := a.KilogramsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerKilogram returned NaN")
		}
	}
}

func TestMassFraction_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassFractionFactory{}
	a, err := factory.CreateMassFraction(90, units.MassFractionDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassFractionDecimalFraction {
		t.Errorf("expected default unit DecimalFraction, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassFractionDecimalFraction
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassFractionDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassFractionDecimalFraction {
		t.Errorf("expected unit DecimalFraction, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassFractionFactory_FromDto(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionDecimalFraction,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MassFractionDto{
        Value: math.NaN(),
        Unit:  units.MassFractionDecimalFraction,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test DecimalFraction conversion
    decimal_fractionsDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionDecimalFraction,
    }
    
    var decimal_fractionsResult *units.MassFraction
    decimal_fractionsResult, err = factory.FromDto(decimal_fractionsDto)
    if err != nil {
        t.Errorf("FromDto() with DecimalFraction returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimal_fractionsResult.Convert(units.MassFractionDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimalFraction = %v, want %v", converted, 100)
    }
    // Test GramPerGram conversion
    grams_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionGramPerGram,
    }
    
    var grams_per_gramResult *units.MassFraction
    grams_per_gramResult, err = factory.FromDto(grams_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_gramResult.Convert(units.MassFractionGramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerGram = %v, want %v", converted, 100)
    }
    // Test GramPerKilogram conversion
    grams_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionGramPerKilogram,
    }
    
    var grams_per_kilogramResult *units.MassFraction
    grams_per_kilogramResult, err = factory.FromDto(grams_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_kilogramResult.Convert(units.MassFractionGramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerKilogram = %v, want %v", converted, 100)
    }
    // Test Percent conversion
    percentDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionPercent,
    }
    
    var percentResult *units.MassFraction
    percentResult, err = factory.FromDto(percentDto)
    if err != nil {
        t.Errorf("FromDto() with Percent returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = percentResult.Convert(units.MassFractionPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Percent = %v, want %v", converted, 100)
    }
    // Test PartPerThousand conversion
    parts_per_thousandDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionPartPerThousand,
    }
    
    var parts_per_thousandResult *units.MassFraction
    parts_per_thousandResult, err = factory.FromDto(parts_per_thousandDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerThousand returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_thousandResult.Convert(units.MassFractionPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerThousand = %v, want %v", converted, 100)
    }
    // Test PartPerMillion conversion
    parts_per_millionDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionPartPerMillion,
    }
    
    var parts_per_millionResult *units.MassFraction
    parts_per_millionResult, err = factory.FromDto(parts_per_millionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerMillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_millionResult.Convert(units.MassFractionPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerMillion = %v, want %v", converted, 100)
    }
    // Test PartPerBillion conversion
    parts_per_billionDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionPartPerBillion,
    }
    
    var parts_per_billionResult *units.MassFraction
    parts_per_billionResult, err = factory.FromDto(parts_per_billionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerBillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_billionResult.Convert(units.MassFractionPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerBillion = %v, want %v", converted, 100)
    }
    // Test PartPerTrillion conversion
    parts_per_trillionDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionPartPerTrillion,
    }
    
    var parts_per_trillionResult *units.MassFraction
    parts_per_trillionResult, err = factory.FromDto(parts_per_trillionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerTrillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_trillionResult.Convert(units.MassFractionPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerTrillion = %v, want %v", converted, 100)
    }
    // Test NanogramPerGram conversion
    nanograms_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionNanogramPerGram,
    }
    
    var nanograms_per_gramResult *units.MassFraction
    nanograms_per_gramResult, err = factory.FromDto(nanograms_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_gramResult.Convert(units.MassFractionNanogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerGram = %v, want %v", converted, 100)
    }
    // Test MicrogramPerGram conversion
    micrograms_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionMicrogramPerGram,
    }
    
    var micrograms_per_gramResult *units.MassFraction
    micrograms_per_gramResult, err = factory.FromDto(micrograms_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_gramResult.Convert(units.MassFractionMicrogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerGram = %v, want %v", converted, 100)
    }
    // Test MilligramPerGram conversion
    milligrams_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionMilligramPerGram,
    }
    
    var milligrams_per_gramResult *units.MassFraction
    milligrams_per_gramResult, err = factory.FromDto(milligrams_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_gramResult.Convert(units.MassFractionMilligramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerGram = %v, want %v", converted, 100)
    }
    // Test CentigramPerGram conversion
    centigrams_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionCentigramPerGram,
    }
    
    var centigrams_per_gramResult *units.MassFraction
    centigrams_per_gramResult, err = factory.FromDto(centigrams_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_gramResult.Convert(units.MassFractionCentigramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerGram = %v, want %v", converted, 100)
    }
    // Test DecigramPerGram conversion
    decigrams_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionDecigramPerGram,
    }
    
    var decigrams_per_gramResult *units.MassFraction
    decigrams_per_gramResult, err = factory.FromDto(decigrams_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_gramResult.Convert(units.MassFractionDecigramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerGram = %v, want %v", converted, 100)
    }
    // Test DecagramPerGram conversion
    decagrams_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionDecagramPerGram,
    }
    
    var decagrams_per_gramResult *units.MassFraction
    decagrams_per_gramResult, err = factory.FromDto(decagrams_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with DecagramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_gramResult.Convert(units.MassFractionDecagramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerGram = %v, want %v", converted, 100)
    }
    // Test HectogramPerGram conversion
    hectograms_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionHectogramPerGram,
    }
    
    var hectograms_per_gramResult *units.MassFraction
    hectograms_per_gramResult, err = factory.FromDto(hectograms_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with HectogramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_gramResult.Convert(units.MassFractionHectogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerGram = %v, want %v", converted, 100)
    }
    // Test KilogramPerGram conversion
    kilograms_per_gramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionKilogramPerGram,
    }
    
    var kilograms_per_gramResult *units.MassFraction
    kilograms_per_gramResult, err = factory.FromDto(kilograms_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_gramResult.Convert(units.MassFractionKilogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerGram = %v, want %v", converted, 100)
    }
    // Test NanogramPerKilogram conversion
    nanograms_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionNanogramPerKilogram,
    }
    
    var nanograms_per_kilogramResult *units.MassFraction
    nanograms_per_kilogramResult, err = factory.FromDto(nanograms_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_kilogramResult.Convert(units.MassFractionNanogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerKilogram = %v, want %v", converted, 100)
    }
    // Test MicrogramPerKilogram conversion
    micrograms_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionMicrogramPerKilogram,
    }
    
    var micrograms_per_kilogramResult *units.MassFraction
    micrograms_per_kilogramResult, err = factory.FromDto(micrograms_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_kilogramResult.Convert(units.MassFractionMicrogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerKilogram = %v, want %v", converted, 100)
    }
    // Test MilligramPerKilogram conversion
    milligrams_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionMilligramPerKilogram,
    }
    
    var milligrams_per_kilogramResult *units.MassFraction
    milligrams_per_kilogramResult, err = factory.FromDto(milligrams_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_kilogramResult.Convert(units.MassFractionMilligramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerKilogram = %v, want %v", converted, 100)
    }
    // Test CentigramPerKilogram conversion
    centigrams_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionCentigramPerKilogram,
    }
    
    var centigrams_per_kilogramResult *units.MassFraction
    centigrams_per_kilogramResult, err = factory.FromDto(centigrams_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_kilogramResult.Convert(units.MassFractionCentigramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerKilogram = %v, want %v", converted, 100)
    }
    // Test DecigramPerKilogram conversion
    decigrams_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionDecigramPerKilogram,
    }
    
    var decigrams_per_kilogramResult *units.MassFraction
    decigrams_per_kilogramResult, err = factory.FromDto(decigrams_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_kilogramResult.Convert(units.MassFractionDecigramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerKilogram = %v, want %v", converted, 100)
    }
    // Test DecagramPerKilogram conversion
    decagrams_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionDecagramPerKilogram,
    }
    
    var decagrams_per_kilogramResult *units.MassFraction
    decagrams_per_kilogramResult, err = factory.FromDto(decagrams_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with DecagramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_kilogramResult.Convert(units.MassFractionDecagramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerKilogram = %v, want %v", converted, 100)
    }
    // Test HectogramPerKilogram conversion
    hectograms_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionHectogramPerKilogram,
    }
    
    var hectograms_per_kilogramResult *units.MassFraction
    hectograms_per_kilogramResult, err = factory.FromDto(hectograms_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with HectogramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_kilogramResult.Convert(units.MassFractionHectogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerKilogram = %v, want %v", converted, 100)
    }
    // Test KilogramPerKilogram conversion
    kilograms_per_kilogramDto := units.MassFractionDto{
        Value: 100,
        Unit:  units.MassFractionKilogramPerKilogram,
    }
    
    var kilograms_per_kilogramResult *units.MassFraction
    kilograms_per_kilogramResult, err = factory.FromDto(kilograms_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_kilogramResult.Convert(units.MassFractionKilogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerKilogram = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MassFractionDto{
        Value: 0,
        Unit:  units.MassFractionDecimalFraction,
    }
    
    var zeroResult *units.MassFraction
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMassFractionFactory_FromDtoJSON(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "DecimalFraction"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "DecimalFraction"}`)
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
    nanJSON, _ := json.Marshal(units.MassFractionDto{
        Value: nanValue,
        Unit:  units.MassFractionDecimalFraction,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with DecimalFraction unit
    decimal_fractionsJSON := []byte(`{"value": 100, "unit": "DecimalFraction"}`)
    decimal_fractionsResult, err := factory.FromDtoJSON(decimal_fractionsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimalFraction unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimal_fractionsResult.Convert(units.MassFractionDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimalFraction = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerGram unit
    grams_per_gramJSON := []byte(`{"value": 100, "unit": "GramPerGram"}`)
    grams_per_gramResult, err := factory.FromDtoJSON(grams_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_gramResult.Convert(units.MassFractionGramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerKilogram unit
    grams_per_kilogramJSON := []byte(`{"value": 100, "unit": "GramPerKilogram"}`)
    grams_per_kilogramResult, err := factory.FromDtoJSON(grams_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_kilogramResult.Convert(units.MassFractionGramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with Percent unit
    percentJSON := []byte(`{"value": 100, "unit": "Percent"}`)
    percentResult, err := factory.FromDtoJSON(percentJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Percent unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = percentResult.Convert(units.MassFractionPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Percent = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerThousand unit
    parts_per_thousandJSON := []byte(`{"value": 100, "unit": "PartPerThousand"}`)
    parts_per_thousandResult, err := factory.FromDtoJSON(parts_per_thousandJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerThousand unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_thousandResult.Convert(units.MassFractionPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerThousand = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerMillion unit
    parts_per_millionJSON := []byte(`{"value": 100, "unit": "PartPerMillion"}`)
    parts_per_millionResult, err := factory.FromDtoJSON(parts_per_millionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerMillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_millionResult.Convert(units.MassFractionPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerMillion = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerBillion unit
    parts_per_billionJSON := []byte(`{"value": 100, "unit": "PartPerBillion"}`)
    parts_per_billionResult, err := factory.FromDtoJSON(parts_per_billionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerBillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_billionResult.Convert(units.MassFractionPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerBillion = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerTrillion unit
    parts_per_trillionJSON := []byte(`{"value": 100, "unit": "PartPerTrillion"}`)
    parts_per_trillionResult, err := factory.FromDtoJSON(parts_per_trillionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerTrillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_trillionResult.Convert(units.MassFractionPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerTrillion = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerGram unit
    nanograms_per_gramJSON := []byte(`{"value": 100, "unit": "NanogramPerGram"}`)
    nanograms_per_gramResult, err := factory.FromDtoJSON(nanograms_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_gramResult.Convert(units.MassFractionNanogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerGram unit
    micrograms_per_gramJSON := []byte(`{"value": 100, "unit": "MicrogramPerGram"}`)
    micrograms_per_gramResult, err := factory.FromDtoJSON(micrograms_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_gramResult.Convert(units.MassFractionMicrogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerGram unit
    milligrams_per_gramJSON := []byte(`{"value": 100, "unit": "MilligramPerGram"}`)
    milligrams_per_gramResult, err := factory.FromDtoJSON(milligrams_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_gramResult.Convert(units.MassFractionMilligramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerGram unit
    centigrams_per_gramJSON := []byte(`{"value": 100, "unit": "CentigramPerGram"}`)
    centigrams_per_gramResult, err := factory.FromDtoJSON(centigrams_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_gramResult.Convert(units.MassFractionCentigramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerGram unit
    decigrams_per_gramJSON := []byte(`{"value": 100, "unit": "DecigramPerGram"}`)
    decigrams_per_gramResult, err := factory.FromDtoJSON(decigrams_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_gramResult.Convert(units.MassFractionDecigramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with DecagramPerGram unit
    decagrams_per_gramJSON := []byte(`{"value": 100, "unit": "DecagramPerGram"}`)
    decagrams_per_gramResult, err := factory.FromDtoJSON(decagrams_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecagramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_gramResult.Convert(units.MassFractionDecagramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with HectogramPerGram unit
    hectograms_per_gramJSON := []byte(`{"value": 100, "unit": "HectogramPerGram"}`)
    hectograms_per_gramResult, err := factory.FromDtoJSON(hectograms_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectogramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_gramResult.Convert(units.MassFractionHectogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerGram unit
    kilograms_per_gramJSON := []byte(`{"value": 100, "unit": "KilogramPerGram"}`)
    kilograms_per_gramResult, err := factory.FromDtoJSON(kilograms_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_gramResult.Convert(units.MassFractionKilogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerGram = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerKilogram unit
    nanograms_per_kilogramJSON := []byte(`{"value": 100, "unit": "NanogramPerKilogram"}`)
    nanograms_per_kilogramResult, err := factory.FromDtoJSON(nanograms_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_kilogramResult.Convert(units.MassFractionNanogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerKilogram unit
    micrograms_per_kilogramJSON := []byte(`{"value": 100, "unit": "MicrogramPerKilogram"}`)
    micrograms_per_kilogramResult, err := factory.FromDtoJSON(micrograms_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_kilogramResult.Convert(units.MassFractionMicrogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerKilogram unit
    milligrams_per_kilogramJSON := []byte(`{"value": 100, "unit": "MilligramPerKilogram"}`)
    milligrams_per_kilogramResult, err := factory.FromDtoJSON(milligrams_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_kilogramResult.Convert(units.MassFractionMilligramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerKilogram unit
    centigrams_per_kilogramJSON := []byte(`{"value": 100, "unit": "CentigramPerKilogram"}`)
    centigrams_per_kilogramResult, err := factory.FromDtoJSON(centigrams_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_kilogramResult.Convert(units.MassFractionCentigramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerKilogram unit
    decigrams_per_kilogramJSON := []byte(`{"value": 100, "unit": "DecigramPerKilogram"}`)
    decigrams_per_kilogramResult, err := factory.FromDtoJSON(decigrams_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_kilogramResult.Convert(units.MassFractionDecigramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with DecagramPerKilogram unit
    decagrams_per_kilogramJSON := []byte(`{"value": 100, "unit": "DecagramPerKilogram"}`)
    decagrams_per_kilogramResult, err := factory.FromDtoJSON(decagrams_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecagramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_kilogramResult.Convert(units.MassFractionDecagramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with HectogramPerKilogram unit
    hectograms_per_kilogramJSON := []byte(`{"value": 100, "unit": "HectogramPerKilogram"}`)
    hectograms_per_kilogramResult, err := factory.FromDtoJSON(hectograms_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectogramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_kilogramResult.Convert(units.MassFractionHectogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerKilogram unit
    kilograms_per_kilogramJSON := []byte(`{"value": 100, "unit": "KilogramPerKilogram"}`)
    kilograms_per_kilogramResult, err := factory.FromDtoJSON(kilograms_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_kilogramResult.Convert(units.MassFractionKilogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerKilogram = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "DecimalFraction"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDecimalFractions function
func TestMassFractionFactory_FromDecimalFractions(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimalFractions(100)
    if err != nil {
        t.Errorf("FromDecimalFractions() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimalFractions() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimalFractions(math.NaN())
    if err == nil {
        t.Error("FromDecimalFractions() with NaN value should return error")
    }

    _, err = factory.FromDecimalFractions(math.Inf(1))
    if err == nil {
        t.Error("FromDecimalFractions() with +Inf value should return error")
    }

    _, err = factory.FromDecimalFractions(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimalFractions() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimalFractions(0)
    if err != nil {
        t.Errorf("FromDecimalFractions() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionDecimalFraction)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimalFractions() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerGram function
func TestMassFractionFactory_FromGramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerGram(100)
    if err != nil {
        t.Errorf("FromGramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionGramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromGramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromGramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerGram(0)
    if err != nil {
        t.Errorf("FromGramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionGramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerKilogram function
func TestMassFractionFactory_FromGramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromGramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionGramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromGramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromGramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromGramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionGramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromPercent function
func TestMassFractionFactory_FromPercent(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPercent(100)
    if err != nil {
        t.Errorf("FromPercent() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPercent() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPercent(math.NaN())
    if err == nil {
        t.Error("FromPercent() with NaN value should return error")
    }

    _, err = factory.FromPercent(math.Inf(1))
    if err == nil {
        t.Error("FromPercent() with +Inf value should return error")
    }

    _, err = factory.FromPercent(math.Inf(-1))
    if err == nil {
        t.Error("FromPercent() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPercent(0)
    if err != nil {
        t.Errorf("FromPercent() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionPercent)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPercent() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerThousand function
func TestMassFractionFactory_FromPartsPerThousand(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerThousand(100)
    if err != nil {
        t.Errorf("FromPartsPerThousand() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerThousand() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerThousand(math.NaN())
    if err == nil {
        t.Error("FromPartsPerThousand() with NaN value should return error")
    }

    _, err = factory.FromPartsPerThousand(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerThousand() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerThousand(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerThousand() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerThousand(0)
    if err != nil {
        t.Errorf("FromPartsPerThousand() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionPartPerThousand)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerThousand() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerMillion function
func TestMassFractionFactory_FromPartsPerMillion(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerMillion(100)
    if err != nil {
        t.Errorf("FromPartsPerMillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerMillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerMillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerMillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerMillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerMillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerMillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerMillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerMillion(0)
    if err != nil {
        t.Errorf("FromPartsPerMillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionPartPerMillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerMillion() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerBillion function
func TestMassFractionFactory_FromPartsPerBillion(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerBillion(100)
    if err != nil {
        t.Errorf("FromPartsPerBillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerBillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerBillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerBillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerBillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerBillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerBillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerBillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerBillion(0)
    if err != nil {
        t.Errorf("FromPartsPerBillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionPartPerBillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerBillion() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerTrillion function
func TestMassFractionFactory_FromPartsPerTrillion(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerTrillion(100)
    if err != nil {
        t.Errorf("FromPartsPerTrillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerTrillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerTrillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerTrillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerTrillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerTrillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerTrillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerTrillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerTrillion(0)
    if err != nil {
        t.Errorf("FromPartsPerTrillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionPartPerTrillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerTrillion() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerGram function
func TestMassFractionFactory_FromNanogramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerGram(100)
    if err != nil {
        t.Errorf("FromNanogramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionNanogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerGram(0)
    if err != nil {
        t.Errorf("FromNanogramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionNanogramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerGram function
func TestMassFractionFactory_FromMicrogramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerGram(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionMicrogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerGram(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionMicrogramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerGram function
func TestMassFractionFactory_FromMilligramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerGram(100)
    if err != nil {
        t.Errorf("FromMilligramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionMilligramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerGram(0)
    if err != nil {
        t.Errorf("FromMilligramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionMilligramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerGram function
func TestMassFractionFactory_FromCentigramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerGram(100)
    if err != nil {
        t.Errorf("FromCentigramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionCentigramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerGram(0)
    if err != nil {
        t.Errorf("FromCentigramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionCentigramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerGram function
func TestMassFractionFactory_FromDecigramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerGram(100)
    if err != nil {
        t.Errorf("FromDecigramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionDecigramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerGram(0)
    if err != nil {
        t.Errorf("FromDecigramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionDecigramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromDecagramsPerGram function
func TestMassFractionFactory_FromDecagramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecagramsPerGram(100)
    if err != nil {
        t.Errorf("FromDecagramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionDecagramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecagramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecagramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromDecagramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromDecagramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromDecagramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromDecagramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromDecagramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecagramsPerGram(0)
    if err != nil {
        t.Errorf("FromDecagramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionDecagramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecagramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromHectogramsPerGram function
func TestMassFractionFactory_FromHectogramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectogramsPerGram(100)
    if err != nil {
        t.Errorf("FromHectogramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionHectogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectogramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectogramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromHectogramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromHectogramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromHectogramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromHectogramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromHectogramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectogramsPerGram(0)
    if err != nil {
        t.Errorf("FromHectogramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionHectogramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectogramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerGram function
func TestMassFractionFactory_FromKilogramsPerGram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerGram(100)
    if err != nil {
        t.Errorf("FromKilogramsPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionKilogramPerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerGram(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerGram() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerGram() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerGram(0)
    if err != nil {
        t.Errorf("FromKilogramsPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionKilogramPerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerKilogram function
func TestMassFractionFactory_FromNanogramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromNanogramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionNanogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromNanogramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionNanogramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerKilogram function
func TestMassFractionFactory_FromMicrogramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionMicrogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionMicrogramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerKilogram function
func TestMassFractionFactory_FromMilligramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromMilligramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionMilligramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromMilligramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionMilligramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerKilogram function
func TestMassFractionFactory_FromCentigramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromCentigramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionCentigramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromCentigramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionCentigramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerKilogram function
func TestMassFractionFactory_FromDecigramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromDecigramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionDecigramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromDecigramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionDecigramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromDecagramsPerKilogram function
func TestMassFractionFactory_FromDecagramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecagramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromDecagramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionDecagramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecagramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecagramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromDecagramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromDecagramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromDecagramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromDecagramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromDecagramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecagramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromDecagramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionDecagramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecagramsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromHectogramsPerKilogram function
func TestMassFractionFactory_FromHectogramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectogramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromHectogramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionHectogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectogramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectogramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromHectogramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromHectogramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromHectogramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromHectogramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromHectogramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectogramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromHectogramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionHectogramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectogramsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerKilogram function
func TestMassFractionFactory_FromKilogramsPerKilogram(t *testing.T) {
    factory := units.MassFractionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerKilogram(100)
    if err != nil {
        t.Errorf("FromKilogramsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFractionKilogramPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerKilogram(0)
    if err != nil {
        t.Errorf("FromKilogramsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFractionKilogramPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerKilogram() with zero value = %v, want 0", converted)
    }
}

func TestMassFractionToString(t *testing.T) {
	factory := units.MassFractionFactory{}
	a, err := factory.CreateMassFraction(45, units.MassFractionDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassFractionDecimalFraction, 2)
	expected := "45.00 " + units.GetMassFractionAbbreviation(units.MassFractionDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassFractionDecimalFraction, -1)
	expected = "45 " + units.GetMassFractionAbbreviation(units.MassFractionDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMassFraction_EqualityAndComparison(t *testing.T) {
	factory := units.MassFractionFactory{}
	a1, _ := factory.CreateMassFraction(60, units.MassFractionDecimalFraction)
	a2, _ := factory.CreateMassFraction(60, units.MassFractionDecimalFraction)
	a3, _ := factory.CreateMassFraction(120, units.MassFractionDecimalFraction)

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

func TestMassFraction_Arithmetic(t *testing.T) {
	factory := units.MassFractionFactory{}
	a1, _ := factory.CreateMassFraction(30, units.MassFractionDecimalFraction)
	a2, _ := factory.CreateMassFraction(45, units.MassFractionDecimalFraction)

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


func TestGetMassFractionAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MassFractionUnits
        want string
    }{
        {
            name: "DecimalFraction abbreviation",
            unit: units.MassFractionDecimalFraction,
            want: "",
        },
        {
            name: "GramPerGram abbreviation",
            unit: units.MassFractionGramPerGram,
            want: "g/g",
        },
        {
            name: "GramPerKilogram abbreviation",
            unit: units.MassFractionGramPerKilogram,
            want: "g/kg",
        },
        {
            name: "Percent abbreviation",
            unit: units.MassFractionPercent,
            want: "%",
        },
        {
            name: "PartPerThousand abbreviation",
            unit: units.MassFractionPartPerThousand,
            want: "‰",
        },
        {
            name: "PartPerMillion abbreviation",
            unit: units.MassFractionPartPerMillion,
            want: "ppm",
        },
        {
            name: "PartPerBillion abbreviation",
            unit: units.MassFractionPartPerBillion,
            want: "ppb",
        },
        {
            name: "PartPerTrillion abbreviation",
            unit: units.MassFractionPartPerTrillion,
            want: "ppt",
        },
        {
            name: "NanogramPerGram abbreviation",
            unit: units.MassFractionNanogramPerGram,
            want: "ng/g",
        },
        {
            name: "MicrogramPerGram abbreviation",
            unit: units.MassFractionMicrogramPerGram,
            want: "μg/g",
        },
        {
            name: "MilligramPerGram abbreviation",
            unit: units.MassFractionMilligramPerGram,
            want: "mg/g",
        },
        {
            name: "CentigramPerGram abbreviation",
            unit: units.MassFractionCentigramPerGram,
            want: "cg/g",
        },
        {
            name: "DecigramPerGram abbreviation",
            unit: units.MassFractionDecigramPerGram,
            want: "dg/g",
        },
        {
            name: "DecagramPerGram abbreviation",
            unit: units.MassFractionDecagramPerGram,
            want: "dag/g",
        },
        {
            name: "HectogramPerGram abbreviation",
            unit: units.MassFractionHectogramPerGram,
            want: "hg/g",
        },
        {
            name: "KilogramPerGram abbreviation",
            unit: units.MassFractionKilogramPerGram,
            want: "kg/g",
        },
        {
            name: "NanogramPerKilogram abbreviation",
            unit: units.MassFractionNanogramPerKilogram,
            want: "ng/kg",
        },
        {
            name: "MicrogramPerKilogram abbreviation",
            unit: units.MassFractionMicrogramPerKilogram,
            want: "μg/kg",
        },
        {
            name: "MilligramPerKilogram abbreviation",
            unit: units.MassFractionMilligramPerKilogram,
            want: "mg/kg",
        },
        {
            name: "CentigramPerKilogram abbreviation",
            unit: units.MassFractionCentigramPerKilogram,
            want: "cg/kg",
        },
        {
            name: "DecigramPerKilogram abbreviation",
            unit: units.MassFractionDecigramPerKilogram,
            want: "dg/kg",
        },
        {
            name: "DecagramPerKilogram abbreviation",
            unit: units.MassFractionDecagramPerKilogram,
            want: "dag/kg",
        },
        {
            name: "HectogramPerKilogram abbreviation",
            unit: units.MassFractionHectogramPerKilogram,
            want: "hg/kg",
        },
        {
            name: "KilogramPerKilogram abbreviation",
            unit: units.MassFractionKilogramPerKilogram,
            want: "kg/kg",
        },
        {
            name: "invalid unit",
            unit: units.MassFractionUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMassFractionAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMassFractionAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMassFraction_String(t *testing.T) {
    factory := units.MassFractionFactory{}
    
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
            unit, err := factory.CreateMassFraction(tt.value, units.MassFractionDecimalFraction)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("MassFraction.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestMassFraction_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.MassFractionFactory{}

	_, err := uf.CreateMassFraction(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}