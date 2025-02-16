// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestForcePerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonPerMeter"}`
	
	factory := units.ForcePerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ForcePerLengthNewtonPerMeter {
		t.Errorf("expected unit %v, got %v", units.ForcePerLengthNewtonPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestForcePerLengthDto_ToJSON(t *testing.T) {
	dto := units.ForcePerLengthDto{
		Value: 45,
		Unit:  units.ForcePerLengthNewtonPerMeter,
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
	if result["unit"].(string) != string(units.ForcePerLengthNewtonPerMeter) {
		t.Errorf("expected unit %s, got %v", units.ForcePerLengthNewtonPerMeter, result["unit"])
	}
}

func TestNewForcePerLength_InvalidValue(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateForcePerLength(math.NaN(), units.ForcePerLengthNewtonPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateForcePerLength(math.Inf(1), units.ForcePerLengthNewtonPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestForcePerLengthConversions(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateForcePerLength(180, units.ForcePerLengthNewtonPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonsPerMeter.
		// No expected conversion value provided for NewtonsPerMeter, verifying result is not NaN.
		result := a.NewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerCentimeter.
		// No expected conversion value provided for NewtonsPerCentimeter, verifying result is not NaN.
		result := a.NewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerMillimeter.
		// No expected conversion value provided for NewtonsPerMillimeter, verifying result is not NaN.
		result := a.NewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerMeter.
		// No expected conversion value provided for KilogramsForcePerMeter, verifying result is not NaN.
		result := a.KilogramsForcePerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerCentimeter.
		// No expected conversion value provided for KilogramsForcePerCentimeter, verifying result is not NaN.
		result := a.KilogramsForcePerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerMillimeter.
		// No expected conversion value provided for KilogramsForcePerMillimeter, verifying result is not NaN.
		result := a.KilogramsForcePerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerMeter.
		// No expected conversion value provided for TonnesForcePerMeter, verifying result is not NaN.
		result := a.TonnesForcePerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerMeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerCentimeter.
		// No expected conversion value provided for TonnesForcePerCentimeter, verifying result is not NaN.
		result := a.TonnesForcePerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerMillimeter.
		// No expected conversion value provided for TonnesForcePerMillimeter, verifying result is not NaN.
		result := a.TonnesForcePerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerFoot.
		// No expected conversion value provided for PoundsForcePerFoot, verifying result is not NaN.
		result := a.PoundsForcePerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerInch.
		// No expected conversion value provided for PoundsForcePerInch, verifying result is not NaN.
		result := a.PoundsForcePerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerYard.
		// No expected conversion value provided for PoundsForcePerYard, verifying result is not NaN.
		result := a.PoundsForcePerYard()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerYard returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerFoot.
		// No expected conversion value provided for KilopoundsForcePerFoot, verifying result is not NaN.
		result := a.KilopoundsForcePerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerInch.
		// No expected conversion value provided for KilopoundsForcePerInch, verifying result is not NaN.
		result := a.KilopoundsForcePerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerInch returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonsPerMeter.
		// No expected conversion value provided for NanonewtonsPerMeter, verifying result is not NaN.
		result := a.NanonewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonsPerMeter.
		// No expected conversion value provided for MicronewtonsPerMeter, verifying result is not NaN.
		result := a.MicronewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonsPerMeter.
		// No expected conversion value provided for MillinewtonsPerMeter, verifying result is not NaN.
		result := a.MillinewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonsPerMeter.
		// No expected conversion value provided for CentinewtonsPerMeter, verifying result is not NaN.
		result := a.CentinewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonsPerMeter.
		// No expected conversion value provided for DecinewtonsPerMeter, verifying result is not NaN.
		result := a.DecinewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerMeter.
		// No expected conversion value provided for DecanewtonsPerMeter, verifying result is not NaN.
		result := a.DecanewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerMeter.
		// No expected conversion value provided for KilonewtonsPerMeter, verifying result is not NaN.
		result := a.KilonewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerMeter.
		// No expected conversion value provided for MeganewtonsPerMeter, verifying result is not NaN.
		result := a.MeganewtonsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonsPerCentimeter.
		// No expected conversion value provided for NanonewtonsPerCentimeter, verifying result is not NaN.
		result := a.NanonewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonsPerCentimeter.
		// No expected conversion value provided for MicronewtonsPerCentimeter, verifying result is not NaN.
		result := a.MicronewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonsPerCentimeter.
		// No expected conversion value provided for MillinewtonsPerCentimeter, verifying result is not NaN.
		result := a.MillinewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonsPerCentimeter.
		// No expected conversion value provided for CentinewtonsPerCentimeter, verifying result is not NaN.
		result := a.CentinewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonsPerCentimeter.
		// No expected conversion value provided for DecinewtonsPerCentimeter, verifying result is not NaN.
		result := a.DecinewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerCentimeter.
		// No expected conversion value provided for DecanewtonsPerCentimeter, verifying result is not NaN.
		result := a.DecanewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerCentimeter.
		// No expected conversion value provided for KilonewtonsPerCentimeter, verifying result is not NaN.
		result := a.KilonewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerCentimeter.
		// No expected conversion value provided for MeganewtonsPerCentimeter, verifying result is not NaN.
		result := a.MeganewtonsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonsPerMillimeter.
		// No expected conversion value provided for NanonewtonsPerMillimeter, verifying result is not NaN.
		result := a.NanonewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonsPerMillimeter.
		// No expected conversion value provided for MicronewtonsPerMillimeter, verifying result is not NaN.
		result := a.MicronewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonsPerMillimeter.
		// No expected conversion value provided for MillinewtonsPerMillimeter, verifying result is not NaN.
		result := a.MillinewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonsPerMillimeter.
		// No expected conversion value provided for CentinewtonsPerMillimeter, verifying result is not NaN.
		result := a.CentinewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonsPerMillimeter.
		// No expected conversion value provided for DecinewtonsPerMillimeter, verifying result is not NaN.
		result := a.DecinewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerMillimeter.
		// No expected conversion value provided for DecanewtonsPerMillimeter, verifying result is not NaN.
		result := a.DecanewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerMillimeter.
		// No expected conversion value provided for KilonewtonsPerMillimeter, verifying result is not NaN.
		result := a.KilonewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerMillimeter.
		// No expected conversion value provided for MeganewtonsPerMillimeter, verifying result is not NaN.
		result := a.MeganewtonsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonsPerMillimeter returned NaN")
		}
	}
}

func TestForcePerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	a, err := factory.CreateForcePerLength(90, units.ForcePerLengthNewtonPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ForcePerLengthNewtonPerMeter {
		t.Errorf("expected default unit NewtonPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ForcePerLengthNewtonPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ForcePerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ForcePerLengthNewtonPerMeter {
		t.Errorf("expected unit NewtonPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestForcePerLengthToString(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	a, err := factory.CreateForcePerLength(45, units.ForcePerLengthNewtonPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ForcePerLengthNewtonPerMeter, 2)
	expected := "45.00 " + units.GetForcePerLengthAbbreviation(units.ForcePerLengthNewtonPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ForcePerLengthNewtonPerMeter, -1)
	expected = "45 " + units.GetForcePerLengthAbbreviation(units.ForcePerLengthNewtonPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestForcePerLength_EqualityAndComparison(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	a1, _ := factory.CreateForcePerLength(60, units.ForcePerLengthNewtonPerMeter)
	a2, _ := factory.CreateForcePerLength(60, units.ForcePerLengthNewtonPerMeter)
	a3, _ := factory.CreateForcePerLength(120, units.ForcePerLengthNewtonPerMeter)

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

func TestForcePerLength_Arithmetic(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	a1, _ := factory.CreateForcePerLength(30, units.ForcePerLengthNewtonPerMeter)
	a2, _ := factory.CreateForcePerLength(45, units.ForcePerLengthNewtonPerMeter)

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