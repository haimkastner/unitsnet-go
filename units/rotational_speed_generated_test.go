// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRotationalSpeedDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "RadianPerSecond"}`
	
	factory := units.RotationalSpeedDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RotationalSpeedRadianPerSecond {
		t.Errorf("expected unit %v, got %v", units.RotationalSpeedRadianPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "RadianPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRotationalSpeedDto_ToJSON(t *testing.T) {
	dto := units.RotationalSpeedDto{
		Value: 45,
		Unit:  units.RotationalSpeedRadianPerSecond,
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
	if result["unit"].(string) != string(units.RotationalSpeedRadianPerSecond) {
		t.Errorf("expected unit %s, got %v", units.RotationalSpeedRadianPerSecond, result["unit"])
	}
}

func TestNewRotationalSpeed_InvalidValue(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRotationalSpeed(math.NaN(), units.RotationalSpeedRadianPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRotationalSpeed(math.Inf(1), units.RotationalSpeedRadianPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRotationalSpeedConversions(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRotationalSpeed(180, units.RotationalSpeedRadianPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to RadiansPerSecond.
		// No expected conversion value provided for RadiansPerSecond, verifying result is not NaN.
		result := a.RadiansPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to RadiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesPerSecond.
		// No expected conversion value provided for DegreesPerSecond, verifying result is not NaN.
		result := a.DegreesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesPerMinute.
		// No expected conversion value provided for DegreesPerMinute, verifying result is not NaN.
		result := a.DegreesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to RevolutionsPerSecond.
		// No expected conversion value provided for RevolutionsPerSecond, verifying result is not NaN.
		result := a.RevolutionsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to RevolutionsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to RevolutionsPerMinute.
		// No expected conversion value provided for RevolutionsPerMinute, verifying result is not NaN.
		result := a.RevolutionsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to RevolutionsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to NanoradiansPerSecond.
		// No expected conversion value provided for NanoradiansPerSecond, verifying result is not NaN.
		result := a.NanoradiansPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanoradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicroradiansPerSecond.
		// No expected conversion value provided for MicroradiansPerSecond, verifying result is not NaN.
		result := a.MicroradiansPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicroradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MilliradiansPerSecond.
		// No expected conversion value provided for MilliradiansPerSecond, verifying result is not NaN.
		result := a.MilliradiansPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentiradiansPerSecond.
		// No expected conversion value provided for CentiradiansPerSecond, verifying result is not NaN.
		result := a.CentiradiansPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentiradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DeciradiansPerSecond.
		// No expected conversion value provided for DeciradiansPerSecond, verifying result is not NaN.
		result := a.DeciradiansPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DeciradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanodegreesPerSecond.
		// No expected conversion value provided for NanodegreesPerSecond, verifying result is not NaN.
		result := a.NanodegreesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanodegreesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrodegreesPerSecond.
		// No expected conversion value provided for MicrodegreesPerSecond, verifying result is not NaN.
		result := a.MicrodegreesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrodegreesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillidegreesPerSecond.
		// No expected conversion value provided for MillidegreesPerSecond, verifying result is not NaN.
		result := a.MillidegreesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillidegreesPerSecond returned NaN")
		}
	}
}

func TestRotationalSpeed_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	a, err := factory.CreateRotationalSpeed(90, units.RotationalSpeedRadianPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RotationalSpeedRadianPerSecond {
		t.Errorf("expected default unit RadianPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RotationalSpeedRadianPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RotationalSpeedDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RotationalSpeedRadianPerSecond {
		t.Errorf("expected unit RadianPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRotationalSpeedToString(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	a, err := factory.CreateRotationalSpeed(45, units.RotationalSpeedRadianPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RotationalSpeedRadianPerSecond, 2)
	expected := "45.00 " + units.GetRotationalSpeedAbbreviation(units.RotationalSpeedRadianPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RotationalSpeedRadianPerSecond, -1)
	expected = "45 " + units.GetRotationalSpeedAbbreviation(units.RotationalSpeedRadianPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRotationalSpeed_EqualityAndComparison(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	a1, _ := factory.CreateRotationalSpeed(60, units.RotationalSpeedRadianPerSecond)
	a2, _ := factory.CreateRotationalSpeed(60, units.RotationalSpeedRadianPerSecond)
	a3, _ := factory.CreateRotationalSpeed(120, units.RotationalSpeedRadianPerSecond)

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

func TestRotationalSpeed_Arithmetic(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	a1, _ := factory.CreateRotationalSpeed(30, units.RotationalSpeedRadianPerSecond)
	a2, _ := factory.CreateRotationalSpeed(45, units.RotationalSpeedRadianPerSecond)

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