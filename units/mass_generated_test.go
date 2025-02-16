// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Kilogram"}`
	
	factory := units.MassDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassKilogram {
		t.Errorf("expected unit %v, got %v", units.MassKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Kilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassDto_ToJSON(t *testing.T) {
	dto := units.MassDto{
		Value: 45,
		Unit:  units.MassKilogram,
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
	if result["unit"].(string) != string(units.MassKilogram) {
		t.Errorf("expected unit %s, got %v", units.MassKilogram, result["unit"])
	}
}

func TestNewMass_InvalidValue(t *testing.T) {
	factory := units.MassFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMass(math.NaN(), units.MassKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMass(math.Inf(1), units.MassKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassConversions(t *testing.T) {
	factory := units.MassFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMass(180, units.MassKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Grams.
		// No expected conversion value provided for Grams, verifying result is not NaN.
		result := a.Grams()
		if math.IsNaN(result) {
			t.Errorf("conversion to Grams returned NaN")
		}
	}
	{
		// Test conversion to Tonnes.
		// No expected conversion value provided for Tonnes, verifying result is not NaN.
		result := a.Tonnes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Tonnes returned NaN")
		}
	}
	{
		// Test conversion to ShortTons.
		// No expected conversion value provided for ShortTons, verifying result is not NaN.
		result := a.ShortTons()
		if math.IsNaN(result) {
			t.Errorf("conversion to ShortTons returned NaN")
		}
	}
	{
		// Test conversion to LongTons.
		// No expected conversion value provided for LongTons, verifying result is not NaN.
		result := a.LongTons()
		if math.IsNaN(result) {
			t.Errorf("conversion to LongTons returned NaN")
		}
	}
	{
		// Test conversion to Pounds.
		// No expected conversion value provided for Pounds, verifying result is not NaN.
		result := a.Pounds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Pounds returned NaN")
		}
	}
	{
		// Test conversion to Ounces.
		// No expected conversion value provided for Ounces, verifying result is not NaN.
		result := a.Ounces()
		if math.IsNaN(result) {
			t.Errorf("conversion to Ounces returned NaN")
		}
	}
	{
		// Test conversion to Slugs.
		// No expected conversion value provided for Slugs, verifying result is not NaN.
		result := a.Slugs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Slugs returned NaN")
		}
	}
	{
		// Test conversion to Stone.
		// No expected conversion value provided for Stone, verifying result is not NaN.
		result := a.Stone()
		if math.IsNaN(result) {
			t.Errorf("conversion to Stone returned NaN")
		}
	}
	{
		// Test conversion to ShortHundredweight.
		// No expected conversion value provided for ShortHundredweight, verifying result is not NaN.
		result := a.ShortHundredweight()
		if math.IsNaN(result) {
			t.Errorf("conversion to ShortHundredweight returned NaN")
		}
	}
	{
		// Test conversion to LongHundredweight.
		// No expected conversion value provided for LongHundredweight, verifying result is not NaN.
		result := a.LongHundredweight()
		if math.IsNaN(result) {
			t.Errorf("conversion to LongHundredweight returned NaN")
		}
	}
	{
		// Test conversion to Grains.
		// No expected conversion value provided for Grains, verifying result is not NaN.
		result := a.Grains()
		if math.IsNaN(result) {
			t.Errorf("conversion to Grains returned NaN")
		}
	}
	{
		// Test conversion to SolarMasses.
		// No expected conversion value provided for SolarMasses, verifying result is not NaN.
		result := a.SolarMasses()
		if math.IsNaN(result) {
			t.Errorf("conversion to SolarMasses returned NaN")
		}
	}
	{
		// Test conversion to EarthMasses.
		// No expected conversion value provided for EarthMasses, verifying result is not NaN.
		result := a.EarthMasses()
		if math.IsNaN(result) {
			t.Errorf("conversion to EarthMasses returned NaN")
		}
	}
	{
		// Test conversion to Femtograms.
		// No expected conversion value provided for Femtograms, verifying result is not NaN.
		result := a.Femtograms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Femtograms returned NaN")
		}
	}
	{
		// Test conversion to Picograms.
		// No expected conversion value provided for Picograms, verifying result is not NaN.
		result := a.Picograms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picograms returned NaN")
		}
	}
	{
		// Test conversion to Nanograms.
		// No expected conversion value provided for Nanograms, verifying result is not NaN.
		result := a.Nanograms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanograms returned NaN")
		}
	}
	{
		// Test conversion to Micrograms.
		// No expected conversion value provided for Micrograms, verifying result is not NaN.
		result := a.Micrograms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micrograms returned NaN")
		}
	}
	{
		// Test conversion to Milligrams.
		// No expected conversion value provided for Milligrams, verifying result is not NaN.
		result := a.Milligrams()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milligrams returned NaN")
		}
	}
	{
		// Test conversion to Centigrams.
		// No expected conversion value provided for Centigrams, verifying result is not NaN.
		result := a.Centigrams()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centigrams returned NaN")
		}
	}
	{
		// Test conversion to Decigrams.
		// No expected conversion value provided for Decigrams, verifying result is not NaN.
		result := a.Decigrams()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decigrams returned NaN")
		}
	}
	{
		// Test conversion to Decagrams.
		// No expected conversion value provided for Decagrams, verifying result is not NaN.
		result := a.Decagrams()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decagrams returned NaN")
		}
	}
	{
		// Test conversion to Hectograms.
		// No expected conversion value provided for Hectograms, verifying result is not NaN.
		result := a.Hectograms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hectograms returned NaN")
		}
	}
	{
		// Test conversion to Kilograms.
		// No expected conversion value provided for Kilograms, verifying result is not NaN.
		result := a.Kilograms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilograms returned NaN")
		}
	}
	{
		// Test conversion to Kilotonnes.
		// No expected conversion value provided for Kilotonnes, verifying result is not NaN.
		result := a.Kilotonnes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilotonnes returned NaN")
		}
	}
	{
		// Test conversion to Megatonnes.
		// No expected conversion value provided for Megatonnes, verifying result is not NaN.
		result := a.Megatonnes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megatonnes returned NaN")
		}
	}
	{
		// Test conversion to Kilopounds.
		// No expected conversion value provided for Kilopounds, verifying result is not NaN.
		result := a.Kilopounds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilopounds returned NaN")
		}
	}
	{
		// Test conversion to Megapounds.
		// No expected conversion value provided for Megapounds, verifying result is not NaN.
		result := a.Megapounds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megapounds returned NaN")
		}
	}
}

func TestMass_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassFactory{}
	a, err := factory.CreateMass(90, units.MassKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassKilogram {
		t.Errorf("expected default unit Kilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassGram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassKilogram {
		t.Errorf("expected unit Kilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassToString(t *testing.T) {
	factory := units.MassFactory{}
	a, err := factory.CreateMass(45, units.MassKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassKilogram, 2)
	expected := "45.00 " + units.GetMassAbbreviation(units.MassKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassKilogram, -1)
	expected = "45 " + units.GetMassAbbreviation(units.MassKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMass_EqualityAndComparison(t *testing.T) {
	factory := units.MassFactory{}
	a1, _ := factory.CreateMass(60, units.MassKilogram)
	a2, _ := factory.CreateMass(60, units.MassKilogram)
	a3, _ := factory.CreateMass(120, units.MassKilogram)

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

func TestMass_Arithmetic(t *testing.T) {
	factory := units.MassFactory{}
	a1, _ := factory.CreateMass(30, units.MassKilogram)
	a2, _ := factory.CreateMass(45, units.MassKilogram)

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