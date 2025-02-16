// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTorquePerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonMeterPerMeter"}`
	
	factory := units.TorquePerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TorquePerLengthNewtonMeterPerMeter {
		t.Errorf("expected unit %v, got %v", units.TorquePerLengthNewtonMeterPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonMeterPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTorquePerLengthDto_ToJSON(t *testing.T) {
	dto := units.TorquePerLengthDto{
		Value: 45,
		Unit:  units.TorquePerLengthNewtonMeterPerMeter,
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
	if result["unit"].(string) != string(units.TorquePerLengthNewtonMeterPerMeter) {
		t.Errorf("expected unit %s, got %v", units.TorquePerLengthNewtonMeterPerMeter, result["unit"])
	}
}

func TestNewTorquePerLength_InvalidValue(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTorquePerLength(math.NaN(), units.TorquePerLengthNewtonMeterPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTorquePerLength(math.Inf(1), units.TorquePerLengthNewtonMeterPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTorquePerLengthConversions(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTorquePerLength(180, units.TorquePerLengthNewtonMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonMillimetersPerMeter.
		// No expected conversion value provided for NewtonMillimetersPerMeter, verifying result is not NaN.
		result := a.NewtonMillimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonCentimetersPerMeter.
		// No expected conversion value provided for NewtonCentimetersPerMeter, verifying result is not NaN.
		result := a.NewtonCentimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonMetersPerMeter.
		// No expected conversion value provided for NewtonMetersPerMeter, verifying result is not NaN.
		result := a.NewtonMetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundForceInchesPerFoot.
		// No expected conversion value provided for PoundForceInchesPerFoot, verifying result is not NaN.
		result := a.PoundForceInchesPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceInchesPerFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeetPerFoot.
		// No expected conversion value provided for PoundForceFeetPerFoot, verifying result is not NaN.
		result := a.PoundForceFeetPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceFeetPerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceMillimetersPerMeter.
		// No expected conversion value provided for KilogramForceMillimetersPerMeter, verifying result is not NaN.
		result := a.KilogramForceMillimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceCentimetersPerMeter.
		// No expected conversion value provided for KilogramForceCentimetersPerMeter, verifying result is not NaN.
		result := a.KilogramForceCentimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceMetersPerMeter.
		// No expected conversion value provided for KilogramForceMetersPerMeter, verifying result is not NaN.
		result := a.KilogramForceMetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to TonneForceMillimetersPerMeter.
		// No expected conversion value provided for TonneForceMillimetersPerMeter, verifying result is not NaN.
		result := a.TonneForceMillimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to TonneForceCentimetersPerMeter.
		// No expected conversion value provided for TonneForceCentimetersPerMeter, verifying result is not NaN.
		result := a.TonneForceCentimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to TonneForceMetersPerMeter.
		// No expected conversion value provided for TonneForceMetersPerMeter, verifying result is not NaN.
		result := a.TonneForceMetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMillimetersPerMeter.
		// No expected conversion value provided for KilonewtonMillimetersPerMeter, verifying result is not NaN.
		result := a.KilonewtonMillimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMillimetersPerMeter.
		// No expected conversion value provided for MeganewtonMillimetersPerMeter, verifying result is not NaN.
		result := a.MeganewtonMillimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonCentimetersPerMeter.
		// No expected conversion value provided for KilonewtonCentimetersPerMeter, verifying result is not NaN.
		result := a.KilonewtonCentimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonCentimetersPerMeter.
		// No expected conversion value provided for MeganewtonCentimetersPerMeter, verifying result is not NaN.
		result := a.MeganewtonCentimetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMetersPerMeter.
		// No expected conversion value provided for KilonewtonMetersPerMeter, verifying result is not NaN.
		result := a.KilonewtonMetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMetersPerMeter.
		// No expected conversion value provided for MeganewtonMetersPerMeter, verifying result is not NaN.
		result := a.MeganewtonMetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceInchesPerFoot.
		// No expected conversion value provided for KilopoundForceInchesPerFoot, verifying result is not NaN.
		result := a.KilopoundForceInchesPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundForceInchesPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MegapoundForceInchesPerFoot.
		// No expected conversion value provided for MegapoundForceInchesPerFoot, verifying result is not NaN.
		result := a.MegapoundForceInchesPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundForceInchesPerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceFeetPerFoot.
		// No expected conversion value provided for KilopoundForceFeetPerFoot, verifying result is not NaN.
		result := a.KilopoundForceFeetPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundForceFeetPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MegapoundForceFeetPerFoot.
		// No expected conversion value provided for MegapoundForceFeetPerFoot, verifying result is not NaN.
		result := a.MegapoundForceFeetPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundForceFeetPerFoot returned NaN")
		}
	}
}

func TestTorquePerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	a, err := factory.CreateTorquePerLength(90, units.TorquePerLengthNewtonMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TorquePerLengthNewtonMeterPerMeter {
		t.Errorf("expected default unit NewtonMeterPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TorquePerLengthNewtonMillimeterPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TorquePerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TorquePerLengthNewtonMeterPerMeter {
		t.Errorf("expected unit NewtonMeterPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTorquePerLengthToString(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	a, err := factory.CreateTorquePerLength(45, units.TorquePerLengthNewtonMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TorquePerLengthNewtonMeterPerMeter, 2)
	expected := "45.00 " + units.GetTorquePerLengthAbbreviation(units.TorquePerLengthNewtonMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TorquePerLengthNewtonMeterPerMeter, -1)
	expected = "45 " + units.GetTorquePerLengthAbbreviation(units.TorquePerLengthNewtonMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTorquePerLength_EqualityAndComparison(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	a1, _ := factory.CreateTorquePerLength(60, units.TorquePerLengthNewtonMeterPerMeter)
	a2, _ := factory.CreateTorquePerLength(60, units.TorquePerLengthNewtonMeterPerMeter)
	a3, _ := factory.CreateTorquePerLength(120, units.TorquePerLengthNewtonMeterPerMeter)

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

func TestTorquePerLength_Arithmetic(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	a1, _ := factory.CreateTorquePerLength(30, units.TorquePerLengthNewtonMeterPerMeter)
	a2, _ := factory.CreateTorquePerLength(45, units.TorquePerLengthNewtonMeterPerMeter)

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