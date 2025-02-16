// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAmountOfSubstanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Mole"}`
	
	factory := units.AmountOfSubstanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AmountOfSubstanceMole {
		t.Errorf("expected unit %v, got %v", units.AmountOfSubstanceMole, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Mole"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAmountOfSubstanceDto_ToJSON(t *testing.T) {
	dto := units.AmountOfSubstanceDto{
		Value: 45,
		Unit:  units.AmountOfSubstanceMole,
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
	if result["unit"].(string) != string(units.AmountOfSubstanceMole) {
		t.Errorf("expected unit %s, got %v", units.AmountOfSubstanceMole, result["unit"])
	}
}

func TestNewAmountOfSubstance_InvalidValue(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAmountOfSubstance(math.NaN(), units.AmountOfSubstanceMole)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAmountOfSubstance(math.Inf(1), units.AmountOfSubstanceMole)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAmountOfSubstanceConversions(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAmountOfSubstance(180, units.AmountOfSubstanceMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Moles.
		// No expected conversion value provided for Moles, verifying result is not NaN.
		result := a.Moles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Moles returned NaN")
		}
	}
	{
		// Test conversion to PoundMoles.
		// No expected conversion value provided for PoundMoles, verifying result is not NaN.
		result := a.PoundMoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundMoles returned NaN")
		}
	}
	{
		// Test conversion to Femtomoles.
		// No expected conversion value provided for Femtomoles, verifying result is not NaN.
		result := a.Femtomoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Femtomoles returned NaN")
		}
	}
	{
		// Test conversion to Picomoles.
		// No expected conversion value provided for Picomoles, verifying result is not NaN.
		result := a.Picomoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picomoles returned NaN")
		}
	}
	{
		// Test conversion to Nanomoles.
		// No expected conversion value provided for Nanomoles, verifying result is not NaN.
		result := a.Nanomoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanomoles returned NaN")
		}
	}
	{
		// Test conversion to Micromoles.
		// No expected conversion value provided for Micromoles, verifying result is not NaN.
		result := a.Micromoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micromoles returned NaN")
		}
	}
	{
		// Test conversion to Millimoles.
		// No expected conversion value provided for Millimoles, verifying result is not NaN.
		result := a.Millimoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millimoles returned NaN")
		}
	}
	{
		// Test conversion to Centimoles.
		// No expected conversion value provided for Centimoles, verifying result is not NaN.
		result := a.Centimoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centimoles returned NaN")
		}
	}
	{
		// Test conversion to Decimoles.
		// No expected conversion value provided for Decimoles, verifying result is not NaN.
		result := a.Decimoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decimoles returned NaN")
		}
	}
	{
		// Test conversion to Kilomoles.
		// No expected conversion value provided for Kilomoles, verifying result is not NaN.
		result := a.Kilomoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilomoles returned NaN")
		}
	}
	{
		// Test conversion to Megamoles.
		// No expected conversion value provided for Megamoles, verifying result is not NaN.
		result := a.Megamoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megamoles returned NaN")
		}
	}
	{
		// Test conversion to NanopoundMoles.
		// No expected conversion value provided for NanopoundMoles, verifying result is not NaN.
		result := a.NanopoundMoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanopoundMoles returned NaN")
		}
	}
	{
		// Test conversion to MicropoundMoles.
		// No expected conversion value provided for MicropoundMoles, verifying result is not NaN.
		result := a.MicropoundMoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicropoundMoles returned NaN")
		}
	}
	{
		// Test conversion to MillipoundMoles.
		// No expected conversion value provided for MillipoundMoles, verifying result is not NaN.
		result := a.MillipoundMoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillipoundMoles returned NaN")
		}
	}
	{
		// Test conversion to CentipoundMoles.
		// No expected conversion value provided for CentipoundMoles, verifying result is not NaN.
		result := a.CentipoundMoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentipoundMoles returned NaN")
		}
	}
	{
		// Test conversion to DecipoundMoles.
		// No expected conversion value provided for DecipoundMoles, verifying result is not NaN.
		result := a.DecipoundMoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecipoundMoles returned NaN")
		}
	}
	{
		// Test conversion to KilopoundMoles.
		// No expected conversion value provided for KilopoundMoles, verifying result is not NaN.
		result := a.KilopoundMoles()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundMoles returned NaN")
		}
	}
}

func TestAmountOfSubstance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	a, err := factory.CreateAmountOfSubstance(90, units.AmountOfSubstanceMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AmountOfSubstanceMole {
		t.Errorf("expected default unit Mole, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AmountOfSubstanceMole
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AmountOfSubstanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AmountOfSubstanceMole {
		t.Errorf("expected unit Mole, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAmountOfSubstanceToString(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	a, err := factory.CreateAmountOfSubstance(45, units.AmountOfSubstanceMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AmountOfSubstanceMole, 2)
	expected := "45.00 " + units.GetAmountOfSubstanceAbbreviation(units.AmountOfSubstanceMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AmountOfSubstanceMole, -1)
	expected = "45 " + units.GetAmountOfSubstanceAbbreviation(units.AmountOfSubstanceMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAmountOfSubstance_EqualityAndComparison(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	a1, _ := factory.CreateAmountOfSubstance(60, units.AmountOfSubstanceMole)
	a2, _ := factory.CreateAmountOfSubstance(60, units.AmountOfSubstanceMole)
	a3, _ := factory.CreateAmountOfSubstance(120, units.AmountOfSubstanceMole)

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

func TestAmountOfSubstance_Arithmetic(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	a1, _ := factory.CreateAmountOfSubstance(30, units.AmountOfSubstanceMole)
	a2, _ := factory.CreateAmountOfSubstance(45, units.AmountOfSubstanceMole)

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