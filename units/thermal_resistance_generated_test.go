// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestThermalResistanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeterKelvinPerKilowatt"}`
	
	factory := units.ThermalResistanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ThermalResistanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected unit %v, got %v", units.ThermalResistanceSquareMeterKelvinPerKilowatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeterKelvinPerKilowatt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestThermalResistanceDto_ToJSON(t *testing.T) {
	dto := units.ThermalResistanceDto{
		Value: 45,
		Unit:  units.ThermalResistanceSquareMeterKelvinPerKilowatt,
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
	if result["unit"].(string) != string(units.ThermalResistanceSquareMeterKelvinPerKilowatt) {
		t.Errorf("expected unit %s, got %v", units.ThermalResistanceSquareMeterKelvinPerKilowatt, result["unit"])
	}
}

func TestNewThermalResistance_InvalidValue(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateThermalResistance(math.NaN(), units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateThermalResistance(math.Inf(1), units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestThermalResistanceConversions(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateThermalResistance(180, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SquareMeterKelvinsPerKilowatt.
		// No expected conversion value provided for SquareMeterKelvinsPerKilowatt, verifying result is not NaN.
		result := a.SquareMeterKelvinsPerKilowatt()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeterKelvinsPerKilowatt returned NaN")
		}
	}
	{
		// Test conversion to SquareMeterKelvinsPerWatt.
		// No expected conversion value provided for SquareMeterKelvinsPerWatt, verifying result is not NaN.
		result := a.SquareMeterKelvinsPerWatt()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeterKelvinsPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareMeterDegreesCelsiusPerWatt.
		// No expected conversion value provided for SquareMeterDegreesCelsiusPerWatt, verifying result is not NaN.
		result := a.SquareMeterDegreesCelsiusPerWatt()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeterDegreesCelsiusPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeterKelvinsPerWatt.
		// No expected conversion value provided for SquareCentimeterKelvinsPerWatt, verifying result is not NaN.
		result := a.SquareCentimeterKelvinsPerWatt()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareCentimeterKelvinsPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeterHourDegreesCelsiusPerKilocalorie.
		// No expected conversion value provided for SquareCentimeterHourDegreesCelsiusPerKilocalorie, verifying result is not NaN.
		result := a.SquareCentimeterHourDegreesCelsiusPerKilocalorie()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareCentimeterHourDegreesCelsiusPerKilocalorie returned NaN")
		}
	}
	{
		// Test conversion to HourSquareFeetDegreesFahrenheitPerBtu.
		// No expected conversion value provided for HourSquareFeetDegreesFahrenheitPerBtu, verifying result is not NaN.
		result := a.HourSquareFeetDegreesFahrenheitPerBtu()
		if math.IsNaN(result) {
			t.Errorf("conversion to HourSquareFeetDegreesFahrenheitPerBtu returned NaN")
		}
	}
}

func TestThermalResistance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a, err := factory.CreateThermalResistance(90, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ThermalResistanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected default unit SquareMeterKelvinPerKilowatt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ThermalResistanceSquareMeterKelvinPerKilowatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ThermalResistanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ThermalResistanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected unit SquareMeterKelvinPerKilowatt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestThermalResistanceToString(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a, err := factory.CreateThermalResistance(45, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ThermalResistanceSquareMeterKelvinPerKilowatt, 2)
	expected := "45.00 " + units.GetThermalResistanceAbbreviation(units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ThermalResistanceSquareMeterKelvinPerKilowatt, -1)
	expected = "45 " + units.GetThermalResistanceAbbreviation(units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestThermalResistance_EqualityAndComparison(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a1, _ := factory.CreateThermalResistance(60, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	a2, _ := factory.CreateThermalResistance(60, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	a3, _ := factory.CreateThermalResistance(120, units.ThermalResistanceSquareMeterKelvinPerKilowatt)

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

func TestThermalResistance_Arithmetic(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a1, _ := factory.CreateThermalResistance(30, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	a2, _ := factory.CreateThermalResistance(45, units.ThermalResistanceSquareMeterKelvinPerKilowatt)

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