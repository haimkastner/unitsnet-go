// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricPotentialChangeRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltPerSecond"}`
	
	factory := units.ElectricPotentialChangeRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricPotentialChangeRateVoltPerSecond {
		t.Errorf("expected unit %v, got %v", units.ElectricPotentialChangeRateVoltPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricPotentialChangeRateDto_ToJSON(t *testing.T) {
	dto := units.ElectricPotentialChangeRateDto{
		Value: 45,
		Unit:  units.ElectricPotentialChangeRateVoltPerSecond,
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
	if result["unit"].(string) != string(units.ElectricPotentialChangeRateVoltPerSecond) {
		t.Errorf("expected unit %s, got %v", units.ElectricPotentialChangeRateVoltPerSecond, result["unit"])
	}
}

func TestNewElectricPotentialChangeRate_InvalidValue(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricPotentialChangeRate(math.NaN(), units.ElectricPotentialChangeRateVoltPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricPotentialChangeRate(math.Inf(1), units.ElectricPotentialChangeRateVoltPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricPotentialChangeRateConversions(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricPotentialChangeRate(180, units.ElectricPotentialChangeRateVoltPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltsPerSeconds.
		// No expected conversion value provided for VoltsPerSeconds, verifying result is not NaN.
		result := a.VoltsPerSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerMicroseconds.
		// No expected conversion value provided for VoltsPerMicroseconds, verifying result is not NaN.
		result := a.VoltsPerMicroseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerMinutes.
		// No expected conversion value provided for VoltsPerMinutes, verifying result is not NaN.
		result := a.VoltsPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerHours.
		// No expected conversion value provided for VoltsPerHours, verifying result is not NaN.
		result := a.VoltsPerHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltsPerHours returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerSeconds.
		// No expected conversion value provided for MicrovoltsPerSeconds, verifying result is not NaN.
		result := a.MicrovoltsPerSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrovoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerSeconds.
		// No expected conversion value provided for MillivoltsPerSeconds, verifying result is not NaN.
		result := a.MillivoltsPerSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillivoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerSeconds.
		// No expected conversion value provided for KilovoltsPerSeconds, verifying result is not NaN.
		result := a.KilovoltsPerSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerSeconds.
		// No expected conversion value provided for MegavoltsPerSeconds, verifying result is not NaN.
		result := a.MegavoltsPerSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerMicroseconds.
		// No expected conversion value provided for MicrovoltsPerMicroseconds, verifying result is not NaN.
		result := a.MicrovoltsPerMicroseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrovoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerMicroseconds.
		// No expected conversion value provided for MillivoltsPerMicroseconds, verifying result is not NaN.
		result := a.MillivoltsPerMicroseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillivoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerMicroseconds.
		// No expected conversion value provided for KilovoltsPerMicroseconds, verifying result is not NaN.
		result := a.KilovoltsPerMicroseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerMicroseconds.
		// No expected conversion value provided for MegavoltsPerMicroseconds, verifying result is not NaN.
		result := a.MegavoltsPerMicroseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerMinutes.
		// No expected conversion value provided for MicrovoltsPerMinutes, verifying result is not NaN.
		result := a.MicrovoltsPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrovoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerMinutes.
		// No expected conversion value provided for MillivoltsPerMinutes, verifying result is not NaN.
		result := a.MillivoltsPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillivoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerMinutes.
		// No expected conversion value provided for KilovoltsPerMinutes, verifying result is not NaN.
		result := a.KilovoltsPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerMinutes.
		// No expected conversion value provided for MegavoltsPerMinutes, verifying result is not NaN.
		result := a.MegavoltsPerMinutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerHours.
		// No expected conversion value provided for MicrovoltsPerHours, verifying result is not NaN.
		result := a.MicrovoltsPerHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrovoltsPerHours returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerHours.
		// No expected conversion value provided for MillivoltsPerHours, verifying result is not NaN.
		result := a.MillivoltsPerHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillivoltsPerHours returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerHours.
		// No expected conversion value provided for KilovoltsPerHours, verifying result is not NaN.
		result := a.KilovoltsPerHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltsPerHours returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerHours.
		// No expected conversion value provided for MegavoltsPerHours, verifying result is not NaN.
		result := a.MegavoltsPerHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltsPerHours returned NaN")
		}
	}
}

func TestElectricPotentialChangeRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	a, err := factory.CreateElectricPotentialChangeRate(90, units.ElectricPotentialChangeRateVoltPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricPotentialChangeRateVoltPerSecond {
		t.Errorf("expected default unit VoltPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricPotentialChangeRateVoltPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricPotentialChangeRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricPotentialChangeRateVoltPerSecond {
		t.Errorf("expected unit VoltPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricPotentialChangeRateToString(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	a, err := factory.CreateElectricPotentialChangeRate(45, units.ElectricPotentialChangeRateVoltPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricPotentialChangeRateVoltPerSecond, 2)
	expected := "45.00 " + units.GetElectricPotentialChangeRateAbbreviation(units.ElectricPotentialChangeRateVoltPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricPotentialChangeRateVoltPerSecond, -1)
	expected = "45 " + units.GetElectricPotentialChangeRateAbbreviation(units.ElectricPotentialChangeRateVoltPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricPotentialChangeRate_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	a1, _ := factory.CreateElectricPotentialChangeRate(60, units.ElectricPotentialChangeRateVoltPerSecond)
	a2, _ := factory.CreateElectricPotentialChangeRate(60, units.ElectricPotentialChangeRateVoltPerSecond)
	a3, _ := factory.CreateElectricPotentialChangeRate(120, units.ElectricPotentialChangeRateVoltPerSecond)

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

func TestElectricPotentialChangeRate_Arithmetic(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	a1, _ := factory.CreateElectricPotentialChangeRate(30, units.ElectricPotentialChangeRateVoltPerSecond)
	a2, _ := factory.CreateElectricPotentialChangeRate(45, units.ElectricPotentialChangeRateVoltPerSecond)

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