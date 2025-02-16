// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricCurrentGradientDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "AmperePerSecond"}`
	
	factory := units.ElectricCurrentGradientDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricCurrentGradientAmperePerSecond {
		t.Errorf("expected unit %v, got %v", units.ElectricCurrentGradientAmperePerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "AmperePerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricCurrentGradientDto_ToJSON(t *testing.T) {
	dto := units.ElectricCurrentGradientDto{
		Value: 45,
		Unit:  units.ElectricCurrentGradientAmperePerSecond,
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
	if result["unit"].(string) != string(units.ElectricCurrentGradientAmperePerSecond) {
		t.Errorf("expected unit %s, got %v", units.ElectricCurrentGradientAmperePerSecond, result["unit"])
	}
}

func TestNewElectricCurrentGradient_InvalidValue(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricCurrentGradient(math.NaN(), units.ElectricCurrentGradientAmperePerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricCurrentGradient(math.Inf(1), units.ElectricCurrentGradientAmperePerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricCurrentGradientConversions(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricCurrentGradient(180, units.ElectricCurrentGradientAmperePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to AmperesPerSecond.
		// No expected conversion value provided for AmperesPerSecond, verifying result is not NaN.
		result := a.AmperesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to AmperesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerMinute.
		// No expected conversion value provided for AmperesPerMinute, verifying result is not NaN.
		result := a.AmperesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to AmperesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerMillisecond.
		// No expected conversion value provided for AmperesPerMillisecond, verifying result is not NaN.
		result := a.AmperesPerMillisecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to AmperesPerMillisecond returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerMicrosecond.
		// No expected conversion value provided for AmperesPerMicrosecond, verifying result is not NaN.
		result := a.AmperesPerMicrosecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to AmperesPerMicrosecond returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerNanosecond.
		// No expected conversion value provided for AmperesPerNanosecond, verifying result is not NaN.
		result := a.AmperesPerNanosecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to AmperesPerNanosecond returned NaN")
		}
	}
	{
		// Test conversion to MilliamperesPerSecond.
		// No expected conversion value provided for MilliamperesPerSecond, verifying result is not NaN.
		result := a.MilliamperesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliamperesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MilliamperesPerMinute.
		// No expected conversion value provided for MilliamperesPerMinute, verifying result is not NaN.
		result := a.MilliamperesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliamperesPerMinute returned NaN")
		}
	}
}

func TestElectricCurrentGradient_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	a, err := factory.CreateElectricCurrentGradient(90, units.ElectricCurrentGradientAmperePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricCurrentGradientAmperePerSecond {
		t.Errorf("expected default unit AmperePerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricCurrentGradientAmperePerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricCurrentGradientDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricCurrentGradientAmperePerSecond {
		t.Errorf("expected unit AmperePerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricCurrentGradientToString(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	a, err := factory.CreateElectricCurrentGradient(45, units.ElectricCurrentGradientAmperePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricCurrentGradientAmperePerSecond, 2)
	expected := "45.00 " + units.GetElectricCurrentGradientAbbreviation(units.ElectricCurrentGradientAmperePerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricCurrentGradientAmperePerSecond, -1)
	expected = "45 " + units.GetElectricCurrentGradientAbbreviation(units.ElectricCurrentGradientAmperePerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricCurrentGradient_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	a1, _ := factory.CreateElectricCurrentGradient(60, units.ElectricCurrentGradientAmperePerSecond)
	a2, _ := factory.CreateElectricCurrentGradient(60, units.ElectricCurrentGradientAmperePerSecond)
	a3, _ := factory.CreateElectricCurrentGradient(120, units.ElectricCurrentGradientAmperePerSecond)

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

func TestElectricCurrentGradient_Arithmetic(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	a1, _ := factory.CreateElectricCurrentGradient(30, units.ElectricCurrentGradientAmperePerSecond)
	a2, _ := factory.CreateElectricCurrentGradient(45, units.ElectricCurrentGradientAmperePerSecond)

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