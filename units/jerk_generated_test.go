// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestJerkDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterPerSecondCubed"}`
	
	factory := units.JerkDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.JerkMeterPerSecondCubed {
		t.Errorf("expected unit %v, got %v", units.JerkMeterPerSecondCubed, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterPerSecondCubed"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestJerkDto_ToJSON(t *testing.T) {
	dto := units.JerkDto{
		Value: 45,
		Unit:  units.JerkMeterPerSecondCubed,
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
	if result["unit"].(string) != string(units.JerkMeterPerSecondCubed) {
		t.Errorf("expected unit %s, got %v", units.JerkMeterPerSecondCubed, result["unit"])
	}
}

func TestNewJerk_InvalidValue(t *testing.T) {
	factory := units.JerkFactory{}
	// NaN value should return an error.
	_, err := factory.CreateJerk(math.NaN(), units.JerkMeterPerSecondCubed)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateJerk(math.Inf(1), units.JerkMeterPerSecondCubed)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestJerkConversions(t *testing.T) {
	factory := units.JerkFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateJerk(180, units.JerkMeterPerSecondCubed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersPerSecondCubed.
		// No expected conversion value provided for MetersPerSecondCubed, verifying result is not NaN.
		result := a.MetersPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to InchesPerSecondCubed.
		// No expected conversion value provided for InchesPerSecondCubed, verifying result is not NaN.
		result := a.InchesPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to FeetPerSecondCubed.
		// No expected conversion value provided for FeetPerSecondCubed, verifying result is not NaN.
		result := a.FeetPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to StandardGravitiesPerSecond.
		// No expected conversion value provided for StandardGravitiesPerSecond, verifying result is not NaN.
		result := a.StandardGravitiesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardGravitiesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanometersPerSecondCubed.
		// No expected conversion value provided for NanometersPerSecondCubed, verifying result is not NaN.
		result := a.NanometersPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanometersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to MicrometersPerSecondCubed.
		// No expected conversion value provided for MicrometersPerSecondCubed, verifying result is not NaN.
		result := a.MicrometersPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrometersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerSecondCubed.
		// No expected conversion value provided for MillimetersPerSecondCubed, verifying result is not NaN.
		result := a.MillimetersPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerSecondCubed.
		// No expected conversion value provided for CentimetersPerSecondCubed, verifying result is not NaN.
		result := a.CentimetersPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to DecimetersPerSecondCubed.
		// No expected conversion value provided for DecimetersPerSecondCubed, verifying result is not NaN.
		result := a.DecimetersPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimetersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerSecondCubed.
		// No expected conversion value provided for KilometersPerSecondCubed, verifying result is not NaN.
		result := a.KilometersPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilometersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to MillistandardGravitiesPerSecond.
		// No expected conversion value provided for MillistandardGravitiesPerSecond, verifying result is not NaN.
		result := a.MillistandardGravitiesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillistandardGravitiesPerSecond returned NaN")
		}
	}
}

func TestJerk_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.JerkFactory{}
	a, err := factory.CreateJerk(90, units.JerkMeterPerSecondCubed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.JerkMeterPerSecondCubed {
		t.Errorf("expected default unit MeterPerSecondCubed, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.JerkMeterPerSecondCubed
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.JerkDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.JerkMeterPerSecondCubed {
		t.Errorf("expected unit MeterPerSecondCubed, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestJerkToString(t *testing.T) {
	factory := units.JerkFactory{}
	a, err := factory.CreateJerk(45, units.JerkMeterPerSecondCubed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.JerkMeterPerSecondCubed, 2)
	expected := "45.00 " + units.GetJerkAbbreviation(units.JerkMeterPerSecondCubed)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.JerkMeterPerSecondCubed, -1)
	expected = "45 " + units.GetJerkAbbreviation(units.JerkMeterPerSecondCubed)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestJerk_EqualityAndComparison(t *testing.T) {
	factory := units.JerkFactory{}
	a1, _ := factory.CreateJerk(60, units.JerkMeterPerSecondCubed)
	a2, _ := factory.CreateJerk(60, units.JerkMeterPerSecondCubed)
	a3, _ := factory.CreateJerk(120, units.JerkMeterPerSecondCubed)

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

func TestJerk_Arithmetic(t *testing.T) {
	factory := units.JerkFactory{}
	a1, _ := factory.CreateJerk(30, units.JerkMeterPerSecondCubed)
	a2, _ := factory.CreateJerk(45, units.JerkMeterPerSecondCubed)

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