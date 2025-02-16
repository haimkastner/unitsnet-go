// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestDurationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Second"}`
	
	factory := units.DurationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.DurationSecond {
		t.Errorf("expected unit %v, got %v", units.DurationSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Second"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestDurationDto_ToJSON(t *testing.T) {
	dto := units.DurationDto{
		Value: 45,
		Unit:  units.DurationSecond,
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
	if result["unit"].(string) != string(units.DurationSecond) {
		t.Errorf("expected unit %s, got %v", units.DurationSecond, result["unit"])
	}
}

func TestNewDuration_InvalidValue(t *testing.T) {
	factory := units.DurationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateDuration(math.NaN(), units.DurationSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateDuration(math.Inf(1), units.DurationSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestDurationConversions(t *testing.T) {
	factory := units.DurationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateDuration(180, units.DurationSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Years365.
		// No expected conversion value provided for Years365, verifying result is not NaN.
		result := a.Years365()
		if math.IsNaN(result) {
			t.Errorf("conversion to Years365 returned NaN")
		}
	}
	{
		// Test conversion to Months30.
		// No expected conversion value provided for Months30, verifying result is not NaN.
		result := a.Months30()
		if math.IsNaN(result) {
			t.Errorf("conversion to Months30 returned NaN")
		}
	}
	{
		// Test conversion to Weeks.
		// No expected conversion value provided for Weeks, verifying result is not NaN.
		result := a.Weeks()
		if math.IsNaN(result) {
			t.Errorf("conversion to Weeks returned NaN")
		}
	}
	{
		// Test conversion to Days.
		// No expected conversion value provided for Days, verifying result is not NaN.
		result := a.Days()
		if math.IsNaN(result) {
			t.Errorf("conversion to Days returned NaN")
		}
	}
	{
		// Test conversion to Hours.
		// No expected conversion value provided for Hours, verifying result is not NaN.
		result := a.Hours()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hours returned NaN")
		}
	}
	{
		// Test conversion to Minutes.
		// No expected conversion value provided for Minutes, verifying result is not NaN.
		result := a.Minutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Minutes returned NaN")
		}
	}
	{
		// Test conversion to Seconds.
		// No expected conversion value provided for Seconds, verifying result is not NaN.
		result := a.Seconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Seconds returned NaN")
		}
	}
	{
		// Test conversion to JulianYears.
		// No expected conversion value provided for JulianYears, verifying result is not NaN.
		result := a.JulianYears()
		if math.IsNaN(result) {
			t.Errorf("conversion to JulianYears returned NaN")
		}
	}
	{
		// Test conversion to Sols.
		// No expected conversion value provided for Sols, verifying result is not NaN.
		result := a.Sols()
		if math.IsNaN(result) {
			t.Errorf("conversion to Sols returned NaN")
		}
	}
	{
		// Test conversion to Nanoseconds.
		// No expected conversion value provided for Nanoseconds, verifying result is not NaN.
		result := a.Nanoseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanoseconds returned NaN")
		}
	}
	{
		// Test conversion to Microseconds.
		// No expected conversion value provided for Microseconds, verifying result is not NaN.
		result := a.Microseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microseconds returned NaN")
		}
	}
	{
		// Test conversion to Milliseconds.
		// No expected conversion value provided for Milliseconds, verifying result is not NaN.
		result := a.Milliseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliseconds returned NaN")
		}
	}
}

func TestDuration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.DurationFactory{}
	a, err := factory.CreateDuration(90, units.DurationSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.DurationSecond {
		t.Errorf("expected default unit Second, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.DurationYear365
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.DurationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.DurationSecond {
		t.Errorf("expected unit Second, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestDurationToString(t *testing.T) {
	factory := units.DurationFactory{}
	a, err := factory.CreateDuration(45, units.DurationSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.DurationSecond, 2)
	expected := "45.00 " + units.GetDurationAbbreviation(units.DurationSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.DurationSecond, -1)
	expected = "45 " + units.GetDurationAbbreviation(units.DurationSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestDuration_EqualityAndComparison(t *testing.T) {
	factory := units.DurationFactory{}
	a1, _ := factory.CreateDuration(60, units.DurationSecond)
	a2, _ := factory.CreateDuration(60, units.DurationSecond)
	a3, _ := factory.CreateDuration(120, units.DurationSecond)

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

func TestDuration_Arithmetic(t *testing.T) {
	factory := units.DurationFactory{}
	a1, _ := factory.CreateDuration(30, units.DurationSecond)
	a2, _ := factory.CreateDuration(45, units.DurationSecond)

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