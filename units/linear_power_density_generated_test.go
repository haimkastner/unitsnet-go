// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLinearPowerDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerMeter"}`
	
	factory := units.LinearPowerDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LinearPowerDensityWattPerMeter {
		t.Errorf("expected unit %v, got %v", units.LinearPowerDensityWattPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLinearPowerDensityDto_ToJSON(t *testing.T) {
	dto := units.LinearPowerDensityDto{
		Value: 45,
		Unit:  units.LinearPowerDensityWattPerMeter,
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
	if result["unit"].(string) != string(units.LinearPowerDensityWattPerMeter) {
		t.Errorf("expected unit %s, got %v", units.LinearPowerDensityWattPerMeter, result["unit"])
	}
}

func TestNewLinearPowerDensity_InvalidValue(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLinearPowerDensity(math.NaN(), units.LinearPowerDensityWattPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLinearPowerDensity(math.Inf(1), units.LinearPowerDensityWattPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLinearPowerDensityConversions(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLinearPowerDensity(180, units.LinearPowerDensityWattPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerMeter.
		// No expected conversion value provided for WattsPerMeter, verifying result is not NaN.
		result := a.WattsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerCentimeter.
		// No expected conversion value provided for WattsPerCentimeter, verifying result is not NaN.
		result := a.WattsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerMillimeter.
		// No expected conversion value provided for WattsPerMillimeter, verifying result is not NaN.
		result := a.WattsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerInch.
		// No expected conversion value provided for WattsPerInch, verifying result is not NaN.
		result := a.WattsPerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to WattsPerFoot.
		// No expected conversion value provided for WattsPerFoot, verifying result is not NaN.
		result := a.WattsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerMeter.
		// No expected conversion value provided for MilliwattsPerMeter, verifying result is not NaN.
		result := a.MilliwattsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerMeter.
		// No expected conversion value provided for KilowattsPerMeter, verifying result is not NaN.
		result := a.KilowattsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerMeter.
		// No expected conversion value provided for MegawattsPerMeter, verifying result is not NaN.
		result := a.MegawattsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerMeter.
		// No expected conversion value provided for GigawattsPerMeter, verifying result is not NaN.
		result := a.GigawattsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerCentimeter.
		// No expected conversion value provided for MilliwattsPerCentimeter, verifying result is not NaN.
		result := a.MilliwattsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerCentimeter.
		// No expected conversion value provided for KilowattsPerCentimeter, verifying result is not NaN.
		result := a.KilowattsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerCentimeter.
		// No expected conversion value provided for MegawattsPerCentimeter, verifying result is not NaN.
		result := a.MegawattsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerCentimeter.
		// No expected conversion value provided for GigawattsPerCentimeter, verifying result is not NaN.
		result := a.GigawattsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerMillimeter.
		// No expected conversion value provided for MilliwattsPerMillimeter, verifying result is not NaN.
		result := a.MilliwattsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerMillimeter.
		// No expected conversion value provided for KilowattsPerMillimeter, verifying result is not NaN.
		result := a.KilowattsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerMillimeter.
		// No expected conversion value provided for MegawattsPerMillimeter, verifying result is not NaN.
		result := a.MegawattsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerMillimeter.
		// No expected conversion value provided for GigawattsPerMillimeter, verifying result is not NaN.
		result := a.GigawattsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerInch.
		// No expected conversion value provided for MilliwattsPerInch, verifying result is not NaN.
		result := a.MilliwattsPerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerInch.
		// No expected conversion value provided for KilowattsPerInch, verifying result is not NaN.
		result := a.KilowattsPerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerInch.
		// No expected conversion value provided for MegawattsPerInch, verifying result is not NaN.
		result := a.MegawattsPerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerInch.
		// No expected conversion value provided for GigawattsPerInch, verifying result is not NaN.
		result := a.GigawattsPerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerFoot.
		// No expected conversion value provided for MilliwattsPerFoot, verifying result is not NaN.
		result := a.MilliwattsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerFoot.
		// No expected conversion value provided for KilowattsPerFoot, verifying result is not NaN.
		result := a.KilowattsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerFoot.
		// No expected conversion value provided for MegawattsPerFoot, verifying result is not NaN.
		result := a.MegawattsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerFoot.
		// No expected conversion value provided for GigawattsPerFoot, verifying result is not NaN.
		result := a.GigawattsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerFoot returned NaN")
		}
	}
}

func TestLinearPowerDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	a, err := factory.CreateLinearPowerDensity(90, units.LinearPowerDensityWattPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LinearPowerDensityWattPerMeter {
		t.Errorf("expected default unit WattPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LinearPowerDensityWattPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LinearPowerDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LinearPowerDensityWattPerMeter {
		t.Errorf("expected unit WattPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLinearPowerDensityToString(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	a, err := factory.CreateLinearPowerDensity(45, units.LinearPowerDensityWattPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LinearPowerDensityWattPerMeter, 2)
	expected := "45.00 " + units.GetLinearPowerDensityAbbreviation(units.LinearPowerDensityWattPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LinearPowerDensityWattPerMeter, -1)
	expected = "45 " + units.GetLinearPowerDensityAbbreviation(units.LinearPowerDensityWattPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLinearPowerDensity_EqualityAndComparison(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	a1, _ := factory.CreateLinearPowerDensity(60, units.LinearPowerDensityWattPerMeter)
	a2, _ := factory.CreateLinearPowerDensity(60, units.LinearPowerDensityWattPerMeter)
	a3, _ := factory.CreateLinearPowerDensity(120, units.LinearPowerDensityWattPerMeter)

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

func TestLinearPowerDensity_Arithmetic(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	a1, _ := factory.CreateLinearPowerDensity(30, units.LinearPowerDensityWattPerMeter)
	a2, _ := factory.CreateLinearPowerDensity(45, units.LinearPowerDensityWattPerMeter)

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