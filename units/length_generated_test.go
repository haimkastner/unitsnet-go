// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Meter"}`
	
	factory := units.LengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LengthMeter {
		t.Errorf("expected unit %v, got %v", units.LengthMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Meter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLengthDto_ToJSON(t *testing.T) {
	dto := units.LengthDto{
		Value: 45,
		Unit:  units.LengthMeter,
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
	if result["unit"].(string) != string(units.LengthMeter) {
		t.Errorf("expected unit %s, got %v", units.LengthMeter, result["unit"])
	}
}

func TestNewLength_InvalidValue(t *testing.T) {
	factory := units.LengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLength(math.NaN(), units.LengthMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLength(math.Inf(1), units.LengthMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLengthConversions(t *testing.T) {
	factory := units.LengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLength(180, units.LengthMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Meters.
		// No expected conversion value provided for Meters, verifying result is not NaN.
		result := a.Meters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Meters returned NaN")
		}
	}
	{
		// Test conversion to Miles.
		// No expected conversion value provided for Miles, verifying result is not NaN.
		result := a.Miles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Miles returned NaN")
		}
	}
	{
		// Test conversion to Yards.
		// No expected conversion value provided for Yards, verifying result is not NaN.
		result := a.Yards()
		if math.IsNaN(result) {
			t.Errorf("conversion to Yards returned NaN")
		}
	}
	{
		// Test conversion to Feet.
		// No expected conversion value provided for Feet, verifying result is not NaN.
		result := a.Feet()
		if math.IsNaN(result) {
			t.Errorf("conversion to Feet returned NaN")
		}
	}
	{
		// Test conversion to UsSurveyFeet.
		// No expected conversion value provided for UsSurveyFeet, verifying result is not NaN.
		result := a.UsSurveyFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsSurveyFeet returned NaN")
		}
	}
	{
		// Test conversion to Inches.
		// No expected conversion value provided for Inches, verifying result is not NaN.
		result := a.Inches()
		if math.IsNaN(result) {
			t.Errorf("conversion to Inches returned NaN")
		}
	}
	{
		// Test conversion to Mils.
		// No expected conversion value provided for Mils, verifying result is not NaN.
		result := a.Mils()
		if math.IsNaN(result) {
			t.Errorf("conversion to Mils returned NaN")
		}
	}
	{
		// Test conversion to NauticalMiles.
		// No expected conversion value provided for NauticalMiles, verifying result is not NaN.
		result := a.NauticalMiles()
		if math.IsNaN(result) {
			t.Errorf("conversion to NauticalMiles returned NaN")
		}
	}
	{
		// Test conversion to Fathoms.
		// No expected conversion value provided for Fathoms, verifying result is not NaN.
		result := a.Fathoms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Fathoms returned NaN")
		}
	}
	{
		// Test conversion to Shackles.
		// No expected conversion value provided for Shackles, verifying result is not NaN.
		result := a.Shackles()
		if math.IsNaN(result) {
			t.Errorf("conversion to Shackles returned NaN")
		}
	}
	{
		// Test conversion to Microinches.
		// No expected conversion value provided for Microinches, verifying result is not NaN.
		result := a.Microinches()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microinches returned NaN")
		}
	}
	{
		// Test conversion to PrinterPoints.
		// No expected conversion value provided for PrinterPoints, verifying result is not NaN.
		result := a.PrinterPoints()
		if math.IsNaN(result) {
			t.Errorf("conversion to PrinterPoints returned NaN")
		}
	}
	{
		// Test conversion to DtpPoints.
		// No expected conversion value provided for DtpPoints, verifying result is not NaN.
		result := a.DtpPoints()
		if math.IsNaN(result) {
			t.Errorf("conversion to DtpPoints returned NaN")
		}
	}
	{
		// Test conversion to PrinterPicas.
		// No expected conversion value provided for PrinterPicas, verifying result is not NaN.
		result := a.PrinterPicas()
		if math.IsNaN(result) {
			t.Errorf("conversion to PrinterPicas returned NaN")
		}
	}
	{
		// Test conversion to DtpPicas.
		// No expected conversion value provided for DtpPicas, verifying result is not NaN.
		result := a.DtpPicas()
		if math.IsNaN(result) {
			t.Errorf("conversion to DtpPicas returned NaN")
		}
	}
	{
		// Test conversion to Twips.
		// No expected conversion value provided for Twips, verifying result is not NaN.
		result := a.Twips()
		if math.IsNaN(result) {
			t.Errorf("conversion to Twips returned NaN")
		}
	}
	{
		// Test conversion to Hands.
		// No expected conversion value provided for Hands, verifying result is not NaN.
		result := a.Hands()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hands returned NaN")
		}
	}
	{
		// Test conversion to AstronomicalUnits.
		// No expected conversion value provided for AstronomicalUnits, verifying result is not NaN.
		result := a.AstronomicalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to AstronomicalUnits returned NaN")
		}
	}
	{
		// Test conversion to Parsecs.
		// No expected conversion value provided for Parsecs, verifying result is not NaN.
		result := a.Parsecs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Parsecs returned NaN")
		}
	}
	{
		// Test conversion to LightYears.
		// No expected conversion value provided for LightYears, verifying result is not NaN.
		result := a.LightYears()
		if math.IsNaN(result) {
			t.Errorf("conversion to LightYears returned NaN")
		}
	}
	{
		// Test conversion to SolarRadiuses.
		// No expected conversion value provided for SolarRadiuses, verifying result is not NaN.
		result := a.SolarRadiuses()
		if math.IsNaN(result) {
			t.Errorf("conversion to SolarRadiuses returned NaN")
		}
	}
	{
		// Test conversion to Chains.
		// No expected conversion value provided for Chains, verifying result is not NaN.
		result := a.Chains()
		if math.IsNaN(result) {
			t.Errorf("conversion to Chains returned NaN")
		}
	}
	{
		// Test conversion to Angstroms.
		// No expected conversion value provided for Angstroms, verifying result is not NaN.
		result := a.Angstroms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Angstroms returned NaN")
		}
	}
	{
		// Test conversion to DataMiles.
		// No expected conversion value provided for DataMiles, verifying result is not NaN.
		result := a.DataMiles()
		if math.IsNaN(result) {
			t.Errorf("conversion to DataMiles returned NaN")
		}
	}
	{
		// Test conversion to Femtometers.
		// No expected conversion value provided for Femtometers, verifying result is not NaN.
		result := a.Femtometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to Femtometers returned NaN")
		}
	}
	{
		// Test conversion to Picometers.
		// No expected conversion value provided for Picometers, verifying result is not NaN.
		result := a.Picometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picometers returned NaN")
		}
	}
	{
		// Test conversion to Nanometers.
		// No expected conversion value provided for Nanometers, verifying result is not NaN.
		result := a.Nanometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanometers returned NaN")
		}
	}
	{
		// Test conversion to Micrometers.
		// No expected conversion value provided for Micrometers, verifying result is not NaN.
		result := a.Micrometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micrometers returned NaN")
		}
	}
	{
		// Test conversion to Millimeters.
		// No expected conversion value provided for Millimeters, verifying result is not NaN.
		result := a.Millimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millimeters returned NaN")
		}
	}
	{
		// Test conversion to Centimeters.
		// No expected conversion value provided for Centimeters, verifying result is not NaN.
		result := a.Centimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centimeters returned NaN")
		}
	}
	{
		// Test conversion to Decimeters.
		// No expected conversion value provided for Decimeters, verifying result is not NaN.
		result := a.Decimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decimeters returned NaN")
		}
	}
	{
		// Test conversion to Decameters.
		// No expected conversion value provided for Decameters, verifying result is not NaN.
		result := a.Decameters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decameters returned NaN")
		}
	}
	{
		// Test conversion to Hectometers.
		// No expected conversion value provided for Hectometers, verifying result is not NaN.
		result := a.Hectometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hectometers returned NaN")
		}
	}
	{
		// Test conversion to Kilometers.
		// No expected conversion value provided for Kilometers, verifying result is not NaN.
		result := a.Kilometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilometers returned NaN")
		}
	}
	{
		// Test conversion to Megameters.
		// No expected conversion value provided for Megameters, verifying result is not NaN.
		result := a.Megameters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megameters returned NaN")
		}
	}
	{
		// Test conversion to Gigameters.
		// No expected conversion value provided for Gigameters, verifying result is not NaN.
		result := a.Gigameters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigameters returned NaN")
		}
	}
	{
		// Test conversion to Kiloyards.
		// No expected conversion value provided for Kiloyards, verifying result is not NaN.
		result := a.Kiloyards()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kiloyards returned NaN")
		}
	}
	{
		// Test conversion to Kilofeet.
		// No expected conversion value provided for Kilofeet, verifying result is not NaN.
		result := a.Kilofeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilofeet returned NaN")
		}
	}
	{
		// Test conversion to Kiloparsecs.
		// No expected conversion value provided for Kiloparsecs, verifying result is not NaN.
		result := a.Kiloparsecs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kiloparsecs returned NaN")
		}
	}
	{
		// Test conversion to Megaparsecs.
		// No expected conversion value provided for Megaparsecs, verifying result is not NaN.
		result := a.Megaparsecs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megaparsecs returned NaN")
		}
	}
	{
		// Test conversion to KilolightYears.
		// No expected conversion value provided for KilolightYears, verifying result is not NaN.
		result := a.KilolightYears()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilolightYears returned NaN")
		}
	}
	{
		// Test conversion to MegalightYears.
		// No expected conversion value provided for MegalightYears, verifying result is not NaN.
		result := a.MegalightYears()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegalightYears returned NaN")
		}
	}
}

func TestLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LengthFactory{}
	a, err := factory.CreateLength(90, units.LengthMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LengthMeter {
		t.Errorf("expected default unit Meter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LengthMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LengthMeter {
		t.Errorf("expected unit Meter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLengthToString(t *testing.T) {
	factory := units.LengthFactory{}
	a, err := factory.CreateLength(45, units.LengthMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LengthMeter, 2)
	expected := "45.00 " + units.GetLengthAbbreviation(units.LengthMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LengthMeter, -1)
	expected = "45 " + units.GetLengthAbbreviation(units.LengthMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLength_EqualityAndComparison(t *testing.T) {
	factory := units.LengthFactory{}
	a1, _ := factory.CreateLength(60, units.LengthMeter)
	a2, _ := factory.CreateLength(60, units.LengthMeter)
	a3, _ := factory.CreateLength(120, units.LengthMeter)

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

func TestLength_Arithmetic(t *testing.T) {
	factory := units.LengthFactory{}
	a1, _ := factory.CreateLength(30, units.LengthMeter)
	a2, _ := factory.CreateLength(45, units.LengthMeter)

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