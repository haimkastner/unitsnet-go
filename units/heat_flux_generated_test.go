// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestHeatFluxDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerSquareMeter"}`
	
	factory := units.HeatFluxDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.HeatFluxWattPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.HeatFluxWattPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestHeatFluxDto_ToJSON(t *testing.T) {
	dto := units.HeatFluxDto{
		Value: 45,
		Unit:  units.HeatFluxWattPerSquareMeter,
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
	if result["unit"].(string) != string(units.HeatFluxWattPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.HeatFluxWattPerSquareMeter, result["unit"])
	}
}

func TestNewHeatFlux_InvalidValue(t *testing.T) {
	factory := units.HeatFluxFactory{}
	// NaN value should return an error.
	_, err := factory.CreateHeatFlux(math.NaN(), units.HeatFluxWattPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateHeatFlux(math.Inf(1), units.HeatFluxWattPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestHeatFluxConversions(t *testing.T) {
	factory := units.HeatFluxFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateHeatFlux(180, units.HeatFluxWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerSquareMeter.
		// No expected conversion value provided for WattsPerSquareMeter, verifying result is not NaN.
		result := a.WattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerSquareInch.
		// No expected conversion value provided for WattsPerSquareInch, verifying result is not NaN.
		result := a.WattsPerSquareInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to WattsPerSquareFoot.
		// No expected conversion value provided for WattsPerSquareFoot, verifying result is not NaN.
		result := a.WattsPerSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to BtusPerSecondSquareInch.
		// No expected conversion value provided for BtusPerSecondSquareInch, verifying result is not NaN.
		result := a.BtusPerSecondSquareInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerSecondSquareInch returned NaN")
		}
	}
	{
		// Test conversion to BtusPerSecondSquareFoot.
		// No expected conversion value provided for BtusPerSecondSquareFoot, verifying result is not NaN.
		result := a.BtusPerSecondSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerSecondSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to BtusPerMinuteSquareFoot.
		// No expected conversion value provided for BtusPerMinuteSquareFoot, verifying result is not NaN.
		result := a.BtusPerMinuteSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerMinuteSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to BtusPerHourSquareFoot.
		// No expected conversion value provided for BtusPerHourSquareFoot, verifying result is not NaN.
		result := a.BtusPerHourSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerHourSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerSecondSquareCentimeter.
		// No expected conversion value provided for CaloriesPerSecondSquareCentimeter, verifying result is not NaN.
		result := a.CaloriesPerSecondSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CaloriesPerSecondSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerHourSquareMeter.
		// No expected conversion value provided for KilocaloriesPerHourSquareMeter, verifying result is not NaN.
		result := a.KilocaloriesPerHourSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocaloriesPerHourSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerFootSecond.
		// No expected conversion value provided for PoundsForcePerFootSecond, verifying result is not NaN.
		result := a.PoundsForcePerFootSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerFootSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerSecondCubed.
		// No expected conversion value provided for PoundsPerSecondCubed, verifying result is not NaN.
		result := a.PoundsPerSecondCubed()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerSquareMeter.
		// No expected conversion value provided for NanowattsPerSquareMeter, verifying result is not NaN.
		result := a.NanowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerSquareMeter.
		// No expected conversion value provided for MicrowattsPerSquareMeter, verifying result is not NaN.
		result := a.MicrowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerSquareMeter.
		// No expected conversion value provided for MilliwattsPerSquareMeter, verifying result is not NaN.
		result := a.MilliwattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to CentiwattsPerSquareMeter.
		// No expected conversion value provided for CentiwattsPerSquareMeter, verifying result is not NaN.
		result := a.CentiwattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentiwattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerSquareMeter.
		// No expected conversion value provided for DeciwattsPerSquareMeter, verifying result is not NaN.
		result := a.DeciwattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DeciwattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerSquareMeter.
		// No expected conversion value provided for KilowattsPerSquareMeter, verifying result is not NaN.
		result := a.KilowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerSecondSquareCentimeter.
		// No expected conversion value provided for KilocaloriesPerSecondSquareCentimeter, verifying result is not NaN.
		result := a.KilocaloriesPerSecondSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocaloriesPerSecondSquareCentimeter returned NaN")
		}
	}
}

func TestHeatFlux_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.HeatFluxFactory{}
	a, err := factory.CreateHeatFlux(90, units.HeatFluxWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.HeatFluxWattPerSquareMeter {
		t.Errorf("expected default unit WattPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.HeatFluxWattPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.HeatFluxDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.HeatFluxWattPerSquareMeter {
		t.Errorf("expected unit WattPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestHeatFluxToString(t *testing.T) {
	factory := units.HeatFluxFactory{}
	a, err := factory.CreateHeatFlux(45, units.HeatFluxWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.HeatFluxWattPerSquareMeter, 2)
	expected := "45.00 " + units.GetHeatFluxAbbreviation(units.HeatFluxWattPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.HeatFluxWattPerSquareMeter, -1)
	expected = "45 " + units.GetHeatFluxAbbreviation(units.HeatFluxWattPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestHeatFlux_EqualityAndComparison(t *testing.T) {
	factory := units.HeatFluxFactory{}
	a1, _ := factory.CreateHeatFlux(60, units.HeatFluxWattPerSquareMeter)
	a2, _ := factory.CreateHeatFlux(60, units.HeatFluxWattPerSquareMeter)
	a3, _ := factory.CreateHeatFlux(120, units.HeatFluxWattPerSquareMeter)

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

func TestHeatFlux_Arithmetic(t *testing.T) {
	factory := units.HeatFluxFactory{}
	a1, _ := factory.CreateHeatFlux(30, units.HeatFluxWattPerSquareMeter)
	a2, _ := factory.CreateHeatFlux(45, units.HeatFluxWattPerSquareMeter)

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