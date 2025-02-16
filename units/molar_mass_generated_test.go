// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarMassDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerMole"}`
	
	factory := units.MolarMassDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarMassKilogramPerMole {
		t.Errorf("expected unit %v, got %v", units.MolarMassKilogramPerMole, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerMole"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarMassDto_ToJSON(t *testing.T) {
	dto := units.MolarMassDto{
		Value: 45,
		Unit:  units.MolarMassKilogramPerMole,
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
	if result["unit"].(string) != string(units.MolarMassKilogramPerMole) {
		t.Errorf("expected unit %s, got %v", units.MolarMassKilogramPerMole, result["unit"])
	}
}

func TestNewMolarMass_InvalidValue(t *testing.T) {
	factory := units.MolarMassFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarMass(math.NaN(), units.MolarMassKilogramPerMole)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarMass(math.Inf(1), units.MolarMassKilogramPerMole)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarMassConversions(t *testing.T) {
	factory := units.MolarMassFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarMass(180, units.MolarMassKilogramPerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerMole.
		// No expected conversion value provided for GramsPerMole, verifying result is not NaN.
		result := a.GramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerKilomole.
		// No expected conversion value provided for KilogramsPerKilomole, verifying result is not NaN.
		result := a.KilogramsPerKilomole()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerKilomole returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerMole.
		// No expected conversion value provided for PoundsPerMole, verifying result is not NaN.
		result := a.PoundsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerMole returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerMole.
		// No expected conversion value provided for NanogramsPerMole, verifying result is not NaN.
		result := a.NanogramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMole.
		// No expected conversion value provided for MicrogramsPerMole, verifying result is not NaN.
		result := a.MicrogramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMole.
		// No expected conversion value provided for MilligramsPerMole, verifying result is not NaN.
		result := a.MilligramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerMole.
		// No expected conversion value provided for CentigramsPerMole, verifying result is not NaN.
		result := a.CentigramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerMole.
		// No expected conversion value provided for DecigramsPerMole, verifying result is not NaN.
		result := a.DecigramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to DecagramsPerMole.
		// No expected conversion value provided for DecagramsPerMole, verifying result is not NaN.
		result := a.DecagramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecagramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to HectogramsPerMole.
		// No expected conversion value provided for HectogramsPerMole, verifying result is not NaN.
		result := a.HectogramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectogramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerMole.
		// No expected conversion value provided for KilogramsPerMole, verifying result is not NaN.
		result := a.KilogramsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerMole.
		// No expected conversion value provided for KilopoundsPerMole, verifying result is not NaN.
		result := a.KilopoundsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsPerMole returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsPerMole.
		// No expected conversion value provided for MegapoundsPerMole, verifying result is not NaN.
		result := a.MegapoundsPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundsPerMole returned NaN")
		}
	}
}

func TestMolarMass_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarMassFactory{}
	a, err := factory.CreateMolarMass(90, units.MolarMassKilogramPerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarMassKilogramPerMole {
		t.Errorf("expected default unit KilogramPerMole, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarMassGramPerMole
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarMassDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarMassKilogramPerMole {
		t.Errorf("expected unit KilogramPerMole, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarMassToString(t *testing.T) {
	factory := units.MolarMassFactory{}
	a, err := factory.CreateMolarMass(45, units.MolarMassKilogramPerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarMassKilogramPerMole, 2)
	expected := "45.00 " + units.GetMolarMassAbbreviation(units.MolarMassKilogramPerMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarMassKilogramPerMole, -1)
	expected = "45 " + units.GetMolarMassAbbreviation(units.MolarMassKilogramPerMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarMass_EqualityAndComparison(t *testing.T) {
	factory := units.MolarMassFactory{}
	a1, _ := factory.CreateMolarMass(60, units.MolarMassKilogramPerMole)
	a2, _ := factory.CreateMolarMass(60, units.MolarMassKilogramPerMole)
	a3, _ := factory.CreateMolarMass(120, units.MolarMassKilogramPerMole)

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

func TestMolarMass_Arithmetic(t *testing.T) {
	factory := units.MolarMassFactory{}
	a1, _ := factory.CreateMolarMass(30, units.MolarMassKilogramPerMole)
	a2, _ := factory.CreateMolarMass(45, units.MolarMassKilogramPerMole)

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