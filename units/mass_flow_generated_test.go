// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassFlowDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "GramPerSecond"}`
	
	factory := units.MassFlowDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassFlowGramPerSecond {
		t.Errorf("expected unit %v, got %v", units.MassFlowGramPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "GramPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassFlowDto_ToJSON(t *testing.T) {
	dto := units.MassFlowDto{
		Value: 45,
		Unit:  units.MassFlowGramPerSecond,
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
	if result["unit"].(string) != string(units.MassFlowGramPerSecond) {
		t.Errorf("expected unit %s, got %v", units.MassFlowGramPerSecond, result["unit"])
	}
}

func TestNewMassFlow_InvalidValue(t *testing.T) {
	factory := units.MassFlowFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMassFlow(math.NaN(), units.MassFlowGramPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMassFlow(math.Inf(1), units.MassFlowGramPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassFlowConversions(t *testing.T) {
	factory := units.MassFlowFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMassFlow(180, units.MassFlowGramPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerSecond.
		// No expected conversion value provided for GramsPerSecond, verifying result is not NaN.
		result := a.GramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GramsPerDay.
		// No expected conversion value provided for GramsPerDay, verifying result is not NaN.
		result := a.GramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to GramsPerHour.
		// No expected conversion value provided for GramsPerHour, verifying result is not NaN.
		result := a.GramsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerHour.
		// No expected conversion value provided for KilogramsPerHour, verifying result is not NaN.
		result := a.KilogramsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerMinute.
		// No expected conversion value provided for KilogramsPerMinute, verifying result is not NaN.
		result := a.KilogramsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerHour.
		// No expected conversion value provided for TonnesPerHour, verifying result is not NaN.
		result := a.TonnesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerHour returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerDay.
		// No expected conversion value provided for PoundsPerDay, verifying result is not NaN.
		result := a.PoundsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerDay returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerHour.
		// No expected conversion value provided for PoundsPerHour, verifying result is not NaN.
		result := a.PoundsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerHour returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerMinute.
		// No expected conversion value provided for PoundsPerMinute, verifying result is not NaN.
		result := a.PoundsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerSecond.
		// No expected conversion value provided for PoundsPerSecond, verifying result is not NaN.
		result := a.PoundsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerDay.
		// No expected conversion value provided for TonnesPerDay, verifying result is not NaN.
		result := a.TonnesPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerDay returned NaN")
		}
	}
	{
		// Test conversion to ShortTonsPerHour.
		// No expected conversion value provided for ShortTonsPerHour, verifying result is not NaN.
		result := a.ShortTonsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to ShortTonsPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerSecond.
		// No expected conversion value provided for NanogramsPerSecond, verifying result is not NaN.
		result := a.NanogramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerSecond.
		// No expected conversion value provided for MicrogramsPerSecond, verifying result is not NaN.
		result := a.MicrogramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerSecond.
		// No expected conversion value provided for MilligramsPerSecond, verifying result is not NaN.
		result := a.MilligramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerSecond.
		// No expected conversion value provided for CentigramsPerSecond, verifying result is not NaN.
		result := a.CentigramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerSecond.
		// No expected conversion value provided for DecigramsPerSecond, verifying result is not NaN.
		result := a.DecigramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecagramsPerSecond.
		// No expected conversion value provided for DecagramsPerSecond, verifying result is not NaN.
		result := a.DecagramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecagramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to HectogramsPerSecond.
		// No expected conversion value provided for HectogramsPerSecond, verifying result is not NaN.
		result := a.HectogramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectogramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerSecond.
		// No expected conversion value provided for KilogramsPerSecond, verifying result is not NaN.
		result := a.KilogramsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerDay.
		// No expected conversion value provided for NanogramsPerDay, verifying result is not NaN.
		result := a.NanogramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerDay.
		// No expected conversion value provided for MicrogramsPerDay, verifying result is not NaN.
		result := a.MicrogramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerDay.
		// No expected conversion value provided for MilligramsPerDay, verifying result is not NaN.
		result := a.MilligramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerDay.
		// No expected conversion value provided for CentigramsPerDay, verifying result is not NaN.
		result := a.CentigramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerDay.
		// No expected conversion value provided for DecigramsPerDay, verifying result is not NaN.
		result := a.DecigramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to DecagramsPerDay.
		// No expected conversion value provided for DecagramsPerDay, verifying result is not NaN.
		result := a.DecagramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecagramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to HectogramsPerDay.
		// No expected conversion value provided for HectogramsPerDay, verifying result is not NaN.
		result := a.HectogramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectogramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerDay.
		// No expected conversion value provided for KilogramsPerDay, verifying result is not NaN.
		result := a.KilogramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegagramsPerDay.
		// No expected conversion value provided for MegagramsPerDay, verifying result is not NaN.
		result := a.MegagramsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegagramsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsPerDay.
		// No expected conversion value provided for MegapoundsPerDay, verifying result is not NaN.
		result := a.MegapoundsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsPerHour.
		// No expected conversion value provided for MegapoundsPerHour, verifying result is not NaN.
		result := a.MegapoundsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsPerMinute.
		// No expected conversion value provided for MegapoundsPerMinute, verifying result is not NaN.
		result := a.MegapoundsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsPerSecond.
		// No expected conversion value provided for MegapoundsPerSecond, verifying result is not NaN.
		result := a.MegapoundsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundsPerSecond returned NaN")
		}
	}
}

func TestMassFlow_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassFlowFactory{}
	a, err := factory.CreateMassFlow(90, units.MassFlowGramPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassFlowGramPerSecond {
		t.Errorf("expected default unit GramPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassFlowGramPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassFlowDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassFlowGramPerSecond {
		t.Errorf("expected unit GramPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassFlowFactory_FromDto(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowGramPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MassFlowDto{
        Value: math.NaN(),
        Unit:  units.MassFlowGramPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GramPerSecond conversion
    grams_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowGramPerSecond,
    }
    
    var grams_per_secondResult *units.MassFlow
    grams_per_secondResult, err = factory.FromDto(grams_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_secondResult.Convert(units.MassFlowGramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSecond = %v, want %v", converted, 100)
    }
    // Test GramPerDay conversion
    grams_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowGramPerDay,
    }
    
    var grams_per_dayResult *units.MassFlow
    grams_per_dayResult, err = factory.FromDto(grams_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_dayResult.Convert(units.MassFlowGramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerDay = %v, want %v", converted, 100)
    }
    // Test GramPerHour conversion
    grams_per_hourDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowGramPerHour,
    }
    
    var grams_per_hourResult *units.MassFlow
    grams_per_hourResult, err = factory.FromDto(grams_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_hourResult.Convert(units.MassFlowGramPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerHour = %v, want %v", converted, 100)
    }
    // Test KilogramPerHour conversion
    kilograms_per_hourDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowKilogramPerHour,
    }
    
    var kilograms_per_hourResult *units.MassFlow
    kilograms_per_hourResult, err = factory.FromDto(kilograms_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_hourResult.Convert(units.MassFlowKilogramPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerHour = %v, want %v", converted, 100)
    }
    // Test KilogramPerMinute conversion
    kilograms_per_minuteDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowKilogramPerMinute,
    }
    
    var kilograms_per_minuteResult *units.MassFlow
    kilograms_per_minuteResult, err = factory.FromDto(kilograms_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_minuteResult.Convert(units.MassFlowKilogramPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerMinute = %v, want %v", converted, 100)
    }
    // Test TonnePerHour conversion
    tonnes_per_hourDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowTonnePerHour,
    }
    
    var tonnes_per_hourResult *units.MassFlow
    tonnes_per_hourResult, err = factory.FromDto(tonnes_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_hourResult.Convert(units.MassFlowTonnePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerHour = %v, want %v", converted, 100)
    }
    // Test PoundPerDay conversion
    pounds_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowPoundPerDay,
    }
    
    var pounds_per_dayResult *units.MassFlow
    pounds_per_dayResult, err = factory.FromDto(pounds_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_dayResult.Convert(units.MassFlowPoundPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerDay = %v, want %v", converted, 100)
    }
    // Test PoundPerHour conversion
    pounds_per_hourDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowPoundPerHour,
    }
    
    var pounds_per_hourResult *units.MassFlow
    pounds_per_hourResult, err = factory.FromDto(pounds_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_hourResult.Convert(units.MassFlowPoundPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerHour = %v, want %v", converted, 100)
    }
    // Test PoundPerMinute conversion
    pounds_per_minuteDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowPoundPerMinute,
    }
    
    var pounds_per_minuteResult *units.MassFlow
    pounds_per_minuteResult, err = factory.FromDto(pounds_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_minuteResult.Convert(units.MassFlowPoundPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerMinute = %v, want %v", converted, 100)
    }
    // Test PoundPerSecond conversion
    pounds_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowPoundPerSecond,
    }
    
    var pounds_per_secondResult *units.MassFlow
    pounds_per_secondResult, err = factory.FromDto(pounds_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_secondResult.Convert(units.MassFlowPoundPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerSecond = %v, want %v", converted, 100)
    }
    // Test TonnePerDay conversion
    tonnes_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowTonnePerDay,
    }
    
    var tonnes_per_dayResult *units.MassFlow
    tonnes_per_dayResult, err = factory.FromDto(tonnes_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_dayResult.Convert(units.MassFlowTonnePerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerDay = %v, want %v", converted, 100)
    }
    // Test ShortTonPerHour conversion
    short_tons_per_hourDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowShortTonPerHour,
    }
    
    var short_tons_per_hourResult *units.MassFlow
    short_tons_per_hourResult, err = factory.FromDto(short_tons_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with ShortTonPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = short_tons_per_hourResult.Convert(units.MassFlowShortTonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ShortTonPerHour = %v, want %v", converted, 100)
    }
    // Test NanogramPerSecond conversion
    nanograms_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowNanogramPerSecond,
    }
    
    var nanograms_per_secondResult *units.MassFlow
    nanograms_per_secondResult, err = factory.FromDto(nanograms_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_secondResult.Convert(units.MassFlowNanogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerSecond = %v, want %v", converted, 100)
    }
    // Test MicrogramPerSecond conversion
    micrograms_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMicrogramPerSecond,
    }
    
    var micrograms_per_secondResult *units.MassFlow
    micrograms_per_secondResult, err = factory.FromDto(micrograms_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_secondResult.Convert(units.MassFlowMicrogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerSecond = %v, want %v", converted, 100)
    }
    // Test MilligramPerSecond conversion
    milligrams_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMilligramPerSecond,
    }
    
    var milligrams_per_secondResult *units.MassFlow
    milligrams_per_secondResult, err = factory.FromDto(milligrams_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_secondResult.Convert(units.MassFlowMilligramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerSecond = %v, want %v", converted, 100)
    }
    // Test CentigramPerSecond conversion
    centigrams_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowCentigramPerSecond,
    }
    
    var centigrams_per_secondResult *units.MassFlow
    centigrams_per_secondResult, err = factory.FromDto(centigrams_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_secondResult.Convert(units.MassFlowCentigramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerSecond = %v, want %v", converted, 100)
    }
    // Test DecigramPerSecond conversion
    decigrams_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowDecigramPerSecond,
    }
    
    var decigrams_per_secondResult *units.MassFlow
    decigrams_per_secondResult, err = factory.FromDto(decigrams_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_secondResult.Convert(units.MassFlowDecigramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerSecond = %v, want %v", converted, 100)
    }
    // Test DecagramPerSecond conversion
    decagrams_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowDecagramPerSecond,
    }
    
    var decagrams_per_secondResult *units.MassFlow
    decagrams_per_secondResult, err = factory.FromDto(decagrams_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecagramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_secondResult.Convert(units.MassFlowDecagramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerSecond = %v, want %v", converted, 100)
    }
    // Test HectogramPerSecond conversion
    hectograms_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowHectogramPerSecond,
    }
    
    var hectograms_per_secondResult *units.MassFlow
    hectograms_per_secondResult, err = factory.FromDto(hectograms_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with HectogramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_secondResult.Convert(units.MassFlowHectogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerSecond = %v, want %v", converted, 100)
    }
    // Test KilogramPerSecond conversion
    kilograms_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowKilogramPerSecond,
    }
    
    var kilograms_per_secondResult *units.MassFlow
    kilograms_per_secondResult, err = factory.FromDto(kilograms_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_secondResult.Convert(units.MassFlowKilogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSecond = %v, want %v", converted, 100)
    }
    // Test NanogramPerDay conversion
    nanograms_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowNanogramPerDay,
    }
    
    var nanograms_per_dayResult *units.MassFlow
    nanograms_per_dayResult, err = factory.FromDto(nanograms_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_dayResult.Convert(units.MassFlowNanogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerDay = %v, want %v", converted, 100)
    }
    // Test MicrogramPerDay conversion
    micrograms_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMicrogramPerDay,
    }
    
    var micrograms_per_dayResult *units.MassFlow
    micrograms_per_dayResult, err = factory.FromDto(micrograms_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_dayResult.Convert(units.MassFlowMicrogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerDay = %v, want %v", converted, 100)
    }
    // Test MilligramPerDay conversion
    milligrams_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMilligramPerDay,
    }
    
    var milligrams_per_dayResult *units.MassFlow
    milligrams_per_dayResult, err = factory.FromDto(milligrams_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_dayResult.Convert(units.MassFlowMilligramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerDay = %v, want %v", converted, 100)
    }
    // Test CentigramPerDay conversion
    centigrams_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowCentigramPerDay,
    }
    
    var centigrams_per_dayResult *units.MassFlow
    centigrams_per_dayResult, err = factory.FromDto(centigrams_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_dayResult.Convert(units.MassFlowCentigramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerDay = %v, want %v", converted, 100)
    }
    // Test DecigramPerDay conversion
    decigrams_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowDecigramPerDay,
    }
    
    var decigrams_per_dayResult *units.MassFlow
    decigrams_per_dayResult, err = factory.FromDto(decigrams_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_dayResult.Convert(units.MassFlowDecigramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerDay = %v, want %v", converted, 100)
    }
    // Test DecagramPerDay conversion
    decagrams_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowDecagramPerDay,
    }
    
    var decagrams_per_dayResult *units.MassFlow
    decagrams_per_dayResult, err = factory.FromDto(decagrams_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with DecagramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_dayResult.Convert(units.MassFlowDecagramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerDay = %v, want %v", converted, 100)
    }
    // Test HectogramPerDay conversion
    hectograms_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowHectogramPerDay,
    }
    
    var hectograms_per_dayResult *units.MassFlow
    hectograms_per_dayResult, err = factory.FromDto(hectograms_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with HectogramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_dayResult.Convert(units.MassFlowHectogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerDay = %v, want %v", converted, 100)
    }
    // Test KilogramPerDay conversion
    kilograms_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowKilogramPerDay,
    }
    
    var kilograms_per_dayResult *units.MassFlow
    kilograms_per_dayResult, err = factory.FromDto(kilograms_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_dayResult.Convert(units.MassFlowKilogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerDay = %v, want %v", converted, 100)
    }
    // Test MegagramPerDay conversion
    megagrams_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMegagramPerDay,
    }
    
    var megagrams_per_dayResult *units.MassFlow
    megagrams_per_dayResult, err = factory.FromDto(megagrams_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MegagramPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megagrams_per_dayResult.Convert(units.MassFlowMegagramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegagramPerDay = %v, want %v", converted, 100)
    }
    // Test MegapoundPerDay conversion
    megapounds_per_dayDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMegapoundPerDay,
    }
    
    var megapounds_per_dayResult *units.MassFlow
    megapounds_per_dayResult, err = factory.FromDto(megapounds_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_dayResult.Convert(units.MassFlowMegapoundPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerDay = %v, want %v", converted, 100)
    }
    // Test MegapoundPerHour conversion
    megapounds_per_hourDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMegapoundPerHour,
    }
    
    var megapounds_per_hourResult *units.MassFlow
    megapounds_per_hourResult, err = factory.FromDto(megapounds_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_hourResult.Convert(units.MassFlowMegapoundPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerHour = %v, want %v", converted, 100)
    }
    // Test MegapoundPerMinute conversion
    megapounds_per_minuteDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMegapoundPerMinute,
    }
    
    var megapounds_per_minuteResult *units.MassFlow
    megapounds_per_minuteResult, err = factory.FromDto(megapounds_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_minuteResult.Convert(units.MassFlowMegapoundPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerMinute = %v, want %v", converted, 100)
    }
    // Test MegapoundPerSecond conversion
    megapounds_per_secondDto := units.MassFlowDto{
        Value: 100,
        Unit:  units.MassFlowMegapoundPerSecond,
    }
    
    var megapounds_per_secondResult *units.MassFlow
    megapounds_per_secondResult, err = factory.FromDto(megapounds_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_secondResult.Convert(units.MassFlowMegapoundPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MassFlowDto{
        Value: 0,
        Unit:  units.MassFlowGramPerSecond,
    }
    
    var zeroResult *units.MassFlow
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMassFlowFactory_FromDtoJSON(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "GramPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "GramPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.MassFlowDto{
        Value: nanValue,
        Unit:  units.MassFlowGramPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with GramPerSecond unit
    grams_per_secondJSON := []byte(`{"value": 100, "unit": "GramPerSecond"}`)
    grams_per_secondResult, err := factory.FromDtoJSON(grams_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_secondResult.Convert(units.MassFlowGramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerDay unit
    grams_per_dayJSON := []byte(`{"value": 100, "unit": "GramPerDay"}`)
    grams_per_dayResult, err := factory.FromDtoJSON(grams_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_dayResult.Convert(units.MassFlowGramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerHour unit
    grams_per_hourJSON := []byte(`{"value": 100, "unit": "GramPerHour"}`)
    grams_per_hourResult, err := factory.FromDtoJSON(grams_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_hourResult.Convert(units.MassFlowGramPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerHour unit
    kilograms_per_hourJSON := []byte(`{"value": 100, "unit": "KilogramPerHour"}`)
    kilograms_per_hourResult, err := factory.FromDtoJSON(kilograms_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_hourResult.Convert(units.MassFlowKilogramPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerMinute unit
    kilograms_per_minuteJSON := []byte(`{"value": 100, "unit": "KilogramPerMinute"}`)
    kilograms_per_minuteResult, err := factory.FromDtoJSON(kilograms_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_minuteResult.Convert(units.MassFlowKilogramPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with TonnePerHour unit
    tonnes_per_hourJSON := []byte(`{"value": 100, "unit": "TonnePerHour"}`)
    tonnes_per_hourResult, err := factory.FromDtoJSON(tonnes_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonnePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_hourResult.Convert(units.MassFlowTonnePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerDay unit
    pounds_per_dayJSON := []byte(`{"value": 100, "unit": "PoundPerDay"}`)
    pounds_per_dayResult, err := factory.FromDtoJSON(pounds_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_dayResult.Convert(units.MassFlowPoundPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerHour unit
    pounds_per_hourJSON := []byte(`{"value": 100, "unit": "PoundPerHour"}`)
    pounds_per_hourResult, err := factory.FromDtoJSON(pounds_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_hourResult.Convert(units.MassFlowPoundPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerMinute unit
    pounds_per_minuteJSON := []byte(`{"value": 100, "unit": "PoundPerMinute"}`)
    pounds_per_minuteResult, err := factory.FromDtoJSON(pounds_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_minuteResult.Convert(units.MassFlowPoundPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerSecond unit
    pounds_per_secondJSON := []byte(`{"value": 100, "unit": "PoundPerSecond"}`)
    pounds_per_secondResult, err := factory.FromDtoJSON(pounds_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_secondResult.Convert(units.MassFlowPoundPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with TonnePerDay unit
    tonnes_per_dayJSON := []byte(`{"value": 100, "unit": "TonnePerDay"}`)
    tonnes_per_dayResult, err := factory.FromDtoJSON(tonnes_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonnePerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_dayResult.Convert(units.MassFlowTonnePerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerDay = %v, want %v", converted, 100)
    }
    // Test JSON with ShortTonPerHour unit
    short_tons_per_hourJSON := []byte(`{"value": 100, "unit": "ShortTonPerHour"}`)
    short_tons_per_hourResult, err := factory.FromDtoJSON(short_tons_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ShortTonPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = short_tons_per_hourResult.Convert(units.MassFlowShortTonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ShortTonPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerSecond unit
    nanograms_per_secondJSON := []byte(`{"value": 100, "unit": "NanogramPerSecond"}`)
    nanograms_per_secondResult, err := factory.FromDtoJSON(nanograms_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_secondResult.Convert(units.MassFlowNanogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerSecond unit
    micrograms_per_secondJSON := []byte(`{"value": 100, "unit": "MicrogramPerSecond"}`)
    micrograms_per_secondResult, err := factory.FromDtoJSON(micrograms_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_secondResult.Convert(units.MassFlowMicrogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerSecond unit
    milligrams_per_secondJSON := []byte(`{"value": 100, "unit": "MilligramPerSecond"}`)
    milligrams_per_secondResult, err := factory.FromDtoJSON(milligrams_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_secondResult.Convert(units.MassFlowMilligramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerSecond unit
    centigrams_per_secondJSON := []byte(`{"value": 100, "unit": "CentigramPerSecond"}`)
    centigrams_per_secondResult, err := factory.FromDtoJSON(centigrams_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_secondResult.Convert(units.MassFlowCentigramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerSecond unit
    decigrams_per_secondJSON := []byte(`{"value": 100, "unit": "DecigramPerSecond"}`)
    decigrams_per_secondResult, err := factory.FromDtoJSON(decigrams_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_secondResult.Convert(units.MassFlowDecigramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecagramPerSecond unit
    decagrams_per_secondJSON := []byte(`{"value": 100, "unit": "DecagramPerSecond"}`)
    decagrams_per_secondResult, err := factory.FromDtoJSON(decagrams_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecagramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_secondResult.Convert(units.MassFlowDecagramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with HectogramPerSecond unit
    hectograms_per_secondJSON := []byte(`{"value": 100, "unit": "HectogramPerSecond"}`)
    hectograms_per_secondResult, err := factory.FromDtoJSON(hectograms_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectogramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_secondResult.Convert(units.MassFlowHectogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerSecond unit
    kilograms_per_secondJSON := []byte(`{"value": 100, "unit": "KilogramPerSecond"}`)
    kilograms_per_secondResult, err := factory.FromDtoJSON(kilograms_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_secondResult.Convert(units.MassFlowKilogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerDay unit
    nanograms_per_dayJSON := []byte(`{"value": 100, "unit": "NanogramPerDay"}`)
    nanograms_per_dayResult, err := factory.FromDtoJSON(nanograms_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_dayResult.Convert(units.MassFlowNanogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerDay unit
    micrograms_per_dayJSON := []byte(`{"value": 100, "unit": "MicrogramPerDay"}`)
    micrograms_per_dayResult, err := factory.FromDtoJSON(micrograms_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_dayResult.Convert(units.MassFlowMicrogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerDay unit
    milligrams_per_dayJSON := []byte(`{"value": 100, "unit": "MilligramPerDay"}`)
    milligrams_per_dayResult, err := factory.FromDtoJSON(milligrams_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_dayResult.Convert(units.MassFlowMilligramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerDay unit
    centigrams_per_dayJSON := []byte(`{"value": 100, "unit": "CentigramPerDay"}`)
    centigrams_per_dayResult, err := factory.FromDtoJSON(centigrams_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_dayResult.Convert(units.MassFlowCentigramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerDay unit
    decigrams_per_dayJSON := []byte(`{"value": 100, "unit": "DecigramPerDay"}`)
    decigrams_per_dayResult, err := factory.FromDtoJSON(decigrams_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_dayResult.Convert(units.MassFlowDecigramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with DecagramPerDay unit
    decagrams_per_dayJSON := []byte(`{"value": 100, "unit": "DecagramPerDay"}`)
    decagrams_per_dayResult, err := factory.FromDtoJSON(decagrams_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecagramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_dayResult.Convert(units.MassFlowDecagramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with HectogramPerDay unit
    hectograms_per_dayJSON := []byte(`{"value": 100, "unit": "HectogramPerDay"}`)
    hectograms_per_dayResult, err := factory.FromDtoJSON(hectograms_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectogramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_dayResult.Convert(units.MassFlowHectogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerDay unit
    kilograms_per_dayJSON := []byte(`{"value": 100, "unit": "KilogramPerDay"}`)
    kilograms_per_dayResult, err := factory.FromDtoJSON(kilograms_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_dayResult.Convert(units.MassFlowKilogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MegagramPerDay unit
    megagrams_per_dayJSON := []byte(`{"value": 100, "unit": "MegagramPerDay"}`)
    megagrams_per_dayResult, err := factory.FromDtoJSON(megagrams_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegagramPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megagrams_per_dayResult.Convert(units.MassFlowMegagramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegagramPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundPerDay unit
    megapounds_per_dayJSON := []byte(`{"value": 100, "unit": "MegapoundPerDay"}`)
    megapounds_per_dayResult, err := factory.FromDtoJSON(megapounds_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_dayResult.Convert(units.MassFlowMegapoundPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundPerHour unit
    megapounds_per_hourJSON := []byte(`{"value": 100, "unit": "MegapoundPerHour"}`)
    megapounds_per_hourResult, err := factory.FromDtoJSON(megapounds_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_hourResult.Convert(units.MassFlowMegapoundPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundPerMinute unit
    megapounds_per_minuteJSON := []byte(`{"value": 100, "unit": "MegapoundPerMinute"}`)
    megapounds_per_minuteResult, err := factory.FromDtoJSON(megapounds_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_minuteResult.Convert(units.MassFlowMegapoundPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundPerSecond unit
    megapounds_per_secondJSON := []byte(`{"value": 100, "unit": "MegapoundPerSecond"}`)
    megapounds_per_secondResult, err := factory.FromDtoJSON(megapounds_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_secondResult.Convert(units.MassFlowMegapoundPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "GramPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGramsPerSecond function
func TestMassFlowFactory_FromGramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerSecond(100)
    if err != nil {
        t.Errorf("FromGramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowGramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromGramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromGramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerSecond(0)
    if err != nil {
        t.Errorf("FromGramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowGramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerDay function
func TestMassFlowFactory_FromGramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerDay(100)
    if err != nil {
        t.Errorf("FromGramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowGramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromGramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromGramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerDay(0)
    if err != nil {
        t.Errorf("FromGramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowGramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerHour function
func TestMassFlowFactory_FromGramsPerHour(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerHour(100)
    if err != nil {
        t.Errorf("FromGramsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowGramPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerHour(math.NaN())
    if err == nil {
        t.Error("FromGramsPerHour() with NaN value should return error")
    }

    _, err = factory.FromGramsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerHour(0)
    if err != nil {
        t.Errorf("FromGramsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowGramPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerHour function
func TestMassFlowFactory_FromKilogramsPerHour(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerHour(100)
    if err != nil {
        t.Errorf("FromKilogramsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowKilogramPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerHour(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerHour() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerHour(0)
    if err != nil {
        t.Errorf("FromKilogramsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowKilogramPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerMinute function
func TestMassFlowFactory_FromKilogramsPerMinute(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerMinute(100)
    if err != nil {
        t.Errorf("FromKilogramsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowKilogramPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerMinute(0)
    if err != nil {
        t.Errorf("FromKilogramsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowKilogramPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerHour function
func TestMassFlowFactory_FromTonnesPerHour(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerHour(100)
    if err != nil {
        t.Errorf("FromTonnesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowTonnePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesPerHour(math.NaN())
    if err == nil {
        t.Error("FromTonnesPerHour() with NaN value should return error")
    }

    _, err = factory.FromTonnesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromTonnesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesPerHour(0)
    if err != nil {
        t.Errorf("FromTonnesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowTonnePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerDay function
func TestMassFlowFactory_FromPoundsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerDay(100)
    if err != nil {
        t.Errorf("FromPoundsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowPoundPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerDay(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerDay() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerDay(0)
    if err != nil {
        t.Errorf("FromPoundsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowPoundPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerHour function
func TestMassFlowFactory_FromPoundsPerHour(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerHour(100)
    if err != nil {
        t.Errorf("FromPoundsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowPoundPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerHour(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerHour() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerHour(0)
    if err != nil {
        t.Errorf("FromPoundsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowPoundPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerMinute function
func TestMassFlowFactory_FromPoundsPerMinute(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerMinute(100)
    if err != nil {
        t.Errorf("FromPoundsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowPoundPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerMinute(0)
    if err != nil {
        t.Errorf("FromPoundsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowPoundPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerSecond function
func TestMassFlowFactory_FromPoundsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerSecond(100)
    if err != nil {
        t.Errorf("FromPoundsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowPoundPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerSecond(0)
    if err != nil {
        t.Errorf("FromPoundsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowPoundPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerDay function
func TestMassFlowFactory_FromTonnesPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerDay(100)
    if err != nil {
        t.Errorf("FromTonnesPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowTonnePerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesPerDay(math.NaN())
    if err == nil {
        t.Error("FromTonnesPerDay() with NaN value should return error")
    }

    _, err = factory.FromTonnesPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesPerDay() with +Inf value should return error")
    }

    _, err = factory.FromTonnesPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesPerDay(0)
    if err != nil {
        t.Errorf("FromTonnesPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowTonnePerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromShortTonsPerHour function
func TestMassFlowFactory_FromShortTonsPerHour(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromShortTonsPerHour(100)
    if err != nil {
        t.Errorf("FromShortTonsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowShortTonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromShortTonsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromShortTonsPerHour(math.NaN())
    if err == nil {
        t.Error("FromShortTonsPerHour() with NaN value should return error")
    }

    _, err = factory.FromShortTonsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromShortTonsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromShortTonsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromShortTonsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromShortTonsPerHour(0)
    if err != nil {
        t.Errorf("FromShortTonsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowShortTonPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromShortTonsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerSecond function
func TestMassFlowFactory_FromNanogramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerSecond(100)
    if err != nil {
        t.Errorf("FromNanogramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowNanogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerSecond(0)
    if err != nil {
        t.Errorf("FromNanogramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowNanogramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerSecond function
func TestMassFlowFactory_FromMicrogramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerSecond(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMicrogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerSecond(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMicrogramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerSecond function
func TestMassFlowFactory_FromMilligramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerSecond(100)
    if err != nil {
        t.Errorf("FromMilligramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMilligramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerSecond(0)
    if err != nil {
        t.Errorf("FromMilligramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMilligramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerSecond function
func TestMassFlowFactory_FromCentigramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerSecond(100)
    if err != nil {
        t.Errorf("FromCentigramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowCentigramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerSecond(0)
    if err != nil {
        t.Errorf("FromCentigramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowCentigramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerSecond function
func TestMassFlowFactory_FromDecigramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerSecond(100)
    if err != nil {
        t.Errorf("FromDecigramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowDecigramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerSecond(0)
    if err != nil {
        t.Errorf("FromDecigramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowDecigramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecagramsPerSecond function
func TestMassFlowFactory_FromDecagramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecagramsPerSecond(100)
    if err != nil {
        t.Errorf("FromDecagramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowDecagramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecagramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecagramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecagramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecagramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecagramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecagramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecagramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecagramsPerSecond(0)
    if err != nil {
        t.Errorf("FromDecagramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowDecagramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecagramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromHectogramsPerSecond function
func TestMassFlowFactory_FromHectogramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectogramsPerSecond(100)
    if err != nil {
        t.Errorf("FromHectogramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowHectogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectogramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectogramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromHectogramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromHectogramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromHectogramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromHectogramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromHectogramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectogramsPerSecond(0)
    if err != nil {
        t.Errorf("FromHectogramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowHectogramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectogramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerSecond function
func TestMassFlowFactory_FromKilogramsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerSecond(100)
    if err != nil {
        t.Errorf("FromKilogramsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowKilogramPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerSecond(0)
    if err != nil {
        t.Errorf("FromKilogramsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowKilogramPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerDay function
func TestMassFlowFactory_FromNanogramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerDay(100)
    if err != nil {
        t.Errorf("FromNanogramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowNanogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerDay(0)
    if err != nil {
        t.Errorf("FromNanogramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowNanogramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerDay function
func TestMassFlowFactory_FromMicrogramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerDay(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMicrogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerDay(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMicrogramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerDay function
func TestMassFlowFactory_FromMilligramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerDay(100)
    if err != nil {
        t.Errorf("FromMilligramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMilligramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerDay(0)
    if err != nil {
        t.Errorf("FromMilligramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMilligramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerDay function
func TestMassFlowFactory_FromCentigramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerDay(100)
    if err != nil {
        t.Errorf("FromCentigramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowCentigramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerDay(0)
    if err != nil {
        t.Errorf("FromCentigramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowCentigramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerDay function
func TestMassFlowFactory_FromDecigramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerDay(100)
    if err != nil {
        t.Errorf("FromDecigramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowDecigramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerDay(0)
    if err != nil {
        t.Errorf("FromDecigramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowDecigramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromDecagramsPerDay function
func TestMassFlowFactory_FromDecagramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecagramsPerDay(100)
    if err != nil {
        t.Errorf("FromDecagramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowDecagramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecagramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecagramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromDecagramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromDecagramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromDecagramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromDecagramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromDecagramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecagramsPerDay(0)
    if err != nil {
        t.Errorf("FromDecagramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowDecagramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecagramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromHectogramsPerDay function
func TestMassFlowFactory_FromHectogramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectogramsPerDay(100)
    if err != nil {
        t.Errorf("FromHectogramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowHectogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectogramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectogramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromHectogramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromHectogramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromHectogramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromHectogramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromHectogramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectogramsPerDay(0)
    if err != nil {
        t.Errorf("FromHectogramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowHectogramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectogramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerDay function
func TestMassFlowFactory_FromKilogramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerDay(100)
    if err != nil {
        t.Errorf("FromKilogramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowKilogramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerDay(0)
    if err != nil {
        t.Errorf("FromKilogramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowKilogramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMegagramsPerDay function
func TestMassFlowFactory_FromMegagramsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegagramsPerDay(100)
    if err != nil {
        t.Errorf("FromMegagramsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMegagramPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegagramsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegagramsPerDay(math.NaN())
    if err == nil {
        t.Error("FromMegagramsPerDay() with NaN value should return error")
    }

    _, err = factory.FromMegagramsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMegagramsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMegagramsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMegagramsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegagramsPerDay(0)
    if err != nil {
        t.Errorf("FromMegagramsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMegagramPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegagramsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundsPerDay function
func TestMassFlowFactory_FromMegapoundsPerDay(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundsPerDay(100)
    if err != nil {
        t.Errorf("FromMegapoundsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMegapoundPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundsPerDay(math.NaN())
    if err == nil {
        t.Error("FromMegapoundsPerDay() with NaN value should return error")
    }

    _, err = factory.FromMegapoundsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundsPerDay(0)
    if err != nil {
        t.Errorf("FromMegapoundsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMegapoundPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundsPerHour function
func TestMassFlowFactory_FromMegapoundsPerHour(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundsPerHour(100)
    if err != nil {
        t.Errorf("FromMegapoundsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMegapoundPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundsPerHour(math.NaN())
    if err == nil {
        t.Error("FromMegapoundsPerHour() with NaN value should return error")
    }

    _, err = factory.FromMegapoundsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundsPerHour(0)
    if err != nil {
        t.Errorf("FromMegapoundsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMegapoundPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundsPerMinute function
func TestMassFlowFactory_FromMegapoundsPerMinute(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundsPerMinute(100)
    if err != nil {
        t.Errorf("FromMegapoundsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMegapoundPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMegapoundsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMegapoundsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundsPerMinute(0)
    if err != nil {
        t.Errorf("FromMegapoundsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMegapoundPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundsPerSecond function
func TestMassFlowFactory_FromMegapoundsPerSecond(t *testing.T) {
    factory := units.MassFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundsPerSecond(100)
    if err != nil {
        t.Errorf("FromMegapoundsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFlowMegapoundPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegapoundsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegapoundsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundsPerSecond(0)
    if err != nil {
        t.Errorf("FromMegapoundsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFlowMegapoundPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundsPerSecond() with zero value = %v, want 0", converted)
    }
}

func TestMassFlowToString(t *testing.T) {
	factory := units.MassFlowFactory{}
	a, err := factory.CreateMassFlow(45, units.MassFlowGramPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassFlowGramPerSecond, 2)
	expected := "45.00 " + units.GetMassFlowAbbreviation(units.MassFlowGramPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassFlowGramPerSecond, -1)
	expected = "45 " + units.GetMassFlowAbbreviation(units.MassFlowGramPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMassFlow_EqualityAndComparison(t *testing.T) {
	factory := units.MassFlowFactory{}
	a1, _ := factory.CreateMassFlow(60, units.MassFlowGramPerSecond)
	a2, _ := factory.CreateMassFlow(60, units.MassFlowGramPerSecond)
	a3, _ := factory.CreateMassFlow(120, units.MassFlowGramPerSecond)

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

func TestMassFlow_Arithmetic(t *testing.T) {
	factory := units.MassFlowFactory{}
	a1, _ := factory.CreateMassFlow(30, units.MassFlowGramPerSecond)
	a2, _ := factory.CreateMassFlow(45, units.MassFlowGramPerSecond)

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