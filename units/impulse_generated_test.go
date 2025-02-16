// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestImpulseDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonSecond"}`
	
	factory := units.ImpulseDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ImpulseNewtonSecond {
		t.Errorf("expected unit %v, got %v", units.ImpulseNewtonSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestImpulseDto_ToJSON(t *testing.T) {
	dto := units.ImpulseDto{
		Value: 45,
		Unit:  units.ImpulseNewtonSecond,
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
	if result["unit"].(string) != string(units.ImpulseNewtonSecond) {
		t.Errorf("expected unit %s, got %v", units.ImpulseNewtonSecond, result["unit"])
	}
}

func TestNewImpulse_InvalidValue(t *testing.T) {
	factory := units.ImpulseFactory{}
	// NaN value should return an error.
	_, err := factory.CreateImpulse(math.NaN(), units.ImpulseNewtonSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateImpulse(math.Inf(1), units.ImpulseNewtonSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestImpulseConversions(t *testing.T) {
	factory := units.ImpulseFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateImpulse(180, units.ImpulseNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to KilogramMetersPerSecond.
		// No expected conversion value provided for KilogramMetersPerSecond, verifying result is not NaN.
		result := a.KilogramMetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NewtonSeconds.
		// No expected conversion value provided for NewtonSeconds, verifying result is not NaN.
		result := a.NewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to PoundFeetPerSecond.
		// No expected conversion value provided for PoundFeetPerSecond, verifying result is not NaN.
		result := a.PoundFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundForceSeconds.
		// No expected conversion value provided for PoundForceSeconds, verifying result is not NaN.
		result := a.PoundForceSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceSeconds returned NaN")
		}
	}
	{
		// Test conversion to SlugFeetPerSecond.
		// No expected conversion value provided for SlugFeetPerSecond, verifying result is not NaN.
		result := a.SlugFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonSeconds.
		// No expected conversion value provided for NanonewtonSeconds, verifying result is not NaN.
		result := a.NanonewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonSeconds.
		// No expected conversion value provided for MicronewtonSeconds, verifying result is not NaN.
		result := a.MicronewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonSeconds.
		// No expected conversion value provided for MillinewtonSeconds, verifying result is not NaN.
		result := a.MillinewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonSeconds.
		// No expected conversion value provided for CentinewtonSeconds, verifying result is not NaN.
		result := a.CentinewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonSeconds.
		// No expected conversion value provided for DecinewtonSeconds, verifying result is not NaN.
		result := a.DecinewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonSeconds.
		// No expected conversion value provided for DecanewtonSeconds, verifying result is not NaN.
		result := a.DecanewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonSeconds.
		// No expected conversion value provided for KilonewtonSeconds, verifying result is not NaN.
		result := a.KilonewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonSeconds.
		// No expected conversion value provided for MeganewtonSeconds, verifying result is not NaN.
		result := a.MeganewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonSeconds returned NaN")
		}
	}
}

func TestImpulse_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ImpulseFactory{}
	a, err := factory.CreateImpulse(90, units.ImpulseNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ImpulseNewtonSecond {
		t.Errorf("expected default unit NewtonSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ImpulseKilogramMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ImpulseDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ImpulseNewtonSecond {
		t.Errorf("expected unit NewtonSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestImpulseToString(t *testing.T) {
	factory := units.ImpulseFactory{}
	a, err := factory.CreateImpulse(45, units.ImpulseNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ImpulseNewtonSecond, 2)
	expected := "45.00 " + units.GetImpulseAbbreviation(units.ImpulseNewtonSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ImpulseNewtonSecond, -1)
	expected = "45 " + units.GetImpulseAbbreviation(units.ImpulseNewtonSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestImpulse_EqualityAndComparison(t *testing.T) {
	factory := units.ImpulseFactory{}
	a1, _ := factory.CreateImpulse(60, units.ImpulseNewtonSecond)
	a2, _ := factory.CreateImpulse(60, units.ImpulseNewtonSecond)
	a3, _ := factory.CreateImpulse(120, units.ImpulseNewtonSecond)

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

func TestImpulse_Arithmetic(t *testing.T) {
	factory := units.ImpulseFactory{}
	a1, _ := factory.CreateImpulse(30, units.ImpulseNewtonSecond)
	a2, _ := factory.CreateImpulse(45, units.ImpulseNewtonSecond)

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