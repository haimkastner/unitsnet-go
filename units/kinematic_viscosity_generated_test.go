// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestKinematicViscosityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeterPerSecond"}`
	
	factory := units.KinematicViscosityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.KinematicViscositySquareMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.KinematicViscositySquareMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestKinematicViscosityDto_ToJSON(t *testing.T) {
	dto := units.KinematicViscosityDto{
		Value: 45,
		Unit:  units.KinematicViscositySquareMeterPerSecond,
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
	if result["unit"].(string) != string(units.KinematicViscositySquareMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.KinematicViscositySquareMeterPerSecond, result["unit"])
	}
}

func TestNewKinematicViscosity_InvalidValue(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateKinematicViscosity(math.NaN(), units.KinematicViscositySquareMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateKinematicViscosity(math.Inf(1), units.KinematicViscositySquareMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestKinematicViscosityConversions(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateKinematicViscosity(180, units.KinematicViscositySquareMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SquareMetersPerSecond.
		// No expected conversion value provided for SquareMetersPerSecond, verifying result is not NaN.
		result := a.SquareMetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to Stokes.
		// No expected conversion value provided for Stokes, verifying result is not NaN.
		result := a.Stokes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Stokes returned NaN")
		}
	}
	{
		// Test conversion to SquareFeetPerSecond.
		// No expected conversion value provided for SquareFeetPerSecond, verifying result is not NaN.
		result := a.SquareFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to Nanostokes.
		// No expected conversion value provided for Nanostokes, verifying result is not NaN.
		result := a.Nanostokes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanostokes returned NaN")
		}
	}
	{
		// Test conversion to Microstokes.
		// No expected conversion value provided for Microstokes, verifying result is not NaN.
		result := a.Microstokes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microstokes returned NaN")
		}
	}
	{
		// Test conversion to Millistokes.
		// No expected conversion value provided for Millistokes, verifying result is not NaN.
		result := a.Millistokes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millistokes returned NaN")
		}
	}
	{
		// Test conversion to Centistokes.
		// No expected conversion value provided for Centistokes, verifying result is not NaN.
		result := a.Centistokes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centistokes returned NaN")
		}
	}
	{
		// Test conversion to Decistokes.
		// No expected conversion value provided for Decistokes, verifying result is not NaN.
		result := a.Decistokes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decistokes returned NaN")
		}
	}
	{
		// Test conversion to Kilostokes.
		// No expected conversion value provided for Kilostokes, verifying result is not NaN.
		result := a.Kilostokes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilostokes returned NaN")
		}
	}
}

func TestKinematicViscosity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	a, err := factory.CreateKinematicViscosity(90, units.KinematicViscositySquareMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.KinematicViscositySquareMeterPerSecond {
		t.Errorf("expected default unit SquareMeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.KinematicViscositySquareMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.KinematicViscosityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.KinematicViscositySquareMeterPerSecond {
		t.Errorf("expected unit SquareMeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestKinematicViscosityToString(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	a, err := factory.CreateKinematicViscosity(45, units.KinematicViscositySquareMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.KinematicViscositySquareMeterPerSecond, 2)
	expected := "45.00 " + units.GetKinematicViscosityAbbreviation(units.KinematicViscositySquareMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.KinematicViscositySquareMeterPerSecond, -1)
	expected = "45 " + units.GetKinematicViscosityAbbreviation(units.KinematicViscositySquareMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestKinematicViscosity_EqualityAndComparison(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	a1, _ := factory.CreateKinematicViscosity(60, units.KinematicViscositySquareMeterPerSecond)
	a2, _ := factory.CreateKinematicViscosity(60, units.KinematicViscositySquareMeterPerSecond)
	a3, _ := factory.CreateKinematicViscosity(120, units.KinematicViscositySquareMeterPerSecond)

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

func TestKinematicViscosity_Arithmetic(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	a1, _ := factory.CreateKinematicViscosity(30, units.KinematicViscositySquareMeterPerSecond)
	a2, _ := factory.CreateKinematicViscosity(45, units.KinematicViscositySquareMeterPerSecond)

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