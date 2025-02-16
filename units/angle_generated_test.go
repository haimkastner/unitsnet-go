// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAngleDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Degree"}`
	
	factory := units.AngleDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AngleDegree {
		t.Errorf("expected unit %v, got %v", units.AngleDegree, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Degree"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAngleDto_ToJSON(t *testing.T) {
	dto := units.AngleDto{
		Value: 45,
		Unit:  units.AngleDegree,
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
	if result["unit"].(string) != string(units.AngleDegree) {
		t.Errorf("expected unit %s, got %v", units.AngleDegree, result["unit"])
	}
}

func TestNewAngle_InvalidValue(t *testing.T) {
	factory := units.AngleFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAngle(math.NaN(), units.AngleDegree)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAngle(math.Inf(1), units.AngleDegree)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAngleConversions(t *testing.T) {
	factory := units.AngleFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAngle(180, units.AngleDegree)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Radians.
		// No expected conversion value provided for Radians, verifying result is not NaN.
		result := a.Radians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Radians returned NaN")
		}
	}
	{
		// Test conversion to Degrees.
		// No expected conversion value provided for Degrees, verifying result is not NaN.
		result := a.Degrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to Degrees returned NaN")
		}
	}
	{
		// Test conversion to Arcminutes.
		// No expected conversion value provided for Arcminutes, verifying result is not NaN.
		result := a.Arcminutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Arcminutes returned NaN")
		}
	}
	{
		// Test conversion to Arcseconds.
		// No expected conversion value provided for Arcseconds, verifying result is not NaN.
		result := a.Arcseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Arcseconds returned NaN")
		}
	}
	{
		// Test conversion to Gradians.
		// No expected conversion value provided for Gradians, verifying result is not NaN.
		result := a.Gradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gradians returned NaN")
		}
	}
	{
		// Test conversion to NatoMils.
		// No expected conversion value provided for NatoMils, verifying result is not NaN.
		result := a.NatoMils()
		if math.IsNaN(result) {
			t.Errorf("conversion to NatoMils returned NaN")
		}
	}
	{
		// Test conversion to Revolutions.
		// No expected conversion value provided for Revolutions, verifying result is not NaN.
		result := a.Revolutions()
		if math.IsNaN(result) {
			t.Errorf("conversion to Revolutions returned NaN")
		}
	}
	{
		// Test conversion to Tilt.
		// No expected conversion value provided for Tilt, verifying result is not NaN.
		result := a.Tilt()
		if math.IsNaN(result) {
			t.Errorf("conversion to Tilt returned NaN")
		}
	}
	{
		// Test conversion to Nanoradians.
		// No expected conversion value provided for Nanoradians, verifying result is not NaN.
		result := a.Nanoradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanoradians returned NaN")
		}
	}
	{
		// Test conversion to Microradians.
		// No expected conversion value provided for Microradians, verifying result is not NaN.
		result := a.Microradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microradians returned NaN")
		}
	}
	{
		// Test conversion to Milliradians.
		// No expected conversion value provided for Milliradians, verifying result is not NaN.
		result := a.Milliradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliradians returned NaN")
		}
	}
	{
		// Test conversion to Centiradians.
		// No expected conversion value provided for Centiradians, verifying result is not NaN.
		result := a.Centiradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centiradians returned NaN")
		}
	}
	{
		// Test conversion to Deciradians.
		// No expected conversion value provided for Deciradians, verifying result is not NaN.
		result := a.Deciradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Deciradians returned NaN")
		}
	}
	{
		// Test conversion to Nanodegrees.
		// No expected conversion value provided for Nanodegrees, verifying result is not NaN.
		result := a.Nanodegrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanodegrees returned NaN")
		}
	}
	{
		// Test conversion to Microdegrees.
		// No expected conversion value provided for Microdegrees, verifying result is not NaN.
		result := a.Microdegrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microdegrees returned NaN")
		}
	}
	{
		// Test conversion to Millidegrees.
		// No expected conversion value provided for Millidegrees, verifying result is not NaN.
		result := a.Millidegrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millidegrees returned NaN")
		}
	}
}

func TestAngle_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AngleFactory{}
	a, err := factory.CreateAngle(90, units.AngleDegree)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AngleDegree {
		t.Errorf("expected default unit Degree, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AngleRadian
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AngleDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AngleDegree {
		t.Errorf("expected unit Degree, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAngleToString(t *testing.T) {
	factory := units.AngleFactory{}
	a, err := factory.CreateAngle(45, units.AngleDegree)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AngleDegree, 2)
	expected := "45.00 " + units.GetAngleAbbreviation(units.AngleDegree)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AngleDegree, -1)
	expected = "45 " + units.GetAngleAbbreviation(units.AngleDegree)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAngle_EqualityAndComparison(t *testing.T) {
	factory := units.AngleFactory{}
	a1, _ := factory.CreateAngle(60, units.AngleDegree)
	a2, _ := factory.CreateAngle(60, units.AngleDegree)
	a3, _ := factory.CreateAngle(120, units.AngleDegree)

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

func TestAngle_Arithmetic(t *testing.T) {
	factory := units.AngleFactory{}
	a1, _ := factory.CreateAngle(30, units.AngleDegree)
	a2, _ := factory.CreateAngle(45, units.AngleDegree)

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