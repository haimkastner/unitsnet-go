// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestInformationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Bit"}`
	
	factory := units.InformationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.InformationBit {
		t.Errorf("expected unit %v, got %v", units.InformationBit, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Bit"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestInformationDto_ToJSON(t *testing.T) {
	dto := units.InformationDto{
		Value: 45,
		Unit:  units.InformationBit,
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
	if result["unit"].(string) != string(units.InformationBit) {
		t.Errorf("expected unit %s, got %v", units.InformationBit, result["unit"])
	}
}

func TestNewInformation_InvalidValue(t *testing.T) {
	factory := units.InformationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateInformation(math.NaN(), units.InformationBit)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateInformation(math.Inf(1), units.InformationBit)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestInformationConversions(t *testing.T) {
	factory := units.InformationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateInformation(180, units.InformationBit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Bytes.
		// No expected conversion value provided for Bytes, verifying result is not NaN.
		result := a.Bytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Bytes returned NaN")
		}
	}
	{
		// Test conversion to Bits.
		// No expected conversion value provided for Bits, verifying result is not NaN.
		result := a.Bits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Bits returned NaN")
		}
	}
	{
		// Test conversion to Kilobytes.
		// No expected conversion value provided for Kilobytes, verifying result is not NaN.
		result := a.Kilobytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilobytes returned NaN")
		}
	}
	{
		// Test conversion to Megabytes.
		// No expected conversion value provided for Megabytes, verifying result is not NaN.
		result := a.Megabytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megabytes returned NaN")
		}
	}
	{
		// Test conversion to Gigabytes.
		// No expected conversion value provided for Gigabytes, verifying result is not NaN.
		result := a.Gigabytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigabytes returned NaN")
		}
	}
	{
		// Test conversion to Terabytes.
		// No expected conversion value provided for Terabytes, verifying result is not NaN.
		result := a.Terabytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terabytes returned NaN")
		}
	}
	{
		// Test conversion to Petabytes.
		// No expected conversion value provided for Petabytes, verifying result is not NaN.
		result := a.Petabytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Petabytes returned NaN")
		}
	}
	{
		// Test conversion to Exabytes.
		// No expected conversion value provided for Exabytes, verifying result is not NaN.
		result := a.Exabytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Exabytes returned NaN")
		}
	}
	{
		// Test conversion to Kibibytes.
		// No expected conversion value provided for Kibibytes, verifying result is not NaN.
		result := a.Kibibytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kibibytes returned NaN")
		}
	}
	{
		// Test conversion to Mebibytes.
		// No expected conversion value provided for Mebibytes, verifying result is not NaN.
		result := a.Mebibytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Mebibytes returned NaN")
		}
	}
	{
		// Test conversion to Gibibytes.
		// No expected conversion value provided for Gibibytes, verifying result is not NaN.
		result := a.Gibibytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gibibytes returned NaN")
		}
	}
	{
		// Test conversion to Tebibytes.
		// No expected conversion value provided for Tebibytes, verifying result is not NaN.
		result := a.Tebibytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Tebibytes returned NaN")
		}
	}
	{
		// Test conversion to Pebibytes.
		// No expected conversion value provided for Pebibytes, verifying result is not NaN.
		result := a.Pebibytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Pebibytes returned NaN")
		}
	}
	{
		// Test conversion to Exbibytes.
		// No expected conversion value provided for Exbibytes, verifying result is not NaN.
		result := a.Exbibytes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Exbibytes returned NaN")
		}
	}
	{
		// Test conversion to Kilobits.
		// No expected conversion value provided for Kilobits, verifying result is not NaN.
		result := a.Kilobits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilobits returned NaN")
		}
	}
	{
		// Test conversion to Megabits.
		// No expected conversion value provided for Megabits, verifying result is not NaN.
		result := a.Megabits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megabits returned NaN")
		}
	}
	{
		// Test conversion to Gigabits.
		// No expected conversion value provided for Gigabits, verifying result is not NaN.
		result := a.Gigabits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigabits returned NaN")
		}
	}
	{
		// Test conversion to Terabits.
		// No expected conversion value provided for Terabits, verifying result is not NaN.
		result := a.Terabits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terabits returned NaN")
		}
	}
	{
		// Test conversion to Petabits.
		// No expected conversion value provided for Petabits, verifying result is not NaN.
		result := a.Petabits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Petabits returned NaN")
		}
	}
	{
		// Test conversion to Exabits.
		// No expected conversion value provided for Exabits, verifying result is not NaN.
		result := a.Exabits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Exabits returned NaN")
		}
	}
	{
		// Test conversion to Kibibits.
		// No expected conversion value provided for Kibibits, verifying result is not NaN.
		result := a.Kibibits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kibibits returned NaN")
		}
	}
	{
		// Test conversion to Mebibits.
		// No expected conversion value provided for Mebibits, verifying result is not NaN.
		result := a.Mebibits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Mebibits returned NaN")
		}
	}
	{
		// Test conversion to Gibibits.
		// No expected conversion value provided for Gibibits, verifying result is not NaN.
		result := a.Gibibits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gibibits returned NaN")
		}
	}
	{
		// Test conversion to Tebibits.
		// No expected conversion value provided for Tebibits, verifying result is not NaN.
		result := a.Tebibits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Tebibits returned NaN")
		}
	}
	{
		// Test conversion to Pebibits.
		// No expected conversion value provided for Pebibits, verifying result is not NaN.
		result := a.Pebibits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Pebibits returned NaN")
		}
	}
	{
		// Test conversion to Exbibits.
		// No expected conversion value provided for Exbibits, verifying result is not NaN.
		result := a.Exbibits()
		if math.IsNaN(result) {
			t.Errorf("conversion to Exbibits returned NaN")
		}
	}
}

func TestInformation_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.InformationFactory{}
	a, err := factory.CreateInformation(90, units.InformationBit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.InformationBit {
		t.Errorf("expected default unit Bit, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.InformationByte
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.InformationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.InformationBit {
		t.Errorf("expected unit Bit, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestInformationToString(t *testing.T) {
	factory := units.InformationFactory{}
	a, err := factory.CreateInformation(45, units.InformationBit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.InformationBit, 2)
	expected := "45.00 " + units.GetInformationAbbreviation(units.InformationBit)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.InformationBit, -1)
	expected = "45 " + units.GetInformationAbbreviation(units.InformationBit)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestInformation_EqualityAndComparison(t *testing.T) {
	factory := units.InformationFactory{}
	a1, _ := factory.CreateInformation(60, units.InformationBit)
	a2, _ := factory.CreateInformation(60, units.InformationBit)
	a3, _ := factory.CreateInformation(120, units.InformationBit)

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

func TestInformation_Arithmetic(t *testing.T) {
	factory := units.InformationFactory{}
	a1, _ := factory.CreateInformation(30, units.InformationBit)
	a2, _ := factory.CreateInformation(45, units.InformationBit)

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