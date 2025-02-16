// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificWeightDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonPerCubicMeter"}`
	
	factory := units.SpecificWeightDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificWeightNewtonPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.SpecificWeightNewtonPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificWeightDto_ToJSON(t *testing.T) {
	dto := units.SpecificWeightDto{
		Value: 45,
		Unit:  units.SpecificWeightNewtonPerCubicMeter,
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
	if result["unit"].(string) != string(units.SpecificWeightNewtonPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.SpecificWeightNewtonPerCubicMeter, result["unit"])
	}
}

func TestNewSpecificWeight_InvalidValue(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificWeight(math.NaN(), units.SpecificWeightNewtonPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificWeight(math.Inf(1), units.SpecificWeightNewtonPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificWeightConversions(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificWeight(180, units.SpecificWeightNewtonPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonsPerCubicMillimeter.
		// No expected conversion value provided for NewtonsPerCubicMillimeter, verifying result is not NaN.
		result := a.NewtonsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerCubicCentimeter.
		// No expected conversion value provided for NewtonsPerCubicCentimeter, verifying result is not NaN.
		result := a.NewtonsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerCubicMeter.
		// No expected conversion value provided for NewtonsPerCubicMeter, verifying result is not NaN.
		result := a.NewtonsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerCubicMillimeter.
		// No expected conversion value provided for KilogramsForcePerCubicMillimeter, verifying result is not NaN.
		result := a.KilogramsForcePerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerCubicCentimeter.
		// No expected conversion value provided for KilogramsForcePerCubicCentimeter, verifying result is not NaN.
		result := a.KilogramsForcePerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerCubicMeter.
		// No expected conversion value provided for KilogramsForcePerCubicMeter, verifying result is not NaN.
		result := a.KilogramsForcePerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerCubicInch.
		// No expected conversion value provided for PoundsForcePerCubicInch, verifying result is not NaN.
		result := a.PoundsForcePerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerCubicFoot.
		// No expected conversion value provided for PoundsForcePerCubicFoot, verifying result is not NaN.
		result := a.PoundsForcePerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerCubicMillimeter.
		// No expected conversion value provided for TonnesForcePerCubicMillimeter, verifying result is not NaN.
		result := a.TonnesForcePerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerCubicCentimeter.
		// No expected conversion value provided for TonnesForcePerCubicCentimeter, verifying result is not NaN.
		result := a.TonnesForcePerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerCubicMeter.
		// No expected conversion value provided for TonnesForcePerCubicMeter, verifying result is not NaN.
		result := a.TonnesForcePerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerCubicMillimeter.
		// No expected conversion value provided for KilonewtonsPerCubicMillimeter, verifying result is not NaN.
		result := a.KilonewtonsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerCubicCentimeter.
		// No expected conversion value provided for KilonewtonsPerCubicCentimeter, verifying result is not NaN.
		result := a.KilonewtonsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerCubicMeter.
		// No expected conversion value provided for KilonewtonsPerCubicMeter, verifying result is not NaN.
		result := a.KilonewtonsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerCubicMeter.
		// No expected conversion value provided for MeganewtonsPerCubicMeter, verifying result is not NaN.
		result := a.MeganewtonsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerCubicInch.
		// No expected conversion value provided for KilopoundsForcePerCubicInch, verifying result is not NaN.
		result := a.KilopoundsForcePerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerCubicFoot.
		// No expected conversion value provided for KilopoundsForcePerCubicFoot, verifying result is not NaN.
		result := a.KilopoundsForcePerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerCubicFoot returned NaN")
		}
	}
}

func TestSpecificWeight_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	a, err := factory.CreateSpecificWeight(90, units.SpecificWeightNewtonPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificWeightNewtonPerCubicMeter {
		t.Errorf("expected default unit NewtonPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificWeightNewtonPerCubicMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificWeightDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificWeightNewtonPerCubicMeter {
		t.Errorf("expected unit NewtonPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificWeightToString(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	a, err := factory.CreateSpecificWeight(45, units.SpecificWeightNewtonPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificWeightNewtonPerCubicMeter, 2)
	expected := "45.00 " + units.GetSpecificWeightAbbreviation(units.SpecificWeightNewtonPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificWeightNewtonPerCubicMeter, -1)
	expected = "45 " + units.GetSpecificWeightAbbreviation(units.SpecificWeightNewtonPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificWeight_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	a1, _ := factory.CreateSpecificWeight(60, units.SpecificWeightNewtonPerCubicMeter)
	a2, _ := factory.CreateSpecificWeight(60, units.SpecificWeightNewtonPerCubicMeter)
	a3, _ := factory.CreateSpecificWeight(120, units.SpecificWeightNewtonPerCubicMeter)

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

func TestSpecificWeight_Arithmetic(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	a1, _ := factory.CreateSpecificWeight(30, units.SpecificWeightNewtonPerCubicMeter)
	a2, _ := factory.CreateSpecificWeight(45, units.SpecificWeightNewtonPerCubicMeter)

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