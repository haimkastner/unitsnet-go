// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificVolumeDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CubicMeterPerKilogram"}`
	
	factory := units.SpecificVolumeDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificVolumeCubicMeterPerKilogram {
		t.Errorf("expected unit %v, got %v", units.SpecificVolumeCubicMeterPerKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CubicMeterPerKilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificVolumeDto_ToJSON(t *testing.T) {
	dto := units.SpecificVolumeDto{
		Value: 45,
		Unit:  units.SpecificVolumeCubicMeterPerKilogram,
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
	if result["unit"].(string) != string(units.SpecificVolumeCubicMeterPerKilogram) {
		t.Errorf("expected unit %s, got %v", units.SpecificVolumeCubicMeterPerKilogram, result["unit"])
	}
}

func TestNewSpecificVolume_InvalidValue(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificVolume(math.NaN(), units.SpecificVolumeCubicMeterPerKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificVolume(math.Inf(1), units.SpecificVolumeCubicMeterPerKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificVolumeConversions(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificVolume(180, units.SpecificVolumeCubicMeterPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CubicMetersPerKilogram.
		// No expected conversion value provided for CubicMetersPerKilogram, verifying result is not NaN.
		result := a.CubicMetersPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMetersPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerPound.
		// No expected conversion value provided for CubicFeetPerPound, verifying result is not NaN.
		result := a.CubicFeetPerPound()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicFeetPerPound returned NaN")
		}
	}
	{
		// Test conversion to MillicubicMetersPerKilogram.
		// No expected conversion value provided for MillicubicMetersPerKilogram, verifying result is not NaN.
		result := a.MillicubicMetersPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillicubicMetersPerKilogram returned NaN")
		}
	}
}

func TestSpecificVolume_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	a, err := factory.CreateSpecificVolume(90, units.SpecificVolumeCubicMeterPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificVolumeCubicMeterPerKilogram {
		t.Errorf("expected default unit CubicMeterPerKilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificVolumeCubicMeterPerKilogram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificVolumeDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificVolumeCubicMeterPerKilogram {
		t.Errorf("expected unit CubicMeterPerKilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificVolumeToString(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	a, err := factory.CreateSpecificVolume(45, units.SpecificVolumeCubicMeterPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificVolumeCubicMeterPerKilogram, 2)
	expected := "45.00 " + units.GetSpecificVolumeAbbreviation(units.SpecificVolumeCubicMeterPerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificVolumeCubicMeterPerKilogram, -1)
	expected = "45 " + units.GetSpecificVolumeAbbreviation(units.SpecificVolumeCubicMeterPerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificVolume_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	a1, _ := factory.CreateSpecificVolume(60, units.SpecificVolumeCubicMeterPerKilogram)
	a2, _ := factory.CreateSpecificVolume(60, units.SpecificVolumeCubicMeterPerKilogram)
	a3, _ := factory.CreateSpecificVolume(120, units.SpecificVolumeCubicMeterPerKilogram)

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

func TestSpecificVolume_Arithmetic(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	a1, _ := factory.CreateSpecificVolume(30, units.SpecificVolumeCubicMeterPerKilogram)
	a2, _ := factory.CreateSpecificVolume(45, units.SpecificVolumeCubicMeterPerKilogram)

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