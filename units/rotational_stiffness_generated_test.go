// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRotationalStiffnessDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonMeterPerRadian"}`
	
	factory := units.RotationalStiffnessDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RotationalStiffnessNewtonMeterPerRadian {
		t.Errorf("expected unit %v, got %v", units.RotationalStiffnessNewtonMeterPerRadian, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonMeterPerRadian"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRotationalStiffnessDto_ToJSON(t *testing.T) {
	dto := units.RotationalStiffnessDto{
		Value: 45,
		Unit:  units.RotationalStiffnessNewtonMeterPerRadian,
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
	if result["unit"].(string) != string(units.RotationalStiffnessNewtonMeterPerRadian) {
		t.Errorf("expected unit %s, got %v", units.RotationalStiffnessNewtonMeterPerRadian, result["unit"])
	}
}

func TestNewRotationalStiffness_InvalidValue(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRotationalStiffness(math.NaN(), units.RotationalStiffnessNewtonMeterPerRadian)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRotationalStiffness(math.Inf(1), units.RotationalStiffnessNewtonMeterPerRadian)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRotationalStiffnessConversions(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRotationalStiffness(180, units.RotationalStiffnessNewtonMeterPerRadian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonMetersPerRadian.
		// No expected conversion value provided for NewtonMetersPerRadian, verifying result is not NaN.
		result := a.NewtonMetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeetPerDegrees.
		// No expected conversion value provided for PoundForceFeetPerDegrees, verifying result is not NaN.
		result := a.PoundForceFeetPerDegrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceFeetPerDegrees returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceFeetPerDegrees.
		// No expected conversion value provided for KilopoundForceFeetPerDegrees, verifying result is not NaN.
		result := a.KilopoundForceFeetPerDegrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundForceFeetPerDegrees returned NaN")
		}
	}
	{
		// Test conversion to NewtonMillimetersPerDegree.
		// No expected conversion value provided for NewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.NewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to NewtonMetersPerDegree.
		// No expected conversion value provided for NewtonMetersPerDegree, verifying result is not NaN.
		result := a.NewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to NewtonMillimetersPerRadian.
		// No expected conversion value provided for NewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.NewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeetPerRadian.
		// No expected conversion value provided for PoundForceFeetPerRadian, verifying result is not NaN.
		result := a.PoundForceFeetPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceFeetPerRadian returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMetersPerRadian.
		// No expected conversion value provided for KilonewtonMetersPerRadian, verifying result is not NaN.
		result := a.KilonewtonMetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMetersPerRadian.
		// No expected conversion value provided for MeganewtonMetersPerRadian, verifying result is not NaN.
		result := a.MeganewtonMetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonMillimetersPerDegree.
		// No expected conversion value provided for NanonewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.NanonewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonMillimetersPerDegree.
		// No expected conversion value provided for MicronewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.MicronewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonMillimetersPerDegree.
		// No expected conversion value provided for MillinewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.MillinewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonMillimetersPerDegree.
		// No expected conversion value provided for CentinewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.CentinewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonMillimetersPerDegree.
		// No expected conversion value provided for DecinewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.DecinewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonMillimetersPerDegree.
		// No expected conversion value provided for DecanewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.DecanewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMillimetersPerDegree.
		// No expected conversion value provided for KilonewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.KilonewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMillimetersPerDegree.
		// No expected conversion value provided for MeganewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.MeganewtonMillimetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonMetersPerDegree.
		// No expected conversion value provided for NanonewtonMetersPerDegree, verifying result is not NaN.
		result := a.NanonewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonMetersPerDegree.
		// No expected conversion value provided for MicronewtonMetersPerDegree, verifying result is not NaN.
		result := a.MicronewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonMetersPerDegree.
		// No expected conversion value provided for MillinewtonMetersPerDegree, verifying result is not NaN.
		result := a.MillinewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonMetersPerDegree.
		// No expected conversion value provided for CentinewtonMetersPerDegree, verifying result is not NaN.
		result := a.CentinewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonMetersPerDegree.
		// No expected conversion value provided for DecinewtonMetersPerDegree, verifying result is not NaN.
		result := a.DecinewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonMetersPerDegree.
		// No expected conversion value provided for DecanewtonMetersPerDegree, verifying result is not NaN.
		result := a.DecanewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMetersPerDegree.
		// No expected conversion value provided for KilonewtonMetersPerDegree, verifying result is not NaN.
		result := a.KilonewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMetersPerDegree.
		// No expected conversion value provided for MeganewtonMetersPerDegree, verifying result is not NaN.
		result := a.MeganewtonMetersPerDegree()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonMillimetersPerRadian.
		// No expected conversion value provided for NanonewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.NanonewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonMillimetersPerRadian.
		// No expected conversion value provided for MicronewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.MicronewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonMillimetersPerRadian.
		// No expected conversion value provided for MillinewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.MillinewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonMillimetersPerRadian.
		// No expected conversion value provided for CentinewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.CentinewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonMillimetersPerRadian.
		// No expected conversion value provided for DecinewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.DecinewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonMillimetersPerRadian.
		// No expected conversion value provided for DecanewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.DecanewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMillimetersPerRadian.
		// No expected conversion value provided for KilonewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.KilonewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMillimetersPerRadian.
		// No expected conversion value provided for MeganewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.MeganewtonMillimetersPerRadian()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMillimetersPerRadian returned NaN")
		}
	}
}

func TestRotationalStiffness_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	a, err := factory.CreateRotationalStiffness(90, units.RotationalStiffnessNewtonMeterPerRadian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RotationalStiffnessNewtonMeterPerRadian {
		t.Errorf("expected default unit NewtonMeterPerRadian, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RotationalStiffnessNewtonMeterPerRadian
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RotationalStiffnessDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RotationalStiffnessNewtonMeterPerRadian {
		t.Errorf("expected unit NewtonMeterPerRadian, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRotationalStiffnessToString(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	a, err := factory.CreateRotationalStiffness(45, units.RotationalStiffnessNewtonMeterPerRadian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RotationalStiffnessNewtonMeterPerRadian, 2)
	expected := "45.00 " + units.GetRotationalStiffnessAbbreviation(units.RotationalStiffnessNewtonMeterPerRadian)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RotationalStiffnessNewtonMeterPerRadian, -1)
	expected = "45 " + units.GetRotationalStiffnessAbbreviation(units.RotationalStiffnessNewtonMeterPerRadian)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRotationalStiffness_EqualityAndComparison(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	a1, _ := factory.CreateRotationalStiffness(60, units.RotationalStiffnessNewtonMeterPerRadian)
	a2, _ := factory.CreateRotationalStiffness(60, units.RotationalStiffnessNewtonMeterPerRadian)
	a3, _ := factory.CreateRotationalStiffness(120, units.RotationalStiffnessNewtonMeterPerRadian)

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

func TestRotationalStiffness_Arithmetic(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	a1, _ := factory.CreateRotationalStiffness(30, units.RotationalStiffnessNewtonMeterPerRadian)
	a2, _ := factory.CreateRotationalStiffness(45, units.RotationalStiffnessNewtonMeterPerRadian)

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