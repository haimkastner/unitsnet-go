// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerCubicMeter"}`
	
	factory := units.DensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.DensityKilogramPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.DensityKilogramPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestDensityDto_ToJSON(t *testing.T) {
	dto := units.DensityDto{
		Value: 45,
		Unit:  units.DensityKilogramPerCubicMeter,
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
	if result["unit"].(string) != string(units.DensityKilogramPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.DensityKilogramPerCubicMeter, result["unit"])
	}
}

func TestNewDensity_InvalidValue(t *testing.T) {
	factory := units.DensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateDensity(math.NaN(), units.DensityKilogramPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateDensity(math.Inf(1), units.DensityKilogramPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestDensityConversions(t *testing.T) {
	factory := units.DensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateDensity(180, units.DensityKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerCubicMillimeter.
		// No expected conversion value provided for GramsPerCubicMillimeter, verifying result is not NaN.
		result := a.GramsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicCentimeter.
		// No expected conversion value provided for GramsPerCubicCentimeter, verifying result is not NaN.
		result := a.GramsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicMeter.
		// No expected conversion value provided for GramsPerCubicMeter, verifying result is not NaN.
		result := a.GramsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicInch.
		// No expected conversion value provided for PoundsPerCubicInch, verifying result is not NaN.
		result := a.PoundsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicFoot.
		// No expected conversion value provided for PoundsPerCubicFoot, verifying result is not NaN.
		result := a.PoundsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicYard.
		// No expected conversion value provided for PoundsPerCubicYard, verifying result is not NaN.
		result := a.PoundsPerCubicYard()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicYard returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicMillimeter.
		// No expected conversion value provided for TonnesPerCubicMillimeter, verifying result is not NaN.
		result := a.TonnesPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicCentimeter.
		// No expected conversion value provided for TonnesPerCubicCentimeter, verifying result is not NaN.
		result := a.TonnesPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicMeter.
		// No expected conversion value provided for TonnesPerCubicMeter, verifying result is not NaN.
		result := a.TonnesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicFoot.
		// No expected conversion value provided for SlugsPerCubicFoot, verifying result is not NaN.
		result := a.SlugsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to GramsPerLiter.
		// No expected conversion value provided for GramsPerLiter, verifying result is not NaN.
		result := a.GramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerDeciLiter.
		// No expected conversion value provided for GramsPerDeciLiter, verifying result is not NaN.
		result := a.GramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerMilliliter.
		// No expected conversion value provided for GramsPerMilliliter, verifying result is not NaN.
		result := a.GramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerUSGallon.
		// No expected conversion value provided for PoundsPerUSGallon, verifying result is not NaN.
		result := a.PoundsPerUSGallon()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerUSGallon returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerImperialGallon.
		// No expected conversion value provided for PoundsPerImperialGallon, verifying result is not NaN.
		result := a.PoundsPerImperialGallon()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerImperialGallon returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerLiter.
		// No expected conversion value provided for KilogramsPerLiter, verifying result is not NaN.
		result := a.KilogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicFoot.
		// No expected conversion value provided for TonnesPerCubicFoot, verifying result is not NaN.
		result := a.TonnesPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicInch.
		// No expected conversion value provided for TonnesPerCubicInch, verifying result is not NaN.
		result := a.TonnesPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicFoot.
		// No expected conversion value provided for GramsPerCubicFoot, verifying result is not NaN.
		result := a.GramsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicInch.
		// No expected conversion value provided for GramsPerCubicInch, verifying result is not NaN.
		result := a.GramsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicMeter.
		// No expected conversion value provided for PoundsPerCubicMeter, verifying result is not NaN.
		result := a.PoundsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicCentimeter.
		// No expected conversion value provided for PoundsPerCubicCentimeter, verifying result is not NaN.
		result := a.PoundsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicMillimeter.
		// No expected conversion value provided for PoundsPerCubicMillimeter, verifying result is not NaN.
		result := a.PoundsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicMeter.
		// No expected conversion value provided for SlugsPerCubicMeter, verifying result is not NaN.
		result := a.SlugsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicCentimeter.
		// No expected conversion value provided for SlugsPerCubicCentimeter, verifying result is not NaN.
		result := a.SlugsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicMillimeter.
		// No expected conversion value provided for SlugsPerCubicMillimeter, verifying result is not NaN.
		result := a.SlugsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicInch.
		// No expected conversion value provided for SlugsPerCubicInch, verifying result is not NaN.
		result := a.SlugsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicMillimeter.
		// No expected conversion value provided for KilogramsPerCubicMillimeter, verifying result is not NaN.
		result := a.KilogramsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicCentimeter.
		// No expected conversion value provided for KilogramsPerCubicCentimeter, verifying result is not NaN.
		result := a.KilogramsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicMeter.
		// No expected conversion value provided for KilogramsPerCubicMeter, verifying result is not NaN.
		result := a.KilogramsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerCubicMeter.
		// No expected conversion value provided for MilligramsPerCubicMeter, verifying result is not NaN.
		result := a.MilligramsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerCubicMeter.
		// No expected conversion value provided for MicrogramsPerCubicMeter, verifying result is not NaN.
		result := a.MicrogramsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerCubicInch.
		// No expected conversion value provided for KilopoundsPerCubicInch, verifying result is not NaN.
		result := a.KilopoundsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerCubicFoot.
		// No expected conversion value provided for KilopoundsPerCubicFoot, verifying result is not NaN.
		result := a.KilopoundsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerCubicYard.
		// No expected conversion value provided for KilopoundsPerCubicYard, verifying result is not NaN.
		result := a.KilopoundsPerCubicYard()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsPerCubicYard returned NaN")
		}
	}
	{
		// Test conversion to FemtogramsPerLiter.
		// No expected conversion value provided for FemtogramsPerLiter, verifying result is not NaN.
		result := a.FemtogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to FemtogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerLiter.
		// No expected conversion value provided for PicogramsPerLiter, verifying result is not NaN.
		result := a.PicogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerLiter.
		// No expected conversion value provided for NanogramsPerLiter, verifying result is not NaN.
		result := a.NanogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerLiter.
		// No expected conversion value provided for MicrogramsPerLiter, verifying result is not NaN.
		result := a.MicrogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerLiter.
		// No expected conversion value provided for MilligramsPerLiter, verifying result is not NaN.
		result := a.MilligramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerLiter.
		// No expected conversion value provided for CentigramsPerLiter, verifying result is not NaN.
		result := a.CentigramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerLiter.
		// No expected conversion value provided for DecigramsPerLiter, verifying result is not NaN.
		result := a.DecigramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to FemtogramsPerDeciLiter.
		// No expected conversion value provided for FemtogramsPerDeciLiter, verifying result is not NaN.
		result := a.FemtogramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to FemtogramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerDeciLiter.
		// No expected conversion value provided for PicogramsPerDeciLiter, verifying result is not NaN.
		result := a.PicogramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicogramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerDeciLiter.
		// No expected conversion value provided for NanogramsPerDeciLiter, verifying result is not NaN.
		result := a.NanogramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerDeciLiter.
		// No expected conversion value provided for MicrogramsPerDeciLiter, verifying result is not NaN.
		result := a.MicrogramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerDeciLiter.
		// No expected conversion value provided for MilligramsPerDeciLiter, verifying result is not NaN.
		result := a.MilligramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerDeciLiter.
		// No expected conversion value provided for CentigramsPerDeciLiter, verifying result is not NaN.
		result := a.CentigramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerDeciLiter.
		// No expected conversion value provided for DecigramsPerDeciLiter, verifying result is not NaN.
		result := a.DecigramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to FemtogramsPerMilliliter.
		// No expected conversion value provided for FemtogramsPerMilliliter, verifying result is not NaN.
		result := a.FemtogramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to FemtogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerMilliliter.
		// No expected conversion value provided for PicogramsPerMilliliter, verifying result is not NaN.
		result := a.PicogramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerMilliliter.
		// No expected conversion value provided for NanogramsPerMilliliter, verifying result is not NaN.
		result := a.NanogramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMilliliter.
		// No expected conversion value provided for MicrogramsPerMilliliter, verifying result is not NaN.
		result := a.MicrogramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMilliliter.
		// No expected conversion value provided for MilligramsPerMilliliter, verifying result is not NaN.
		result := a.MilligramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerMilliliter.
		// No expected conversion value provided for CentigramsPerMilliliter, verifying result is not NaN.
		result := a.CentigramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerMilliliter.
		// No expected conversion value provided for DecigramsPerMilliliter, verifying result is not NaN.
		result := a.DecigramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerMilliliter returned NaN")
		}
	}
}

func TestDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.DensityFactory{}
	a, err := factory.CreateDensity(90, units.DensityKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.DensityKilogramPerCubicMeter {
		t.Errorf("expected default unit KilogramPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.DensityGramPerCubicMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.DensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.DensityKilogramPerCubicMeter {
		t.Errorf("expected unit KilogramPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestDensityToString(t *testing.T) {
	factory := units.DensityFactory{}
	a, err := factory.CreateDensity(45, units.DensityKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.DensityKilogramPerCubicMeter, 2)
	expected := "45.00 " + units.GetDensityAbbreviation(units.DensityKilogramPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.DensityKilogramPerCubicMeter, -1)
	expected = "45 " + units.GetDensityAbbreviation(units.DensityKilogramPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestDensity_EqualityAndComparison(t *testing.T) {
	factory := units.DensityFactory{}
	a1, _ := factory.CreateDensity(60, units.DensityKilogramPerCubicMeter)
	a2, _ := factory.CreateDensity(60, units.DensityKilogramPerCubicMeter)
	a3, _ := factory.CreateDensity(120, units.DensityKilogramPerCubicMeter)

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

func TestDensity_Arithmetic(t *testing.T) {
	factory := units.DensityFactory{}
	a1, _ := factory.CreateDensity(30, units.DensityKilogramPerCubicMeter)
	a2, _ := factory.CreateDensity(45, units.DensityKilogramPerCubicMeter)

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