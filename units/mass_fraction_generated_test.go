// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

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
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimalFractions returned NaN")
		}
	}
	{
		// Test conversion to GramsPerGram.
		// No expected conversion value provided for GramsPerGram, verifying result is not NaN.
		result := a.GramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to GramsPerKilogram.
		// No expected conversion value provided for GramsPerKilogram, verifying result is not NaN.
		result := a.GramsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to Percent.
		// No expected conversion value provided for Percent, verifying result is not NaN.
		result := a.Percent()
		if math.IsNaN(result) {
			t.Errorf("conversion to Percent returned NaN")
		}
	}
	{
		// Test conversion to PartsPerThousand.
		// No expected conversion value provided for PartsPerThousand, verifying result is not NaN.
		result := a.PartsPerThousand()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerThousand returned NaN")
		}
	}
	{
		// Test conversion to PartsPerMillion.
		// No expected conversion value provided for PartsPerMillion, verifying result is not NaN.
		result := a.PartsPerMillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerMillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerBillion.
		// No expected conversion value provided for PartsPerBillion, verifying result is not NaN.
		result := a.PartsPerBillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerBillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerTrillion.
		// No expected conversion value provided for PartsPerTrillion, verifying result is not NaN.
		result := a.PartsPerTrillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerTrillion returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerGram.
		// No expected conversion value provided for NanogramsPerGram, verifying result is not NaN.
		result := a.NanogramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerGram.
		// No expected conversion value provided for MicrogramsPerGram, verifying result is not NaN.
		result := a.MicrogramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerGram.
		// No expected conversion value provided for MilligramsPerGram, verifying result is not NaN.
		result := a.MilligramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerGram.
		// No expected conversion value provided for CentigramsPerGram, verifying result is not NaN.
		result := a.CentigramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerGram.
		// No expected conversion value provided for DecigramsPerGram, verifying result is not NaN.
		result := a.DecigramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to DecagramsPerGram.
		// No expected conversion value provided for DecagramsPerGram, verifying result is not NaN.
		result := a.DecagramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecagramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to HectogramsPerGram.
		// No expected conversion value provided for HectogramsPerGram, verifying result is not NaN.
		result := a.HectogramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectogramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerGram.
		// No expected conversion value provided for KilogramsPerGram, verifying result is not NaN.
		result := a.KilogramsPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerGram returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerKilogram.
		// No expected conversion value provided for NanogramsPerKilogram, verifying result is not NaN.
		result := a.NanogramsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerKilogram.
		// No expected conversion value provided for MicrogramsPerKilogram, verifying result is not NaN.
		result := a.MicrogramsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerKilogram.
		// No expected conversion value provided for MilligramsPerKilogram, verifying result is not NaN.
		result := a.MilligramsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerKilogram.
		// No expected conversion value provided for CentigramsPerKilogram, verifying result is not NaN.
		result := a.CentigramsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerKilogram.
		// No expected conversion value provided for DecigramsPerKilogram, verifying result is not NaN.
		result := a.DecigramsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to DecagramsPerKilogram.
		// No expected conversion value provided for DecagramsPerKilogram, verifying result is not NaN.
		result := a.DecagramsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecagramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to HectogramsPerKilogram.
		// No expected conversion value provided for HectogramsPerKilogram, verifying result is not NaN.
		result := a.HectogramsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectogramsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerKilogram.
		// No expected conversion value provided for KilogramsPerKilogram, verifying result is not NaN.
		result := a.KilogramsPerKilogram()
		if math.IsNaN(result) {
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