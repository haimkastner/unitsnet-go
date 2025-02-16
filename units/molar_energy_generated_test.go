// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerMole"}`
	
	factory := units.MolarEnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarEnergyJoulePerMole {
		t.Errorf("expected unit %v, got %v", units.MolarEnergyJoulePerMole, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerMole"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarEnergyDto_ToJSON(t *testing.T) {
	dto := units.MolarEnergyDto{
		Value: 45,
		Unit:  units.MolarEnergyJoulePerMole,
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
	if result["unit"].(string) != string(units.MolarEnergyJoulePerMole) {
		t.Errorf("expected unit %s, got %v", units.MolarEnergyJoulePerMole, result["unit"])
	}
}

func TestNewMolarEnergy_InvalidValue(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarEnergy(math.NaN(), units.MolarEnergyJoulePerMole)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarEnergy(math.Inf(1), units.MolarEnergyJoulePerMole)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarEnergyConversions(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarEnergy(180, units.MolarEnergyJoulePerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerMole.
		// No expected conversion value provided for JoulesPerMole, verifying result is not NaN.
		result := a.JoulesPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerMole returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerMole.
		// No expected conversion value provided for KilojoulesPerMole, verifying result is not NaN.
		result := a.KilojoulesPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerMole returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerMole.
		// No expected conversion value provided for MegajoulesPerMole, verifying result is not NaN.
		result := a.MegajoulesPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerMole returned NaN")
		}
	}
}

func TestMolarEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	a, err := factory.CreateMolarEnergy(90, units.MolarEnergyJoulePerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarEnergyJoulePerMole {
		t.Errorf("expected default unit JoulePerMole, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarEnergyJoulePerMole
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarEnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarEnergyJoulePerMole {
		t.Errorf("expected unit JoulePerMole, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarEnergyToString(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	a, err := factory.CreateMolarEnergy(45, units.MolarEnergyJoulePerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarEnergyJoulePerMole, 2)
	expected := "45.00 " + units.GetMolarEnergyAbbreviation(units.MolarEnergyJoulePerMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarEnergyJoulePerMole, -1)
	expected = "45 " + units.GetMolarEnergyAbbreviation(units.MolarEnergyJoulePerMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	a1, _ := factory.CreateMolarEnergy(60, units.MolarEnergyJoulePerMole)
	a2, _ := factory.CreateMolarEnergy(60, units.MolarEnergyJoulePerMole)
	a3, _ := factory.CreateMolarEnergy(120, units.MolarEnergyJoulePerMole)

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

func TestMolarEnergy_Arithmetic(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	a1, _ := factory.CreateMolarEnergy(30, units.MolarEnergyJoulePerMole)
	a2, _ := factory.CreateMolarEnergy(45, units.MolarEnergyJoulePerMole)

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