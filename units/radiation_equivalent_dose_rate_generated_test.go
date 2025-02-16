// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRadiationEquivalentDoseRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SievertPerSecond"}`
	
	factory := units.RadiationEquivalentDoseRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RadiationEquivalentDoseRateSievertPerSecond {
		t.Errorf("expected unit %v, got %v", units.RadiationEquivalentDoseRateSievertPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SievertPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRadiationEquivalentDoseRateDto_ToJSON(t *testing.T) {
	dto := units.RadiationEquivalentDoseRateDto{
		Value: 45,
		Unit:  units.RadiationEquivalentDoseRateSievertPerSecond,
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
	if result["unit"].(string) != string(units.RadiationEquivalentDoseRateSievertPerSecond) {
		t.Errorf("expected unit %s, got %v", units.RadiationEquivalentDoseRateSievertPerSecond, result["unit"])
	}
}

func TestNewRadiationEquivalentDoseRate_InvalidValue(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRadiationEquivalentDoseRate(math.NaN(), units.RadiationEquivalentDoseRateSievertPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRadiationEquivalentDoseRate(math.Inf(1), units.RadiationEquivalentDoseRateSievertPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRadiationEquivalentDoseRateConversions(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRadiationEquivalentDoseRate(180, units.RadiationEquivalentDoseRateSievertPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SievertsPerHour.
		// No expected conversion value provided for SievertsPerHour, verifying result is not NaN.
		result := a.SievertsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to SievertsPerHour returned NaN")
		}
	}
	{
		// Test conversion to SievertsPerSecond.
		// No expected conversion value provided for SievertsPerSecond, verifying result is not NaN.
		result := a.SievertsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to SievertsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to RoentgensEquivalentManPerHour.
		// No expected conversion value provided for RoentgensEquivalentManPerHour, verifying result is not NaN.
		result := a.RoentgensEquivalentManPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to RoentgensEquivalentManPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanosievertsPerHour.
		// No expected conversion value provided for NanosievertsPerHour, verifying result is not NaN.
		result := a.NanosievertsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanosievertsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MicrosievertsPerHour.
		// No expected conversion value provided for MicrosievertsPerHour, verifying result is not NaN.
		result := a.MicrosievertsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrosievertsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MillisievertsPerHour.
		// No expected conversion value provided for MillisievertsPerHour, verifying result is not NaN.
		result := a.MillisievertsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillisievertsPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanosievertsPerSecond.
		// No expected conversion value provided for NanosievertsPerSecond, verifying result is not NaN.
		result := a.NanosievertsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanosievertsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrosievertsPerSecond.
		// No expected conversion value provided for MicrosievertsPerSecond, verifying result is not NaN.
		result := a.MicrosievertsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrosievertsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillisievertsPerSecond.
		// No expected conversion value provided for MillisievertsPerSecond, verifying result is not NaN.
		result := a.MillisievertsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillisievertsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MilliroentgensEquivalentManPerHour.
		// No expected conversion value provided for MilliroentgensEquivalentManPerHour, verifying result is not NaN.
		result := a.MilliroentgensEquivalentManPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliroentgensEquivalentManPerHour returned NaN")
		}
	}
}

func TestRadiationEquivalentDoseRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	a, err := factory.CreateRadiationEquivalentDoseRate(90, units.RadiationEquivalentDoseRateSievertPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RadiationEquivalentDoseRateSievertPerSecond {
		t.Errorf("expected default unit SievertPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RadiationEquivalentDoseRateSievertPerHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RadiationEquivalentDoseRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RadiationEquivalentDoseRateSievertPerSecond {
		t.Errorf("expected unit SievertPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRadiationEquivalentDoseRateToString(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	a, err := factory.CreateRadiationEquivalentDoseRate(45, units.RadiationEquivalentDoseRateSievertPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RadiationEquivalentDoseRateSievertPerSecond, 2)
	expected := "45.00 " + units.GetRadiationEquivalentDoseRateAbbreviation(units.RadiationEquivalentDoseRateSievertPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RadiationEquivalentDoseRateSievertPerSecond, -1)
	expected = "45 " + units.GetRadiationEquivalentDoseRateAbbreviation(units.RadiationEquivalentDoseRateSievertPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRadiationEquivalentDoseRate_EqualityAndComparison(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	a1, _ := factory.CreateRadiationEquivalentDoseRate(60, units.RadiationEquivalentDoseRateSievertPerSecond)
	a2, _ := factory.CreateRadiationEquivalentDoseRate(60, units.RadiationEquivalentDoseRateSievertPerSecond)
	a3, _ := factory.CreateRadiationEquivalentDoseRate(120, units.RadiationEquivalentDoseRateSievertPerSecond)

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

func TestRadiationEquivalentDoseRate_Arithmetic(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	a1, _ := factory.CreateRadiationEquivalentDoseRate(30, units.RadiationEquivalentDoseRateSievertPerSecond)
	a2, _ := factory.CreateRadiationEquivalentDoseRate(45, units.RadiationEquivalentDoseRateSievertPerSecond)

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