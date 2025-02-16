// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRatioDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecimalFraction"}`
	
	factory := units.RatioDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RatioDecimalFraction {
		t.Errorf("expected unit %v, got %v", units.RatioDecimalFraction, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecimalFraction"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRatioDto_ToJSON(t *testing.T) {
	dto := units.RatioDto{
		Value: 45,
		Unit:  units.RatioDecimalFraction,
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
	if result["unit"].(string) != string(units.RatioDecimalFraction) {
		t.Errorf("expected unit %s, got %v", units.RatioDecimalFraction, result["unit"])
	}
}

func TestNewRatio_InvalidValue(t *testing.T) {
	factory := units.RatioFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRatio(math.NaN(), units.RatioDecimalFraction)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRatio(math.Inf(1), units.RatioDecimalFraction)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRatioConversions(t *testing.T) {
	factory := units.RatioFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRatio(180, units.RatioDecimalFraction)
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
}

func TestRatio_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RatioFactory{}
	a, err := factory.CreateRatio(90, units.RatioDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RatioDecimalFraction {
		t.Errorf("expected default unit DecimalFraction, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RatioDecimalFraction
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RatioDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RatioDecimalFraction {
		t.Errorf("expected unit DecimalFraction, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRatioToString(t *testing.T) {
	factory := units.RatioFactory{}
	a, err := factory.CreateRatio(45, units.RatioDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RatioDecimalFraction, 2)
	expected := "45.00 " + units.GetRatioAbbreviation(units.RatioDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RatioDecimalFraction, -1)
	expected = "45 " + units.GetRatioAbbreviation(units.RatioDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRatio_EqualityAndComparison(t *testing.T) {
	factory := units.RatioFactory{}
	a1, _ := factory.CreateRatio(60, units.RatioDecimalFraction)
	a2, _ := factory.CreateRatio(60, units.RatioDecimalFraction)
	a3, _ := factory.CreateRatio(120, units.RatioDecimalFraction)

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

func TestRatio_Arithmetic(t *testing.T) {
	factory := units.RatioFactory{}
	a1, _ := factory.CreateRatio(30, units.RatioDecimalFraction)
	a2, _ := factory.CreateRatio(45, units.RatioDecimalFraction)

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