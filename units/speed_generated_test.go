// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpeedDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterPerSecond"}`
	
	factory := units.SpeedDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpeedMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.SpeedMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpeedDto_ToJSON(t *testing.T) {
	dto := units.SpeedDto{
		Value: 45,
		Unit:  units.SpeedMeterPerSecond,
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
	if result["unit"].(string) != string(units.SpeedMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.SpeedMeterPerSecond, result["unit"])
	}
}

func TestNewSpeed_InvalidValue(t *testing.T) {
	factory := units.SpeedFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpeed(math.NaN(), units.SpeedMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpeed(math.Inf(1), units.SpeedMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpeedConversions(t *testing.T) {
	factory := units.SpeedFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpeed(180, units.SpeedMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersPerSecond.
		// No expected conversion value provided for MetersPerSecond, verifying result is not NaN.
		result := a.MetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MetersPerMinutes.
		// No expected conversion value provided for MetersPerMinutes, verifying result is not NaN.
		result := a.MetersPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MetersPerHour.
		// No expected conversion value provided for MetersPerHour, verifying result is not NaN.
		result := a.MetersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to FeetPerSecond.
		// No expected conversion value provided for FeetPerSecond, verifying result is not NaN.
		result := a.FeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to FeetPerMinute.
		// No expected conversion value provided for FeetPerMinute, verifying result is not NaN.
		result := a.FeetPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to FeetPerHour.
		// No expected conversion value provided for FeetPerHour, verifying result is not NaN.
		result := a.FeetPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetPerHour returned NaN")
		}
	}
	{
		// Test conversion to UsSurveyFeetPerSecond.
		// No expected conversion value provided for UsSurveyFeetPerSecond, verifying result is not NaN.
		result := a.UsSurveyFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsSurveyFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to UsSurveyFeetPerMinute.
		// No expected conversion value provided for UsSurveyFeetPerMinute, verifying result is not NaN.
		result := a.UsSurveyFeetPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsSurveyFeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to UsSurveyFeetPerHour.
		// No expected conversion value provided for UsSurveyFeetPerHour, verifying result is not NaN.
		result := a.UsSurveyFeetPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsSurveyFeetPerHour returned NaN")
		}
	}
	{
		// Test conversion to InchesPerSecond.
		// No expected conversion value provided for InchesPerSecond, verifying result is not NaN.
		result := a.InchesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to InchesPerMinute.
		// No expected conversion value provided for InchesPerMinute, verifying result is not NaN.
		result := a.InchesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to InchesPerHour.
		// No expected conversion value provided for InchesPerHour, verifying result is not NaN.
		result := a.InchesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesPerHour returned NaN")
		}
	}
	{
		// Test conversion to YardsPerSecond.
		// No expected conversion value provided for YardsPerSecond, verifying result is not NaN.
		result := a.YardsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to YardsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to YardsPerMinute.
		// No expected conversion value provided for YardsPerMinute, verifying result is not NaN.
		result := a.YardsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to YardsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to YardsPerHour.
		// No expected conversion value provided for YardsPerHour, verifying result is not NaN.
		result := a.YardsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to YardsPerHour returned NaN")
		}
	}
	{
		// Test conversion to Knots.
		// No expected conversion value provided for Knots, verifying result is not NaN.
		result := a.Knots()
		if math.IsNaN(result) {
			t.Errorf("conversion to Knots returned NaN")
		}
	}
	{
		// Test conversion to MilesPerHour.
		// No expected conversion value provided for MilesPerHour, verifying result is not NaN.
		result := a.MilesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilesPerHour returned NaN")
		}
	}
	{
		// Test conversion to Mach.
		// No expected conversion value provided for Mach, verifying result is not NaN.
		result := a.Mach()
		if math.IsNaN(result) {
			t.Errorf("conversion to Mach returned NaN")
		}
	}
	{
		// Test conversion to NanometersPerSecond.
		// No expected conversion value provided for NanometersPerSecond, verifying result is not NaN.
		result := a.NanometersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanometersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrometersPerSecond.
		// No expected conversion value provided for MicrometersPerSecond, verifying result is not NaN.
		result := a.MicrometersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrometersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerSecond.
		// No expected conversion value provided for MillimetersPerSecond, verifying result is not NaN.
		result := a.MillimetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerSecond.
		// No expected conversion value provided for CentimetersPerSecond, verifying result is not NaN.
		result := a.CentimetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecimetersPerSecond.
		// No expected conversion value provided for DecimetersPerSecond, verifying result is not NaN.
		result := a.DecimetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerSecond.
		// No expected conversion value provided for KilometersPerSecond, verifying result is not NaN.
		result := a.KilometersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilometersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanometersPerMinutes.
		// No expected conversion value provided for NanometersPerMinutes, verifying result is not NaN.
		result := a.NanometersPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanometersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MicrometersPerMinutes.
		// No expected conversion value provided for MicrometersPerMinutes, verifying result is not NaN.
		result := a.MicrometersPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrometersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerMinutes.
		// No expected conversion value provided for MillimetersPerMinutes, verifying result is not NaN.
		result := a.MillimetersPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerMinutes.
		// No expected conversion value provided for CentimetersPerMinutes, verifying result is not NaN.
		result := a.CentimetersPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to DecimetersPerMinutes.
		// No expected conversion value provided for DecimetersPerMinutes, verifying result is not NaN.
		result := a.DecimetersPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimetersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerMinutes.
		// No expected conversion value provided for KilometersPerMinutes, verifying result is not NaN.
		result := a.KilometersPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilometersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerHour.
		// No expected conversion value provided for MillimetersPerHour, verifying result is not NaN.
		result := a.MillimetersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerHour.
		// No expected conversion value provided for CentimetersPerHour, verifying result is not NaN.
		result := a.CentimetersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerHour.
		// No expected conversion value provided for KilometersPerHour, verifying result is not NaN.
		result := a.KilometersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilometersPerHour returned NaN")
		}
	}
}

func TestSpeed_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpeedFactory{}
	a, err := factory.CreateSpeed(90, units.SpeedMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpeedMeterPerSecond {
		t.Errorf("expected default unit MeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpeedMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpeedDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpeedMeterPerSecond {
		t.Errorf("expected unit MeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpeedToString(t *testing.T) {
	factory := units.SpeedFactory{}
	a, err := factory.CreateSpeed(45, units.SpeedMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpeedMeterPerSecond, 2)
	expected := "45.00 " + units.GetSpeedAbbreviation(units.SpeedMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpeedMeterPerSecond, -1)
	expected = "45 " + units.GetSpeedAbbreviation(units.SpeedMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpeed_EqualityAndComparison(t *testing.T) {
	factory := units.SpeedFactory{}
	a1, _ := factory.CreateSpeed(60, units.SpeedMeterPerSecond)
	a2, _ := factory.CreateSpeed(60, units.SpeedMeterPerSecond)
	a3, _ := factory.CreateSpeed(120, units.SpeedMeterPerSecond)

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

func TestSpeed_Arithmetic(t *testing.T) {
	factory := units.SpeedFactory{}
	a1, _ := factory.CreateSpeed(30, units.SpeedMeterPerSecond)
	a2, _ := factory.CreateSpeed(45, units.SpeedMeterPerSecond)

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