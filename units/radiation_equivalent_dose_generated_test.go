// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRadiationEquivalentDoseDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Sievert"}`
	
	factory := units.RadiationEquivalentDoseDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RadiationEquivalentDoseSievert {
		t.Errorf("expected unit %v, got %v", units.RadiationEquivalentDoseSievert, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Sievert"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRadiationEquivalentDoseDto_ToJSON(t *testing.T) {
	dto := units.RadiationEquivalentDoseDto{
		Value: 45,
		Unit:  units.RadiationEquivalentDoseSievert,
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
	if result["unit"].(string) != string(units.RadiationEquivalentDoseSievert) {
		t.Errorf("expected unit %s, got %v", units.RadiationEquivalentDoseSievert, result["unit"])
	}
}

func TestNewRadiationEquivalentDose_InvalidValue(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRadiationEquivalentDose(math.NaN(), units.RadiationEquivalentDoseSievert)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRadiationEquivalentDose(math.Inf(1), units.RadiationEquivalentDoseSievert)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRadiationEquivalentDoseConversions(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRadiationEquivalentDose(180, units.RadiationEquivalentDoseSievert)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Sieverts.
		// No expected conversion value provided for Sieverts, verifying result is not NaN.
		result := a.Sieverts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Sieverts returned NaN")
		}
	}
	{
		// Test conversion to RoentgensEquivalentMan.
		// No expected conversion value provided for RoentgensEquivalentMan, verifying result is not NaN.
		result := a.RoentgensEquivalentMan()
		if math.IsNaN(result) {
			t.Errorf("conversion to RoentgensEquivalentMan returned NaN")
		}
	}
	{
		// Test conversion to Nanosieverts.
		// No expected conversion value provided for Nanosieverts, verifying result is not NaN.
		result := a.Nanosieverts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanosieverts returned NaN")
		}
	}
	{
		// Test conversion to Microsieverts.
		// No expected conversion value provided for Microsieverts, verifying result is not NaN.
		result := a.Microsieverts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microsieverts returned NaN")
		}
	}
	{
		// Test conversion to Millisieverts.
		// No expected conversion value provided for Millisieverts, verifying result is not NaN.
		result := a.Millisieverts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millisieverts returned NaN")
		}
	}
	{
		// Test conversion to MilliroentgensEquivalentMan.
		// No expected conversion value provided for MilliroentgensEquivalentMan, verifying result is not NaN.
		result := a.MilliroentgensEquivalentMan()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliroentgensEquivalentMan returned NaN")
		}
	}
}

func TestRadiationEquivalentDose_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	a, err := factory.CreateRadiationEquivalentDose(90, units.RadiationEquivalentDoseSievert)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RadiationEquivalentDoseSievert {
		t.Errorf("expected default unit Sievert, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RadiationEquivalentDoseSievert
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RadiationEquivalentDoseDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RadiationEquivalentDoseSievert {
		t.Errorf("expected unit Sievert, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRadiationEquivalentDoseToString(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	a, err := factory.CreateRadiationEquivalentDose(45, units.RadiationEquivalentDoseSievert)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RadiationEquivalentDoseSievert, 2)
	expected := "45.00 " + units.GetRadiationEquivalentDoseAbbreviation(units.RadiationEquivalentDoseSievert)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RadiationEquivalentDoseSievert, -1)
	expected = "45 " + units.GetRadiationEquivalentDoseAbbreviation(units.RadiationEquivalentDoseSievert)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRadiationEquivalentDose_EqualityAndComparison(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	a1, _ := factory.CreateRadiationEquivalentDose(60, units.RadiationEquivalentDoseSievert)
	a2, _ := factory.CreateRadiationEquivalentDose(60, units.RadiationEquivalentDoseSievert)
	a3, _ := factory.CreateRadiationEquivalentDose(120, units.RadiationEquivalentDoseSievert)

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

func TestRadiationEquivalentDose_Arithmetic(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	a1, _ := factory.CreateRadiationEquivalentDose(30, units.RadiationEquivalentDoseSievert)
	a2, _ := factory.CreateRadiationEquivalentDose(45, units.RadiationEquivalentDoseSievert)

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