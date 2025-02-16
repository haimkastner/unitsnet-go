// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRotationalAccelerationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "RadianPerSecondSquared"}`
	
	factory := units.RotationalAccelerationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RotationalAccelerationRadianPerSecondSquared {
		t.Errorf("expected unit %v, got %v", units.RotationalAccelerationRadianPerSecondSquared, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "RadianPerSecondSquared"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRotationalAccelerationDto_ToJSON(t *testing.T) {
	dto := units.RotationalAccelerationDto{
		Value: 45,
		Unit:  units.RotationalAccelerationRadianPerSecondSquared,
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
	if result["unit"].(string) != string(units.RotationalAccelerationRadianPerSecondSquared) {
		t.Errorf("expected unit %s, got %v", units.RotationalAccelerationRadianPerSecondSquared, result["unit"])
	}
}

func TestNewRotationalAcceleration_InvalidValue(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRotationalAcceleration(math.NaN(), units.RotationalAccelerationRadianPerSecondSquared)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRotationalAcceleration(math.Inf(1), units.RotationalAccelerationRadianPerSecondSquared)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRotationalAccelerationConversions(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRotationalAcceleration(180, units.RotationalAccelerationRadianPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to RadiansPerSecondSquared.
		// No expected conversion value provided for RadiansPerSecondSquared, verifying result is not NaN.
		result := a.RadiansPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to RadiansPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to DegreesPerSecondSquared.
		// No expected conversion value provided for DegreesPerSecondSquared, verifying result is not NaN.
		result := a.DegreesPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to RevolutionsPerMinutePerSecond.
		// No expected conversion value provided for RevolutionsPerMinutePerSecond, verifying result is not NaN.
		result := a.RevolutionsPerMinutePerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to RevolutionsPerMinutePerSecond returned NaN")
		}
	}
	{
		// Test conversion to RevolutionsPerSecondSquared.
		// No expected conversion value provided for RevolutionsPerSecondSquared, verifying result is not NaN.
		result := a.RevolutionsPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to RevolutionsPerSecondSquared returned NaN")
		}
	}
}

func TestRotationalAcceleration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	a, err := factory.CreateRotationalAcceleration(90, units.RotationalAccelerationRadianPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RotationalAccelerationRadianPerSecondSquared {
		t.Errorf("expected default unit RadianPerSecondSquared, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RotationalAccelerationRadianPerSecondSquared
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RotationalAccelerationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RotationalAccelerationRadianPerSecondSquared {
		t.Errorf("expected unit RadianPerSecondSquared, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRotationalAccelerationToString(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	a, err := factory.CreateRotationalAcceleration(45, units.RotationalAccelerationRadianPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RotationalAccelerationRadianPerSecondSquared, 2)
	expected := "45.00 " + units.GetRotationalAccelerationAbbreviation(units.RotationalAccelerationRadianPerSecondSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RotationalAccelerationRadianPerSecondSquared, -1)
	expected = "45 " + units.GetRotationalAccelerationAbbreviation(units.RotationalAccelerationRadianPerSecondSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRotationalAcceleration_EqualityAndComparison(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	a1, _ := factory.CreateRotationalAcceleration(60, units.RotationalAccelerationRadianPerSecondSquared)
	a2, _ := factory.CreateRotationalAcceleration(60, units.RotationalAccelerationRadianPerSecondSquared)
	a3, _ := factory.CreateRotationalAcceleration(120, units.RotationalAccelerationRadianPerSecondSquared)

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

func TestRotationalAcceleration_Arithmetic(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	a1, _ := factory.CreateRotationalAcceleration(30, units.RotationalAccelerationRadianPerSecondSquared)
	a2, _ := factory.CreateRotationalAcceleration(45, units.RotationalAccelerationRadianPerSecondSquared)

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