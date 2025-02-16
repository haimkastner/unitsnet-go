// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAccelerationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterPerSecondSquared"}`
	
	factory := units.AccelerationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AccelerationMeterPerSecondSquared {
		t.Errorf("expected unit %v, got %v", units.AccelerationMeterPerSecondSquared, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterPerSecondSquared"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAccelerationDto_ToJSON(t *testing.T) {
	dto := units.AccelerationDto{
		Value: 45,
		Unit:  units.AccelerationMeterPerSecondSquared,
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
	if result["unit"].(string) != string(units.AccelerationMeterPerSecondSquared) {
		t.Errorf("expected unit %s, got %v", units.AccelerationMeterPerSecondSquared, result["unit"])
	}
}

func TestNewAcceleration_InvalidValue(t *testing.T) {
	factory := units.AccelerationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAcceleration(math.NaN(), units.AccelerationMeterPerSecondSquared)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAcceleration(math.Inf(1), units.AccelerationMeterPerSecondSquared)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAccelerationConversions(t *testing.T) {
	factory := units.AccelerationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAcceleration(180, units.AccelerationMeterPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersPerSecondSquared.
		// No expected conversion value provided for MetersPerSecondSquared, verifying result is not NaN.
		result := a.MetersPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to InchesPerSecondSquared.
		// No expected conversion value provided for InchesPerSecondSquared, verifying result is not NaN.
		result := a.InchesPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to FeetPerSecondSquared.
		// No expected conversion value provided for FeetPerSecondSquared, verifying result is not NaN.
		result := a.FeetPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to KnotsPerSecond.
		// No expected conversion value provided for KnotsPerSecond, verifying result is not NaN.
		result := a.KnotsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KnotsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KnotsPerMinute.
		// No expected conversion value provided for KnotsPerMinute, verifying result is not NaN.
		result := a.KnotsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KnotsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KnotsPerHour.
		// No expected conversion value provided for KnotsPerHour, verifying result is not NaN.
		result := a.KnotsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KnotsPerHour returned NaN")
		}
	}
	{
		// Test conversion to StandardGravity.
		// No expected conversion value provided for StandardGravity, verifying result is not NaN.
		result := a.StandardGravity()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardGravity returned NaN")
		}
	}
	{
		// Test conversion to NanometersPerSecondSquared.
		// No expected conversion value provided for NanometersPerSecondSquared, verifying result is not NaN.
		result := a.NanometersPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanometersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to MicrometersPerSecondSquared.
		// No expected conversion value provided for MicrometersPerSecondSquared, verifying result is not NaN.
		result := a.MicrometersPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrometersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerSecondSquared.
		// No expected conversion value provided for MillimetersPerSecondSquared, verifying result is not NaN.
		result := a.MillimetersPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerSecondSquared.
		// No expected conversion value provided for CentimetersPerSecondSquared, verifying result is not NaN.
		result := a.CentimetersPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to DecimetersPerSecondSquared.
		// No expected conversion value provided for DecimetersPerSecondSquared, verifying result is not NaN.
		result := a.DecimetersPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimetersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerSecondSquared.
		// No expected conversion value provided for KilometersPerSecondSquared, verifying result is not NaN.
		result := a.KilometersPerSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilometersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to MillistandardGravity.
		// No expected conversion value provided for MillistandardGravity, verifying result is not NaN.
		result := a.MillistandardGravity()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillistandardGravity returned NaN")
		}
	}
}

func TestAcceleration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AccelerationFactory{}
	a, err := factory.CreateAcceleration(90, units.AccelerationMeterPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AccelerationMeterPerSecondSquared {
		t.Errorf("expected default unit MeterPerSecondSquared, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AccelerationMeterPerSecondSquared
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AccelerationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AccelerationMeterPerSecondSquared {
		t.Errorf("expected unit MeterPerSecondSquared, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAccelerationToString(t *testing.T) {
	factory := units.AccelerationFactory{}
	a, err := factory.CreateAcceleration(45, units.AccelerationMeterPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AccelerationMeterPerSecondSquared, 2)
	expected := "45.00 " + units.GetAccelerationAbbreviation(units.AccelerationMeterPerSecondSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AccelerationMeterPerSecondSquared, -1)
	expected = "45 " + units.GetAccelerationAbbreviation(units.AccelerationMeterPerSecondSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAcceleration_EqualityAndComparison(t *testing.T) {
	factory := units.AccelerationFactory{}
	a1, _ := factory.CreateAcceleration(60, units.AccelerationMeterPerSecondSquared)
	a2, _ := factory.CreateAcceleration(60, units.AccelerationMeterPerSecondSquared)
	a3, _ := factory.CreateAcceleration(120, units.AccelerationMeterPerSecondSquared)

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

func TestAcceleration_Arithmetic(t *testing.T) {
	factory := units.AccelerationFactory{}
	a1, _ := factory.CreateAcceleration(30, units.AccelerationMeterPerSecondSquared)
	a2, _ := factory.CreateAcceleration(45, units.AccelerationMeterPerSecondSquared)

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