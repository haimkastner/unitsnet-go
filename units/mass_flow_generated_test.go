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