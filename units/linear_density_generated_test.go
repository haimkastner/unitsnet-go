// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLinearDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerMeter"}`
	
	factory := units.LinearDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LinearDensityKilogramPerMeter {
		t.Errorf("expected unit %v, got %v", units.LinearDensityKilogramPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLinearDensityDto_ToJSON(t *testing.T) {
	dto := units.LinearDensityDto{
		Value: 45,
		Unit:  units.LinearDensityKilogramPerMeter,
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
	if result["unit"].(string) != string(units.LinearDensityKilogramPerMeter) {
		t.Errorf("expected unit %s, got %v", units.LinearDensityKilogramPerMeter, result["unit"])
	}
}

func TestNewLinearDensity_InvalidValue(t *testing.T) {
	factory := units.LinearDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLinearDensity(math.NaN(), units.LinearDensityKilogramPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLinearDensity(math.Inf(1), units.LinearDensityKilogramPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLinearDensityConversions(t *testing.T) {
	factory := units.LinearDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLinearDensity(180, units.LinearDensityKilogramPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerMillimeter.
		// No expected conversion value provided for GramsPerMillimeter, verifying result is not NaN.
		result := a.GramsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCentimeter.
		// No expected conversion value provided for GramsPerCentimeter, verifying result is not NaN.
		result := a.GramsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerMeter.
		// No expected conversion value provided for GramsPerMeter, verifying result is not NaN.
		result := a.GramsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerInch.
		// No expected conversion value provided for PoundsPerInch, verifying result is not NaN.
		result := a.PoundsPerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerFoot.
		// No expected conversion value provided for PoundsPerFoot, verifying result is not NaN.
		result := a.PoundsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to GramsPerFoot.
		// No expected conversion value provided for GramsPerFoot, verifying result is not NaN.
		result := a.GramsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMillimeter.
		// No expected conversion value provided for MicrogramsPerMillimeter, verifying result is not NaN.
		result := a.MicrogramsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMillimeter.
		// No expected conversion value provided for MilligramsPerMillimeter, verifying result is not NaN.
		result := a.MilligramsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerMillimeter.
		// No expected conversion value provided for KilogramsPerMillimeter, verifying result is not NaN.
		result := a.KilogramsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerCentimeter.
		// No expected conversion value provided for MicrogramsPerCentimeter, verifying result is not NaN.
		result := a.MicrogramsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerCentimeter.
		// No expected conversion value provided for MilligramsPerCentimeter, verifying result is not NaN.
		result := a.MilligramsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCentimeter.
		// No expected conversion value provided for KilogramsPerCentimeter, verifying result is not NaN.
		result := a.KilogramsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMeter.
		// No expected conversion value provided for MicrogramsPerMeter, verifying result is not NaN.
		result := a.MicrogramsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMeter.
		// No expected conversion value provided for MilligramsPerMeter, verifying result is not NaN.
		result := a.MilligramsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerMeter.
		// No expected conversion value provided for KilogramsPerMeter, verifying result is not NaN.
		result := a.KilogramsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerFoot.
		// No expected conversion value provided for MicrogramsPerFoot, verifying result is not NaN.
		result := a.MicrogramsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerFoot.
		// No expected conversion value provided for MilligramsPerFoot, verifying result is not NaN.
		result := a.MilligramsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerFoot.
		// No expected conversion value provided for KilogramsPerFoot, verifying result is not NaN.
		result := a.KilogramsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerFoot returned NaN")
		}
	}
}

func TestLinearDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LinearDensityFactory{}
	a, err := factory.CreateLinearDensity(90, units.LinearDensityKilogramPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LinearDensityKilogramPerMeter {
		t.Errorf("expected default unit KilogramPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LinearDensityGramPerMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LinearDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LinearDensityKilogramPerMeter {
		t.Errorf("expected unit KilogramPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLinearDensityToString(t *testing.T) {
	factory := units.LinearDensityFactory{}
	a, err := factory.CreateLinearDensity(45, units.LinearDensityKilogramPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LinearDensityKilogramPerMeter, 2)
	expected := "45.00 " + units.GetLinearDensityAbbreviation(units.LinearDensityKilogramPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LinearDensityKilogramPerMeter, -1)
	expected = "45 " + units.GetLinearDensityAbbreviation(units.LinearDensityKilogramPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLinearDensity_EqualityAndComparison(t *testing.T) {
	factory := units.LinearDensityFactory{}
	a1, _ := factory.CreateLinearDensity(60, units.LinearDensityKilogramPerMeter)
	a2, _ := factory.CreateLinearDensity(60, units.LinearDensityKilogramPerMeter)
	a3, _ := factory.CreateLinearDensity(120, units.LinearDensityKilogramPerMeter)

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

func TestLinearDensity_Arithmetic(t *testing.T) {
	factory := units.LinearDensityFactory{}
	a1, _ := factory.CreateLinearDensity(30, units.LinearDensityKilogramPerMeter)
	a2, _ := factory.CreateLinearDensity(45, units.LinearDensityKilogramPerMeter)

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