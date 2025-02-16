// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPressureDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Pascal"}`
	
	factory := units.PressureDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PressurePascal {
		t.Errorf("expected unit %v, got %v", units.PressurePascal, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Pascal"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPressureDto_ToJSON(t *testing.T) {
	dto := units.PressureDto{
		Value: 45,
		Unit:  units.PressurePascal,
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
	if result["unit"].(string) != string(units.PressurePascal) {
		t.Errorf("expected unit %s, got %v", units.PressurePascal, result["unit"])
	}
}

func TestNewPressure_InvalidValue(t *testing.T) {
	factory := units.PressureFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePressure(math.NaN(), units.PressurePascal)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePressure(math.Inf(1), units.PressurePascal)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPressureConversions(t *testing.T) {
	factory := units.PressureFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePressure(180, units.PressurePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Pascals.
		// No expected conversion value provided for Pascals, verifying result is not NaN.
		result := a.Pascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Pascals returned NaN")
		}
	}
	{
		// Test conversion to Atmospheres.
		// No expected conversion value provided for Atmospheres, verifying result is not NaN.
		result := a.Atmospheres()
		if math.IsNaN(result) {
			t.Errorf("conversion to Atmospheres returned NaN")
		}
	}
	{
		// Test conversion to Bars.
		// No expected conversion value provided for Bars, verifying result is not NaN.
		result := a.Bars()
		if math.IsNaN(result) {
			t.Errorf("conversion to Bars returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerSquareMeter.
		// No expected conversion value provided for KilogramsForcePerSquareMeter, verifying result is not NaN.
		result := a.KilogramsForcePerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerSquareCentimeter.
		// No expected conversion value provided for KilogramsForcePerSquareCentimeter, verifying result is not NaN.
		result := a.KilogramsForcePerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerSquareMillimeter.
		// No expected conversion value provided for KilogramsForcePerSquareMillimeter, verifying result is not NaN.
		result := a.KilogramsForcePerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForcePerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerSquareMeter.
		// No expected conversion value provided for NewtonsPerSquareMeter, verifying result is not NaN.
		result := a.NewtonsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerSquareCentimeter.
		// No expected conversion value provided for NewtonsPerSquareCentimeter, verifying result is not NaN.
		result := a.NewtonsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerSquareMillimeter.
		// No expected conversion value provided for NewtonsPerSquareMillimeter, verifying result is not NaN.
		result := a.NewtonsPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonsPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TechnicalAtmospheres.
		// No expected conversion value provided for TechnicalAtmospheres, verifying result is not NaN.
		result := a.TechnicalAtmospheres()
		if math.IsNaN(result) {
			t.Errorf("conversion to TechnicalAtmospheres returned NaN")
		}
	}
	{
		// Test conversion to Torrs.
		// No expected conversion value provided for Torrs, verifying result is not NaN.
		result := a.Torrs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Torrs returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareInch.
		// No expected conversion value provided for PoundsForcePerSquareInch, verifying result is not NaN.
		result := a.PoundsForcePerSquareInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareMil.
		// No expected conversion value provided for PoundsForcePerSquareMil, verifying result is not NaN.
		result := a.PoundsForcePerSquareMil()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerSquareMil returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareFoot.
		// No expected conversion value provided for PoundsForcePerSquareFoot, verifying result is not NaN.
		result := a.PoundsForcePerSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForcePerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerSquareMillimeter.
		// No expected conversion value provided for TonnesForcePerSquareMillimeter, verifying result is not NaN.
		result := a.TonnesForcePerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerSquareMeter.
		// No expected conversion value provided for TonnesForcePerSquareMeter, verifying result is not NaN.
		result := a.TonnesForcePerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MetersOfHead.
		// No expected conversion value provided for MetersOfHead, verifying result is not NaN.
		result := a.MetersOfHead()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersOfHead returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerSquareCentimeter.
		// No expected conversion value provided for TonnesForcePerSquareCentimeter, verifying result is not NaN.
		result := a.TonnesForcePerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForcePerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to FeetOfHead.
		// No expected conversion value provided for FeetOfHead, verifying result is not NaN.
		result := a.FeetOfHead()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetOfHead returned NaN")
		}
	}
	{
		// Test conversion to MillimetersOfMercury.
		// No expected conversion value provided for MillimetersOfMercury, verifying result is not NaN.
		result := a.MillimetersOfMercury()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersOfMercury returned NaN")
		}
	}
	{
		// Test conversion to InchesOfMercury.
		// No expected conversion value provided for InchesOfMercury, verifying result is not NaN.
		result := a.InchesOfMercury()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesOfMercury returned NaN")
		}
	}
	{
		// Test conversion to DynesPerSquareCentimeter.
		// No expected conversion value provided for DynesPerSquareCentimeter, verifying result is not NaN.
		result := a.DynesPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DynesPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerInchSecondSquared.
		// No expected conversion value provided for PoundsPerInchSecondSquared, verifying result is not NaN.
		result := a.PoundsPerInchSecondSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerInchSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to MetersOfWaterColumn.
		// No expected conversion value provided for MetersOfWaterColumn, verifying result is not NaN.
		result := a.MetersOfWaterColumn()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersOfWaterColumn returned NaN")
		}
	}
	{
		// Test conversion to InchesOfWaterColumn.
		// No expected conversion value provided for InchesOfWaterColumn, verifying result is not NaN.
		result := a.InchesOfWaterColumn()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesOfWaterColumn returned NaN")
		}
	}
	{
		// Test conversion to MetersOfElevation.
		// No expected conversion value provided for MetersOfElevation, verifying result is not NaN.
		result := a.MetersOfElevation()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersOfElevation returned NaN")
		}
	}
	{
		// Test conversion to FeetOfElevation.
		// No expected conversion value provided for FeetOfElevation, verifying result is not NaN.
		result := a.FeetOfElevation()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetOfElevation returned NaN")
		}
	}
	{
		// Test conversion to Micropascals.
		// No expected conversion value provided for Micropascals, verifying result is not NaN.
		result := a.Micropascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micropascals returned NaN")
		}
	}
	{
		// Test conversion to Millipascals.
		// No expected conversion value provided for Millipascals, verifying result is not NaN.
		result := a.Millipascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millipascals returned NaN")
		}
	}
	{
		// Test conversion to Decapascals.
		// No expected conversion value provided for Decapascals, verifying result is not NaN.
		result := a.Decapascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decapascals returned NaN")
		}
	}
	{
		// Test conversion to Hectopascals.
		// No expected conversion value provided for Hectopascals, verifying result is not NaN.
		result := a.Hectopascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hectopascals returned NaN")
		}
	}
	{
		// Test conversion to Kilopascals.
		// No expected conversion value provided for Kilopascals, verifying result is not NaN.
		result := a.Kilopascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilopascals returned NaN")
		}
	}
	{
		// Test conversion to Megapascals.
		// No expected conversion value provided for Megapascals, verifying result is not NaN.
		result := a.Megapascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megapascals returned NaN")
		}
	}
	{
		// Test conversion to Gigapascals.
		// No expected conversion value provided for Gigapascals, verifying result is not NaN.
		result := a.Gigapascals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigapascals returned NaN")
		}
	}
	{
		// Test conversion to Microbars.
		// No expected conversion value provided for Microbars, verifying result is not NaN.
		result := a.Microbars()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microbars returned NaN")
		}
	}
	{
		// Test conversion to Millibars.
		// No expected conversion value provided for Millibars, verifying result is not NaN.
		result := a.Millibars()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millibars returned NaN")
		}
	}
	{
		// Test conversion to Centibars.
		// No expected conversion value provided for Centibars, verifying result is not NaN.
		result := a.Centibars()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centibars returned NaN")
		}
	}
	{
		// Test conversion to Decibars.
		// No expected conversion value provided for Decibars, verifying result is not NaN.
		result := a.Decibars()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decibars returned NaN")
		}
	}
	{
		// Test conversion to Kilobars.
		// No expected conversion value provided for Kilobars, verifying result is not NaN.
		result := a.Kilobars()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilobars returned NaN")
		}
	}
	{
		// Test conversion to Megabars.
		// No expected conversion value provided for Megabars, verifying result is not NaN.
		result := a.Megabars()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megabars returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerSquareMeter.
		// No expected conversion value provided for KilonewtonsPerSquareMeter, verifying result is not NaN.
		result := a.KilonewtonsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerSquareMeter.
		// No expected conversion value provided for MeganewtonsPerSquareMeter, verifying result is not NaN.
		result := a.MeganewtonsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerSquareCentimeter.
		// No expected conversion value provided for KilonewtonsPerSquareCentimeter, verifying result is not NaN.
		result := a.KilonewtonsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerSquareMillimeter.
		// No expected conversion value provided for KilonewtonsPerSquareMillimeter, verifying result is not NaN.
		result := a.KilonewtonsPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonsPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareInch.
		// No expected conversion value provided for KilopoundsForcePerSquareInch, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareMil.
		// No expected conversion value provided for KilopoundsForcePerSquareMil, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareMil()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerSquareMil returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareFoot.
		// No expected conversion value provided for KilopoundsForcePerSquareFoot, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForcePerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to MillimetersOfWaterColumn.
		// No expected conversion value provided for MillimetersOfWaterColumn, verifying result is not NaN.
		result := a.MillimetersOfWaterColumn()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersOfWaterColumn returned NaN")
		}
	}
	{
		// Test conversion to CentimetersOfWaterColumn.
		// No expected conversion value provided for CentimetersOfWaterColumn, verifying result is not NaN.
		result := a.CentimetersOfWaterColumn()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersOfWaterColumn returned NaN")
		}
	}
}

func TestPressure_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PressureFactory{}
	a, err := factory.CreatePressure(90, units.PressurePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PressurePascal {
		t.Errorf("expected default unit Pascal, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PressurePascal
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PressureDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PressurePascal {
		t.Errorf("expected unit Pascal, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPressureToString(t *testing.T) {
	factory := units.PressureFactory{}
	a, err := factory.CreatePressure(45, units.PressurePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PressurePascal, 2)
	expected := "45.00 " + units.GetPressureAbbreviation(units.PressurePascal)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PressurePascal, -1)
	expected = "45 " + units.GetPressureAbbreviation(units.PressurePascal)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPressure_EqualityAndComparison(t *testing.T) {
	factory := units.PressureFactory{}
	a1, _ := factory.CreatePressure(60, units.PressurePascal)
	a2, _ := factory.CreatePressure(60, units.PressurePascal)
	a3, _ := factory.CreatePressure(120, units.PressurePascal)

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

func TestPressure_Arithmetic(t *testing.T) {
	factory := units.PressureFactory{}
	a1, _ := factory.CreatePressure(30, units.PressurePascal)
	a2, _ := factory.CreatePressure(45, units.PressurePascal)

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