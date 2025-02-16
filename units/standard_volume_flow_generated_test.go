// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestStandardVolumeFlowDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "StandardCubicMeterPerSecond"}`
	
	factory := units.StandardVolumeFlowDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.StandardVolumeFlowStandardCubicMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.StandardVolumeFlowStandardCubicMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "StandardCubicMeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestStandardVolumeFlowDto_ToJSON(t *testing.T) {
	dto := units.StandardVolumeFlowDto{
		Value: 45,
		Unit:  units.StandardVolumeFlowStandardCubicMeterPerSecond,
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
	if result["unit"].(string) != string(units.StandardVolumeFlowStandardCubicMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.StandardVolumeFlowStandardCubicMeterPerSecond, result["unit"])
	}
}

func TestNewStandardVolumeFlow_InvalidValue(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	// NaN value should return an error.
	_, err := factory.CreateStandardVolumeFlow(math.NaN(), units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateStandardVolumeFlow(math.Inf(1), units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestStandardVolumeFlowConversions(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateStandardVolumeFlow(180, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to StandardCubicMetersPerSecond.
		// No expected conversion value provided for StandardCubicMetersPerSecond, verifying result is not NaN.
		result := a.StandardCubicMetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicMetersPerMinute.
		// No expected conversion value provided for StandardCubicMetersPerMinute, verifying result is not NaN.
		result := a.StandardCubicMetersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicMetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicMetersPerHour.
		// No expected conversion value provided for StandardCubicMetersPerHour, verifying result is not NaN.
		result := a.StandardCubicMetersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicMetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicMetersPerDay.
		// No expected conversion value provided for StandardCubicMetersPerDay, verifying result is not NaN.
		result := a.StandardCubicMetersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicMetersPerDay returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicCentimetersPerMinute.
		// No expected conversion value provided for StandardCubicCentimetersPerMinute, verifying result is not NaN.
		result := a.StandardCubicCentimetersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicCentimetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to StandardLitersPerMinute.
		// No expected conversion value provided for StandardLitersPerMinute, verifying result is not NaN.
		result := a.StandardLitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardLitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicFeetPerSecond.
		// No expected conversion value provided for StandardCubicFeetPerSecond, verifying result is not NaN.
		result := a.StandardCubicFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicFeetPerMinute.
		// No expected conversion value provided for StandardCubicFeetPerMinute, verifying result is not NaN.
		result := a.StandardCubicFeetPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicFeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicFeetPerHour.
		// No expected conversion value provided for StandardCubicFeetPerHour, verifying result is not NaN.
		result := a.StandardCubicFeetPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicFeetPerHour returned NaN")
		}
	}
}

func TestStandardVolumeFlow_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	a, err := factory.CreateStandardVolumeFlow(90, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.StandardVolumeFlowStandardCubicMeterPerSecond {
		t.Errorf("expected default unit StandardCubicMeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.StandardVolumeFlowStandardCubicMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.StandardVolumeFlowDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.StandardVolumeFlowStandardCubicMeterPerSecond {
		t.Errorf("expected unit StandardCubicMeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestStandardVolumeFlowToString(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	a, err := factory.CreateStandardVolumeFlow(45, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.StandardVolumeFlowStandardCubicMeterPerSecond, 2)
	expected := "45.00 " + units.GetStandardVolumeFlowAbbreviation(units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.StandardVolumeFlowStandardCubicMeterPerSecond, -1)
	expected = "45 " + units.GetStandardVolumeFlowAbbreviation(units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestStandardVolumeFlow_EqualityAndComparison(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	a1, _ := factory.CreateStandardVolumeFlow(60, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	a2, _ := factory.CreateStandardVolumeFlow(60, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	a3, _ := factory.CreateStandardVolumeFlow(120, units.StandardVolumeFlowStandardCubicMeterPerSecond)

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

func TestStandardVolumeFlow_Arithmetic(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	a1, _ := factory.CreateStandardVolumeFlow(30, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	a2, _ := factory.CreateStandardVolumeFlow(45, units.StandardVolumeFlowStandardCubicMeterPerSecond)

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