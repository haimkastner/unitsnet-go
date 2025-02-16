// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricConductivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SiemensPerMeter"}`
	
	factory := units.ElectricConductivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricConductivitySiemensPerMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricConductivitySiemensPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SiemensPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricConductivityDto_ToJSON(t *testing.T) {
	dto := units.ElectricConductivityDto{
		Value: 45,
		Unit:  units.ElectricConductivitySiemensPerMeter,
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
	if result["unit"].(string) != string(units.ElectricConductivitySiemensPerMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricConductivitySiemensPerMeter, result["unit"])
	}
}

func TestNewElectricConductivity_InvalidValue(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricConductivity(math.NaN(), units.ElectricConductivitySiemensPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricConductivity(math.Inf(1), units.ElectricConductivitySiemensPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricConductivityConversions(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricConductivity(180, units.ElectricConductivitySiemensPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SiemensPerMeter.
		// No expected conversion value provided for SiemensPerMeter, verifying result is not NaN.
		result := a.SiemensPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to SiemensPerMeter returned NaN")
		}
	}
	{
		// Test conversion to SiemensPerInch.
		// No expected conversion value provided for SiemensPerInch, verifying result is not NaN.
		result := a.SiemensPerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to SiemensPerInch returned NaN")
		}
	}
	{
		// Test conversion to SiemensPerFoot.
		// No expected conversion value provided for SiemensPerFoot, verifying result is not NaN.
		result := a.SiemensPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to SiemensPerFoot returned NaN")
		}
	}
	{
		// Test conversion to SiemensPerCentimeter.
		// No expected conversion value provided for SiemensPerCentimeter, verifying result is not NaN.
		result := a.SiemensPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to SiemensPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicrosiemensPerCentimeter.
		// No expected conversion value provided for MicrosiemensPerCentimeter, verifying result is not NaN.
		result := a.MicrosiemensPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrosiemensPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MillisiemensPerCentimeter.
		// No expected conversion value provided for MillisiemensPerCentimeter, verifying result is not NaN.
		result := a.MillisiemensPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillisiemensPerCentimeter returned NaN")
		}
	}
}

func TestElectricConductivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	a, err := factory.CreateElectricConductivity(90, units.ElectricConductivitySiemensPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricConductivitySiemensPerMeter {
		t.Errorf("expected default unit SiemensPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricConductivitySiemensPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricConductivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricConductivitySiemensPerMeter {
		t.Errorf("expected unit SiemensPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricConductivityToString(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	a, err := factory.CreateElectricConductivity(45, units.ElectricConductivitySiemensPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricConductivitySiemensPerMeter, 2)
	expected := "45.00 " + units.GetElectricConductivityAbbreviation(units.ElectricConductivitySiemensPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricConductivitySiemensPerMeter, -1)
	expected = "45 " + units.GetElectricConductivityAbbreviation(units.ElectricConductivitySiemensPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricConductivity_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	a1, _ := factory.CreateElectricConductivity(60, units.ElectricConductivitySiemensPerMeter)
	a2, _ := factory.CreateElectricConductivity(60, units.ElectricConductivitySiemensPerMeter)
	a3, _ := factory.CreateElectricConductivity(120, units.ElectricConductivitySiemensPerMeter)

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

func TestElectricConductivity_Arithmetic(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	a1, _ := factory.CreateElectricConductivity(30, units.ElectricConductivitySiemensPerMeter)
	a2, _ := factory.CreateElectricConductivity(45, units.ElectricConductivitySiemensPerMeter)

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