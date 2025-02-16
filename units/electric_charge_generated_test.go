// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricChargeDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Coulomb"}`
	
	factory := units.ElectricChargeDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricChargeCoulomb {
		t.Errorf("expected unit %v, got %v", units.ElectricChargeCoulomb, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Coulomb"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricChargeDto_ToJSON(t *testing.T) {
	dto := units.ElectricChargeDto{
		Value: 45,
		Unit:  units.ElectricChargeCoulomb,
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
	if result["unit"].(string) != string(units.ElectricChargeCoulomb) {
		t.Errorf("expected unit %s, got %v", units.ElectricChargeCoulomb, result["unit"])
	}
}

func TestNewElectricCharge_InvalidValue(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricCharge(math.NaN(), units.ElectricChargeCoulomb)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricCharge(math.Inf(1), units.ElectricChargeCoulomb)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricChargeConversions(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricCharge(180, units.ElectricChargeCoulomb)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Coulombs.
		// No expected conversion value provided for Coulombs, verifying result is not NaN.
		result := a.Coulombs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Coulombs returned NaN")
		}
	}
	{
		// Test conversion to AmpereHours.
		// No expected conversion value provided for AmpereHours, verifying result is not NaN.
		result := a.AmpereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to AmpereHours returned NaN")
		}
	}
	{
		// Test conversion to Picocoulombs.
		// No expected conversion value provided for Picocoulombs, verifying result is not NaN.
		result := a.Picocoulombs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picocoulombs returned NaN")
		}
	}
	{
		// Test conversion to Nanocoulombs.
		// No expected conversion value provided for Nanocoulombs, verifying result is not NaN.
		result := a.Nanocoulombs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanocoulombs returned NaN")
		}
	}
	{
		// Test conversion to Microcoulombs.
		// No expected conversion value provided for Microcoulombs, verifying result is not NaN.
		result := a.Microcoulombs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microcoulombs returned NaN")
		}
	}
	{
		// Test conversion to Millicoulombs.
		// No expected conversion value provided for Millicoulombs, verifying result is not NaN.
		result := a.Millicoulombs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millicoulombs returned NaN")
		}
	}
	{
		// Test conversion to Kilocoulombs.
		// No expected conversion value provided for Kilocoulombs, verifying result is not NaN.
		result := a.Kilocoulombs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilocoulombs returned NaN")
		}
	}
	{
		// Test conversion to Megacoulombs.
		// No expected conversion value provided for Megacoulombs, verifying result is not NaN.
		result := a.Megacoulombs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megacoulombs returned NaN")
		}
	}
	{
		// Test conversion to MilliampereHours.
		// No expected conversion value provided for MilliampereHours, verifying result is not NaN.
		result := a.MilliampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliampereHours returned NaN")
		}
	}
	{
		// Test conversion to KiloampereHours.
		// No expected conversion value provided for KiloampereHours, verifying result is not NaN.
		result := a.KiloampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to KiloampereHours returned NaN")
		}
	}
	{
		// Test conversion to MegaampereHours.
		// No expected conversion value provided for MegaampereHours, verifying result is not NaN.
		result := a.MegaampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegaampereHours returned NaN")
		}
	}
}

func TestElectricCharge_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	a, err := factory.CreateElectricCharge(90, units.ElectricChargeCoulomb)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricChargeCoulomb {
		t.Errorf("expected default unit Coulomb, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricChargeCoulomb
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricChargeDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricChargeCoulomb {
		t.Errorf("expected unit Coulomb, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricChargeToString(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	a, err := factory.CreateElectricCharge(45, units.ElectricChargeCoulomb)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricChargeCoulomb, 2)
	expected := "45.00 " + units.GetElectricChargeAbbreviation(units.ElectricChargeCoulomb)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricChargeCoulomb, -1)
	expected = "45 " + units.GetElectricChargeAbbreviation(units.ElectricChargeCoulomb)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricCharge_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	a1, _ := factory.CreateElectricCharge(60, units.ElectricChargeCoulomb)
	a2, _ := factory.CreateElectricCharge(60, units.ElectricChargeCoulomb)
	a3, _ := factory.CreateElectricCharge(120, units.ElectricChargeCoulomb)

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

func TestElectricCharge_Arithmetic(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	a1, _ := factory.CreateElectricCharge(30, units.ElectricChargeCoulomb)
	a2, _ := factory.CreateElectricCharge(45, units.ElectricChargeCoulomb)

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