// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMagneticFluxDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Weber"}`
	
	factory := units.MagneticFluxDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MagneticFluxWeber {
		t.Errorf("expected unit %v, got %v", units.MagneticFluxWeber, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Weber"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMagneticFluxDto_ToJSON(t *testing.T) {
	dto := units.MagneticFluxDto{
		Value: 45,
		Unit:  units.MagneticFluxWeber,
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
	if result["unit"].(string) != string(units.MagneticFluxWeber) {
		t.Errorf("expected unit %s, got %v", units.MagneticFluxWeber, result["unit"])
	}
}

func TestNewMagneticFlux_InvalidValue(t *testing.T) {
	factory := units.MagneticFluxFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMagneticFlux(math.NaN(), units.MagneticFluxWeber)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMagneticFlux(math.Inf(1), units.MagneticFluxWeber)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMagneticFluxConversions(t *testing.T) {
	factory := units.MagneticFluxFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMagneticFlux(180, units.MagneticFluxWeber)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Webers.
		// No expected conversion value provided for Webers, verifying result is not NaN.
		result := a.Webers()
		if math.IsNaN(result) {
			t.Errorf("conversion to Webers returned NaN")
		}
	}
}

func TestMagneticFlux_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MagneticFluxFactory{}
	a, err := factory.CreateMagneticFlux(90, units.MagneticFluxWeber)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MagneticFluxWeber {
		t.Errorf("expected default unit Weber, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MagneticFluxWeber
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MagneticFluxDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MagneticFluxWeber {
		t.Errorf("expected unit Weber, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMagneticFluxToString(t *testing.T) {
	factory := units.MagneticFluxFactory{}
	a, err := factory.CreateMagneticFlux(45, units.MagneticFluxWeber)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MagneticFluxWeber, 2)
	expected := "45.00 " + units.GetMagneticFluxAbbreviation(units.MagneticFluxWeber)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MagneticFluxWeber, -1)
	expected = "45 " + units.GetMagneticFluxAbbreviation(units.MagneticFluxWeber)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMagneticFlux_EqualityAndComparison(t *testing.T) {
	factory := units.MagneticFluxFactory{}
	a1, _ := factory.CreateMagneticFlux(60, units.MagneticFluxWeber)
	a2, _ := factory.CreateMagneticFlux(60, units.MagneticFluxWeber)
	a3, _ := factory.CreateMagneticFlux(120, units.MagneticFluxWeber)

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

func TestMagneticFlux_Arithmetic(t *testing.T) {
	factory := units.MagneticFluxFactory{}
	a1, _ := factory.CreateMagneticFlux(30, units.MagneticFluxWeber)
	a2, _ := factory.CreateMagneticFlux(45, units.MagneticFluxWeber)

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