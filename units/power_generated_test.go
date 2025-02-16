// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Watt"}`
	
	factory := units.PowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PowerWatt {
		t.Errorf("expected unit %v, got %v", units.PowerWatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Watt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPowerDto_ToJSON(t *testing.T) {
	dto := units.PowerDto{
		Value: 45,
		Unit:  units.PowerWatt,
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
	if result["unit"].(string) != string(units.PowerWatt) {
		t.Errorf("expected unit %s, got %v", units.PowerWatt, result["unit"])
	}
}

func TestNewPower_InvalidValue(t *testing.T) {
	factory := units.PowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePower(math.NaN(), units.PowerWatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePower(math.Inf(1), units.PowerWatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPowerConversions(t *testing.T) {
	factory := units.PowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePower(180, units.PowerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Watts.
		// No expected conversion value provided for Watts, verifying result is not NaN.
		result := a.Watts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Watts returned NaN")
		}
	}
	{
		// Test conversion to MechanicalHorsepower.
		// No expected conversion value provided for MechanicalHorsepower, verifying result is not NaN.
		result := a.MechanicalHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to MechanicalHorsepower returned NaN")
		}
	}
	{
		// Test conversion to MetricHorsepower.
		// No expected conversion value provided for MetricHorsepower, verifying result is not NaN.
		result := a.MetricHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetricHorsepower returned NaN")
		}
	}
	{
		// Test conversion to ElectricalHorsepower.
		// No expected conversion value provided for ElectricalHorsepower, verifying result is not NaN.
		result := a.ElectricalHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to ElectricalHorsepower returned NaN")
		}
	}
	{
		// Test conversion to BoilerHorsepower.
		// No expected conversion value provided for BoilerHorsepower, verifying result is not NaN.
		result := a.BoilerHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to BoilerHorsepower returned NaN")
		}
	}
	{
		// Test conversion to HydraulicHorsepower.
		// No expected conversion value provided for HydraulicHorsepower, verifying result is not NaN.
		result := a.HydraulicHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to HydraulicHorsepower returned NaN")
		}
	}
	{
		// Test conversion to BritishThermalUnitsPerHour.
		// No expected conversion value provided for BritishThermalUnitsPerHour, verifying result is not NaN.
		result := a.BritishThermalUnitsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to BritishThermalUnitsPerHour returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerHour.
		// No expected conversion value provided for JoulesPerHour, verifying result is not NaN.
		result := a.JoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerHour returned NaN")
		}
	}
	{
		// Test conversion to TonsOfRefrigeration.
		// No expected conversion value provided for TonsOfRefrigeration, verifying result is not NaN.
		result := a.TonsOfRefrigeration()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonsOfRefrigeration returned NaN")
		}
	}
	{
		// Test conversion to Femtowatts.
		// No expected conversion value provided for Femtowatts, verifying result is not NaN.
		result := a.Femtowatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Femtowatts returned NaN")
		}
	}
	{
		// Test conversion to Picowatts.
		// No expected conversion value provided for Picowatts, verifying result is not NaN.
		result := a.Picowatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picowatts returned NaN")
		}
	}
	{
		// Test conversion to Nanowatts.
		// No expected conversion value provided for Nanowatts, verifying result is not NaN.
		result := a.Nanowatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanowatts returned NaN")
		}
	}
	{
		// Test conversion to Microwatts.
		// No expected conversion value provided for Microwatts, verifying result is not NaN.
		result := a.Microwatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microwatts returned NaN")
		}
	}
	{
		// Test conversion to Milliwatts.
		// No expected conversion value provided for Milliwatts, verifying result is not NaN.
		result := a.Milliwatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliwatts returned NaN")
		}
	}
	{
		// Test conversion to Deciwatts.
		// No expected conversion value provided for Deciwatts, verifying result is not NaN.
		result := a.Deciwatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Deciwatts returned NaN")
		}
	}
	{
		// Test conversion to Decawatts.
		// No expected conversion value provided for Decawatts, verifying result is not NaN.
		result := a.Decawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decawatts returned NaN")
		}
	}
	{
		// Test conversion to Kilowatts.
		// No expected conversion value provided for Kilowatts, verifying result is not NaN.
		result := a.Kilowatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilowatts returned NaN")
		}
	}
	{
		// Test conversion to Megawatts.
		// No expected conversion value provided for Megawatts, verifying result is not NaN.
		result := a.Megawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megawatts returned NaN")
		}
	}
	{
		// Test conversion to Gigawatts.
		// No expected conversion value provided for Gigawatts, verifying result is not NaN.
		result := a.Gigawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigawatts returned NaN")
		}
	}
	{
		// Test conversion to Terawatts.
		// No expected conversion value provided for Terawatts, verifying result is not NaN.
		result := a.Terawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terawatts returned NaN")
		}
	}
	{
		// Test conversion to Petawatts.
		// No expected conversion value provided for Petawatts, verifying result is not NaN.
		result := a.Petawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Petawatts returned NaN")
		}
	}
	{
		// Test conversion to KilobritishThermalUnitsPerHour.
		// No expected conversion value provided for KilobritishThermalUnitsPerHour, verifying result is not NaN.
		result := a.KilobritishThermalUnitsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilobritishThermalUnitsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MegabritishThermalUnitsPerHour.
		// No expected conversion value provided for MegabritishThermalUnitsPerHour, verifying result is not NaN.
		result := a.MegabritishThermalUnitsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegabritishThermalUnitsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MillijoulesPerHour.
		// No expected conversion value provided for MillijoulesPerHour, verifying result is not NaN.
		result := a.MillijoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillijoulesPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerHour.
		// No expected conversion value provided for KilojoulesPerHour, verifying result is not NaN.
		result := a.KilojoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerHour returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerHour.
		// No expected conversion value provided for MegajoulesPerHour, verifying result is not NaN.
		result := a.MegajoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerHour returned NaN")
		}
	}
	{
		// Test conversion to GigajoulesPerHour.
		// No expected conversion value provided for GigajoulesPerHour, verifying result is not NaN.
		result := a.GigajoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigajoulesPerHour returned NaN")
		}
	}
}

func TestPower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PowerFactory{}
	a, err := factory.CreatePower(90, units.PowerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PowerWatt {
		t.Errorf("expected default unit Watt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PowerWatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PowerWatt {
		t.Errorf("expected unit Watt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPowerToString(t *testing.T) {
	factory := units.PowerFactory{}
	a, err := factory.CreatePower(45, units.PowerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PowerWatt, 2)
	expected := "45.00 " + units.GetPowerAbbreviation(units.PowerWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PowerWatt, -1)
	expected = "45 " + units.GetPowerAbbreviation(units.PowerWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPower_EqualityAndComparison(t *testing.T) {
	factory := units.PowerFactory{}
	a1, _ := factory.CreatePower(60, units.PowerWatt)
	a2, _ := factory.CreatePower(60, units.PowerWatt)
	a3, _ := factory.CreatePower(120, units.PowerWatt)

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

func TestPower_Arithmetic(t *testing.T) {
	factory := units.PowerFactory{}
	a1, _ := factory.CreatePower(30, units.PowerWatt)
	a2, _ := factory.CreatePower(45, units.PowerWatt)

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