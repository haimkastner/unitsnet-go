// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRadiationExposureDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CoulombPerKilogram"}`
	
	factory := units.RadiationExposureDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RadiationExposureCoulombPerKilogram {
		t.Errorf("expected unit %v, got %v", units.RadiationExposureCoulombPerKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CoulombPerKilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRadiationExposureDto_ToJSON(t *testing.T) {
	dto := units.RadiationExposureDto{
		Value: 45,
		Unit:  units.RadiationExposureCoulombPerKilogram,
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
	if result["unit"].(string) != string(units.RadiationExposureCoulombPerKilogram) {
		t.Errorf("expected unit %s, got %v", units.RadiationExposureCoulombPerKilogram, result["unit"])
	}
}

func TestNewRadiationExposure_InvalidValue(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRadiationExposure(math.NaN(), units.RadiationExposureCoulombPerKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRadiationExposure(math.Inf(1), units.RadiationExposureCoulombPerKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRadiationExposureConversions(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRadiationExposure(180, units.RadiationExposureCoulombPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CoulombsPerKilogram.
		// No expected conversion value provided for CoulombsPerKilogram, verifying result is not NaN.
		result := a.CoulombsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to CoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to Roentgens.
		// No expected conversion value provided for Roentgens, verifying result is not NaN.
		result := a.Roentgens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Roentgens returned NaN")
		}
	}
	{
		// Test conversion to PicocoulombsPerKilogram.
		// No expected conversion value provided for PicocoulombsPerKilogram, verifying result is not NaN.
		result := a.PicocoulombsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicocoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to NanocoulombsPerKilogram.
		// No expected conversion value provided for NanocoulombsPerKilogram, verifying result is not NaN.
		result := a.NanocoulombsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanocoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MicrocoulombsPerKilogram.
		// No expected conversion value provided for MicrocoulombsPerKilogram, verifying result is not NaN.
		result := a.MicrocoulombsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrocoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MillicoulombsPerKilogram.
		// No expected conversion value provided for MillicoulombsPerKilogram, verifying result is not NaN.
		result := a.MillicoulombsPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillicoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to Microroentgens.
		// No expected conversion value provided for Microroentgens, verifying result is not NaN.
		result := a.Microroentgens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microroentgens returned NaN")
		}
	}
	{
		// Test conversion to Milliroentgens.
		// No expected conversion value provided for Milliroentgens, verifying result is not NaN.
		result := a.Milliroentgens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliroentgens returned NaN")
		}
	}
}

func TestRadiationExposure_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	a, err := factory.CreateRadiationExposure(90, units.RadiationExposureCoulombPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RadiationExposureCoulombPerKilogram {
		t.Errorf("expected default unit CoulombPerKilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RadiationExposureCoulombPerKilogram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RadiationExposureDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RadiationExposureCoulombPerKilogram {
		t.Errorf("expected unit CoulombPerKilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRadiationExposureToString(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	a, err := factory.CreateRadiationExposure(45, units.RadiationExposureCoulombPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RadiationExposureCoulombPerKilogram, 2)
	expected := "45.00 " + units.GetRadiationExposureAbbreviation(units.RadiationExposureCoulombPerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RadiationExposureCoulombPerKilogram, -1)
	expected = "45 " + units.GetRadiationExposureAbbreviation(units.RadiationExposureCoulombPerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRadiationExposure_EqualityAndComparison(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	a1, _ := factory.CreateRadiationExposure(60, units.RadiationExposureCoulombPerKilogram)
	a2, _ := factory.CreateRadiationExposure(60, units.RadiationExposureCoulombPerKilogram)
	a3, _ := factory.CreateRadiationExposure(120, units.RadiationExposureCoulombPerKilogram)

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

func TestRadiationExposure_Arithmetic(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	a1, _ := factory.CreateRadiationExposure(30, units.RadiationExposureCoulombPerKilogram)
	a2, _ := factory.CreateRadiationExposure(45, units.RadiationExposureCoulombPerKilogram)

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