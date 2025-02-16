// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumeConcentrationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecimalFraction"}`
	
	factory := units.VolumeConcentrationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumeConcentrationDecimalFraction {
		t.Errorf("expected unit %v, got %v", units.VolumeConcentrationDecimalFraction, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecimalFraction"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumeConcentrationDto_ToJSON(t *testing.T) {
	dto := units.VolumeConcentrationDto{
		Value: 45,
		Unit:  units.VolumeConcentrationDecimalFraction,
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
	if result["unit"].(string) != string(units.VolumeConcentrationDecimalFraction) {
		t.Errorf("expected unit %s, got %v", units.VolumeConcentrationDecimalFraction, result["unit"])
	}
}

func TestNewVolumeConcentration_InvalidValue(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolumeConcentration(math.NaN(), units.VolumeConcentrationDecimalFraction)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolumeConcentration(math.Inf(1), units.VolumeConcentrationDecimalFraction)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumeConcentrationConversions(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolumeConcentration(180, units.VolumeConcentrationDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to DecimalFractions.
		// No expected conversion value provided for DecimalFractions, verifying result is not NaN.
		result := a.DecimalFractions()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimalFractions returned NaN")
		}
	}
	{
		// Test conversion to LitersPerLiter.
		// No expected conversion value provided for LitersPerLiter, verifying result is not NaN.
		result := a.LitersPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to LitersPerMililiter.
		// No expected conversion value provided for LitersPerMililiter, verifying result is not NaN.
		result := a.LitersPerMililiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to Percent.
		// No expected conversion value provided for Percent, verifying result is not NaN.
		result := a.Percent()
		if math.IsNaN(result) {
			t.Errorf("conversion to Percent returned NaN")
		}
	}
	{
		// Test conversion to PartsPerThousand.
		// No expected conversion value provided for PartsPerThousand, verifying result is not NaN.
		result := a.PartsPerThousand()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerThousand returned NaN")
		}
	}
	{
		// Test conversion to PartsPerMillion.
		// No expected conversion value provided for PartsPerMillion, verifying result is not NaN.
		result := a.PartsPerMillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerMillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerBillion.
		// No expected conversion value provided for PartsPerBillion, verifying result is not NaN.
		result := a.PartsPerBillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerBillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerTrillion.
		// No expected conversion value provided for PartsPerTrillion, verifying result is not NaN.
		result := a.PartsPerTrillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerTrillion returned NaN")
		}
	}
	{
		// Test conversion to PicolitersPerLiter.
		// No expected conversion value provided for PicolitersPerLiter, verifying result is not NaN.
		result := a.PicolitersPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicolitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerLiter.
		// No expected conversion value provided for NanolitersPerLiter, verifying result is not NaN.
		result := a.NanolitersPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanolitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerLiter.
		// No expected conversion value provided for MicrolitersPerLiter, verifying result is not NaN.
		result := a.MicrolitersPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrolitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerLiter.
		// No expected conversion value provided for MillilitersPerLiter, verifying result is not NaN.
		result := a.MillilitersPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillilitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerLiter.
		// No expected conversion value provided for CentilitersPerLiter, verifying result is not NaN.
		result := a.CentilitersPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentilitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerLiter.
		// No expected conversion value provided for DecilitersPerLiter, verifying result is not NaN.
		result := a.DecilitersPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecilitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PicolitersPerMililiter.
		// No expected conversion value provided for PicolitersPerMililiter, verifying result is not NaN.
		result := a.PicolitersPerMililiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicolitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerMililiter.
		// No expected conversion value provided for NanolitersPerMililiter, verifying result is not NaN.
		result := a.NanolitersPerMililiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanolitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerMililiter.
		// No expected conversion value provided for MicrolitersPerMililiter, verifying result is not NaN.
		result := a.MicrolitersPerMililiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrolitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerMililiter.
		// No expected conversion value provided for MillilitersPerMililiter, verifying result is not NaN.
		result := a.MillilitersPerMililiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillilitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerMililiter.
		// No expected conversion value provided for CentilitersPerMililiter, verifying result is not NaN.
		result := a.CentilitersPerMililiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentilitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerMililiter.
		// No expected conversion value provided for DecilitersPerMililiter, verifying result is not NaN.
		result := a.DecilitersPerMililiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecilitersPerMililiter returned NaN")
		}
	}
}

func TestVolumeConcentration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	a, err := factory.CreateVolumeConcentration(90, units.VolumeConcentrationDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumeConcentrationDecimalFraction {
		t.Errorf("expected default unit DecimalFraction, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumeConcentrationDecimalFraction
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumeConcentrationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumeConcentrationDecimalFraction {
		t.Errorf("expected unit DecimalFraction, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumeConcentrationToString(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	a, err := factory.CreateVolumeConcentration(45, units.VolumeConcentrationDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumeConcentrationDecimalFraction, 2)
	expected := "45.00 " + units.GetVolumeConcentrationAbbreviation(units.VolumeConcentrationDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumeConcentrationDecimalFraction, -1)
	expected = "45 " + units.GetVolumeConcentrationAbbreviation(units.VolumeConcentrationDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolumeConcentration_EqualityAndComparison(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	a1, _ := factory.CreateVolumeConcentration(60, units.VolumeConcentrationDecimalFraction)
	a2, _ := factory.CreateVolumeConcentration(60, units.VolumeConcentrationDecimalFraction)
	a3, _ := factory.CreateVolumeConcentration(120, units.VolumeConcentrationDecimalFraction)

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

func TestVolumeConcentration_Arithmetic(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	a1, _ := factory.CreateVolumeConcentration(30, units.VolumeConcentrationDecimalFraction)
	a2, _ := factory.CreateVolumeConcentration(45, units.VolumeConcentrationDecimalFraction)

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