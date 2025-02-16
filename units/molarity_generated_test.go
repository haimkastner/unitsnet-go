// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MolePerCubicMeter"}`
	
	factory := units.MolarityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarityMolePerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.MolarityMolePerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MolePerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarityDto_ToJSON(t *testing.T) {
	dto := units.MolarityDto{
		Value: 45,
		Unit:  units.MolarityMolePerCubicMeter,
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
	if result["unit"].(string) != string(units.MolarityMolePerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.MolarityMolePerCubicMeter, result["unit"])
	}
}

func TestNewMolarity_InvalidValue(t *testing.T) {
	factory := units.MolarityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarity(math.NaN(), units.MolarityMolePerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarity(math.Inf(1), units.MolarityMolePerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarityConversions(t *testing.T) {
	factory := units.MolarityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarity(180, units.MolarityMolePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MolesPerCubicMeter.
		// No expected conversion value provided for MolesPerCubicMeter, verifying result is not NaN.
		result := a.MolesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MolesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MolesPerLiter.
		// No expected conversion value provided for MolesPerLiter, verifying result is not NaN.
		result := a.MolesPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PoundMolesPerCubicFoot.
		// No expected conversion value provided for PoundMolesPerCubicFoot, verifying result is not NaN.
		result := a.PoundMolesPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundMolesPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to KilomolesPerCubicMeter.
		// No expected conversion value provided for KilomolesPerCubicMeter, verifying result is not NaN.
		result := a.KilomolesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilomolesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to FemtomolesPerLiter.
		// No expected conversion value provided for FemtomolesPerLiter, verifying result is not NaN.
		result := a.FemtomolesPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to FemtomolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PicomolesPerLiter.
		// No expected conversion value provided for PicomolesPerLiter, verifying result is not NaN.
		result := a.PicomolesPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicomolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanomolesPerLiter.
		// No expected conversion value provided for NanomolesPerLiter, verifying result is not NaN.
		result := a.NanomolesPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanomolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicromolesPerLiter.
		// No expected conversion value provided for MicromolesPerLiter, verifying result is not NaN.
		result := a.MicromolesPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicromolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MillimolesPerLiter.
		// No expected conversion value provided for MillimolesPerLiter, verifying result is not NaN.
		result := a.MillimolesPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to CentimolesPerLiter.
		// No expected conversion value provided for CentimolesPerLiter, verifying result is not NaN.
		result := a.CentimolesPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecimolesPerLiter.
		// No expected conversion value provided for DecimolesPerLiter, verifying result is not NaN.
		result := a.DecimolesPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimolesPerLiter returned NaN")
		}
	}
}

func TestMolarity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarityFactory{}
	a, err := factory.CreateMolarity(90, units.MolarityMolePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarityMolePerCubicMeter {
		t.Errorf("expected default unit MolePerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarityMolePerCubicMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarityMolePerCubicMeter {
		t.Errorf("expected unit MolePerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarityToString(t *testing.T) {
	factory := units.MolarityFactory{}
	a, err := factory.CreateMolarity(45, units.MolarityMolePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarityMolePerCubicMeter, 2)
	expected := "45.00 " + units.GetMolarityAbbreviation(units.MolarityMolePerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarityMolePerCubicMeter, -1)
	expected = "45 " + units.GetMolarityAbbreviation(units.MolarityMolePerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarity_EqualityAndComparison(t *testing.T) {
	factory := units.MolarityFactory{}
	a1, _ := factory.CreateMolarity(60, units.MolarityMolePerCubicMeter)
	a2, _ := factory.CreateMolarity(60, units.MolarityMolePerCubicMeter)
	a3, _ := factory.CreateMolarity(120, units.MolarityMolePerCubicMeter)

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

func TestMolarity_Arithmetic(t *testing.T) {
	factory := units.MolarityFactory{}
	a1, _ := factory.CreateMolarity(30, units.MolarityMolePerCubicMeter)
	a2, _ := factory.CreateMolarity(45, units.MolarityMolePerCubicMeter)

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