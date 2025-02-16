// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricResistivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "OhmMeter"}`
	
	factory := units.ElectricResistivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricResistivityOhmMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricResistivityOhmMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "OhmMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricResistivityDto_ToJSON(t *testing.T) {
	dto := units.ElectricResistivityDto{
		Value: 45,
		Unit:  units.ElectricResistivityOhmMeter,
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
	if result["unit"].(string) != string(units.ElectricResistivityOhmMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricResistivityOhmMeter, result["unit"])
	}
}

func TestNewElectricResistivity_InvalidValue(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricResistivity(math.NaN(), units.ElectricResistivityOhmMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricResistivity(math.Inf(1), units.ElectricResistivityOhmMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricResistivityConversions(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricResistivity(180, units.ElectricResistivityOhmMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to OhmMeters.
		// No expected conversion value provided for OhmMeters, verifying result is not NaN.
		result := a.OhmMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to OhmMeters returned NaN")
		}
	}
	{
		// Test conversion to OhmsCentimeter.
		// No expected conversion value provided for OhmsCentimeter, verifying result is not NaN.
		result := a.OhmsCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to OhmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PicoohmMeters.
		// No expected conversion value provided for PicoohmMeters, verifying result is not NaN.
		result := a.PicoohmMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicoohmMeters returned NaN")
		}
	}
	{
		// Test conversion to NanoohmMeters.
		// No expected conversion value provided for NanoohmMeters, verifying result is not NaN.
		result := a.NanoohmMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanoohmMeters returned NaN")
		}
	}
	{
		// Test conversion to MicroohmMeters.
		// No expected conversion value provided for MicroohmMeters, verifying result is not NaN.
		result := a.MicroohmMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicroohmMeters returned NaN")
		}
	}
	{
		// Test conversion to MilliohmMeters.
		// No expected conversion value provided for MilliohmMeters, verifying result is not NaN.
		result := a.MilliohmMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliohmMeters returned NaN")
		}
	}
	{
		// Test conversion to KiloohmMeters.
		// No expected conversion value provided for KiloohmMeters, verifying result is not NaN.
		result := a.KiloohmMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KiloohmMeters returned NaN")
		}
	}
	{
		// Test conversion to MegaohmMeters.
		// No expected conversion value provided for MegaohmMeters, verifying result is not NaN.
		result := a.MegaohmMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegaohmMeters returned NaN")
		}
	}
	{
		// Test conversion to PicoohmsCentimeter.
		// No expected conversion value provided for PicoohmsCentimeter, verifying result is not NaN.
		result := a.PicoohmsCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicoohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NanoohmsCentimeter.
		// No expected conversion value provided for NanoohmsCentimeter, verifying result is not NaN.
		result := a.NanoohmsCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanoohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicroohmsCentimeter.
		// No expected conversion value provided for MicroohmsCentimeter, verifying result is not NaN.
		result := a.MicroohmsCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicroohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MilliohmsCentimeter.
		// No expected conversion value provided for MilliohmsCentimeter, verifying result is not NaN.
		result := a.MilliohmsCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KiloohmsCentimeter.
		// No expected conversion value provided for KiloohmsCentimeter, verifying result is not NaN.
		result := a.KiloohmsCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KiloohmsCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MegaohmsCentimeter.
		// No expected conversion value provided for MegaohmsCentimeter, verifying result is not NaN.
		result := a.MegaohmsCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegaohmsCentimeter returned NaN")
		}
	}
}

func TestElectricResistivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	a, err := factory.CreateElectricResistivity(90, units.ElectricResistivityOhmMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricResistivityOhmMeter {
		t.Errorf("expected default unit OhmMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricResistivityOhmMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricResistivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricResistivityOhmMeter {
		t.Errorf("expected unit OhmMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricResistivityToString(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	a, err := factory.CreateElectricResistivity(45, units.ElectricResistivityOhmMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricResistivityOhmMeter, 2)
	expected := "45.00 " + units.GetElectricResistivityAbbreviation(units.ElectricResistivityOhmMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricResistivityOhmMeter, -1)
	expected = "45 " + units.GetElectricResistivityAbbreviation(units.ElectricResistivityOhmMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricResistivity_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	a1, _ := factory.CreateElectricResistivity(60, units.ElectricResistivityOhmMeter)
	a2, _ := factory.CreateElectricResistivity(60, units.ElectricResistivityOhmMeter)
	a3, _ := factory.CreateElectricResistivity(120, units.ElectricResistivityOhmMeter)

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

func TestElectricResistivity_Arithmetic(t *testing.T) {
	factory := units.ElectricResistivityFactory{}
	a1, _ := factory.CreateElectricResistivity(30, units.ElectricResistivityOhmMeter)
	a2, _ := factory.CreateElectricResistivity(45, units.ElectricResistivityOhmMeter)

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