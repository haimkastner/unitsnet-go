// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestCompressibilityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "InversePascal"}`
	
	factory := units.CompressibilityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.CompressibilityInversePascal {
		t.Errorf("expected unit %v, got %v", units.CompressibilityInversePascal, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "InversePascal"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestCompressibilityDto_ToJSON(t *testing.T) {
	dto := units.CompressibilityDto{
		Value: 45,
		Unit:  units.CompressibilityInversePascal,
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
	if result["unit"].(string) != string(units.CompressibilityInversePascal) {
		t.Errorf("expected unit %s, got %v", units.CompressibilityInversePascal, result["unit"])
	}
}

func TestNewCompressibility_InvalidValue(t *testing.T) {
	factory := units.CompressibilityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateCompressibility(math.NaN(), units.CompressibilityInversePascal)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateCompressibility(math.Inf(1), units.CompressibilityInversePascal)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestCompressibilityConversions(t *testing.T) {
	factory := units.CompressibilityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateCompressibility(180, units.CompressibilityInversePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to InversePascals.
		// No expected conversion value provided for InversePascals, verifying result is not NaN.
		result := a.InversePascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to InversePascals returned NaN")
		}
	}
	{
		// Test conversion to InverseKilopascals.
		// No expected conversion value provided for InverseKilopascals, verifying result is not NaN.
		result := a.InverseKilopascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseKilopascals returned NaN")
		}
	}
	{
		// Test conversion to InverseMegapascals.
		// No expected conversion value provided for InverseMegapascals, verifying result is not NaN.
		result := a.InverseMegapascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMegapascals returned NaN")
		}
	}
	{
		// Test conversion to InverseAtmospheres.
		// No expected conversion value provided for InverseAtmospheres, verifying result is not NaN.
		result := a.InverseAtmospheres()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseAtmospheres returned NaN")
		}
	}
	{
		// Test conversion to InverseMillibars.
		// No expected conversion value provided for InverseMillibars, verifying result is not NaN.
		result := a.InverseMillibars()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMillibars returned NaN")
		}
	}
	{
		// Test conversion to InverseBars.
		// No expected conversion value provided for InverseBars, verifying result is not NaN.
		result := a.InverseBars()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseBars returned NaN")
		}
	}
	{
		// Test conversion to InversePoundsForcePerSquareInch.
		// No expected conversion value provided for InversePoundsForcePerSquareInch, verifying result is not NaN.
		result := a.InversePoundsForcePerSquareInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to InversePoundsForcePerSquareInch returned NaN")
		}
	}
}

func TestCompressibility_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.CompressibilityFactory{}
	a, err := factory.CreateCompressibility(90, units.CompressibilityInversePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.CompressibilityInversePascal {
		t.Errorf("expected default unit InversePascal, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.CompressibilityInversePascal
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.CompressibilityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.CompressibilityInversePascal {
		t.Errorf("expected unit InversePascal, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestCompressibilityToString(t *testing.T) {
	factory := units.CompressibilityFactory{}
	a, err := factory.CreateCompressibility(45, units.CompressibilityInversePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.CompressibilityInversePascal, 2)
	expected := "45.00 " + units.GetCompressibilityAbbreviation(units.CompressibilityInversePascal)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.CompressibilityInversePascal, -1)
	expected = "45 " + units.GetCompressibilityAbbreviation(units.CompressibilityInversePascal)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestCompressibility_EqualityAndComparison(t *testing.T) {
	factory := units.CompressibilityFactory{}
	a1, _ := factory.CreateCompressibility(60, units.CompressibilityInversePascal)
	a2, _ := factory.CreateCompressibility(60, units.CompressibilityInversePascal)
	a3, _ := factory.CreateCompressibility(120, units.CompressibilityInversePascal)

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

func TestCompressibility_Arithmetic(t *testing.T) {
	factory := units.CompressibilityFactory{}
	a1, _ := factory.CreateCompressibility(30, units.CompressibilityInversePascal)
	a2, _ := factory.CreateCompressibility(45, units.CompressibilityInversePascal)

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