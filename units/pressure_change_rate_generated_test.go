// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPressureChangeRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "PascalPerSecond"}`
	
	factory := units.PressureChangeRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PressureChangeRatePascalPerSecond {
		t.Errorf("expected unit %v, got %v", units.PressureChangeRatePascalPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "PascalPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPressureChangeRateDto_ToJSON(t *testing.T) {
	dto := units.PressureChangeRateDto{
		Value: 45,
		Unit:  units.PressureChangeRatePascalPerSecond,
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
	if result["unit"].(string) != string(units.PressureChangeRatePascalPerSecond) {
		t.Errorf("expected unit %s, got %v", units.PressureChangeRatePascalPerSecond, result["unit"])
	}
}

func TestNewPressureChangeRate_InvalidValue(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePressureChangeRate(math.NaN(), units.PressureChangeRatePascalPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePressureChangeRate(math.Inf(1), units.PressureChangeRatePascalPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPressureChangeRateConversions(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePressureChangeRate(180, units.PressureChangeRatePascalPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PascalsPerSecond.
		// No expected conversion value provided for PascalsPerSecond, verifying result is not NaN.
		result := a.PascalsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PascalsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PascalsPerMinute.
		// No expected conversion value provided for PascalsPerMinute, verifying result is not NaN.
		result := a.PascalsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to PascalsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MillimetersOfMercuryPerSecond.
		// No expected conversion value provided for MillimetersOfMercuryPerSecond, verifying result is not NaN.
		result := a.MillimetersOfMercuryPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersOfMercuryPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AtmospheresPerSecond.
		// No expected conversion value provided for AtmospheresPerSecond, verifying result is not NaN.
		result := a.AtmospheresPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to AtmospheresPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareInchPerSecond.
		// No expected conversion value provided for PoundsForcePerSquareInchPerSecond, verifying result is not NaN.
		result := a.PoundsForcePerSquareInchPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerSquareInchPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareInchPerMinute.
		// No expected conversion value provided for PoundsForcePerSquareInchPerMinute, verifying result is not NaN.
		result := a.PoundsForcePerSquareInchPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerSquareInchPerMinute returned NaN")
		}
	}
	{
		// Test conversion to BarsPerSecond.
		// No expected conversion value provided for BarsPerSecond, verifying result is not NaN.
		result := a.BarsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to BarsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to BarsPerMinute.
		// No expected conversion value provided for BarsPerMinute, verifying result is not NaN.
		result := a.BarsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to BarsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilopascalsPerSecond.
		// No expected conversion value provided for KilopascalsPerSecond, verifying result is not NaN.
		result := a.KilopascalsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopascalsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegapascalsPerSecond.
		// No expected conversion value provided for MegapascalsPerSecond, verifying result is not NaN.
		result := a.MegapascalsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapascalsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilopascalsPerMinute.
		// No expected conversion value provided for KilopascalsPerMinute, verifying result is not NaN.
		result := a.KilopascalsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopascalsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegapascalsPerMinute.
		// No expected conversion value provided for MegapascalsPerMinute, verifying result is not NaN.
		result := a.MegapascalsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapascalsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareInchPerSecond.
		// No expected conversion value provided for KilopoundsForcePerSquareInchPerSecond, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareInchPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerSquareInchPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsForcePerSquareInchPerSecond.
		// No expected conversion value provided for MegapoundsForcePerSquareInchPerSecond, verifying result is not NaN.
		result := a.MegapoundsForcePerSquareInchPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundsForcePerSquareInchPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareInchPerMinute.
		// No expected conversion value provided for KilopoundsForcePerSquareInchPerMinute, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareInchPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerSquareInchPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsForcePerSquareInchPerMinute.
		// No expected conversion value provided for MegapoundsForcePerSquareInchPerMinute, verifying result is not NaN.
		result := a.MegapoundsForcePerSquareInchPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundsForcePerSquareInchPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MillibarsPerSecond.
		// No expected conversion value provided for MillibarsPerSecond, verifying result is not NaN.
		result := a.MillibarsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillibarsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillibarsPerMinute.
		// No expected conversion value provided for MillibarsPerMinute, verifying result is not NaN.
		result := a.MillibarsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillibarsPerMinute returned NaN")
		}
	}
}

func TestPressureChangeRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	a, err := factory.CreatePressureChangeRate(90, units.PressureChangeRatePascalPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PressureChangeRatePascalPerSecond {
		t.Errorf("expected default unit PascalPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PressureChangeRatePascalPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PressureChangeRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PressureChangeRatePascalPerSecond {
		t.Errorf("expected unit PascalPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPressureChangeRateToString(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	a, err := factory.CreatePressureChangeRate(45, units.PressureChangeRatePascalPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PressureChangeRatePascalPerSecond, 2)
	expected := "45.00 " + units.GetPressureChangeRateAbbreviation(units.PressureChangeRatePascalPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PressureChangeRatePascalPerSecond, -1)
	expected = "45 " + units.GetPressureChangeRateAbbreviation(units.PressureChangeRatePascalPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPressureChangeRate_EqualityAndComparison(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	a1, _ := factory.CreatePressureChangeRate(60, units.PressureChangeRatePascalPerSecond)
	a2, _ := factory.CreatePressureChangeRate(60, units.PressureChangeRatePascalPerSecond)
	a3, _ := factory.CreatePressureChangeRate(120, units.PressureChangeRatePascalPerSecond)

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

func TestPressureChangeRate_Arithmetic(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	a1, _ := factory.CreatePressureChangeRate(30, units.PressureChangeRatePascalPerSecond)
	a2, _ := factory.CreatePressureChangeRate(45, units.PressureChangeRatePascalPerSecond)

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