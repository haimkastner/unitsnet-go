// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestIrradiationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerSquareMeter"}`
	
	factory := units.IrradiationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.IrradiationJoulePerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.IrradiationJoulePerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestIrradiationDto_ToJSON(t *testing.T) {
	dto := units.IrradiationDto{
		Value: 45,
		Unit:  units.IrradiationJoulePerSquareMeter,
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
	if result["unit"].(string) != string(units.IrradiationJoulePerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.IrradiationJoulePerSquareMeter, result["unit"])
	}
}

func TestNewIrradiation_InvalidValue(t *testing.T) {
	factory := units.IrradiationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateIrradiation(math.NaN(), units.IrradiationJoulePerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateIrradiation(math.Inf(1), units.IrradiationJoulePerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestIrradiationConversions(t *testing.T) {
	factory := units.IrradiationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateIrradiation(180, units.IrradiationJoulePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerSquareMeter.
		// No expected conversion value provided for JoulesPerSquareMeter, verifying result is not NaN.
		result := a.JoulesPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerSquareCentimeter.
		// No expected conversion value provided for JoulesPerSquareCentimeter, verifying result is not NaN.
		result := a.JoulesPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerSquareMillimeter.
		// No expected conversion value provided for JoulesPerSquareMillimeter, verifying result is not NaN.
		result := a.JoulesPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to WattHoursPerSquareMeter.
		// No expected conversion value provided for WattHoursPerSquareMeter, verifying result is not NaN.
		result := a.WattHoursPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattHoursPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to BtusPerSquareFoot.
		// No expected conversion value provided for BtusPerSquareFoot, verifying result is not NaN.
		result := a.BtusPerSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerSquareMeter.
		// No expected conversion value provided for KilojoulesPerSquareMeter, verifying result is not NaN.
		result := a.KilojoulesPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MillijoulesPerSquareCentimeter.
		// No expected conversion value provided for MillijoulesPerSquareCentimeter, verifying result is not NaN.
		result := a.MillijoulesPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillijoulesPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattHoursPerSquareMeter.
		// No expected conversion value provided for KilowattHoursPerSquareMeter, verifying result is not NaN.
		result := a.KilowattHoursPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattHoursPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilobtusPerSquareFoot.
		// No expected conversion value provided for KilobtusPerSquareFoot, verifying result is not NaN.
		result := a.KilobtusPerSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilobtusPerSquareFoot returned NaN")
		}
	}
}

func TestIrradiation_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.IrradiationFactory{}
	a, err := factory.CreateIrradiation(90, units.IrradiationJoulePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.IrradiationJoulePerSquareMeter {
		t.Errorf("expected default unit JoulePerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.IrradiationJoulePerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.IrradiationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.IrradiationJoulePerSquareMeter {
		t.Errorf("expected unit JoulePerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestIrradiationToString(t *testing.T) {
	factory := units.IrradiationFactory{}
	a, err := factory.CreateIrradiation(45, units.IrradiationJoulePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.IrradiationJoulePerSquareMeter, 2)
	expected := "45.00 " + units.GetIrradiationAbbreviation(units.IrradiationJoulePerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.IrradiationJoulePerSquareMeter, -1)
	expected = "45 " + units.GetIrradiationAbbreviation(units.IrradiationJoulePerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestIrradiation_EqualityAndComparison(t *testing.T) {
	factory := units.IrradiationFactory{}
	a1, _ := factory.CreateIrradiation(60, units.IrradiationJoulePerSquareMeter)
	a2, _ := factory.CreateIrradiation(60, units.IrradiationJoulePerSquareMeter)
	a3, _ := factory.CreateIrradiation(120, units.IrradiationJoulePerSquareMeter)

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

func TestIrradiation_Arithmetic(t *testing.T) {
	factory := units.IrradiationFactory{}
	a1, _ := factory.CreateIrradiation(30, units.IrradiationJoulePerSquareMeter)
	a2, _ := factory.CreateIrradiation(45, units.IrradiationJoulePerSquareMeter)

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