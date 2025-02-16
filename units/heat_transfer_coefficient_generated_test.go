// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestHeatTransferCoefficientDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerSquareMeterKelvin"}`
	
	factory := units.HeatTransferCoefficientDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.HeatTransferCoefficientWattPerSquareMeterKelvin {
		t.Errorf("expected unit %v, got %v", units.HeatTransferCoefficientWattPerSquareMeterKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerSquareMeterKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestHeatTransferCoefficientDto_ToJSON(t *testing.T) {
	dto := units.HeatTransferCoefficientDto{
		Value: 45,
		Unit:  units.HeatTransferCoefficientWattPerSquareMeterKelvin,
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
	if result["unit"].(string) != string(units.HeatTransferCoefficientWattPerSquareMeterKelvin) {
		t.Errorf("expected unit %s, got %v", units.HeatTransferCoefficientWattPerSquareMeterKelvin, result["unit"])
	}
}

func TestNewHeatTransferCoefficient_InvalidValue(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	// NaN value should return an error.
	_, err := factory.CreateHeatTransferCoefficient(math.NaN(), units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateHeatTransferCoefficient(math.Inf(1), units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestHeatTransferCoefficientConversions(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateHeatTransferCoefficient(180, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerSquareMeterKelvin.
		// No expected conversion value provided for WattsPerSquareMeterKelvin, verifying result is not NaN.
		result := a.WattsPerSquareMeterKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to WattsPerSquareMeterCelsius.
		// No expected conversion value provided for WattsPerSquareMeterCelsius, verifying result is not NaN.
		result := a.WattsPerSquareMeterCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareMeterCelsius returned NaN")
		}
	}
	{
		// Test conversion to BtusPerHourSquareFootDegreeFahrenheit.
		// No expected conversion value provided for BtusPerHourSquareFootDegreeFahrenheit, verifying result is not NaN.
		result := a.BtusPerHourSquareFootDegreeFahrenheit()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerHourSquareFootDegreeFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerHourSquareMeterDegreeCelsius.
		// No expected conversion value provided for CaloriesPerHourSquareMeterDegreeCelsius, verifying result is not NaN.
		result := a.CaloriesPerHourSquareMeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to CaloriesPerHourSquareMeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerHourSquareMeterDegreeCelsius.
		// No expected conversion value provided for KilocaloriesPerHourSquareMeterDegreeCelsius, verifying result is not NaN.
		result := a.KilocaloriesPerHourSquareMeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocaloriesPerHourSquareMeterDegreeCelsius returned NaN")
		}
	}
}

func TestHeatTransferCoefficient_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	a, err := factory.CreateHeatTransferCoefficient(90, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.HeatTransferCoefficientWattPerSquareMeterKelvin {
		t.Errorf("expected default unit WattPerSquareMeterKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.HeatTransferCoefficientWattPerSquareMeterKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.HeatTransferCoefficientDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.HeatTransferCoefficientWattPerSquareMeterKelvin {
		t.Errorf("expected unit WattPerSquareMeterKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestHeatTransferCoefficientToString(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	a, err := factory.CreateHeatTransferCoefficient(45, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.HeatTransferCoefficientWattPerSquareMeterKelvin, 2)
	expected := "45.00 " + units.GetHeatTransferCoefficientAbbreviation(units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.HeatTransferCoefficientWattPerSquareMeterKelvin, -1)
	expected = "45 " + units.GetHeatTransferCoefficientAbbreviation(units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestHeatTransferCoefficient_EqualityAndComparison(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	a1, _ := factory.CreateHeatTransferCoefficient(60, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	a2, _ := factory.CreateHeatTransferCoefficient(60, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	a3, _ := factory.CreateHeatTransferCoefficient(120, units.HeatTransferCoefficientWattPerSquareMeterKelvin)

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

func TestHeatTransferCoefficient_Arithmetic(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	a1, _ := factory.CreateHeatTransferCoefficient(30, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	a2, _ := factory.CreateHeatTransferCoefficient(45, units.HeatTransferCoefficientWattPerSquareMeterKelvin)

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