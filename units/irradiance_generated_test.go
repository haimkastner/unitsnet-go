// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestIrradianceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerSquareMeter"}`
	
	factory := units.IrradianceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.IrradianceWattPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.IrradianceWattPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestIrradianceDto_ToJSON(t *testing.T) {
	dto := units.IrradianceDto{
		Value: 45,
		Unit:  units.IrradianceWattPerSquareMeter,
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
	if result["unit"].(string) != string(units.IrradianceWattPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.IrradianceWattPerSquareMeter, result["unit"])
	}
}

func TestNewIrradiance_InvalidValue(t *testing.T) {
	factory := units.IrradianceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateIrradiance(math.NaN(), units.IrradianceWattPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateIrradiance(math.Inf(1), units.IrradianceWattPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestIrradianceConversions(t *testing.T) {
	factory := units.IrradianceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateIrradiance(180, units.IrradianceWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerSquareMeter.
		// No expected conversion value provided for WattsPerSquareMeter, verifying result is not NaN.
		result := a.WattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerSquareCentimeter.
		// No expected conversion value provided for WattsPerSquareCentimeter, verifying result is not NaN.
		result := a.WattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerSquareMeter.
		// No expected conversion value provided for PicowattsPerSquareMeter, verifying result is not NaN.
		result := a.PicowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerSquareMeter.
		// No expected conversion value provided for NanowattsPerSquareMeter, verifying result is not NaN.
		result := a.NanowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerSquareMeter.
		// No expected conversion value provided for MicrowattsPerSquareMeter, verifying result is not NaN.
		result := a.MicrowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerSquareMeter.
		// No expected conversion value provided for MilliwattsPerSquareMeter, verifying result is not NaN.
		result := a.MilliwattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerSquareMeter.
		// No expected conversion value provided for KilowattsPerSquareMeter, verifying result is not NaN.
		result := a.KilowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerSquareMeter.
		// No expected conversion value provided for MegawattsPerSquareMeter, verifying result is not NaN.
		result := a.MegawattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerSquareCentimeter.
		// No expected conversion value provided for PicowattsPerSquareCentimeter, verifying result is not NaN.
		result := a.PicowattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicowattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerSquareCentimeter.
		// No expected conversion value provided for NanowattsPerSquareCentimeter, verifying result is not NaN.
		result := a.NanowattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerSquareCentimeter.
		// No expected conversion value provided for MicrowattsPerSquareCentimeter, verifying result is not NaN.
		result := a.MicrowattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerSquareCentimeter.
		// No expected conversion value provided for MilliwattsPerSquareCentimeter, verifying result is not NaN.
		result := a.MilliwattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerSquareCentimeter.
		// No expected conversion value provided for KilowattsPerSquareCentimeter, verifying result is not NaN.
		result := a.KilowattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerSquareCentimeter.
		// No expected conversion value provided for MegawattsPerSquareCentimeter, verifying result is not NaN.
		result := a.MegawattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerSquareCentimeter returned NaN")
		}
	}
}

func TestIrradiance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.IrradianceFactory{}
	a, err := factory.CreateIrradiance(90, units.IrradianceWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.IrradianceWattPerSquareMeter {
		t.Errorf("expected default unit WattPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.IrradianceWattPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.IrradianceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.IrradianceWattPerSquareMeter {
		t.Errorf("expected unit WattPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestIrradianceToString(t *testing.T) {
	factory := units.IrradianceFactory{}
	a, err := factory.CreateIrradiance(45, units.IrradianceWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.IrradianceWattPerSquareMeter, 2)
	expected := "45.00 " + units.GetIrradianceAbbreviation(units.IrradianceWattPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.IrradianceWattPerSquareMeter, -1)
	expected = "45 " + units.GetIrradianceAbbreviation(units.IrradianceWattPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestIrradiance_EqualityAndComparison(t *testing.T) {
	factory := units.IrradianceFactory{}
	a1, _ := factory.CreateIrradiance(60, units.IrradianceWattPerSquareMeter)
	a2, _ := factory.CreateIrradiance(60, units.IrradianceWattPerSquareMeter)
	a3, _ := factory.CreateIrradiance(120, units.IrradianceWattPerSquareMeter)

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

func TestIrradiance_Arithmetic(t *testing.T) {
	factory := units.IrradianceFactory{}
	a1, _ := factory.CreateIrradiance(30, units.IrradianceWattPerSquareMeter)
	a2, _ := factory.CreateIrradiance(45, units.IrradianceWattPerSquareMeter)

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