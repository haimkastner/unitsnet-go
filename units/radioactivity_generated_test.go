// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRadioactivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Becquerel"}`
	
	factory := units.RadioactivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RadioactivityBecquerel {
		t.Errorf("expected unit %v, got %v", units.RadioactivityBecquerel, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Becquerel"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRadioactivityDto_ToJSON(t *testing.T) {
	dto := units.RadioactivityDto{
		Value: 45,
		Unit:  units.RadioactivityBecquerel,
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
	if result["unit"].(string) != string(units.RadioactivityBecquerel) {
		t.Errorf("expected unit %s, got %v", units.RadioactivityBecquerel, result["unit"])
	}
}

func TestNewRadioactivity_InvalidValue(t *testing.T) {
	factory := units.RadioactivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRadioactivity(math.NaN(), units.RadioactivityBecquerel)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRadioactivity(math.Inf(1), units.RadioactivityBecquerel)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRadioactivityConversions(t *testing.T) {
	factory := units.RadioactivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRadioactivity(180, units.RadioactivityBecquerel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Becquerels.
		// No expected conversion value provided for Becquerels, verifying result is not NaN.
		result := a.Becquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Becquerels returned NaN")
		}
	}
	{
		// Test conversion to Curies.
		// No expected conversion value provided for Curies, verifying result is not NaN.
		result := a.Curies()
		if math.IsNaN(result) {
			t.Errorf("conversion to Curies returned NaN")
		}
	}
	{
		// Test conversion to Rutherfords.
		// No expected conversion value provided for Rutherfords, verifying result is not NaN.
		result := a.Rutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Rutherfords returned NaN")
		}
	}
	{
		// Test conversion to Picobecquerels.
		// No expected conversion value provided for Picobecquerels, verifying result is not NaN.
		result := a.Picobecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picobecquerels returned NaN")
		}
	}
	{
		// Test conversion to Nanobecquerels.
		// No expected conversion value provided for Nanobecquerels, verifying result is not NaN.
		result := a.Nanobecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanobecquerels returned NaN")
		}
	}
	{
		// Test conversion to Microbecquerels.
		// No expected conversion value provided for Microbecquerels, verifying result is not NaN.
		result := a.Microbecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microbecquerels returned NaN")
		}
	}
	{
		// Test conversion to Millibecquerels.
		// No expected conversion value provided for Millibecquerels, verifying result is not NaN.
		result := a.Millibecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millibecquerels returned NaN")
		}
	}
	{
		// Test conversion to Kilobecquerels.
		// No expected conversion value provided for Kilobecquerels, verifying result is not NaN.
		result := a.Kilobecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilobecquerels returned NaN")
		}
	}
	{
		// Test conversion to Megabecquerels.
		// No expected conversion value provided for Megabecquerels, verifying result is not NaN.
		result := a.Megabecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Gigabecquerels.
		// No expected conversion value provided for Gigabecquerels, verifying result is not NaN.
		result := a.Gigabecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Terabecquerels.
		// No expected conversion value provided for Terabecquerels, verifying result is not NaN.
		result := a.Terabecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Petabecquerels.
		// No expected conversion value provided for Petabecquerels, verifying result is not NaN.
		result := a.Petabecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Petabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Exabecquerels.
		// No expected conversion value provided for Exabecquerels, verifying result is not NaN.
		result := a.Exabecquerels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Exabecquerels returned NaN")
		}
	}
	{
		// Test conversion to Picocuries.
		// No expected conversion value provided for Picocuries, verifying result is not NaN.
		result := a.Picocuries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picocuries returned NaN")
		}
	}
	{
		// Test conversion to Nanocuries.
		// No expected conversion value provided for Nanocuries, verifying result is not NaN.
		result := a.Nanocuries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanocuries returned NaN")
		}
	}
	{
		// Test conversion to Microcuries.
		// No expected conversion value provided for Microcuries, verifying result is not NaN.
		result := a.Microcuries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microcuries returned NaN")
		}
	}
	{
		// Test conversion to Millicuries.
		// No expected conversion value provided for Millicuries, verifying result is not NaN.
		result := a.Millicuries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millicuries returned NaN")
		}
	}
	{
		// Test conversion to Kilocuries.
		// No expected conversion value provided for Kilocuries, verifying result is not NaN.
		result := a.Kilocuries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilocuries returned NaN")
		}
	}
	{
		// Test conversion to Megacuries.
		// No expected conversion value provided for Megacuries, verifying result is not NaN.
		result := a.Megacuries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megacuries returned NaN")
		}
	}
	{
		// Test conversion to Gigacuries.
		// No expected conversion value provided for Gigacuries, verifying result is not NaN.
		result := a.Gigacuries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigacuries returned NaN")
		}
	}
	{
		// Test conversion to Teracuries.
		// No expected conversion value provided for Teracuries, verifying result is not NaN.
		result := a.Teracuries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Teracuries returned NaN")
		}
	}
	{
		// Test conversion to Picorutherfords.
		// No expected conversion value provided for Picorutherfords, verifying result is not NaN.
		result := a.Picorutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picorutherfords returned NaN")
		}
	}
	{
		// Test conversion to Nanorutherfords.
		// No expected conversion value provided for Nanorutherfords, verifying result is not NaN.
		result := a.Nanorutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanorutherfords returned NaN")
		}
	}
	{
		// Test conversion to Microrutherfords.
		// No expected conversion value provided for Microrutherfords, verifying result is not NaN.
		result := a.Microrutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microrutherfords returned NaN")
		}
	}
	{
		// Test conversion to Millirutherfords.
		// No expected conversion value provided for Millirutherfords, verifying result is not NaN.
		result := a.Millirutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millirutherfords returned NaN")
		}
	}
	{
		// Test conversion to Kilorutherfords.
		// No expected conversion value provided for Kilorutherfords, verifying result is not NaN.
		result := a.Kilorutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilorutherfords returned NaN")
		}
	}
	{
		// Test conversion to Megarutherfords.
		// No expected conversion value provided for Megarutherfords, verifying result is not NaN.
		result := a.Megarutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megarutherfords returned NaN")
		}
	}
	{
		// Test conversion to Gigarutherfords.
		// No expected conversion value provided for Gigarutherfords, verifying result is not NaN.
		result := a.Gigarutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigarutherfords returned NaN")
		}
	}
	{
		// Test conversion to Terarutherfords.
		// No expected conversion value provided for Terarutherfords, verifying result is not NaN.
		result := a.Terarutherfords()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terarutherfords returned NaN")
		}
	}
}

func TestRadioactivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RadioactivityFactory{}
	a, err := factory.CreateRadioactivity(90, units.RadioactivityBecquerel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RadioactivityBecquerel {
		t.Errorf("expected default unit Becquerel, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RadioactivityBecquerel
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RadioactivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RadioactivityBecquerel {
		t.Errorf("expected unit Becquerel, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRadioactivityToString(t *testing.T) {
	factory := units.RadioactivityFactory{}
	a, err := factory.CreateRadioactivity(45, units.RadioactivityBecquerel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RadioactivityBecquerel, 2)
	expected := "45.00 " + units.GetRadioactivityAbbreviation(units.RadioactivityBecquerel)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RadioactivityBecquerel, -1)
	expected = "45 " + units.GetRadioactivityAbbreviation(units.RadioactivityBecquerel)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRadioactivity_EqualityAndComparison(t *testing.T) {
	factory := units.RadioactivityFactory{}
	a1, _ := factory.CreateRadioactivity(60, units.RadioactivityBecquerel)
	a2, _ := factory.CreateRadioactivity(60, units.RadioactivityBecquerel)
	a3, _ := factory.CreateRadioactivity(120, units.RadioactivityBecquerel)

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

func TestRadioactivity_Arithmetic(t *testing.T) {
	factory := units.RadioactivityFactory{}
	a1, _ := factory.CreateRadioactivity(30, units.RadioactivityBecquerel)
	a2, _ := factory.CreateRadioactivity(45, units.RadioactivityBecquerel)

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