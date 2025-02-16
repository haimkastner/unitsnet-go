// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTemperatureGradientDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KelvinPerMeter"}`
	
	factory := units.TemperatureGradientDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TemperatureGradientKelvinPerMeter {
		t.Errorf("expected unit %v, got %v", units.TemperatureGradientKelvinPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KelvinPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTemperatureGradientDto_ToJSON(t *testing.T) {
	dto := units.TemperatureGradientDto{
		Value: 45,
		Unit:  units.TemperatureGradientKelvinPerMeter,
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
	if result["unit"].(string) != string(units.TemperatureGradientKelvinPerMeter) {
		t.Errorf("expected unit %s, got %v", units.TemperatureGradientKelvinPerMeter, result["unit"])
	}
}

func TestNewTemperatureGradient_InvalidValue(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTemperatureGradient(math.NaN(), units.TemperatureGradientKelvinPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTemperatureGradient(math.Inf(1), units.TemperatureGradientKelvinPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTemperatureGradientConversions(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTemperatureGradient(180, units.TemperatureGradientKelvinPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to KelvinsPerMeter.
		// No expected conversion value provided for KelvinsPerMeter, verifying result is not NaN.
		result := a.KelvinsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KelvinsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelciusPerMeter.
		// No expected conversion value provided for DegreesCelciusPerMeter, verifying result is not NaN.
		result := a.DegreesCelciusPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesCelciusPerMeter returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheitPerFoot.
		// No expected conversion value provided for DegreesFahrenheitPerFoot, verifying result is not NaN.
		result := a.DegreesFahrenheitPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesFahrenheitPerFoot returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelciusPerKilometer.
		// No expected conversion value provided for DegreesCelciusPerKilometer, verifying result is not NaN.
		result := a.DegreesCelciusPerKilometer()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesCelciusPerKilometer returned NaN")
		}
	}
}

func TestTemperatureGradient_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	a, err := factory.CreateTemperatureGradient(90, units.TemperatureGradientKelvinPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TemperatureGradientKelvinPerMeter {
		t.Errorf("expected default unit KelvinPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TemperatureGradientKelvinPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TemperatureGradientDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TemperatureGradientKelvinPerMeter {
		t.Errorf("expected unit KelvinPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTemperatureGradientToString(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	a, err := factory.CreateTemperatureGradient(45, units.TemperatureGradientKelvinPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TemperatureGradientKelvinPerMeter, 2)
	expected := "45.00 " + units.GetTemperatureGradientAbbreviation(units.TemperatureGradientKelvinPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TemperatureGradientKelvinPerMeter, -1)
	expected = "45 " + units.GetTemperatureGradientAbbreviation(units.TemperatureGradientKelvinPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTemperatureGradient_EqualityAndComparison(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	a1, _ := factory.CreateTemperatureGradient(60, units.TemperatureGradientKelvinPerMeter)
	a2, _ := factory.CreateTemperatureGradient(60, units.TemperatureGradientKelvinPerMeter)
	a3, _ := factory.CreateTemperatureGradient(120, units.TemperatureGradientKelvinPerMeter)

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

func TestTemperatureGradient_Arithmetic(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	a1, _ := factory.CreateTemperatureGradient(30, units.TemperatureGradientKelvinPerMeter)
	a2, _ := factory.CreateTemperatureGradient(45, units.TemperatureGradientKelvinPerMeter)

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