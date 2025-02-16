// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAbsorbedDoseOfIonizingRadiationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Gray"}`
	
	factory := units.AbsorbedDoseOfIonizingRadiationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AbsorbedDoseOfIonizingRadiationGray {
		t.Errorf("expected unit %v, got %v", units.AbsorbedDoseOfIonizingRadiationGray, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Gray"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAbsorbedDoseOfIonizingRadiationDto_ToJSON(t *testing.T) {
	dto := units.AbsorbedDoseOfIonizingRadiationDto{
		Value: 45,
		Unit:  units.AbsorbedDoseOfIonizingRadiationGray,
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
	if result["unit"].(string) != string(units.AbsorbedDoseOfIonizingRadiationGray) {
		t.Errorf("expected unit %s, got %v", units.AbsorbedDoseOfIonizingRadiationGray, result["unit"])
	}
}

func TestNewAbsorbedDoseOfIonizingRadiation_InvalidValue(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAbsorbedDoseOfIonizingRadiation(math.NaN(), units.AbsorbedDoseOfIonizingRadiationGray)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAbsorbedDoseOfIonizingRadiation(math.Inf(1), units.AbsorbedDoseOfIonizingRadiationGray)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAbsorbedDoseOfIonizingRadiationConversions(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAbsorbedDoseOfIonizingRadiation(180, units.AbsorbedDoseOfIonizingRadiationGray)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Grays.
		// No expected conversion value provided for Grays, verifying result is not NaN.
		result := a.Grays()
		if math.IsNaN(result) {
			t.Errorf("conversion to Grays returned NaN")
		}
	}
	{
		// Test conversion to Rads.
		// No expected conversion value provided for Rads, verifying result is not NaN.
		result := a.Rads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Rads returned NaN")
		}
	}
	{
		// Test conversion to Femtograys.
		// No expected conversion value provided for Femtograys, verifying result is not NaN.
		result := a.Femtograys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Femtograys returned NaN")
		}
	}
	{
		// Test conversion to Picograys.
		// No expected conversion value provided for Picograys, verifying result is not NaN.
		result := a.Picograys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picograys returned NaN")
		}
	}
	{
		// Test conversion to Nanograys.
		// No expected conversion value provided for Nanograys, verifying result is not NaN.
		result := a.Nanograys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanograys returned NaN")
		}
	}
	{
		// Test conversion to Micrograys.
		// No expected conversion value provided for Micrograys, verifying result is not NaN.
		result := a.Micrograys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micrograys returned NaN")
		}
	}
	{
		// Test conversion to Milligrays.
		// No expected conversion value provided for Milligrays, verifying result is not NaN.
		result := a.Milligrays()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milligrays returned NaN")
		}
	}
	{
		// Test conversion to Centigrays.
		// No expected conversion value provided for Centigrays, verifying result is not NaN.
		result := a.Centigrays()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centigrays returned NaN")
		}
	}
	{
		// Test conversion to Kilograys.
		// No expected conversion value provided for Kilograys, verifying result is not NaN.
		result := a.Kilograys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilograys returned NaN")
		}
	}
	{
		// Test conversion to Megagrays.
		// No expected conversion value provided for Megagrays, verifying result is not NaN.
		result := a.Megagrays()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megagrays returned NaN")
		}
	}
	{
		// Test conversion to Gigagrays.
		// No expected conversion value provided for Gigagrays, verifying result is not NaN.
		result := a.Gigagrays()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigagrays returned NaN")
		}
	}
	{
		// Test conversion to Teragrays.
		// No expected conversion value provided for Teragrays, verifying result is not NaN.
		result := a.Teragrays()
		if math.IsNaN(result) {
			t.Errorf("conversion to Teragrays returned NaN")
		}
	}
	{
		// Test conversion to Petagrays.
		// No expected conversion value provided for Petagrays, verifying result is not NaN.
		result := a.Petagrays()
		if math.IsNaN(result) {
			t.Errorf("conversion to Petagrays returned NaN")
		}
	}
	{
		// Test conversion to Millirads.
		// No expected conversion value provided for Millirads, verifying result is not NaN.
		result := a.Millirads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millirads returned NaN")
		}
	}
	{
		// Test conversion to Kilorads.
		// No expected conversion value provided for Kilorads, verifying result is not NaN.
		result := a.Kilorads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilorads returned NaN")
		}
	}
	{
		// Test conversion to Megarads.
		// No expected conversion value provided for Megarads, verifying result is not NaN.
		result := a.Megarads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megarads returned NaN")
		}
	}
}

func TestAbsorbedDoseOfIonizingRadiation_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	a, err := factory.CreateAbsorbedDoseOfIonizingRadiation(90, units.AbsorbedDoseOfIonizingRadiationGray)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AbsorbedDoseOfIonizingRadiationGray {
		t.Errorf("expected default unit Gray, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AbsorbedDoseOfIonizingRadiationGray
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AbsorbedDoseOfIonizingRadiationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AbsorbedDoseOfIonizingRadiationGray {
		t.Errorf("expected unit Gray, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAbsorbedDoseOfIonizingRadiationToString(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	a, err := factory.CreateAbsorbedDoseOfIonizingRadiation(45, units.AbsorbedDoseOfIonizingRadiationGray)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AbsorbedDoseOfIonizingRadiationGray, 2)
	expected := "45.00 " + units.GetAbsorbedDoseOfIonizingRadiationAbbreviation(units.AbsorbedDoseOfIonizingRadiationGray)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AbsorbedDoseOfIonizingRadiationGray, -1)
	expected = "45 " + units.GetAbsorbedDoseOfIonizingRadiationAbbreviation(units.AbsorbedDoseOfIonizingRadiationGray)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAbsorbedDoseOfIonizingRadiation_EqualityAndComparison(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	a1, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(60, units.AbsorbedDoseOfIonizingRadiationGray)
	a2, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(60, units.AbsorbedDoseOfIonizingRadiationGray)
	a3, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(120, units.AbsorbedDoseOfIonizingRadiationGray)

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

func TestAbsorbedDoseOfIonizingRadiation_Arithmetic(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	a1, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(30, units.AbsorbedDoseOfIonizingRadiationGray)
	a2, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(45, units.AbsorbedDoseOfIonizingRadiationGray)

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