// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricSusceptanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Siemens"}`
	
	factory := units.ElectricSusceptanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricSusceptanceSiemens {
		t.Errorf("expected unit %v, got %v", units.ElectricSusceptanceSiemens, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Siemens"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricSusceptanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricSusceptanceDto{
		Value: 45,
		Unit:  units.ElectricSusceptanceSiemens,
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
	if result["unit"].(string) != string(units.ElectricSusceptanceSiemens) {
		t.Errorf("expected unit %s, got %v", units.ElectricSusceptanceSiemens, result["unit"])
	}
}

func TestNewElectricSusceptance_InvalidValue(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricSusceptance(math.NaN(), units.ElectricSusceptanceSiemens)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricSusceptance(math.Inf(1), units.ElectricSusceptanceSiemens)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricSusceptanceConversions(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricSusceptance(180, units.ElectricSusceptanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Siemens.
		// No expected conversion value provided for Siemens, verifying result is not NaN.
		result := a.Siemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Siemens returned NaN")
		}
	}
	{
		// Test conversion to Mhos.
		// No expected conversion value provided for Mhos, verifying result is not NaN.
		result := a.Mhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Mhos returned NaN")
		}
	}
	{
		// Test conversion to Nanosiemens.
		// No expected conversion value provided for Nanosiemens, verifying result is not NaN.
		result := a.Nanosiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanosiemens returned NaN")
		}
	}
	{
		// Test conversion to Microsiemens.
		// No expected conversion value provided for Microsiemens, verifying result is not NaN.
		result := a.Microsiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microsiemens returned NaN")
		}
	}
	{
		// Test conversion to Millisiemens.
		// No expected conversion value provided for Millisiemens, verifying result is not NaN.
		result := a.Millisiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millisiemens returned NaN")
		}
	}
	{
		// Test conversion to Kilosiemens.
		// No expected conversion value provided for Kilosiemens, verifying result is not NaN.
		result := a.Kilosiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilosiemens returned NaN")
		}
	}
	{
		// Test conversion to Megasiemens.
		// No expected conversion value provided for Megasiemens, verifying result is not NaN.
		result := a.Megasiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megasiemens returned NaN")
		}
	}
	{
		// Test conversion to Gigasiemens.
		// No expected conversion value provided for Gigasiemens, verifying result is not NaN.
		result := a.Gigasiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigasiemens returned NaN")
		}
	}
	{
		// Test conversion to Terasiemens.
		// No expected conversion value provided for Terasiemens, verifying result is not NaN.
		result := a.Terasiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terasiemens returned NaN")
		}
	}
	{
		// Test conversion to Nanomhos.
		// No expected conversion value provided for Nanomhos, verifying result is not NaN.
		result := a.Nanomhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanomhos returned NaN")
		}
	}
	{
		// Test conversion to Micromhos.
		// No expected conversion value provided for Micromhos, verifying result is not NaN.
		result := a.Micromhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micromhos returned NaN")
		}
	}
	{
		// Test conversion to Millimhos.
		// No expected conversion value provided for Millimhos, verifying result is not NaN.
		result := a.Millimhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millimhos returned NaN")
		}
	}
	{
		// Test conversion to Kilomhos.
		// No expected conversion value provided for Kilomhos, verifying result is not NaN.
		result := a.Kilomhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilomhos returned NaN")
		}
	}
	{
		// Test conversion to Megamhos.
		// No expected conversion value provided for Megamhos, verifying result is not NaN.
		result := a.Megamhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megamhos returned NaN")
		}
	}
	{
		// Test conversion to Gigamhos.
		// No expected conversion value provided for Gigamhos, verifying result is not NaN.
		result := a.Gigamhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigamhos returned NaN")
		}
	}
	{
		// Test conversion to Teramhos.
		// No expected conversion value provided for Teramhos, verifying result is not NaN.
		result := a.Teramhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Teramhos returned NaN")
		}
	}
}

func TestElectricSusceptance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	a, err := factory.CreateElectricSusceptance(90, units.ElectricSusceptanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricSusceptanceSiemens {
		t.Errorf("expected default unit Siemens, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricSusceptanceSiemens
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricSusceptanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricSusceptanceSiemens {
		t.Errorf("expected unit Siemens, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricSusceptanceToString(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	a, err := factory.CreateElectricSusceptance(45, units.ElectricSusceptanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricSusceptanceSiemens, 2)
	expected := "45.00 " + units.GetElectricSusceptanceAbbreviation(units.ElectricSusceptanceSiemens)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricSusceptanceSiemens, -1)
	expected = "45 " + units.GetElectricSusceptanceAbbreviation(units.ElectricSusceptanceSiemens)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricSusceptance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	a1, _ := factory.CreateElectricSusceptance(60, units.ElectricSusceptanceSiemens)
	a2, _ := factory.CreateElectricSusceptance(60, units.ElectricSusceptanceSiemens)
	a3, _ := factory.CreateElectricSusceptance(120, units.ElectricSusceptanceSiemens)

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

func TestElectricSusceptance_Arithmetic(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	a1, _ := factory.CreateElectricSusceptance(30, units.ElectricSusceptanceSiemens)
	a2, _ := factory.CreateElectricSusceptance(45, units.ElectricSusceptanceSiemens)

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