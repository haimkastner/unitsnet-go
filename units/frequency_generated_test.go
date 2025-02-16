// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestFrequencyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Hertz"}`
	
	factory := units.FrequencyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.FrequencyHertz {
		t.Errorf("expected unit %v, got %v", units.FrequencyHertz, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Hertz"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestFrequencyDto_ToJSON(t *testing.T) {
	dto := units.FrequencyDto{
		Value: 45,
		Unit:  units.FrequencyHertz,
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
	if result["unit"].(string) != string(units.FrequencyHertz) {
		t.Errorf("expected unit %s, got %v", units.FrequencyHertz, result["unit"])
	}
}

func TestNewFrequency_InvalidValue(t *testing.T) {
	factory := units.FrequencyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateFrequency(math.NaN(), units.FrequencyHertz)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateFrequency(math.Inf(1), units.FrequencyHertz)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestFrequencyConversions(t *testing.T) {
	factory := units.FrequencyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateFrequency(180, units.FrequencyHertz)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Hertz.
		// No expected conversion value provided for Hertz, verifying result is not NaN.
		result := a.Hertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hertz returned NaN")
		}
	}
	{
		// Test conversion to RadiansPerSecond.
		// No expected conversion value provided for RadiansPerSecond, verifying result is not NaN.
		result := a.RadiansPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to RadiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CyclesPerMinute.
		// No expected conversion value provided for CyclesPerMinute, verifying result is not NaN.
		result := a.CyclesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to CyclesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CyclesPerHour.
		// No expected conversion value provided for CyclesPerHour, verifying result is not NaN.
		result := a.CyclesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to CyclesPerHour returned NaN")
		}
	}
	{
		// Test conversion to BeatsPerMinute.
		// No expected conversion value provided for BeatsPerMinute, verifying result is not NaN.
		result := a.BeatsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to BeatsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to PerSecond.
		// No expected conversion value provided for PerSecond, verifying result is not NaN.
		result := a.PerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PerSecond returned NaN")
		}
	}
	{
		// Test conversion to BUnits.
		// No expected conversion value provided for BUnits, verifying result is not NaN.
		result := a.BUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to BUnits returned NaN")
		}
	}
	{
		// Test conversion to Microhertz.
		// No expected conversion value provided for Microhertz, verifying result is not NaN.
		result := a.Microhertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microhertz returned NaN")
		}
	}
	{
		// Test conversion to Millihertz.
		// No expected conversion value provided for Millihertz, verifying result is not NaN.
		result := a.Millihertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millihertz returned NaN")
		}
	}
	{
		// Test conversion to Kilohertz.
		// No expected conversion value provided for Kilohertz, verifying result is not NaN.
		result := a.Kilohertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilohertz returned NaN")
		}
	}
	{
		// Test conversion to Megahertz.
		// No expected conversion value provided for Megahertz, verifying result is not NaN.
		result := a.Megahertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megahertz returned NaN")
		}
	}
	{
		// Test conversion to Gigahertz.
		// No expected conversion value provided for Gigahertz, verifying result is not NaN.
		result := a.Gigahertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigahertz returned NaN")
		}
	}
	{
		// Test conversion to Terahertz.
		// No expected conversion value provided for Terahertz, verifying result is not NaN.
		result := a.Terahertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terahertz returned NaN")
		}
	}
}

func TestFrequency_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.FrequencyFactory{}
	a, err := factory.CreateFrequency(90, units.FrequencyHertz)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.FrequencyHertz {
		t.Errorf("expected default unit Hertz, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.FrequencyHertz
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.FrequencyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.FrequencyHertz {
		t.Errorf("expected unit Hertz, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestFrequencyToString(t *testing.T) {
	factory := units.FrequencyFactory{}
	a, err := factory.CreateFrequency(45, units.FrequencyHertz)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.FrequencyHertz, 2)
	expected := "45.00 " + units.GetFrequencyAbbreviation(units.FrequencyHertz)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.FrequencyHertz, -1)
	expected = "45 " + units.GetFrequencyAbbreviation(units.FrequencyHertz)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestFrequency_EqualityAndComparison(t *testing.T) {
	factory := units.FrequencyFactory{}
	a1, _ := factory.CreateFrequency(60, units.FrequencyHertz)
	a2, _ := factory.CreateFrequency(60, units.FrequencyHertz)
	a3, _ := factory.CreateFrequency(120, units.FrequencyHertz)

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

func TestFrequency_Arithmetic(t *testing.T) {
	factory := units.FrequencyFactory{}
	a1, _ := factory.CreateFrequency(30, units.FrequencyHertz)
	a2, _ := factory.CreateFrequency(45, units.FrequencyHertz)

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