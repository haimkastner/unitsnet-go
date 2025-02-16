// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTorqueDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonMeter"}`
	
	factory := units.TorqueDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TorqueNewtonMeter {
		t.Errorf("expected unit %v, got %v", units.TorqueNewtonMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTorqueDto_ToJSON(t *testing.T) {
	dto := units.TorqueDto{
		Value: 45,
		Unit:  units.TorqueNewtonMeter,
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
	if result["unit"].(string) != string(units.TorqueNewtonMeter) {
		t.Errorf("expected unit %s, got %v", units.TorqueNewtonMeter, result["unit"])
	}
}

func TestNewTorque_InvalidValue(t *testing.T) {
	factory := units.TorqueFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTorque(math.NaN(), units.TorqueNewtonMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTorque(math.Inf(1), units.TorqueNewtonMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTorqueConversions(t *testing.T) {
	factory := units.TorqueFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTorque(180, units.TorqueNewtonMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonMillimeters.
		// No expected conversion value provided for NewtonMillimeters, verifying result is not NaN.
		result := a.NewtonMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMillimeters returned NaN")
		}
	}
	{
		// Test conversion to NewtonCentimeters.
		// No expected conversion value provided for NewtonCentimeters, verifying result is not NaN.
		result := a.NewtonCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonCentimeters returned NaN")
		}
	}
	{
		// Test conversion to NewtonMeters.
		// No expected conversion value provided for NewtonMeters, verifying result is not NaN.
		result := a.NewtonMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMeters returned NaN")
		}
	}
	{
		// Test conversion to PoundalFeet.
		// No expected conversion value provided for PoundalFeet, verifying result is not NaN.
		result := a.PoundalFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundalFeet returned NaN")
		}
	}
	{
		// Test conversion to PoundForceInches.
		// No expected conversion value provided for PoundForceInches, verifying result is not NaN.
		result := a.PoundForceInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceInches returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeet.
		// No expected conversion value provided for PoundForceFeet, verifying result is not NaN.
		result := a.PoundForceFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceFeet returned NaN")
		}
	}
	{
		// Test conversion to GramForceMillimeters.
		// No expected conversion value provided for GramForceMillimeters, verifying result is not NaN.
		result := a.GramForceMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramForceMillimeters returned NaN")
		}
	}
	{
		// Test conversion to GramForceCentimeters.
		// No expected conversion value provided for GramForceCentimeters, verifying result is not NaN.
		result := a.GramForceCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramForceCentimeters returned NaN")
		}
	}
	{
		// Test conversion to GramForceMeters.
		// No expected conversion value provided for GramForceMeters, verifying result is not NaN.
		result := a.GramForceMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramForceMeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceMillimeters.
		// No expected conversion value provided for KilogramForceMillimeters, verifying result is not NaN.
		result := a.KilogramForceMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceMillimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceCentimeters.
		// No expected conversion value provided for KilogramForceCentimeters, verifying result is not NaN.
		result := a.KilogramForceCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceCentimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceMeters.
		// No expected conversion value provided for KilogramForceMeters, verifying result is not NaN.
		result := a.KilogramForceMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceMeters returned NaN")
		}
	}
	{
		// Test conversion to TonneForceMillimeters.
		// No expected conversion value provided for TonneForceMillimeters, verifying result is not NaN.
		result := a.TonneForceMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceMillimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneForceCentimeters.
		// No expected conversion value provided for TonneForceCentimeters, verifying result is not NaN.
		result := a.TonneForceCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceCentimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneForceMeters.
		// No expected conversion value provided for TonneForceMeters, verifying result is not NaN.
		result := a.TonneForceMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceMeters returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMillimeters.
		// No expected conversion value provided for KilonewtonMillimeters, verifying result is not NaN.
		result := a.KilonewtonMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMillimeters returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMillimeters.
		// No expected conversion value provided for MeganewtonMillimeters, verifying result is not NaN.
		result := a.MeganewtonMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMillimeters returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonCentimeters.
		// No expected conversion value provided for KilonewtonCentimeters, verifying result is not NaN.
		result := a.KilonewtonCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonCentimeters.
		// No expected conversion value provided for MeganewtonCentimeters, verifying result is not NaN.
		result := a.MeganewtonCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonCentimeters returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMeters.
		// No expected conversion value provided for KilonewtonMeters, verifying result is not NaN.
		result := a.KilonewtonMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMeters returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMeters.
		// No expected conversion value provided for MeganewtonMeters, verifying result is not NaN.
		result := a.MeganewtonMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMeters returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceInches.
		// No expected conversion value provided for KilopoundForceInches, verifying result is not NaN.
		result := a.KilopoundForceInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundForceInches returned NaN")
		}
	}
	{
		// Test conversion to MegapoundForceInches.
		// No expected conversion value provided for MegapoundForceInches, verifying result is not NaN.
		result := a.MegapoundForceInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundForceInches returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceFeet.
		// No expected conversion value provided for KilopoundForceFeet, verifying result is not NaN.
		result := a.KilopoundForceFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundForceFeet returned NaN")
		}
	}
	{
		// Test conversion to MegapoundForceFeet.
		// No expected conversion value provided for MegapoundForceFeet, verifying result is not NaN.
		result := a.MegapoundForceFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundForceFeet returned NaN")
		}
	}
}

func TestTorque_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TorqueFactory{}
	a, err := factory.CreateTorque(90, units.TorqueNewtonMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TorqueNewtonMeter {
		t.Errorf("expected default unit NewtonMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TorqueNewtonMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TorqueDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TorqueNewtonMeter {
		t.Errorf("expected unit NewtonMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTorqueToString(t *testing.T) {
	factory := units.TorqueFactory{}
	a, err := factory.CreateTorque(45, units.TorqueNewtonMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TorqueNewtonMeter, 2)
	expected := "45.00 " + units.GetTorqueAbbreviation(units.TorqueNewtonMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TorqueNewtonMeter, -1)
	expected = "45 " + units.GetTorqueAbbreviation(units.TorqueNewtonMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTorque_EqualityAndComparison(t *testing.T) {
	factory := units.TorqueFactory{}
	a1, _ := factory.CreateTorque(60, units.TorqueNewtonMeter)
	a2, _ := factory.CreateTorque(60, units.TorqueNewtonMeter)
	a3, _ := factory.CreateTorque(120, units.TorqueNewtonMeter)

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

func TestTorque_Arithmetic(t *testing.T) {
	factory := units.TorqueFactory{}
	a1, _ := factory.CreateTorque(30, units.TorqueNewtonMeter)
	a2, _ := factory.CreateTorque(45, units.TorqueNewtonMeter)

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