// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestForceChangeRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonPerSecond"}`
	
	factory := units.ForceChangeRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ForceChangeRateNewtonPerSecond {
		t.Errorf("expected unit %v, got %v", units.ForceChangeRateNewtonPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestForceChangeRateDto_ToJSON(t *testing.T) {
	dto := units.ForceChangeRateDto{
		Value: 45,
		Unit:  units.ForceChangeRateNewtonPerSecond,
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
	if result["unit"].(string) != string(units.ForceChangeRateNewtonPerSecond) {
		t.Errorf("expected unit %s, got %v", units.ForceChangeRateNewtonPerSecond, result["unit"])
	}
}

func TestNewForceChangeRate_InvalidValue(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateForceChangeRate(math.NaN(), units.ForceChangeRateNewtonPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateForceChangeRate(math.Inf(1), units.ForceChangeRateNewtonPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestForceChangeRateConversions(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateForceChangeRate(180, units.ForceChangeRateNewtonPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonsPerMinute.
		// No expected conversion value provided for NewtonsPerMinute, verifying result is not NaN.
		result := a.NewtonsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerSecond.
		// No expected conversion value provided for NewtonsPerSecond, verifying result is not NaN.
		result := a.NewtonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerMinute.
		// No expected conversion value provided for PoundsForcePerMinute, verifying result is not NaN.
		result := a.PoundsForcePerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerMinute returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSecond.
		// No expected conversion value provided for PoundsForcePerSecond, verifying result is not NaN.
		result := a.PoundsForcePerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerMinute.
		// No expected conversion value provided for DecanewtonsPerMinute, verifying result is not NaN.
		result := a.DecanewtonsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerMinute.
		// No expected conversion value provided for KilonewtonsPerMinute, verifying result is not NaN.
		result := a.KilonewtonsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonsPerSecond.
		// No expected conversion value provided for NanonewtonsPerSecond, verifying result is not NaN.
		result := a.NanonewtonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonsPerSecond.
		// No expected conversion value provided for MicronewtonsPerSecond, verifying result is not NaN.
		result := a.MicronewtonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonsPerSecond.
		// No expected conversion value provided for MillinewtonsPerSecond, verifying result is not NaN.
		result := a.MillinewtonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonsPerSecond.
		// No expected conversion value provided for CentinewtonsPerSecond, verifying result is not NaN.
		result := a.CentinewtonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonsPerSecond.
		// No expected conversion value provided for DecinewtonsPerSecond, verifying result is not NaN.
		result := a.DecinewtonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerSecond.
		// No expected conversion value provided for DecanewtonsPerSecond, verifying result is not NaN.
		result := a.DecanewtonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerSecond.
		// No expected conversion value provided for KilonewtonsPerSecond, verifying result is not NaN.
		result := a.KilonewtonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerMinute.
		// No expected conversion value provided for KilopoundsForcePerMinute, verifying result is not NaN.
		result := a.KilopoundsForcePerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSecond.
		// No expected conversion value provided for KilopoundsForcePerSecond, verifying result is not NaN.
		result := a.KilopoundsForcePerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerSecond returned NaN")
		}
	}
}

func TestForceChangeRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	a, err := factory.CreateForceChangeRate(90, units.ForceChangeRateNewtonPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ForceChangeRateNewtonPerSecond {
		t.Errorf("expected default unit NewtonPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ForceChangeRateNewtonPerMinute
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ForceChangeRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ForceChangeRateNewtonPerSecond {
		t.Errorf("expected unit NewtonPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestForceChangeRateToString(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	a, err := factory.CreateForceChangeRate(45, units.ForceChangeRateNewtonPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ForceChangeRateNewtonPerSecond, 2)
	expected := "45.00 " + units.GetForceChangeRateAbbreviation(units.ForceChangeRateNewtonPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ForceChangeRateNewtonPerSecond, -1)
	expected = "45 " + units.GetForceChangeRateAbbreviation(units.ForceChangeRateNewtonPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestForceChangeRate_EqualityAndComparison(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	a1, _ := factory.CreateForceChangeRate(60, units.ForceChangeRateNewtonPerSecond)
	a2, _ := factory.CreateForceChangeRate(60, units.ForceChangeRateNewtonPerSecond)
	a3, _ := factory.CreateForceChangeRate(120, units.ForceChangeRateNewtonPerSecond)

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

func TestForceChangeRate_Arithmetic(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	a1, _ := factory.CreateForceChangeRate(30, units.ForceChangeRateNewtonPerSecond)
	a2, _ := factory.CreateForceChangeRate(45, units.ForceChangeRateNewtonPerSecond)

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