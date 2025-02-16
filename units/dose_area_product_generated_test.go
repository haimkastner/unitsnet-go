// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestDoseAreaProductDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "GraySquareMeter"}`
	
	factory := units.DoseAreaProductDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.DoseAreaProductGraySquareMeter {
		t.Errorf("expected unit %v, got %v", units.DoseAreaProductGraySquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "GraySquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestDoseAreaProductDto_ToJSON(t *testing.T) {
	dto := units.DoseAreaProductDto{
		Value: 45,
		Unit:  units.DoseAreaProductGraySquareMeter,
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
	if result["unit"].(string) != string(units.DoseAreaProductGraySquareMeter) {
		t.Errorf("expected unit %s, got %v", units.DoseAreaProductGraySquareMeter, result["unit"])
	}
}

func TestNewDoseAreaProduct_InvalidValue(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	// NaN value should return an error.
	_, err := factory.CreateDoseAreaProduct(math.NaN(), units.DoseAreaProductGraySquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateDoseAreaProduct(math.Inf(1), units.DoseAreaProductGraySquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestDoseAreaProductConversions(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateDoseAreaProduct(180, units.DoseAreaProductGraySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GraySquareMeters.
		// No expected conversion value provided for GraySquareMeters, verifying result is not NaN.
		result := a.GraySquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GraySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to GraySquareDecimeters.
		// No expected conversion value provided for GraySquareDecimeters, verifying result is not NaN.
		result := a.GraySquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GraySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to GraySquareCentimeters.
		// No expected conversion value provided for GraySquareCentimeters, verifying result is not NaN.
		result := a.GraySquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GraySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to GraySquareMillimeters.
		// No expected conversion value provided for GraySquareMillimeters, verifying result is not NaN.
		result := a.GraySquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GraySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareMeters.
		// No expected conversion value provided for MicrograySquareMeters, verifying result is not NaN.
		result := a.MicrograySquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrograySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareMeters.
		// No expected conversion value provided for MilligraySquareMeters, verifying result is not NaN.
		result := a.MilligraySquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligraySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareMeters.
		// No expected conversion value provided for CentigraySquareMeters, verifying result is not NaN.
		result := a.CentigraySquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigraySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareMeters.
		// No expected conversion value provided for DecigraySquareMeters, verifying result is not NaN.
		result := a.DecigraySquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigraySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareDecimeters.
		// No expected conversion value provided for MicrograySquareDecimeters, verifying result is not NaN.
		result := a.MicrograySquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrograySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareDecimeters.
		// No expected conversion value provided for MilligraySquareDecimeters, verifying result is not NaN.
		result := a.MilligraySquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligraySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareDecimeters.
		// No expected conversion value provided for CentigraySquareDecimeters, verifying result is not NaN.
		result := a.CentigraySquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigraySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareDecimeters.
		// No expected conversion value provided for DecigraySquareDecimeters, verifying result is not NaN.
		result := a.DecigraySquareDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigraySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareCentimeters.
		// No expected conversion value provided for MicrograySquareCentimeters, verifying result is not NaN.
		result := a.MicrograySquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrograySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareCentimeters.
		// No expected conversion value provided for MilligraySquareCentimeters, verifying result is not NaN.
		result := a.MilligraySquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligraySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareCentimeters.
		// No expected conversion value provided for CentigraySquareCentimeters, verifying result is not NaN.
		result := a.CentigraySquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigraySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareCentimeters.
		// No expected conversion value provided for DecigraySquareCentimeters, verifying result is not NaN.
		result := a.DecigraySquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigraySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareMillimeters.
		// No expected conversion value provided for MicrograySquareMillimeters, verifying result is not NaN.
		result := a.MicrograySquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrograySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareMillimeters.
		// No expected conversion value provided for MilligraySquareMillimeters, verifying result is not NaN.
		result := a.MilligraySquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligraySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareMillimeters.
		// No expected conversion value provided for CentigraySquareMillimeters, verifying result is not NaN.
		result := a.CentigraySquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigraySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareMillimeters.
		// No expected conversion value provided for DecigraySquareMillimeters, verifying result is not NaN.
		result := a.DecigraySquareMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigraySquareMillimeters returned NaN")
		}
	}
}

func TestDoseAreaProduct_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	a, err := factory.CreateDoseAreaProduct(90, units.DoseAreaProductGraySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.DoseAreaProductGraySquareMeter {
		t.Errorf("expected default unit GraySquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.DoseAreaProductGraySquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.DoseAreaProductDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.DoseAreaProductGraySquareMeter {
		t.Errorf("expected unit GraySquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestDoseAreaProductToString(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	a, err := factory.CreateDoseAreaProduct(45, units.DoseAreaProductGraySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.DoseAreaProductGraySquareMeter, 2)
	expected := "45.00 " + units.GetDoseAreaProductAbbreviation(units.DoseAreaProductGraySquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.DoseAreaProductGraySquareMeter, -1)
	expected = "45 " + units.GetDoseAreaProductAbbreviation(units.DoseAreaProductGraySquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestDoseAreaProduct_EqualityAndComparison(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	a1, _ := factory.CreateDoseAreaProduct(60, units.DoseAreaProductGraySquareMeter)
	a2, _ := factory.CreateDoseAreaProduct(60, units.DoseAreaProductGraySquareMeter)
	a3, _ := factory.CreateDoseAreaProduct(120, units.DoseAreaProductGraySquareMeter)

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

func TestDoseAreaProduct_Arithmetic(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	a1, _ := factory.CreateDoseAreaProduct(30, units.DoseAreaProductGraySquareMeter)
	a2, _ := factory.CreateDoseAreaProduct(45, units.DoseAreaProductGraySquareMeter)

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