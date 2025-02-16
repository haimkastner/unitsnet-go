// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerKilogram"}`
	
	factory := units.SpecificEnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificEnergyJoulePerKilogram {
		t.Errorf("expected unit %v, got %v", units.SpecificEnergyJoulePerKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerKilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificEnergyDto_ToJSON(t *testing.T) {
	dto := units.SpecificEnergyDto{
		Value: 45,
		Unit:  units.SpecificEnergyJoulePerKilogram,
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
	if result["unit"].(string) != string(units.SpecificEnergyJoulePerKilogram) {
		t.Errorf("expected unit %s, got %v", units.SpecificEnergyJoulePerKilogram, result["unit"])
	}
}

func TestNewSpecificEnergy_InvalidValue(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificEnergy(math.NaN(), units.SpecificEnergyJoulePerKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificEnergy(math.Inf(1), units.SpecificEnergyJoulePerKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificEnergyConversions(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificEnergy(180, units.SpecificEnergyJoulePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerKilogram.
		// No expected conversion value provided for JoulesPerKilogram, verifying result is not NaN.
		result := a.JoulesPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MegaJoulesPerTonne.
		// No expected conversion value provided for MegaJoulesPerTonne, verifying result is not NaN.
		result := a.MegaJoulesPerTonne()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegaJoulesPerTonne returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerGram.
		// No expected conversion value provided for CaloriesPerGram, verifying result is not NaN.
		result := a.CaloriesPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to CaloriesPerGram returned NaN")
		}
	}
	{
		// Test conversion to WattHoursPerKilogram.
		// No expected conversion value provided for WattHoursPerKilogram, verifying result is not NaN.
		result := a.WattHoursPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattHoursPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to WattDaysPerKilogram.
		// No expected conversion value provided for WattDaysPerKilogram, verifying result is not NaN.
		result := a.WattDaysPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to WattDaysPerTonne.
		// No expected conversion value provided for WattDaysPerTonne, verifying result is not NaN.
		result := a.WattDaysPerTonne()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to WattDaysPerShortTon.
		// No expected conversion value provided for WattDaysPerShortTon, verifying result is not NaN.
		result := a.WattDaysPerShortTon()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to WattHoursPerPound.
		// No expected conversion value provided for WattHoursPerPound, verifying result is not NaN.
		result := a.WattHoursPerPound()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattHoursPerPound returned NaN")
		}
	}
	{
		// Test conversion to BtuPerPound.
		// No expected conversion value provided for BtuPerPound, verifying result is not NaN.
		result := a.BtuPerPound()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtuPerPound returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerKilogram.
		// No expected conversion value provided for KilojoulesPerKilogram, verifying result is not NaN.
		result := a.KilojoulesPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerKilogram.
		// No expected conversion value provided for MegajoulesPerKilogram, verifying result is not NaN.
		result := a.MegajoulesPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerGram.
		// No expected conversion value provided for KilocaloriesPerGram, verifying result is not NaN.
		result := a.KilocaloriesPerGram()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocaloriesPerGram returned NaN")
		}
	}
	{
		// Test conversion to KilowattHoursPerKilogram.
		// No expected conversion value provided for KilowattHoursPerKilogram, verifying result is not NaN.
		result := a.KilowattHoursPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattHoursPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MegawattHoursPerKilogram.
		// No expected conversion value provided for MegawattHoursPerKilogram, verifying result is not NaN.
		result := a.MegawattHoursPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattHoursPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to GigawattHoursPerKilogram.
		// No expected conversion value provided for GigawattHoursPerKilogram, verifying result is not NaN.
		result := a.GigawattHoursPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattHoursPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to KilowattDaysPerKilogram.
		// No expected conversion value provided for KilowattDaysPerKilogram, verifying result is not NaN.
		result := a.KilowattDaysPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MegawattDaysPerKilogram.
		// No expected conversion value provided for MegawattDaysPerKilogram, verifying result is not NaN.
		result := a.MegawattDaysPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to GigawattDaysPerKilogram.
		// No expected conversion value provided for GigawattDaysPerKilogram, verifying result is not NaN.
		result := a.GigawattDaysPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to TerawattDaysPerKilogram.
		// No expected conversion value provided for TerawattDaysPerKilogram, verifying result is not NaN.
		result := a.TerawattDaysPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to KilowattDaysPerTonne.
		// No expected conversion value provided for KilowattDaysPerTonne, verifying result is not NaN.
		result := a.KilowattDaysPerTonne()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to MegawattDaysPerTonne.
		// No expected conversion value provided for MegawattDaysPerTonne, verifying result is not NaN.
		result := a.MegawattDaysPerTonne()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to GigawattDaysPerTonne.
		// No expected conversion value provided for GigawattDaysPerTonne, verifying result is not NaN.
		result := a.GigawattDaysPerTonne()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to TerawattDaysPerTonne.
		// No expected conversion value provided for TerawattDaysPerTonne, verifying result is not NaN.
		result := a.TerawattDaysPerTonne()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to KilowattDaysPerShortTon.
		// No expected conversion value provided for KilowattDaysPerShortTon, verifying result is not NaN.
		result := a.KilowattDaysPerShortTon()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to MegawattDaysPerShortTon.
		// No expected conversion value provided for MegawattDaysPerShortTon, verifying result is not NaN.
		result := a.MegawattDaysPerShortTon()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to GigawattDaysPerShortTon.
		// No expected conversion value provided for GigawattDaysPerShortTon, verifying result is not NaN.
		result := a.GigawattDaysPerShortTon()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to TerawattDaysPerShortTon.
		// No expected conversion value provided for TerawattDaysPerShortTon, verifying result is not NaN.
		result := a.TerawattDaysPerShortTon()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to KilowattHoursPerPound.
		// No expected conversion value provided for KilowattHoursPerPound, verifying result is not NaN.
		result := a.KilowattHoursPerPound()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattHoursPerPound returned NaN")
		}
	}
	{
		// Test conversion to MegawattHoursPerPound.
		// No expected conversion value provided for MegawattHoursPerPound, verifying result is not NaN.
		result := a.MegawattHoursPerPound()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattHoursPerPound returned NaN")
		}
	}
	{
		// Test conversion to GigawattHoursPerPound.
		// No expected conversion value provided for GigawattHoursPerPound, verifying result is not NaN.
		result := a.GigawattHoursPerPound()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattHoursPerPound returned NaN")
		}
	}
}

func TestSpecificEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	a, err := factory.CreateSpecificEnergy(90, units.SpecificEnergyJoulePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificEnergyJoulePerKilogram {
		t.Errorf("expected default unit JoulePerKilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificEnergyJoulePerKilogram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificEnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificEnergyJoulePerKilogram {
		t.Errorf("expected unit JoulePerKilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificEnergyToString(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	a, err := factory.CreateSpecificEnergy(45, units.SpecificEnergyJoulePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificEnergyJoulePerKilogram, 2)
	expected := "45.00 " + units.GetSpecificEnergyAbbreviation(units.SpecificEnergyJoulePerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificEnergyJoulePerKilogram, -1)
	expected = "45 " + units.GetSpecificEnergyAbbreviation(units.SpecificEnergyJoulePerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	a1, _ := factory.CreateSpecificEnergy(60, units.SpecificEnergyJoulePerKilogram)
	a2, _ := factory.CreateSpecificEnergy(60, units.SpecificEnergyJoulePerKilogram)
	a3, _ := factory.CreateSpecificEnergy(120, units.SpecificEnergyJoulePerKilogram)

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

func TestSpecificEnergy_Arithmetic(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	a1, _ := factory.CreateSpecificEnergy(30, units.SpecificEnergyJoulePerKilogram)
	a2, _ := factory.CreateSpecificEnergy(45, units.SpecificEnergyJoulePerKilogram)

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