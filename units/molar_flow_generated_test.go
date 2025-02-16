// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarFlowDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MolePerSecond"}`
	
	factory := units.MolarFlowDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarFlowMolePerSecond {
		t.Errorf("expected unit %v, got %v", units.MolarFlowMolePerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MolePerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarFlowDto_ToJSON(t *testing.T) {
	dto := units.MolarFlowDto{
		Value: 45,
		Unit:  units.MolarFlowMolePerSecond,
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
	if result["unit"].(string) != string(units.MolarFlowMolePerSecond) {
		t.Errorf("expected unit %s, got %v", units.MolarFlowMolePerSecond, result["unit"])
	}
}

func TestNewMolarFlow_InvalidValue(t *testing.T) {
	factory := units.MolarFlowFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarFlow(math.NaN(), units.MolarFlowMolePerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarFlow(math.Inf(1), units.MolarFlowMolePerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarFlowConversions(t *testing.T) {
	factory := units.MolarFlowFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarFlow(180, units.MolarFlowMolePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MolesPerSecond.
		// No expected conversion value provided for MolesPerSecond, verifying result is not NaN.
		result := a.MolesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MolesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MolesPerMinute.
		// No expected conversion value provided for MolesPerMinute, verifying result is not NaN.
		result := a.MolesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MolesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MolesPerHour.
		// No expected conversion value provided for MolesPerHour, verifying result is not NaN.
		result := a.MolesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MolesPerHour returned NaN")
		}
	}
	{
		// Test conversion to PoundMolesPerSecond.
		// No expected conversion value provided for PoundMolesPerSecond, verifying result is not NaN.
		result := a.PoundMolesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundMolesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundMolesPerMinute.
		// No expected conversion value provided for PoundMolesPerMinute, verifying result is not NaN.
		result := a.PoundMolesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundMolesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to PoundMolesPerHour.
		// No expected conversion value provided for PoundMolesPerHour, verifying result is not NaN.
		result := a.PoundMolesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundMolesPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilomolesPerSecond.
		// No expected conversion value provided for KilomolesPerSecond, verifying result is not NaN.
		result := a.KilomolesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilomolesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilomolesPerMinute.
		// No expected conversion value provided for KilomolesPerMinute, verifying result is not NaN.
		result := a.KilomolesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilomolesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilomolesPerHour.
		// No expected conversion value provided for KilomolesPerHour, verifying result is not NaN.
		result := a.KilomolesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilomolesPerHour returned NaN")
		}
	}
}

func TestMolarFlow_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarFlowFactory{}
	a, err := factory.CreateMolarFlow(90, units.MolarFlowMolePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarFlowMolePerSecond {
		t.Errorf("expected default unit MolePerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarFlowMolePerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarFlowDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarFlowMolePerSecond {
		t.Errorf("expected unit MolePerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarFlowToString(t *testing.T) {
	factory := units.MolarFlowFactory{}
	a, err := factory.CreateMolarFlow(45, units.MolarFlowMolePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarFlowMolePerSecond, 2)
	expected := "45.00 " + units.GetMolarFlowAbbreviation(units.MolarFlowMolePerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarFlowMolePerSecond, -1)
	expected = "45 " + units.GetMolarFlowAbbreviation(units.MolarFlowMolePerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarFlow_EqualityAndComparison(t *testing.T) {
	factory := units.MolarFlowFactory{}
	a1, _ := factory.CreateMolarFlow(60, units.MolarFlowMolePerSecond)
	a2, _ := factory.CreateMolarFlow(60, units.MolarFlowMolePerSecond)
	a3, _ := factory.CreateMolarFlow(120, units.MolarFlowMolePerSecond)

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

func TestMolarFlow_Arithmetic(t *testing.T) {
	factory := units.MolarFlowFactory{}
	a1, _ := factory.CreateMolarFlow(30, units.MolarFlowMolePerSecond)
	a2, _ := factory.CreateMolarFlow(45, units.MolarFlowMolePerSecond)

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