// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassFluxDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerSecondPerSquareMeter"}`
	
	factory := units.MassFluxDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassFluxKilogramPerSecondPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.MassFluxKilogramPerSecondPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerSecondPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassFluxDto_ToJSON(t *testing.T) {
	dto := units.MassFluxDto{
		Value: 45,
		Unit:  units.MassFluxKilogramPerSecondPerSquareMeter,
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
	if result["unit"].(string) != string(units.MassFluxKilogramPerSecondPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.MassFluxKilogramPerSecondPerSquareMeter, result["unit"])
	}
}

func TestNewMassFlux_InvalidValue(t *testing.T) {
	factory := units.MassFluxFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMassFlux(math.NaN(), units.MassFluxKilogramPerSecondPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMassFlux(math.Inf(1), units.MassFluxKilogramPerSecondPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassFluxConversions(t *testing.T) {
	factory := units.MassFluxFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMassFlux(180, units.MassFluxKilogramPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerSecondPerSquareMeter.
		// No expected conversion value provided for GramsPerSecondPerSquareMeter, verifying result is not NaN.
		result := a.GramsPerSecondPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerSecondPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerSecondPerSquareCentimeter.
		// No expected conversion value provided for GramsPerSecondPerSquareCentimeter, verifying result is not NaN.
		result := a.GramsPerSecondPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerSecondPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerSecondPerSquareMillimeter.
		// No expected conversion value provided for GramsPerSecondPerSquareMillimeter, verifying result is not NaN.
		result := a.GramsPerSecondPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerSecondPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerHourPerSquareMeter.
		// No expected conversion value provided for GramsPerHourPerSquareMeter, verifying result is not NaN.
		result := a.GramsPerHourPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerHourPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerHourPerSquareCentimeter.
		// No expected conversion value provided for GramsPerHourPerSquareCentimeter, verifying result is not NaN.
		result := a.GramsPerHourPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerHourPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerHourPerSquareMillimeter.
		// No expected conversion value provided for GramsPerHourPerSquareMillimeter, verifying result is not NaN.
		result := a.GramsPerHourPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerHourPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerSecondPerSquareMeter.
		// No expected conversion value provided for KilogramsPerSecondPerSquareMeter, verifying result is not NaN.
		result := a.KilogramsPerSecondPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerSecondPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerSecondPerSquareCentimeter.
		// No expected conversion value provided for KilogramsPerSecondPerSquareCentimeter, verifying result is not NaN.
		result := a.KilogramsPerSecondPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerSecondPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerSecondPerSquareMillimeter.
		// No expected conversion value provided for KilogramsPerSecondPerSquareMillimeter, verifying result is not NaN.
		result := a.KilogramsPerSecondPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerSecondPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerHourPerSquareMeter.
		// No expected conversion value provided for KilogramsPerHourPerSquareMeter, verifying result is not NaN.
		result := a.KilogramsPerHourPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerHourPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerHourPerSquareCentimeter.
		// No expected conversion value provided for KilogramsPerHourPerSquareCentimeter, verifying result is not NaN.
		result := a.KilogramsPerHourPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerHourPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerHourPerSquareMillimeter.
		// No expected conversion value provided for KilogramsPerHourPerSquareMillimeter, verifying result is not NaN.
		result := a.KilogramsPerHourPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerHourPerSquareMillimeter returned NaN")
		}
	}
}

func TestMassFlux_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassFluxFactory{}
	a, err := factory.CreateMassFlux(90, units.MassFluxKilogramPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassFluxKilogramPerSecondPerSquareMeter {
		t.Errorf("expected default unit KilogramPerSecondPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassFluxGramPerSecondPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassFluxDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassFluxKilogramPerSecondPerSquareMeter {
		t.Errorf("expected unit KilogramPerSecondPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassFluxToString(t *testing.T) {
	factory := units.MassFluxFactory{}
	a, err := factory.CreateMassFlux(45, units.MassFluxKilogramPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassFluxKilogramPerSecondPerSquareMeter, 2)
	expected := "45.00 " + units.GetMassFluxAbbreviation(units.MassFluxKilogramPerSecondPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassFluxKilogramPerSecondPerSquareMeter, -1)
	expected = "45 " + units.GetMassFluxAbbreviation(units.MassFluxKilogramPerSecondPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMassFlux_EqualityAndComparison(t *testing.T) {
	factory := units.MassFluxFactory{}
	a1, _ := factory.CreateMassFlux(60, units.MassFluxKilogramPerSecondPerSquareMeter)
	a2, _ := factory.CreateMassFlux(60, units.MassFluxKilogramPerSecondPerSquareMeter)
	a3, _ := factory.CreateMassFlux(120, units.MassFluxKilogramPerSecondPerSquareMeter)

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

func TestMassFlux_Arithmetic(t *testing.T) {
	factory := units.MassFluxFactory{}
	a1, _ := factory.CreateMassFlux(30, units.MassFluxKilogramPerSecondPerSquareMeter)
	a2, _ := factory.CreateMassFlux(45, units.MassFluxKilogramPerSecondPerSquareMeter)

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