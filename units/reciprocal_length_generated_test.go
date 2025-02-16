// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestReciprocalLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "InverseMeter"}`
	
	factory := units.ReciprocalLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ReciprocalLengthInverseMeter {
		t.Errorf("expected unit %v, got %v", units.ReciprocalLengthInverseMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "InverseMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestReciprocalLengthDto_ToJSON(t *testing.T) {
	dto := units.ReciprocalLengthDto{
		Value: 45,
		Unit:  units.ReciprocalLengthInverseMeter,
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
	if result["unit"].(string) != string(units.ReciprocalLengthInverseMeter) {
		t.Errorf("expected unit %s, got %v", units.ReciprocalLengthInverseMeter, result["unit"])
	}
}

func TestNewReciprocalLength_InvalidValue(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateReciprocalLength(math.NaN(), units.ReciprocalLengthInverseMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateReciprocalLength(math.Inf(1), units.ReciprocalLengthInverseMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestReciprocalLengthConversions(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateReciprocalLength(180, units.ReciprocalLengthInverseMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to InverseMeters.
		// No expected conversion value provided for InverseMeters, verifying result is not NaN.
		result := a.InverseMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMeters returned NaN")
		}
	}
	{
		// Test conversion to InverseCentimeters.
		// No expected conversion value provided for InverseCentimeters, verifying result is not NaN.
		result := a.InverseCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseCentimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseMillimeters.
		// No expected conversion value provided for InverseMillimeters, verifying result is not NaN.
		result := a.InverseMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMillimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseMiles.
		// No expected conversion value provided for InverseMiles, verifying result is not NaN.
		result := a.InverseMiles()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMiles returned NaN")
		}
	}
	{
		// Test conversion to InverseYards.
		// No expected conversion value provided for InverseYards, verifying result is not NaN.
		result := a.InverseYards()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseYards returned NaN")
		}
	}
	{
		// Test conversion to InverseFeet.
		// No expected conversion value provided for InverseFeet, verifying result is not NaN.
		result := a.InverseFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseFeet returned NaN")
		}
	}
	{
		// Test conversion to InverseUsSurveyFeet.
		// No expected conversion value provided for InverseUsSurveyFeet, verifying result is not NaN.
		result := a.InverseUsSurveyFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseUsSurveyFeet returned NaN")
		}
	}
	{
		// Test conversion to InverseInches.
		// No expected conversion value provided for InverseInches, verifying result is not NaN.
		result := a.InverseInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseInches returned NaN")
		}
	}
	{
		// Test conversion to InverseMils.
		// No expected conversion value provided for InverseMils, verifying result is not NaN.
		result := a.InverseMils()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMils returned NaN")
		}
	}
	{
		// Test conversion to InverseMicroinches.
		// No expected conversion value provided for InverseMicroinches, verifying result is not NaN.
		result := a.InverseMicroinches()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMicroinches returned NaN")
		}
	}
}

func TestReciprocalLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	a, err := factory.CreateReciprocalLength(90, units.ReciprocalLengthInverseMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ReciprocalLengthInverseMeter {
		t.Errorf("expected default unit InverseMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ReciprocalLengthInverseMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ReciprocalLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ReciprocalLengthInverseMeter {
		t.Errorf("expected unit InverseMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestReciprocalLengthToString(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	a, err := factory.CreateReciprocalLength(45, units.ReciprocalLengthInverseMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ReciprocalLengthInverseMeter, 2)
	expected := "45.00 " + units.GetReciprocalLengthAbbreviation(units.ReciprocalLengthInverseMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ReciprocalLengthInverseMeter, -1)
	expected = "45 " + units.GetReciprocalLengthAbbreviation(units.ReciprocalLengthInverseMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestReciprocalLength_EqualityAndComparison(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	a1, _ := factory.CreateReciprocalLength(60, units.ReciprocalLengthInverseMeter)
	a2, _ := factory.CreateReciprocalLength(60, units.ReciprocalLengthInverseMeter)
	a3, _ := factory.CreateReciprocalLength(120, units.ReciprocalLengthInverseMeter)

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

func TestReciprocalLength_Arithmetic(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	a1, _ := factory.CreateReciprocalLength(30, units.ReciprocalLengthInverseMeter)
	a2, _ := factory.CreateReciprocalLength(45, units.ReciprocalLengthInverseMeter)

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