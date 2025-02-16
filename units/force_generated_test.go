// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestForceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Newton"}`
	
	factory := units.ForceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ForceNewton {
		t.Errorf("expected unit %v, got %v", units.ForceNewton, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Newton"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestForceDto_ToJSON(t *testing.T) {
	dto := units.ForceDto{
		Value: 45,
		Unit:  units.ForceNewton,
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
	if result["unit"].(string) != string(units.ForceNewton) {
		t.Errorf("expected unit %s, got %v", units.ForceNewton, result["unit"])
	}
}

func TestNewForce_InvalidValue(t *testing.T) {
	factory := units.ForceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateForce(math.NaN(), units.ForceNewton)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateForce(math.Inf(1), units.ForceNewton)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestForceConversions(t *testing.T) {
	factory := units.ForceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateForce(180, units.ForceNewton)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Dyne.
		// No expected conversion value provided for Dyne, verifying result is not NaN.
		result := a.Dyne()
		if math.IsNaN(result) {
			t.Errorf("conversion to Dyne returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForce.
		// No expected conversion value provided for KilogramsForce, verifying result is not NaN.
		result := a.KilogramsForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForce returned NaN")
		}
	}
	{
		// Test conversion to TonnesForce.
		// No expected conversion value provided for TonnesForce, verifying result is not NaN.
		result := a.TonnesForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForce returned NaN")
		}
	}
	{
		// Test conversion to Newtons.
		// No expected conversion value provided for Newtons, verifying result is not NaN.
		result := a.Newtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Newtons returned NaN")
		}
	}
	{
		// Test conversion to KiloPonds.
		// No expected conversion value provided for KiloPonds, verifying result is not NaN.
		result := a.KiloPonds()
		if math.IsNaN(result) {
			t.Errorf("conversion to KiloPonds returned NaN")
		}
	}
	{
		// Test conversion to Poundals.
		// No expected conversion value provided for Poundals, verifying result is not NaN.
		result := a.Poundals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Poundals returned NaN")
		}
	}
	{
		// Test conversion to PoundsForce.
		// No expected conversion value provided for PoundsForce, verifying result is not NaN.
		result := a.PoundsForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForce returned NaN")
		}
	}
	{
		// Test conversion to OunceForce.
		// No expected conversion value provided for OunceForce, verifying result is not NaN.
		result := a.OunceForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to OunceForce returned NaN")
		}
	}
	{
		// Test conversion to ShortTonsForce.
		// No expected conversion value provided for ShortTonsForce, verifying result is not NaN.
		result := a.ShortTonsForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to ShortTonsForce returned NaN")
		}
	}
	{
		// Test conversion to Micronewtons.
		// No expected conversion value provided for Micronewtons, verifying result is not NaN.
		result := a.Micronewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micronewtons returned NaN")
		}
	}
	{
		// Test conversion to Millinewtons.
		// No expected conversion value provided for Millinewtons, verifying result is not NaN.
		result := a.Millinewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millinewtons returned NaN")
		}
	}
	{
		// Test conversion to Decanewtons.
		// No expected conversion value provided for Decanewtons, verifying result is not NaN.
		result := a.Decanewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decanewtons returned NaN")
		}
	}
	{
		// Test conversion to Kilonewtons.
		// No expected conversion value provided for Kilonewtons, verifying result is not NaN.
		result := a.Kilonewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilonewtons returned NaN")
		}
	}
	{
		// Test conversion to Meganewtons.
		// No expected conversion value provided for Meganewtons, verifying result is not NaN.
		result := a.Meganewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Meganewtons returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForce.
		// No expected conversion value provided for KilopoundsForce, verifying result is not NaN.
		result := a.KilopoundsForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForce returned NaN")
		}
	}
}

func TestForce_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ForceFactory{}
	a, err := factory.CreateForce(90, units.ForceNewton)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ForceNewton {
		t.Errorf("expected default unit Newton, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ForceDyn
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ForceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ForceNewton {
		t.Errorf("expected unit Newton, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestForceToString(t *testing.T) {
	factory := units.ForceFactory{}
	a, err := factory.CreateForce(45, units.ForceNewton)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ForceNewton, 2)
	expected := "45.00 " + units.GetForceAbbreviation(units.ForceNewton)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ForceNewton, -1)
	expected = "45 " + units.GetForceAbbreviation(units.ForceNewton)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestForce_EqualityAndComparison(t *testing.T) {
	factory := units.ForceFactory{}
	a1, _ := factory.CreateForce(60, units.ForceNewton)
	a2, _ := factory.CreateForce(60, units.ForceNewton)
	a3, _ := factory.CreateForce(120, units.ForceNewton)

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

func TestForce_Arithmetic(t *testing.T) {
	factory := units.ForceFactory{}
	a1, _ := factory.CreateForce(30, units.ForceNewton)
	a2, _ := factory.CreateForce(45, units.ForceNewton)

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