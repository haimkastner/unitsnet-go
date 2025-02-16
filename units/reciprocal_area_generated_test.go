// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestReciprocalAreaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "InverseSquareMeter"}`
	
	factory := units.ReciprocalAreaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ReciprocalAreaInverseSquareMeter {
		t.Errorf("expected unit %v, got %v", units.ReciprocalAreaInverseSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "InverseSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestReciprocalAreaDto_ToJSON(t *testing.T) {
	dto := units.ReciprocalAreaDto{
		Value: 45,
		Unit:  units.ReciprocalAreaInverseSquareMeter,
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
	if result["unit"].(string) != string(units.ReciprocalAreaInverseSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.ReciprocalAreaInverseSquareMeter, result["unit"])
	}
}

func TestNewReciprocalArea_InvalidValue(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateReciprocalArea(math.NaN(), units.ReciprocalAreaInverseSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateReciprocalArea(math.Inf(1), units.ReciprocalAreaInverseSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestReciprocalAreaConversions(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateReciprocalArea(180, units.ReciprocalAreaInverseSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to InverseSquareMeters.
		// No expected conversion value provided for InverseSquareMeters, verifying result is not NaN.
		result := a.InverseSquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareKilometers.
		// No expected conversion value provided for InverseSquareKilometers, verifying result is not NaN.
		result := a.InverseSquareKilometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareKilometers returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareDecimeters.
		// No expected conversion value provided for InverseSquareDecimeters, verifying result is not NaN.
		result := a.InverseSquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareCentimeters.
		// No expected conversion value provided for InverseSquareCentimeters, verifying result is not NaN.
		result := a.InverseSquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareMillimeters.
		// No expected conversion value provided for InverseSquareMillimeters, verifying result is not NaN.
		result := a.InverseSquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareMicrometers.
		// No expected conversion value provided for InverseSquareMicrometers, verifying result is not NaN.
		result := a.InverseSquareMicrometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareMicrometers returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareMiles.
		// No expected conversion value provided for InverseSquareMiles, verifying result is not NaN.
		result := a.InverseSquareMiles()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareMiles returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareYards.
		// No expected conversion value provided for InverseSquareYards, verifying result is not NaN.
		result := a.InverseSquareYards()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareYards returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareFeet.
		// No expected conversion value provided for InverseSquareFeet, verifying result is not NaN.
		result := a.InverseSquareFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareFeet returned NaN")
		}
	}
	{
		// Test conversion to InverseUsSurveySquareFeet.
		// No expected conversion value provided for InverseUsSurveySquareFeet, verifying result is not NaN.
		result := a.InverseUsSurveySquareFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseUsSurveySquareFeet returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareInches.
		// No expected conversion value provided for InverseSquareInches, verifying result is not NaN.
		result := a.InverseSquareInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseSquareInches returned NaN")
		}
	}
}

func TestReciprocalArea_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	a, err := factory.CreateReciprocalArea(90, units.ReciprocalAreaInverseSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ReciprocalAreaInverseSquareMeter {
		t.Errorf("expected default unit InverseSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ReciprocalAreaInverseSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ReciprocalAreaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ReciprocalAreaInverseSquareMeter {
		t.Errorf("expected unit InverseSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestReciprocalAreaToString(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	a, err := factory.CreateReciprocalArea(45, units.ReciprocalAreaInverseSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ReciprocalAreaInverseSquareMeter, 2)
	expected := "45.00 " + units.GetReciprocalAreaAbbreviation(units.ReciprocalAreaInverseSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ReciprocalAreaInverseSquareMeter, -1)
	expected = "45 " + units.GetReciprocalAreaAbbreviation(units.ReciprocalAreaInverseSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestReciprocalArea_EqualityAndComparison(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	a1, _ := factory.CreateReciprocalArea(60, units.ReciprocalAreaInverseSquareMeter)
	a2, _ := factory.CreateReciprocalArea(60, units.ReciprocalAreaInverseSquareMeter)
	a3, _ := factory.CreateReciprocalArea(120, units.ReciprocalAreaInverseSquareMeter)

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

func TestReciprocalArea_Arithmetic(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	a1, _ := factory.CreateReciprocalArea(30, units.ReciprocalAreaInverseSquareMeter)
	a2, _ := factory.CreateReciprocalArea(45, units.ReciprocalAreaInverseSquareMeter)

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