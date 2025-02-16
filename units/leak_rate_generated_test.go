// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLeakRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "PascalCubicMeterPerSecond"}`
	
	factory := units.LeakRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LeakRatePascalCubicMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.LeakRatePascalCubicMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "PascalCubicMeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLeakRateDto_ToJSON(t *testing.T) {
	dto := units.LeakRateDto{
		Value: 45,
		Unit:  units.LeakRatePascalCubicMeterPerSecond,
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
	if result["unit"].(string) != string(units.LeakRatePascalCubicMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.LeakRatePascalCubicMeterPerSecond, result["unit"])
	}
}

func TestNewLeakRate_InvalidValue(t *testing.T) {
	factory := units.LeakRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLeakRate(math.NaN(), units.LeakRatePascalCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLeakRate(math.Inf(1), units.LeakRatePascalCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLeakRateConversions(t *testing.T) {
	factory := units.LeakRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLeakRate(180, units.LeakRatePascalCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PascalCubicMetersPerSecond.
		// No expected conversion value provided for PascalCubicMetersPerSecond, verifying result is not NaN.
		result := a.PascalCubicMetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PascalCubicMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillibarLitersPerSecond.
		// No expected conversion value provided for MillibarLitersPerSecond, verifying result is not NaN.
		result := a.MillibarLitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillibarLitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TorrLitersPerSecond.
		// No expected conversion value provided for TorrLitersPerSecond, verifying result is not NaN.
		result := a.TorrLitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to TorrLitersPerSecond returned NaN")
		}
	}
}

func TestLeakRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LeakRateFactory{}
	a, err := factory.CreateLeakRate(90, units.LeakRatePascalCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LeakRatePascalCubicMeterPerSecond {
		t.Errorf("expected default unit PascalCubicMeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LeakRatePascalCubicMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LeakRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LeakRatePascalCubicMeterPerSecond {
		t.Errorf("expected unit PascalCubicMeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLeakRateToString(t *testing.T) {
	factory := units.LeakRateFactory{}
	a, err := factory.CreateLeakRate(45, units.LeakRatePascalCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LeakRatePascalCubicMeterPerSecond, 2)
	expected := "45.00 " + units.GetLeakRateAbbreviation(units.LeakRatePascalCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LeakRatePascalCubicMeterPerSecond, -1)
	expected = "45 " + units.GetLeakRateAbbreviation(units.LeakRatePascalCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLeakRate_EqualityAndComparison(t *testing.T) {
	factory := units.LeakRateFactory{}
	a1, _ := factory.CreateLeakRate(60, units.LeakRatePascalCubicMeterPerSecond)
	a2, _ := factory.CreateLeakRate(60, units.LeakRatePascalCubicMeterPerSecond)
	a3, _ := factory.CreateLeakRate(120, units.LeakRatePascalCubicMeterPerSecond)

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

func TestLeakRate_Arithmetic(t *testing.T) {
	factory := units.LeakRateFactory{}
	a1, _ := factory.CreateLeakRate(30, units.LeakRatePascalCubicMeterPerSecond)
	a2, _ := factory.CreateLeakRate(45, units.LeakRatePascalCubicMeterPerSecond)

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