// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumePerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CubicMeterPerMeter"}`
	
	factory := units.VolumePerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumePerLengthCubicMeterPerMeter {
		t.Errorf("expected unit %v, got %v", units.VolumePerLengthCubicMeterPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CubicMeterPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumePerLengthDto_ToJSON(t *testing.T) {
	dto := units.VolumePerLengthDto{
		Value: 45,
		Unit:  units.VolumePerLengthCubicMeterPerMeter,
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
	if result["unit"].(string) != string(units.VolumePerLengthCubicMeterPerMeter) {
		t.Errorf("expected unit %s, got %v", units.VolumePerLengthCubicMeterPerMeter, result["unit"])
	}
}

func TestNewVolumePerLength_InvalidValue(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolumePerLength(math.NaN(), units.VolumePerLengthCubicMeterPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolumePerLength(math.Inf(1), units.VolumePerLengthCubicMeterPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumePerLengthConversions(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolumePerLength(180, units.VolumePerLengthCubicMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CubicMetersPerMeter.
		// No expected conversion value provided for CubicMetersPerMeter, verifying result is not NaN.
		result := a.CubicMetersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to LitersPerMeter.
		// No expected conversion value provided for LitersPerMeter, verifying result is not NaN.
		result := a.LitersPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to LitersPerKilometer.
		// No expected conversion value provided for LitersPerKilometer, verifying result is not NaN.
		result := a.LitersPerKilometer()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerKilometer returned NaN")
		}
	}
	{
		// Test conversion to LitersPerMillimeter.
		// No expected conversion value provided for LitersPerMillimeter, verifying result is not NaN.
		result := a.LitersPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerFoot.
		// No expected conversion value provided for OilBarrelsPerFoot, verifying result is not NaN.
		result := a.OilBarrelsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to OilBarrelsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerFoot.
		// No expected conversion value provided for CubicYardsPerFoot, verifying result is not NaN.
		result := a.CubicYardsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicYardsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerUsSurveyFoot.
		// No expected conversion value provided for CubicYardsPerUsSurveyFoot, verifying result is not NaN.
		result := a.CubicYardsPerUsSurveyFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicYardsPerUsSurveyFoot returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerMile.
		// No expected conversion value provided for UsGallonsPerMile, verifying result is not NaN.
		result := a.UsGallonsPerMile()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsGallonsPerMile returned NaN")
		}
	}
	{
		// Test conversion to ImperialGallonsPerMile.
		// No expected conversion value provided for ImperialGallonsPerMile, verifying result is not NaN.
		result := a.ImperialGallonsPerMile()
		if math.IsNaN(result) {
			t.Errorf("conversion to ImperialGallonsPerMile returned NaN")
		}
	}
}

func TestVolumePerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	a, err := factory.CreateVolumePerLength(90, units.VolumePerLengthCubicMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumePerLengthCubicMeterPerMeter {
		t.Errorf("expected default unit CubicMeterPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumePerLengthCubicMeterPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumePerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumePerLengthCubicMeterPerMeter {
		t.Errorf("expected unit CubicMeterPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumePerLengthToString(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	a, err := factory.CreateVolumePerLength(45, units.VolumePerLengthCubicMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumePerLengthCubicMeterPerMeter, 2)
	expected := "45.00 " + units.GetVolumePerLengthAbbreviation(units.VolumePerLengthCubicMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumePerLengthCubicMeterPerMeter, -1)
	expected = "45 " + units.GetVolumePerLengthAbbreviation(units.VolumePerLengthCubicMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolumePerLength_EqualityAndComparison(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	a1, _ := factory.CreateVolumePerLength(60, units.VolumePerLengthCubicMeterPerMeter)
	a2, _ := factory.CreateVolumePerLength(60, units.VolumePerLengthCubicMeterPerMeter)
	a3, _ := factory.CreateVolumePerLength(120, units.VolumePerLengthCubicMeterPerMeter)

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

func TestVolumePerLength_Arithmetic(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	a1, _ := factory.CreateVolumePerLength(30, units.VolumePerLengthCubicMeterPerMeter)
	a2, _ := factory.CreateVolumePerLength(45, units.VolumePerLengthCubicMeterPerMeter)

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