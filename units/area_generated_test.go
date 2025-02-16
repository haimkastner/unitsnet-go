// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAreaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeter"}`
	
	factory := units.AreaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AreaSquareMeter {
		t.Errorf("expected unit %v, got %v", units.AreaSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAreaDto_ToJSON(t *testing.T) {
	dto := units.AreaDto{
		Value: 45,
		Unit:  units.AreaSquareMeter,
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
	if result["unit"].(string) != string(units.AreaSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.AreaSquareMeter, result["unit"])
	}
}

func TestNewArea_InvalidValue(t *testing.T) {
	factory := units.AreaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateArea(math.NaN(), units.AreaSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateArea(math.Inf(1), units.AreaSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAreaConversions(t *testing.T) {
	factory := units.AreaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateArea(180, units.AreaSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SquareKilometers.
		// No expected conversion value provided for SquareKilometers, verifying result is not NaN.
		result := a.SquareKilometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareKilometers returned NaN")
		}
	}
	{
		// Test conversion to SquareMeters.
		// No expected conversion value provided for SquareMeters, verifying result is not NaN.
		result := a.SquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeters returned NaN")
		}
	}
	{
		// Test conversion to SquareDecimeters.
		// No expected conversion value provided for SquareDecimeters, verifying result is not NaN.
		result := a.SquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeters.
		// No expected conversion value provided for SquareCentimeters, verifying result is not NaN.
		result := a.SquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to SquareMillimeters.
		// No expected conversion value provided for SquareMillimeters, verifying result is not NaN.
		result := a.SquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to SquareMicrometers.
		// No expected conversion value provided for SquareMicrometers, verifying result is not NaN.
		result := a.SquareMicrometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMicrometers returned NaN")
		}
	}
	{
		// Test conversion to SquareMiles.
		// No expected conversion value provided for SquareMiles, verifying result is not NaN.
		result := a.SquareMiles()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMiles returned NaN")
		}
	}
	{
		// Test conversion to SquareYards.
		// No expected conversion value provided for SquareYards, verifying result is not NaN.
		result := a.SquareYards()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareYards returned NaN")
		}
	}
	{
		// Test conversion to SquareFeet.
		// No expected conversion value provided for SquareFeet, verifying result is not NaN.
		result := a.SquareFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareFeet returned NaN")
		}
	}
	{
		// Test conversion to UsSurveySquareFeet.
		// No expected conversion value provided for UsSurveySquareFeet, verifying result is not NaN.
		result := a.UsSurveySquareFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsSurveySquareFeet returned NaN")
		}
	}
	{
		// Test conversion to SquareInches.
		// No expected conversion value provided for SquareInches, verifying result is not NaN.
		result := a.SquareInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareInches returned NaN")
		}
	}
	{
		// Test conversion to Acres.
		// No expected conversion value provided for Acres, verifying result is not NaN.
		result := a.Acres()
		if math.IsNaN(result) {
			t.Errorf("conversion to Acres returned NaN")
		}
	}
	{
		// Test conversion to Hectares.
		// No expected conversion value provided for Hectares, verifying result is not NaN.
		result := a.Hectares()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hectares returned NaN")
		}
	}
	{
		// Test conversion to SquareNauticalMiles.
		// No expected conversion value provided for SquareNauticalMiles, verifying result is not NaN.
		result := a.SquareNauticalMiles()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareNauticalMiles returned NaN")
		}
	}
}

func TestArea_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AreaFactory{}
	a, err := factory.CreateArea(90, units.AreaSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AreaSquareMeter {
		t.Errorf("expected default unit SquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AreaSquareKilometer
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AreaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AreaSquareMeter {
		t.Errorf("expected unit SquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAreaToString(t *testing.T) {
	factory := units.AreaFactory{}
	a, err := factory.CreateArea(45, units.AreaSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AreaSquareMeter, 2)
	expected := "45.00 " + units.GetAreaAbbreviation(units.AreaSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AreaSquareMeter, -1)
	expected = "45 " + units.GetAreaAbbreviation(units.AreaSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestArea_EqualityAndComparison(t *testing.T) {
	factory := units.AreaFactory{}
	a1, _ := factory.CreateArea(60, units.AreaSquareMeter)
	a2, _ := factory.CreateArea(60, units.AreaSquareMeter)
	a3, _ := factory.CreateArea(120, units.AreaSquareMeter)

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

func TestArea_Arithmetic(t *testing.T) {
	factory := units.AreaFactory{}
	a1, _ := factory.CreateArea(30, units.AreaSquareMeter)
	a2, _ := factory.CreateArea(45, units.AreaSquareMeter)

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