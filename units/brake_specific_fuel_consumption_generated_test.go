// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestBrakeSpecificFuelConsumptionDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerJoule"}`
	
	factory := units.BrakeSpecificFuelConsumptionDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.BrakeSpecificFuelConsumptionKilogramPerJoule {
		t.Errorf("expected unit %v, got %v", units.BrakeSpecificFuelConsumptionKilogramPerJoule, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerJoule"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestBrakeSpecificFuelConsumptionDto_ToJSON(t *testing.T) {
	dto := units.BrakeSpecificFuelConsumptionDto{
		Value: 45,
		Unit:  units.BrakeSpecificFuelConsumptionKilogramPerJoule,
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
	if result["unit"].(string) != string(units.BrakeSpecificFuelConsumptionKilogramPerJoule) {
		t.Errorf("expected unit %s, got %v", units.BrakeSpecificFuelConsumptionKilogramPerJoule, result["unit"])
	}
}

func TestNewBrakeSpecificFuelConsumption_InvalidValue(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	// NaN value should return an error.
	_, err := factory.CreateBrakeSpecificFuelConsumption(math.NaN(), units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateBrakeSpecificFuelConsumption(math.Inf(1), units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestBrakeSpecificFuelConsumptionConversions(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateBrakeSpecificFuelConsumption(180, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerKiloWattHour.
		// No expected conversion value provided for GramsPerKiloWattHour, verifying result is not NaN.
		result := a.GramsPerKiloWattHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerKiloWattHour returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerJoule.
		// No expected conversion value provided for KilogramsPerJoule, verifying result is not NaN.
		result := a.KilogramsPerJoule()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerJoule returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerMechanicalHorsepowerHour.
		// No expected conversion value provided for PoundsPerMechanicalHorsepowerHour, verifying result is not NaN.
		result := a.PoundsPerMechanicalHorsepowerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerMechanicalHorsepowerHour returned NaN")
		}
	}
}

func TestBrakeSpecificFuelConsumption_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	a, err := factory.CreateBrakeSpecificFuelConsumption(90, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.BrakeSpecificFuelConsumptionKilogramPerJoule {
		t.Errorf("expected default unit KilogramPerJoule, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.BrakeSpecificFuelConsumptionGramPerKiloWattHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.BrakeSpecificFuelConsumptionDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.BrakeSpecificFuelConsumptionKilogramPerJoule {
		t.Errorf("expected unit KilogramPerJoule, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestBrakeSpecificFuelConsumptionToString(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	a, err := factory.CreateBrakeSpecificFuelConsumption(45, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.BrakeSpecificFuelConsumptionKilogramPerJoule, 2)
	expected := "45.00 " + units.GetBrakeSpecificFuelConsumptionAbbreviation(units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.BrakeSpecificFuelConsumptionKilogramPerJoule, -1)
	expected = "45 " + units.GetBrakeSpecificFuelConsumptionAbbreviation(units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestBrakeSpecificFuelConsumption_EqualityAndComparison(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	a1, _ := factory.CreateBrakeSpecificFuelConsumption(60, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	a2, _ := factory.CreateBrakeSpecificFuelConsumption(60, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	a3, _ := factory.CreateBrakeSpecificFuelConsumption(120, units.BrakeSpecificFuelConsumptionKilogramPerJoule)

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

func TestBrakeSpecificFuelConsumption_Arithmetic(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	a1, _ := factory.CreateBrakeSpecificFuelConsumption(30, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	a2, _ := factory.CreateBrakeSpecificFuelConsumption(45, units.BrakeSpecificFuelConsumptionKilogramPerJoule)

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