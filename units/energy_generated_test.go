// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Joule"}`
	
	factory := units.EnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.EnergyJoule {
		t.Errorf("expected unit %v, got %v", units.EnergyJoule, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Joule"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestEnergyDto_ToJSON(t *testing.T) {
	dto := units.EnergyDto{
		Value: 45,
		Unit:  units.EnergyJoule,
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
	if result["unit"].(string) != string(units.EnergyJoule) {
		t.Errorf("expected unit %s, got %v", units.EnergyJoule, result["unit"])
	}
}

func TestNewEnergy_InvalidValue(t *testing.T) {
	factory := units.EnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateEnergy(math.NaN(), units.EnergyJoule)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateEnergy(math.Inf(1), units.EnergyJoule)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestEnergyConversions(t *testing.T) {
	factory := units.EnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateEnergy(180, units.EnergyJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Joules.
		// No expected conversion value provided for Joules, verifying result is not NaN.
		result := a.Joules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Joules returned NaN")
		}
	}
	{
		// Test conversion to Calories.
		// No expected conversion value provided for Calories, verifying result is not NaN.
		result := a.Calories()
		if math.IsNaN(result) {
			t.Errorf("conversion to Calories returned NaN")
		}
	}
	{
		// Test conversion to BritishThermalUnits.
		// No expected conversion value provided for BritishThermalUnits, verifying result is not NaN.
		result := a.BritishThermalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to BritishThermalUnits returned NaN")
		}
	}
	{
		// Test conversion to ElectronVolts.
		// No expected conversion value provided for ElectronVolts, verifying result is not NaN.
		result := a.ElectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to ElectronVolts returned NaN")
		}
	}
	{
		// Test conversion to FootPounds.
		// No expected conversion value provided for FootPounds, verifying result is not NaN.
		result := a.FootPounds()
		if math.IsNaN(result) {
			t.Errorf("conversion to FootPounds returned NaN")
		}
	}
	{
		// Test conversion to Ergs.
		// No expected conversion value provided for Ergs, verifying result is not NaN.
		result := a.Ergs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Ergs returned NaN")
		}
	}
	{
		// Test conversion to WattHours.
		// No expected conversion value provided for WattHours, verifying result is not NaN.
		result := a.WattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattHours returned NaN")
		}
	}
	{
		// Test conversion to WattDays.
		// No expected conversion value provided for WattDays, verifying result is not NaN.
		result := a.WattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattDays returned NaN")
		}
	}
	{
		// Test conversion to ThermsEc.
		// No expected conversion value provided for ThermsEc, verifying result is not NaN.
		result := a.ThermsEc()
		if math.IsNaN(result) {
			t.Errorf("conversion to ThermsEc returned NaN")
		}
	}
	{
		// Test conversion to ThermsUs.
		// No expected conversion value provided for ThermsUs, verifying result is not NaN.
		result := a.ThermsUs()
		if math.IsNaN(result) {
			t.Errorf("conversion to ThermsUs returned NaN")
		}
	}
	{
		// Test conversion to ThermsImperial.
		// No expected conversion value provided for ThermsImperial, verifying result is not NaN.
		result := a.ThermsImperial()
		if math.IsNaN(result) {
			t.Errorf("conversion to ThermsImperial returned NaN")
		}
	}
	{
		// Test conversion to HorsepowerHours.
		// No expected conversion value provided for HorsepowerHours, verifying result is not NaN.
		result := a.HorsepowerHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to HorsepowerHours returned NaN")
		}
	}
	{
		// Test conversion to Nanojoules.
		// No expected conversion value provided for Nanojoules, verifying result is not NaN.
		result := a.Nanojoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanojoules returned NaN")
		}
	}
	{
		// Test conversion to Microjoules.
		// No expected conversion value provided for Microjoules, verifying result is not NaN.
		result := a.Microjoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microjoules returned NaN")
		}
	}
	{
		// Test conversion to Millijoules.
		// No expected conversion value provided for Millijoules, verifying result is not NaN.
		result := a.Millijoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millijoules returned NaN")
		}
	}
	{
		// Test conversion to Kilojoules.
		// No expected conversion value provided for Kilojoules, verifying result is not NaN.
		result := a.Kilojoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilojoules returned NaN")
		}
	}
	{
		// Test conversion to Megajoules.
		// No expected conversion value provided for Megajoules, verifying result is not NaN.
		result := a.Megajoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megajoules returned NaN")
		}
	}
	{
		// Test conversion to Gigajoules.
		// No expected conversion value provided for Gigajoules, verifying result is not NaN.
		result := a.Gigajoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigajoules returned NaN")
		}
	}
	{
		// Test conversion to Terajoules.
		// No expected conversion value provided for Terajoules, verifying result is not NaN.
		result := a.Terajoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terajoules returned NaN")
		}
	}
	{
		// Test conversion to Petajoules.
		// No expected conversion value provided for Petajoules, verifying result is not NaN.
		result := a.Petajoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Petajoules returned NaN")
		}
	}
	{
		// Test conversion to Kilocalories.
		// No expected conversion value provided for Kilocalories, verifying result is not NaN.
		result := a.Kilocalories()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilocalories returned NaN")
		}
	}
	{
		// Test conversion to Megacalories.
		// No expected conversion value provided for Megacalories, verifying result is not NaN.
		result := a.Megacalories()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megacalories returned NaN")
		}
	}
	{
		// Test conversion to KilobritishThermalUnits.
		// No expected conversion value provided for KilobritishThermalUnits, verifying result is not NaN.
		result := a.KilobritishThermalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilobritishThermalUnits returned NaN")
		}
	}
	{
		// Test conversion to MegabritishThermalUnits.
		// No expected conversion value provided for MegabritishThermalUnits, verifying result is not NaN.
		result := a.MegabritishThermalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegabritishThermalUnits returned NaN")
		}
	}
	{
		// Test conversion to GigabritishThermalUnits.
		// No expected conversion value provided for GigabritishThermalUnits, verifying result is not NaN.
		result := a.GigabritishThermalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigabritishThermalUnits returned NaN")
		}
	}
	{
		// Test conversion to KiloelectronVolts.
		// No expected conversion value provided for KiloelectronVolts, verifying result is not NaN.
		result := a.KiloelectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to KiloelectronVolts returned NaN")
		}
	}
	{
		// Test conversion to MegaelectronVolts.
		// No expected conversion value provided for MegaelectronVolts, verifying result is not NaN.
		result := a.MegaelectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegaelectronVolts returned NaN")
		}
	}
	{
		// Test conversion to GigaelectronVolts.
		// No expected conversion value provided for GigaelectronVolts, verifying result is not NaN.
		result := a.GigaelectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigaelectronVolts returned NaN")
		}
	}
	{
		// Test conversion to TeraelectronVolts.
		// No expected conversion value provided for TeraelectronVolts, verifying result is not NaN.
		result := a.TeraelectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to TeraelectronVolts returned NaN")
		}
	}
	{
		// Test conversion to KilowattHours.
		// No expected conversion value provided for KilowattHours, verifying result is not NaN.
		result := a.KilowattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattHours returned NaN")
		}
	}
	{
		// Test conversion to MegawattHours.
		// No expected conversion value provided for MegawattHours, verifying result is not NaN.
		result := a.MegawattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattHours returned NaN")
		}
	}
	{
		// Test conversion to GigawattHours.
		// No expected conversion value provided for GigawattHours, verifying result is not NaN.
		result := a.GigawattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattHours returned NaN")
		}
	}
	{
		// Test conversion to TerawattHours.
		// No expected conversion value provided for TerawattHours, verifying result is not NaN.
		result := a.TerawattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattHours returned NaN")
		}
	}
	{
		// Test conversion to KilowattDays.
		// No expected conversion value provided for KilowattDays, verifying result is not NaN.
		result := a.KilowattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattDays returned NaN")
		}
	}
	{
		// Test conversion to MegawattDays.
		// No expected conversion value provided for MegawattDays, verifying result is not NaN.
		result := a.MegawattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattDays returned NaN")
		}
	}
	{
		// Test conversion to GigawattDays.
		// No expected conversion value provided for GigawattDays, verifying result is not NaN.
		result := a.GigawattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattDays returned NaN")
		}
	}
	{
		// Test conversion to TerawattDays.
		// No expected conversion value provided for TerawattDays, verifying result is not NaN.
		result := a.TerawattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattDays returned NaN")
		}
	}
	{
		// Test conversion to DecathermsEc.
		// No expected conversion value provided for DecathermsEc, verifying result is not NaN.
		result := a.DecathermsEc()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecathermsEc returned NaN")
		}
	}
	{
		// Test conversion to DecathermsUs.
		// No expected conversion value provided for DecathermsUs, verifying result is not NaN.
		result := a.DecathermsUs()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecathermsUs returned NaN")
		}
	}
	{
		// Test conversion to DecathermsImperial.
		// No expected conversion value provided for DecathermsImperial, verifying result is not NaN.
		result := a.DecathermsImperial()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecathermsImperial returned NaN")
		}
	}
}

func TestEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.EnergyFactory{}
	a, err := factory.CreateEnergy(90, units.EnergyJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.EnergyJoule {
		t.Errorf("expected default unit Joule, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.EnergyJoule
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.EnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.EnergyJoule {
		t.Errorf("expected unit Joule, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestEnergyToString(t *testing.T) {
	factory := units.EnergyFactory{}
	a, err := factory.CreateEnergy(45, units.EnergyJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.EnergyJoule, 2)
	expected := "45.00 " + units.GetEnergyAbbreviation(units.EnergyJoule)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.EnergyJoule, -1)
	expected = "45 " + units.GetEnergyAbbreviation(units.EnergyJoule)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.EnergyFactory{}
	a1, _ := factory.CreateEnergy(60, units.EnergyJoule)
	a2, _ := factory.CreateEnergy(60, units.EnergyJoule)
	a3, _ := factory.CreateEnergy(120, units.EnergyJoule)

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

func TestEnergy_Arithmetic(t *testing.T) {
	factory := units.EnergyFactory{}
	a1, _ := factory.CreateEnergy(30, units.EnergyJoule)
	a2, _ := factory.CreateEnergy(45, units.EnergyJoule)

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