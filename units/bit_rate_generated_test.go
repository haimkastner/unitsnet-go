// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestBitRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "BitPerSecond"}`
	
	factory := units.BitRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.BitRateBitPerSecond {
		t.Errorf("expected unit %v, got %v", units.BitRateBitPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "BitPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestBitRateDto_ToJSON(t *testing.T) {
	dto := units.BitRateDto{
		Value: 45,
		Unit:  units.BitRateBitPerSecond,
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
	if result["unit"].(string) != string(units.BitRateBitPerSecond) {
		t.Errorf("expected unit %s, got %v", units.BitRateBitPerSecond, result["unit"])
	}
}

func TestNewBitRate_InvalidValue(t *testing.T) {
	factory := units.BitRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateBitRate(math.NaN(), units.BitRateBitPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateBitRate(math.Inf(1), units.BitRateBitPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestBitRateConversions(t *testing.T) {
	factory := units.BitRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateBitRate(180, units.BitRateBitPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to BitsPerSecond.
		// No expected conversion value provided for BitsPerSecond, verifying result is not NaN.
		result := a.BitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to BitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to BytesPerSecond.
		// No expected conversion value provided for BytesPerSecond, verifying result is not NaN.
		result := a.BytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to BytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilobitsPerSecond.
		// No expected conversion value provided for KilobitsPerSecond, verifying result is not NaN.
		result := a.KilobitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilobitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegabitsPerSecond.
		// No expected conversion value provided for MegabitsPerSecond, verifying result is not NaN.
		result := a.MegabitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GigabitsPerSecond.
		// No expected conversion value provided for GigabitsPerSecond, verifying result is not NaN.
		result := a.GigabitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TerabitsPerSecond.
		// No expected conversion value provided for TerabitsPerSecond, verifying result is not NaN.
		result := a.TerabitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PetabitsPerSecond.
		// No expected conversion value provided for PetabitsPerSecond, verifying result is not NaN.
		result := a.PetabitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PetabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExabitsPerSecond.
		// No expected conversion value provided for ExabitsPerSecond, verifying result is not NaN.
		result := a.ExabitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to ExabitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KibibitsPerSecond.
		// No expected conversion value provided for KibibitsPerSecond, verifying result is not NaN.
		result := a.KibibitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KibibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MebibitsPerSecond.
		// No expected conversion value provided for MebibitsPerSecond, verifying result is not NaN.
		result := a.MebibitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MebibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GibibitsPerSecond.
		// No expected conversion value provided for GibibitsPerSecond, verifying result is not NaN.
		result := a.GibibitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to GibibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TebibitsPerSecond.
		// No expected conversion value provided for TebibitsPerSecond, verifying result is not NaN.
		result := a.TebibitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to TebibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PebibitsPerSecond.
		// No expected conversion value provided for PebibitsPerSecond, verifying result is not NaN.
		result := a.PebibitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PebibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExbibitsPerSecond.
		// No expected conversion value provided for ExbibitsPerSecond, verifying result is not NaN.
		result := a.ExbibitsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to ExbibitsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilobytesPerSecond.
		// No expected conversion value provided for KilobytesPerSecond, verifying result is not NaN.
		result := a.KilobytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilobytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegabytesPerSecond.
		// No expected conversion value provided for MegabytesPerSecond, verifying result is not NaN.
		result := a.MegabytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GigabytesPerSecond.
		// No expected conversion value provided for GigabytesPerSecond, verifying result is not NaN.
		result := a.GigabytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TerabytesPerSecond.
		// No expected conversion value provided for TerabytesPerSecond, verifying result is not NaN.
		result := a.TerabytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PetabytesPerSecond.
		// No expected conversion value provided for PetabytesPerSecond, verifying result is not NaN.
		result := a.PetabytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PetabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExabytesPerSecond.
		// No expected conversion value provided for ExabytesPerSecond, verifying result is not NaN.
		result := a.ExabytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to ExabytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KibibytesPerSecond.
		// No expected conversion value provided for KibibytesPerSecond, verifying result is not NaN.
		result := a.KibibytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KibibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MebibytesPerSecond.
		// No expected conversion value provided for MebibytesPerSecond, verifying result is not NaN.
		result := a.MebibytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MebibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to GibibytesPerSecond.
		// No expected conversion value provided for GibibytesPerSecond, verifying result is not NaN.
		result := a.GibibytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to GibibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TebibytesPerSecond.
		// No expected conversion value provided for TebibytesPerSecond, verifying result is not NaN.
		result := a.TebibytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to TebibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PebibytesPerSecond.
		// No expected conversion value provided for PebibytesPerSecond, verifying result is not NaN.
		result := a.PebibytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PebibytesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to ExbibytesPerSecond.
		// No expected conversion value provided for ExbibytesPerSecond, verifying result is not NaN.
		result := a.ExbibytesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to ExbibytesPerSecond returned NaN")
		}
	}
}

func TestBitRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.BitRateFactory{}
	a, err := factory.CreateBitRate(90, units.BitRateBitPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.BitRateBitPerSecond {
		t.Errorf("expected default unit BitPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.BitRateBitPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.BitRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.BitRateBitPerSecond {
		t.Errorf("expected unit BitPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestBitRateToString(t *testing.T) {
	factory := units.BitRateFactory{}
	a, err := factory.CreateBitRate(45, units.BitRateBitPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.BitRateBitPerSecond, 2)
	expected := "45.00 " + units.GetBitRateAbbreviation(units.BitRateBitPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.BitRateBitPerSecond, -1)
	expected = "45 " + units.GetBitRateAbbreviation(units.BitRateBitPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestBitRate_EqualityAndComparison(t *testing.T) {
	factory := units.BitRateFactory{}
	a1, _ := factory.CreateBitRate(60, units.BitRateBitPerSecond)
	a2, _ := factory.CreateBitRate(60, units.BitRateBitPerSecond)
	a3, _ := factory.CreateBitRate(120, units.BitRateBitPerSecond)

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

func TestBitRate_Arithmetic(t *testing.T) {
	factory := units.BitRateFactory{}
	a1, _ := factory.CreateBitRate(30, units.BitRateBitPerSecond)
	a2, _ := factory.CreateBitRate(45, units.BitRateBitPerSecond)

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