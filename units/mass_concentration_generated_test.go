// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassConcentrationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerCubicMeter"}`
	
	factory := units.MassConcentrationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassConcentrationKilogramPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.MassConcentrationKilogramPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassConcentrationDto_ToJSON(t *testing.T) {
	dto := units.MassConcentrationDto{
		Value: 45,
		Unit:  units.MassConcentrationKilogramPerCubicMeter,
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
	if result["unit"].(string) != string(units.MassConcentrationKilogramPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.MassConcentrationKilogramPerCubicMeter, result["unit"])
	}
}

func TestNewMassConcentration_InvalidValue(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMassConcentration(math.NaN(), units.MassConcentrationKilogramPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMassConcentration(math.Inf(1), units.MassConcentrationKilogramPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassConcentrationConversions(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMassConcentration(180, units.MassConcentrationKilogramPerCubicMeter)
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
		// Test conversion to GramsPerMicroliter.
		// No expected conversion value provided for GramsPerMicroliter, verifying result is not NaN.
		result := a.GramsPerMicroliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerMicroliter returned NaN")
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
		// Test conversion to GramsPerDeciliter.
		// No expected conversion value provided for GramsPerDeciliter, verifying result is not NaN.
		result := a.GramsPerDeciliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerDeciliter returned NaN")
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
		// Test conversion to SlugsPerCubicFoot.
		// No expected conversion value provided for SlugsPerCubicFoot, verifying result is not NaN.
		result := a.SlugsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicFoot returned NaN")
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
		// Test conversion to OuncesPerUSGallon.
		// No expected conversion value provided for OuncesPerUSGallon, verifying result is not NaN.
		result := a.OuncesPerUSGallon()
		if math.IsNaN(result) {
			t.Errorf("conversion to OuncesPerUSGallon returned NaN")
		}
	}
	{
		// Test conversion to OuncesPerImperialGallon.
		// No expected conversion value provided for OuncesPerImperialGallon, verifying result is not NaN.
		result := a.OuncesPerImperialGallon()
		if math.IsNaN(result) {
			t.Errorf("conversion to OuncesPerImperialGallon returned NaN")
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
		// Test conversion to PicogramsPerMicroliter.
		// No expected conversion value provided for PicogramsPerMicroliter, verifying result is not NaN.
		result := a.PicogramsPerMicroliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicogramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerMicroliter.
		// No expected conversion value provided for NanogramsPerMicroliter, verifying result is not NaN.
		result := a.NanogramsPerMicroliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMicroliter.
		// No expected conversion value provided for MicrogramsPerMicroliter, verifying result is not NaN.
		result := a.MicrogramsPerMicroliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMicroliter.
		// No expected conversion value provided for MilligramsPerMicroliter, verifying result is not NaN.
		result := a.MilligramsPerMicroliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerMicroliter.
		// No expected conversion value provided for CentigramsPerMicroliter, verifying result is not NaN.
		result := a.CentigramsPerMicroliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerMicroliter.
		// No expected conversion value provided for DecigramsPerMicroliter, verifying result is not NaN.
		result := a.DecigramsPerMicroliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerMicroliter returned NaN")
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
	{
		// Test conversion to PicogramsPerDeciliter.
		// No expected conversion value provided for PicogramsPerDeciliter, verifying result is not NaN.
		result := a.PicogramsPerDeciliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicogramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerDeciliter.
		// No expected conversion value provided for NanogramsPerDeciliter, verifying result is not NaN.
		result := a.NanogramsPerDeciliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerDeciliter.
		// No expected conversion value provided for MicrogramsPerDeciliter, verifying result is not NaN.
		result := a.MicrogramsPerDeciliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerDeciliter.
		// No expected conversion value provided for MilligramsPerDeciliter, verifying result is not NaN.
		result := a.MilligramsPerDeciliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerDeciliter.
		// No expected conversion value provided for CentigramsPerDeciliter, verifying result is not NaN.
		result := a.CentigramsPerDeciliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerDeciliter.
		// No expected conversion value provided for DecigramsPerDeciliter, verifying result is not NaN.
		result := a.DecigramsPerDeciliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerDeciliter returned NaN")
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
		// Test conversion to KilogramsPerLiter.
		// No expected conversion value provided for KilogramsPerLiter, verifying result is not NaN.
		result := a.KilogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerLiter returned NaN")
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
}

func TestMassConcentration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	a, err := factory.CreateMassConcentration(90, units.MassConcentrationKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassConcentrationKilogramPerCubicMeter {
		t.Errorf("expected default unit KilogramPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassConcentrationGramPerCubicMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassConcentrationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassConcentrationKilogramPerCubicMeter {
		t.Errorf("expected unit KilogramPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassConcentrationToString(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	a, err := factory.CreateMassConcentration(45, units.MassConcentrationKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassConcentrationKilogramPerCubicMeter, 2)
	expected := "45.00 " + units.GetMassConcentrationAbbreviation(units.MassConcentrationKilogramPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassConcentrationKilogramPerCubicMeter, -1)
	expected = "45 " + units.GetMassConcentrationAbbreviation(units.MassConcentrationKilogramPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMassConcentration_EqualityAndComparison(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	a1, _ := factory.CreateMassConcentration(60, units.MassConcentrationKilogramPerCubicMeter)
	a2, _ := factory.CreateMassConcentration(60, units.MassConcentrationKilogramPerCubicMeter)
	a3, _ := factory.CreateMassConcentration(120, units.MassConcentrationKilogramPerCubicMeter)

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

func TestMassConcentration_Arithmetic(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	a1, _ := factory.CreateMassConcentration(30, units.MassConcentrationKilogramPerCubicMeter)
	a2, _ := factory.CreateMassConcentration(45, units.MassConcentrationKilogramPerCubicMeter)

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