// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLuminanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CandelaPerSquareMeter"}`
	
	factory := units.LuminanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LuminanceCandelaPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.LuminanceCandelaPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CandelaPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLuminanceDto_ToJSON(t *testing.T) {
	dto := units.LuminanceDto{
		Value: 45,
		Unit:  units.LuminanceCandelaPerSquareMeter,
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
	if result["unit"].(string) != string(units.LuminanceCandelaPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.LuminanceCandelaPerSquareMeter, result["unit"])
	}
}

func TestNewLuminance_InvalidValue(t *testing.T) {
	factory := units.LuminanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLuminance(math.NaN(), units.LuminanceCandelaPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLuminance(math.Inf(1), units.LuminanceCandelaPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLuminanceConversions(t *testing.T) {
	factory := units.LuminanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLuminance(180, units.LuminanceCandelaPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CandelasPerSquareMeter.
		// No expected conversion value provided for CandelasPerSquareMeter, verifying result is not NaN.
		result := a.CandelasPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to CandelasPerSquareFoot.
		// No expected conversion value provided for CandelasPerSquareFoot, verifying result is not NaN.
		result := a.CandelasPerSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to CandelasPerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to CandelasPerSquareInch.
		// No expected conversion value provided for CandelasPerSquareInch, verifying result is not NaN.
		result := a.CandelasPerSquareInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to CandelasPerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to Nits.
		// No expected conversion value provided for Nits, verifying result is not NaN.
		result := a.Nits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nits returned NaN")
		}
	}
	{
		// Test conversion to NanocandelasPerSquareMeter.
		// No expected conversion value provided for NanocandelasPerSquareMeter, verifying result is not NaN.
		result := a.NanocandelasPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanocandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrocandelasPerSquareMeter.
		// No expected conversion value provided for MicrocandelasPerSquareMeter, verifying result is not NaN.
		result := a.MicrocandelasPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrocandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MillicandelasPerSquareMeter.
		// No expected conversion value provided for MillicandelasPerSquareMeter, verifying result is not NaN.
		result := a.MillicandelasPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillicandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to CenticandelasPerSquareMeter.
		// No expected conversion value provided for CenticandelasPerSquareMeter, verifying result is not NaN.
		result := a.CenticandelasPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CenticandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to DecicandelasPerSquareMeter.
		// No expected conversion value provided for DecicandelasPerSquareMeter, verifying result is not NaN.
		result := a.DecicandelasPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecicandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilocandelasPerSquareMeter.
		// No expected conversion value provided for KilocandelasPerSquareMeter, verifying result is not NaN.
		result := a.KilocandelasPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocandelasPerSquareMeter returned NaN")
		}
	}
}

func TestLuminance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LuminanceFactory{}
	a, err := factory.CreateLuminance(90, units.LuminanceCandelaPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LuminanceCandelaPerSquareMeter {
		t.Errorf("expected default unit CandelaPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LuminanceCandelaPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LuminanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LuminanceCandelaPerSquareMeter {
		t.Errorf("expected unit CandelaPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLuminanceToString(t *testing.T) {
	factory := units.LuminanceFactory{}
	a, err := factory.CreateLuminance(45, units.LuminanceCandelaPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LuminanceCandelaPerSquareMeter, 2)
	expected := "45.00 " + units.GetLuminanceAbbreviation(units.LuminanceCandelaPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LuminanceCandelaPerSquareMeter, -1)
	expected = "45 " + units.GetLuminanceAbbreviation(units.LuminanceCandelaPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLuminance_EqualityAndComparison(t *testing.T) {
	factory := units.LuminanceFactory{}
	a1, _ := factory.CreateLuminance(60, units.LuminanceCandelaPerSquareMeter)
	a2, _ := factory.CreateLuminance(60, units.LuminanceCandelaPerSquareMeter)
	a3, _ := factory.CreateLuminance(120, units.LuminanceCandelaPerSquareMeter)

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

func TestLuminance_Arithmetic(t *testing.T) {
	factory := units.LuminanceFactory{}
	a1, _ := factory.CreateLuminance(30, units.LuminanceCandelaPerSquareMeter)
	a2, _ := factory.CreateLuminance(45, units.LuminanceCandelaPerSquareMeter)

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