// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassMomentOfInertiaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramSquareMeter"}`
	
	factory := units.MassMomentOfInertiaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassMomentOfInertiaKilogramSquareMeter {
		t.Errorf("expected unit %v, got %v", units.MassMomentOfInertiaKilogramSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassMomentOfInertiaDto_ToJSON(t *testing.T) {
	dto := units.MassMomentOfInertiaDto{
		Value: 45,
		Unit:  units.MassMomentOfInertiaKilogramSquareMeter,
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
	if result["unit"].(string) != string(units.MassMomentOfInertiaKilogramSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.MassMomentOfInertiaKilogramSquareMeter, result["unit"])
	}
}

func TestNewMassMomentOfInertia_InvalidValue(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMassMomentOfInertia(math.NaN(), units.MassMomentOfInertiaKilogramSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMassMomentOfInertia(math.Inf(1), units.MassMomentOfInertiaKilogramSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassMomentOfInertiaConversions(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMassMomentOfInertia(180, units.MassMomentOfInertiaKilogramSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramSquareMeters.
		// No expected conversion value provided for GramSquareMeters, verifying result is not NaN.
		result := a.GramSquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to GramSquareDecimeters.
		// No expected conversion value provided for GramSquareDecimeters, verifying result is not NaN.
		result := a.GramSquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to GramSquareCentimeters.
		// No expected conversion value provided for GramSquareCentimeters, verifying result is not NaN.
		result := a.GramSquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to GramSquareMillimeters.
		// No expected conversion value provided for GramSquareMillimeters, verifying result is not NaN.
		result := a.GramSquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramSquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneSquareMeters.
		// No expected conversion value provided for TonneSquareMeters, verifying result is not NaN.
		result := a.TonneSquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to TonneSquareDecimeters.
		// No expected conversion value provided for TonneSquareDecimeters, verifying result is not NaN.
		result := a.TonneSquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneSquareCentimeters.
		// No expected conversion value provided for TonneSquareCentimeters, verifying result is not NaN.
		result := a.TonneSquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneSquareMilimeters.
		// No expected conversion value provided for TonneSquareMilimeters, verifying result is not NaN.
		result := a.TonneSquareMilimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneSquareMilimeters returned NaN")
		}
	}
	{
		// Test conversion to PoundSquareFeet.
		// No expected conversion value provided for PoundSquareFeet, verifying result is not NaN.
		result := a.PoundSquareFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundSquareFeet returned NaN")
		}
	}
	{
		// Test conversion to PoundSquareInches.
		// No expected conversion value provided for PoundSquareInches, verifying result is not NaN.
		result := a.PoundSquareInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundSquareInches returned NaN")
		}
	}
	{
		// Test conversion to SlugSquareFeet.
		// No expected conversion value provided for SlugSquareFeet, verifying result is not NaN.
		result := a.SlugSquareFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugSquareFeet returned NaN")
		}
	}
	{
		// Test conversion to SlugSquareInches.
		// No expected conversion value provided for SlugSquareInches, verifying result is not NaN.
		result := a.SlugSquareInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugSquareInches returned NaN")
		}
	}
	{
		// Test conversion to MilligramSquareMeters.
		// No expected conversion value provided for MilligramSquareMeters, verifying result is not NaN.
		result := a.MilligramSquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramSquareMeters.
		// No expected conversion value provided for KilogramSquareMeters, verifying result is not NaN.
		result := a.KilogramSquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to MilligramSquareDecimeters.
		// No expected conversion value provided for MilligramSquareDecimeters, verifying result is not NaN.
		result := a.MilligramSquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramSquareDecimeters.
		// No expected conversion value provided for KilogramSquareDecimeters, verifying result is not NaN.
		result := a.KilogramSquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligramSquareCentimeters.
		// No expected conversion value provided for MilligramSquareCentimeters, verifying result is not NaN.
		result := a.MilligramSquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramSquareCentimeters.
		// No expected conversion value provided for KilogramSquareCentimeters, verifying result is not NaN.
		result := a.KilogramSquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligramSquareMillimeters.
		// No expected conversion value provided for MilligramSquareMillimeters, verifying result is not NaN.
		result := a.MilligramSquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramSquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramSquareMillimeters.
		// No expected conversion value provided for KilogramSquareMillimeters, verifying result is not NaN.
		result := a.KilogramSquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramSquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to KilotonneSquareMeters.
		// No expected conversion value provided for KilotonneSquareMeters, verifying result is not NaN.
		result := a.KilotonneSquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilotonneSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to MegatonneSquareMeters.
		// No expected conversion value provided for MegatonneSquareMeters, verifying result is not NaN.
		result := a.MegatonneSquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegatonneSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to KilotonneSquareDecimeters.
		// No expected conversion value provided for KilotonneSquareDecimeters, verifying result is not NaN.
		result := a.KilotonneSquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilotonneSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to MegatonneSquareDecimeters.
		// No expected conversion value provided for MegatonneSquareDecimeters, verifying result is not NaN.
		result := a.MegatonneSquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegatonneSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to KilotonneSquareCentimeters.
		// No expected conversion value provided for KilotonneSquareCentimeters, verifying result is not NaN.
		result := a.KilotonneSquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilotonneSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MegatonneSquareCentimeters.
		// No expected conversion value provided for MegatonneSquareCentimeters, verifying result is not NaN.
		result := a.MegatonneSquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegatonneSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to KilotonneSquareMilimeters.
		// No expected conversion value provided for KilotonneSquareMilimeters, verifying result is not NaN.
		result := a.KilotonneSquareMilimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilotonneSquareMilimeters returned NaN")
		}
	}
	{
		// Test conversion to MegatonneSquareMilimeters.
		// No expected conversion value provided for MegatonneSquareMilimeters, verifying result is not NaN.
		result := a.MegatonneSquareMilimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegatonneSquareMilimeters returned NaN")
		}
	}
}

func TestMassMomentOfInertia_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	a, err := factory.CreateMassMomentOfInertia(90, units.MassMomentOfInertiaKilogramSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassMomentOfInertiaKilogramSquareMeter {
		t.Errorf("expected default unit KilogramSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassMomentOfInertiaGramSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassMomentOfInertiaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassMomentOfInertiaKilogramSquareMeter {
		t.Errorf("expected unit KilogramSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassMomentOfInertiaToString(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	a, err := factory.CreateMassMomentOfInertia(45, units.MassMomentOfInertiaKilogramSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassMomentOfInertiaKilogramSquareMeter, 2)
	expected := "45.00 " + units.GetMassMomentOfInertiaAbbreviation(units.MassMomentOfInertiaKilogramSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassMomentOfInertiaKilogramSquareMeter, -1)
	expected = "45 " + units.GetMassMomentOfInertiaAbbreviation(units.MassMomentOfInertiaKilogramSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMassMomentOfInertia_EqualityAndComparison(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	a1, _ := factory.CreateMassMomentOfInertia(60, units.MassMomentOfInertiaKilogramSquareMeter)
	a2, _ := factory.CreateMassMomentOfInertia(60, units.MassMomentOfInertiaKilogramSquareMeter)
	a3, _ := factory.CreateMassMomentOfInertia(120, units.MassMomentOfInertiaKilogramSquareMeter)

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

func TestMassMomentOfInertia_Arithmetic(t *testing.T) {
	factory := units.MassMomentOfInertiaFactory{}
	a1, _ := factory.CreateMassMomentOfInertia(30, units.MassMomentOfInertiaKilogramSquareMeter)
	a2, _ := factory.CreateMassMomentOfInertia(45, units.MassMomentOfInertiaKilogramSquareMeter)

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