// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPowerDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerCubicMeter"}`
	
	factory := units.PowerDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PowerDensityWattPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.PowerDensityWattPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPowerDensityDto_ToJSON(t *testing.T) {
	dto := units.PowerDensityDto{
		Value: 45,
		Unit:  units.PowerDensityWattPerCubicMeter,
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
	if result["unit"].(string) != string(units.PowerDensityWattPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.PowerDensityWattPerCubicMeter, result["unit"])
	}
}

func TestNewPowerDensity_InvalidValue(t *testing.T) {
	factory := units.PowerDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePowerDensity(math.NaN(), units.PowerDensityWattPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePowerDensity(math.Inf(1), units.PowerDensityWattPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPowerDensityConversions(t *testing.T) {
	factory := units.PowerDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePowerDensity(180, units.PowerDensityWattPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerCubicMeter.
		// No expected conversion value provided for WattsPerCubicMeter, verifying result is not NaN.
		result := a.WattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerCubicInch.
		// No expected conversion value provided for WattsPerCubicInch, verifying result is not NaN.
		result := a.WattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to WattsPerCubicFoot.
		// No expected conversion value provided for WattsPerCubicFoot, verifying result is not NaN.
		result := a.WattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to WattsPerLiter.
		// No expected conversion value provided for WattsPerLiter, verifying result is not NaN.
		result := a.WattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerCubicMeter.
		// No expected conversion value provided for PicowattsPerCubicMeter, verifying result is not NaN.
		result := a.PicowattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicowattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerCubicMeter.
		// No expected conversion value provided for NanowattsPerCubicMeter, verifying result is not NaN.
		result := a.NanowattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerCubicMeter.
		// No expected conversion value provided for MicrowattsPerCubicMeter, verifying result is not NaN.
		result := a.MicrowattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerCubicMeter.
		// No expected conversion value provided for MilliwattsPerCubicMeter, verifying result is not NaN.
		result := a.MilliwattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerCubicMeter.
		// No expected conversion value provided for DeciwattsPerCubicMeter, verifying result is not NaN.
		result := a.DeciwattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DeciwattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to DecawattsPerCubicMeter.
		// No expected conversion value provided for DecawattsPerCubicMeter, verifying result is not NaN.
		result := a.DecawattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecawattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerCubicMeter.
		// No expected conversion value provided for KilowattsPerCubicMeter, verifying result is not NaN.
		result := a.KilowattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerCubicMeter.
		// No expected conversion value provided for MegawattsPerCubicMeter, verifying result is not NaN.
		result := a.MegawattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerCubicMeter.
		// No expected conversion value provided for GigawattsPerCubicMeter, verifying result is not NaN.
		result := a.GigawattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to TerawattsPerCubicMeter.
		// No expected conversion value provided for TerawattsPerCubicMeter, verifying result is not NaN.
		result := a.TerawattsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerCubicInch.
		// No expected conversion value provided for PicowattsPerCubicInch, verifying result is not NaN.
		result := a.PicowattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicowattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerCubicInch.
		// No expected conversion value provided for NanowattsPerCubicInch, verifying result is not NaN.
		result := a.NanowattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerCubicInch.
		// No expected conversion value provided for MicrowattsPerCubicInch, verifying result is not NaN.
		result := a.MicrowattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerCubicInch.
		// No expected conversion value provided for MilliwattsPerCubicInch, verifying result is not NaN.
		result := a.MilliwattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerCubicInch.
		// No expected conversion value provided for DeciwattsPerCubicInch, verifying result is not NaN.
		result := a.DeciwattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to DeciwattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to DecawattsPerCubicInch.
		// No expected conversion value provided for DecawattsPerCubicInch, verifying result is not NaN.
		result := a.DecawattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecawattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerCubicInch.
		// No expected conversion value provided for KilowattsPerCubicInch, verifying result is not NaN.
		result := a.KilowattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerCubicInch.
		// No expected conversion value provided for MegawattsPerCubicInch, verifying result is not NaN.
		result := a.MegawattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerCubicInch.
		// No expected conversion value provided for GigawattsPerCubicInch, verifying result is not NaN.
		result := a.GigawattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to TerawattsPerCubicInch.
		// No expected conversion value provided for TerawattsPerCubicInch, verifying result is not NaN.
		result := a.TerawattsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerCubicFoot.
		// No expected conversion value provided for PicowattsPerCubicFoot, verifying result is not NaN.
		result := a.PicowattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicowattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerCubicFoot.
		// No expected conversion value provided for NanowattsPerCubicFoot, verifying result is not NaN.
		result := a.NanowattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerCubicFoot.
		// No expected conversion value provided for MicrowattsPerCubicFoot, verifying result is not NaN.
		result := a.MicrowattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerCubicFoot.
		// No expected conversion value provided for MilliwattsPerCubicFoot, verifying result is not NaN.
		result := a.MilliwattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerCubicFoot.
		// No expected conversion value provided for DeciwattsPerCubicFoot, verifying result is not NaN.
		result := a.DeciwattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to DeciwattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to DecawattsPerCubicFoot.
		// No expected conversion value provided for DecawattsPerCubicFoot, verifying result is not NaN.
		result := a.DecawattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecawattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerCubicFoot.
		// No expected conversion value provided for KilowattsPerCubicFoot, verifying result is not NaN.
		result := a.KilowattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerCubicFoot.
		// No expected conversion value provided for MegawattsPerCubicFoot, verifying result is not NaN.
		result := a.MegawattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerCubicFoot.
		// No expected conversion value provided for GigawattsPerCubicFoot, verifying result is not NaN.
		result := a.GigawattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to TerawattsPerCubicFoot.
		// No expected conversion value provided for TerawattsPerCubicFoot, verifying result is not NaN.
		result := a.TerawattsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerLiter.
		// No expected conversion value provided for PicowattsPerLiter, verifying result is not NaN.
		result := a.PicowattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicowattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerLiter.
		// No expected conversion value provided for NanowattsPerLiter, verifying result is not NaN.
		result := a.NanowattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerLiter.
		// No expected conversion value provided for MicrowattsPerLiter, verifying result is not NaN.
		result := a.MicrowattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerLiter.
		// No expected conversion value provided for MilliwattsPerLiter, verifying result is not NaN.
		result := a.MilliwattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerLiter.
		// No expected conversion value provided for DeciwattsPerLiter, verifying result is not NaN.
		result := a.DeciwattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DeciwattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecawattsPerLiter.
		// No expected conversion value provided for DecawattsPerLiter, verifying result is not NaN.
		result := a.DecawattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecawattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerLiter.
		// No expected conversion value provided for KilowattsPerLiter, verifying result is not NaN.
		result := a.KilowattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerLiter.
		// No expected conversion value provided for MegawattsPerLiter, verifying result is not NaN.
		result := a.MegawattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerLiter.
		// No expected conversion value provided for GigawattsPerLiter, verifying result is not NaN.
		result := a.GigawattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to TerawattsPerLiter.
		// No expected conversion value provided for TerawattsPerLiter, verifying result is not NaN.
		result := a.TerawattsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattsPerLiter returned NaN")
		}
	}
}

func TestPowerDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PowerDensityFactory{}
	a, err := factory.CreatePowerDensity(90, units.PowerDensityWattPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PowerDensityWattPerCubicMeter {
		t.Errorf("expected default unit WattPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PowerDensityWattPerCubicMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PowerDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PowerDensityWattPerCubicMeter {
		t.Errorf("expected unit WattPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPowerDensityToString(t *testing.T) {
	factory := units.PowerDensityFactory{}
	a, err := factory.CreatePowerDensity(45, units.PowerDensityWattPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PowerDensityWattPerCubicMeter, 2)
	expected := "45.00 " + units.GetPowerDensityAbbreviation(units.PowerDensityWattPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PowerDensityWattPerCubicMeter, -1)
	expected = "45 " + units.GetPowerDensityAbbreviation(units.PowerDensityWattPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPowerDensity_EqualityAndComparison(t *testing.T) {
	factory := units.PowerDensityFactory{}
	a1, _ := factory.CreatePowerDensity(60, units.PowerDensityWattPerCubicMeter)
	a2, _ := factory.CreatePowerDensity(60, units.PowerDensityWattPerCubicMeter)
	a3, _ := factory.CreatePowerDensity(120, units.PowerDensityWattPerCubicMeter)

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

func TestPowerDensity_Arithmetic(t *testing.T) {
	factory := units.PowerDensityFactory{}
	a1, _ := factory.CreatePowerDensity(30, units.PowerDensityWattPerCubicMeter)
	a2, _ := factory.CreatePowerDensity(45, units.PowerDensityWattPerCubicMeter)

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