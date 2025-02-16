// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPorousMediumPermeabilityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeter"}`
	
	factory := units.PorousMediumPermeabilityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PorousMediumPermeabilitySquareMeter {
		t.Errorf("expected unit %v, got %v", units.PorousMediumPermeabilitySquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPorousMediumPermeabilityDto_ToJSON(t *testing.T) {
	dto := units.PorousMediumPermeabilityDto{
		Value: 45,
		Unit:  units.PorousMediumPermeabilitySquareMeter,
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
	if result["unit"].(string) != string(units.PorousMediumPermeabilitySquareMeter) {
		t.Errorf("expected unit %s, got %v", units.PorousMediumPermeabilitySquareMeter, result["unit"])
	}
}

func TestNewPorousMediumPermeability_InvalidValue(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePorousMediumPermeability(math.NaN(), units.PorousMediumPermeabilitySquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePorousMediumPermeability(math.Inf(1), units.PorousMediumPermeabilitySquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPorousMediumPermeabilityConversions(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePorousMediumPermeability(180, units.PorousMediumPermeabilitySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Darcys.
		// No expected conversion value provided for Darcys, verifying result is not NaN.
		result := a.Darcys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Darcys returned NaN")
		}
	}
	{
		// Test conversion to SquareMeters.
		// No expected conversion value provided for SquareMeters, verifying result is not NaN.
		result := a.SquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeters returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeters.
		// No expected conversion value provided for SquareCentimeters, verifying result is not NaN.
		result := a.SquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to Microdarcys.
		// No expected conversion value provided for Microdarcys, verifying result is not NaN.
		result := a.Microdarcys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microdarcys returned NaN")
		}
	}
	{
		// Test conversion to Millidarcys.
		// No expected conversion value provided for Millidarcys, verifying result is not NaN.
		result := a.Millidarcys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millidarcys returned NaN")
		}
	}
}

func TestPorousMediumPermeability_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	a, err := factory.CreatePorousMediumPermeability(90, units.PorousMediumPermeabilitySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PorousMediumPermeabilitySquareMeter {
		t.Errorf("expected default unit SquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PorousMediumPermeabilityDarcy
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PorousMediumPermeabilityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PorousMediumPermeabilitySquareMeter {
		t.Errorf("expected unit SquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPorousMediumPermeabilityToString(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	a, err := factory.CreatePorousMediumPermeability(45, units.PorousMediumPermeabilitySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PorousMediumPermeabilitySquareMeter, 2)
	expected := "45.00 " + units.GetPorousMediumPermeabilityAbbreviation(units.PorousMediumPermeabilitySquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PorousMediumPermeabilitySquareMeter, -1)
	expected = "45 " + units.GetPorousMediumPermeabilityAbbreviation(units.PorousMediumPermeabilitySquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPorousMediumPermeability_EqualityAndComparison(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	a1, _ := factory.CreatePorousMediumPermeability(60, units.PorousMediumPermeabilitySquareMeter)
	a2, _ := factory.CreatePorousMediumPermeability(60, units.PorousMediumPermeabilitySquareMeter)
	a3, _ := factory.CreatePorousMediumPermeability(120, units.PorousMediumPermeabilitySquareMeter)

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

func TestPorousMediumPermeability_Arithmetic(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	a1, _ := factory.CreatePorousMediumPermeability(30, units.PorousMediumPermeabilitySquareMeter)
	a2, _ := factory.CreatePorousMediumPermeability(45, units.PorousMediumPermeabilitySquareMeter)

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