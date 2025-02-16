// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumeDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CubicMeter"}`
	
	factory := units.VolumeDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumeCubicMeter {
		t.Errorf("expected unit %v, got %v", units.VolumeCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumeDto_ToJSON(t *testing.T) {
	dto := units.VolumeDto{
		Value: 45,
		Unit:  units.VolumeCubicMeter,
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
	if result["unit"].(string) != string(units.VolumeCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.VolumeCubicMeter, result["unit"])
	}
}

func TestNewVolume_InvalidValue(t *testing.T) {
	factory := units.VolumeFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolume(math.NaN(), units.VolumeCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolume(math.Inf(1), units.VolumeCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumeConversions(t *testing.T) {
	factory := units.VolumeFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolume(180, units.VolumeCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Liters.
		// No expected conversion value provided for Liters, verifying result is not NaN.
		result := a.Liters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Liters returned NaN")
		}
	}
	{
		// Test conversion to CubicMeters.
		// No expected conversion value provided for CubicMeters, verifying result is not NaN.
		result := a.CubicMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMeters returned NaN")
		}
	}
	{
		// Test conversion to CubicKilometers.
		// No expected conversion value provided for CubicKilometers, verifying result is not NaN.
		result := a.CubicKilometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicKilometers returned NaN")
		}
	}
	{
		// Test conversion to CubicHectometers.
		// No expected conversion value provided for CubicHectometers, verifying result is not NaN.
		result := a.CubicHectometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicHectometers returned NaN")
		}
	}
	{
		// Test conversion to CubicDecimeters.
		// No expected conversion value provided for CubicDecimeters, verifying result is not NaN.
		result := a.CubicDecimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicDecimeters returned NaN")
		}
	}
	{
		// Test conversion to CubicCentimeters.
		// No expected conversion value provided for CubicCentimeters, verifying result is not NaN.
		result := a.CubicCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicCentimeters returned NaN")
		}
	}
	{
		// Test conversion to CubicMillimeters.
		// No expected conversion value provided for CubicMillimeters, verifying result is not NaN.
		result := a.CubicMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMillimeters returned NaN")
		}
	}
	{
		// Test conversion to CubicMicrometers.
		// No expected conversion value provided for CubicMicrometers, verifying result is not NaN.
		result := a.CubicMicrometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMicrometers returned NaN")
		}
	}
	{
		// Test conversion to CubicMiles.
		// No expected conversion value provided for CubicMiles, verifying result is not NaN.
		result := a.CubicMiles()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMiles returned NaN")
		}
	}
	{
		// Test conversion to CubicYards.
		// No expected conversion value provided for CubicYards, verifying result is not NaN.
		result := a.CubicYards()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicYards returned NaN")
		}
	}
	{
		// Test conversion to CubicFeet.
		// No expected conversion value provided for CubicFeet, verifying result is not NaN.
		result := a.CubicFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicFeet returned NaN")
		}
	}
	{
		// Test conversion to CubicInches.
		// No expected conversion value provided for CubicInches, verifying result is not NaN.
		result := a.CubicInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicInches returned NaN")
		}
	}
	{
		// Test conversion to ImperialGallons.
		// No expected conversion value provided for ImperialGallons, verifying result is not NaN.
		result := a.ImperialGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to ImperialGallons returned NaN")
		}
	}
	{
		// Test conversion to ImperialOunces.
		// No expected conversion value provided for ImperialOunces, verifying result is not NaN.
		result := a.ImperialOunces()
		if math.IsNaN(result) {
			t.Errorf("conversion to ImperialOunces returned NaN")
		}
	}
	{
		// Test conversion to UsGallons.
		// No expected conversion value provided for UsGallons, verifying result is not NaN.
		result := a.UsGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsGallons returned NaN")
		}
	}
	{
		// Test conversion to UsOunces.
		// No expected conversion value provided for UsOunces, verifying result is not NaN.
		result := a.UsOunces()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsOunces returned NaN")
		}
	}
	{
		// Test conversion to UsTablespoons.
		// No expected conversion value provided for UsTablespoons, verifying result is not NaN.
		result := a.UsTablespoons()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsTablespoons returned NaN")
		}
	}
	{
		// Test conversion to AuTablespoons.
		// No expected conversion value provided for AuTablespoons, verifying result is not NaN.
		result := a.AuTablespoons()
		if math.IsNaN(result) {
			t.Errorf("conversion to AuTablespoons returned NaN")
		}
	}
	{
		// Test conversion to UkTablespoons.
		// No expected conversion value provided for UkTablespoons, verifying result is not NaN.
		result := a.UkTablespoons()
		if math.IsNaN(result) {
			t.Errorf("conversion to UkTablespoons returned NaN")
		}
	}
	{
		// Test conversion to MetricTeaspoons.
		// No expected conversion value provided for MetricTeaspoons, verifying result is not NaN.
		result := a.MetricTeaspoons()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetricTeaspoons returned NaN")
		}
	}
	{
		// Test conversion to UsTeaspoons.
		// No expected conversion value provided for UsTeaspoons, verifying result is not NaN.
		result := a.UsTeaspoons()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsTeaspoons returned NaN")
		}
	}
	{
		// Test conversion to MetricCups.
		// No expected conversion value provided for MetricCups, verifying result is not NaN.
		result := a.MetricCups()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetricCups returned NaN")
		}
	}
	{
		// Test conversion to UsCustomaryCups.
		// No expected conversion value provided for UsCustomaryCups, verifying result is not NaN.
		result := a.UsCustomaryCups()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsCustomaryCups returned NaN")
		}
	}
	{
		// Test conversion to UsLegalCups.
		// No expected conversion value provided for UsLegalCups, verifying result is not NaN.
		result := a.UsLegalCups()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsLegalCups returned NaN")
		}
	}
	{
		// Test conversion to OilBarrels.
		// No expected conversion value provided for OilBarrels, verifying result is not NaN.
		result := a.OilBarrels()
		if math.IsNaN(result) {
			t.Errorf("conversion to OilBarrels returned NaN")
		}
	}
	{
		// Test conversion to UsBeerBarrels.
		// No expected conversion value provided for UsBeerBarrels, verifying result is not NaN.
		result := a.UsBeerBarrels()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsBeerBarrels returned NaN")
		}
	}
	{
		// Test conversion to ImperialBeerBarrels.
		// No expected conversion value provided for ImperialBeerBarrels, verifying result is not NaN.
		result := a.ImperialBeerBarrels()
		if math.IsNaN(result) {
			t.Errorf("conversion to ImperialBeerBarrels returned NaN")
		}
	}
	{
		// Test conversion to UsQuarts.
		// No expected conversion value provided for UsQuarts, verifying result is not NaN.
		result := a.UsQuarts()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsQuarts returned NaN")
		}
	}
	{
		// Test conversion to ImperialQuarts.
		// No expected conversion value provided for ImperialQuarts, verifying result is not NaN.
		result := a.ImperialQuarts()
		if math.IsNaN(result) {
			t.Errorf("conversion to ImperialQuarts returned NaN")
		}
	}
	{
		// Test conversion to UsPints.
		// No expected conversion value provided for UsPints, verifying result is not NaN.
		result := a.UsPints()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsPints returned NaN")
		}
	}
	{
		// Test conversion to AcreFeet.
		// No expected conversion value provided for AcreFeet, verifying result is not NaN.
		result := a.AcreFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to AcreFeet returned NaN")
		}
	}
	{
		// Test conversion to ImperialPints.
		// No expected conversion value provided for ImperialPints, verifying result is not NaN.
		result := a.ImperialPints()
		if math.IsNaN(result) {
			t.Errorf("conversion to ImperialPints returned NaN")
		}
	}
	{
		// Test conversion to BoardFeet.
		// No expected conversion value provided for BoardFeet, verifying result is not NaN.
		result := a.BoardFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to BoardFeet returned NaN")
		}
	}
	{
		// Test conversion to Nanoliters.
		// No expected conversion value provided for Nanoliters, verifying result is not NaN.
		result := a.Nanoliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanoliters returned NaN")
		}
	}
	{
		// Test conversion to Microliters.
		// No expected conversion value provided for Microliters, verifying result is not NaN.
		result := a.Microliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microliters returned NaN")
		}
	}
	{
		// Test conversion to Milliliters.
		// No expected conversion value provided for Milliliters, verifying result is not NaN.
		result := a.Milliliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliliters returned NaN")
		}
	}
	{
		// Test conversion to Centiliters.
		// No expected conversion value provided for Centiliters, verifying result is not NaN.
		result := a.Centiliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centiliters returned NaN")
		}
	}
	{
		// Test conversion to Deciliters.
		// No expected conversion value provided for Deciliters, verifying result is not NaN.
		result := a.Deciliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Deciliters returned NaN")
		}
	}
	{
		// Test conversion to Decaliters.
		// No expected conversion value provided for Decaliters, verifying result is not NaN.
		result := a.Decaliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decaliters returned NaN")
		}
	}
	{
		// Test conversion to Hectoliters.
		// No expected conversion value provided for Hectoliters, verifying result is not NaN.
		result := a.Hectoliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hectoliters returned NaN")
		}
	}
	{
		// Test conversion to Kiloliters.
		// No expected conversion value provided for Kiloliters, verifying result is not NaN.
		result := a.Kiloliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kiloliters returned NaN")
		}
	}
	{
		// Test conversion to Megaliters.
		// No expected conversion value provided for Megaliters, verifying result is not NaN.
		result := a.Megaliters()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megaliters returned NaN")
		}
	}
	{
		// Test conversion to HectocubicMeters.
		// No expected conversion value provided for HectocubicMeters, verifying result is not NaN.
		result := a.HectocubicMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectocubicMeters returned NaN")
		}
	}
	{
		// Test conversion to KilocubicMeters.
		// No expected conversion value provided for KilocubicMeters, verifying result is not NaN.
		result := a.KilocubicMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocubicMeters returned NaN")
		}
	}
	{
		// Test conversion to HectocubicFeet.
		// No expected conversion value provided for HectocubicFeet, verifying result is not NaN.
		result := a.HectocubicFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectocubicFeet returned NaN")
		}
	}
	{
		// Test conversion to KilocubicFeet.
		// No expected conversion value provided for KilocubicFeet, verifying result is not NaN.
		result := a.KilocubicFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocubicFeet returned NaN")
		}
	}
	{
		// Test conversion to MegacubicFeet.
		// No expected conversion value provided for MegacubicFeet, verifying result is not NaN.
		result := a.MegacubicFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegacubicFeet returned NaN")
		}
	}
	{
		// Test conversion to KiloimperialGallons.
		// No expected conversion value provided for KiloimperialGallons, verifying result is not NaN.
		result := a.KiloimperialGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to KiloimperialGallons returned NaN")
		}
	}
	{
		// Test conversion to MegaimperialGallons.
		// No expected conversion value provided for MegaimperialGallons, verifying result is not NaN.
		result := a.MegaimperialGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegaimperialGallons returned NaN")
		}
	}
	{
		// Test conversion to DecausGallons.
		// No expected conversion value provided for DecausGallons, verifying result is not NaN.
		result := a.DecausGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecausGallons returned NaN")
		}
	}
	{
		// Test conversion to DeciusGallons.
		// No expected conversion value provided for DeciusGallons, verifying result is not NaN.
		result := a.DeciusGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to DeciusGallons returned NaN")
		}
	}
	{
		// Test conversion to HectousGallons.
		// No expected conversion value provided for HectousGallons, verifying result is not NaN.
		result := a.HectousGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectousGallons returned NaN")
		}
	}
	{
		// Test conversion to KilousGallons.
		// No expected conversion value provided for KilousGallons, verifying result is not NaN.
		result := a.KilousGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilousGallons returned NaN")
		}
	}
	{
		// Test conversion to MegausGallons.
		// No expected conversion value provided for MegausGallons, verifying result is not NaN.
		result := a.MegausGallons()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegausGallons returned NaN")
		}
	}
}

func TestVolume_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumeFactory{}
	a, err := factory.CreateVolume(90, units.VolumeCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumeCubicMeter {
		t.Errorf("expected default unit CubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumeLiter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumeDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumeCubicMeter {
		t.Errorf("expected unit CubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumeToString(t *testing.T) {
	factory := units.VolumeFactory{}
	a, err := factory.CreateVolume(45, units.VolumeCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumeCubicMeter, 2)
	expected := "45.00 " + units.GetVolumeAbbreviation(units.VolumeCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumeCubicMeter, -1)
	expected = "45 " + units.GetVolumeAbbreviation(units.VolumeCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolume_EqualityAndComparison(t *testing.T) {
	factory := units.VolumeFactory{}
	a1, _ := factory.CreateVolume(60, units.VolumeCubicMeter)
	a2, _ := factory.CreateVolume(60, units.VolumeCubicMeter)
	a3, _ := factory.CreateVolume(120, units.VolumeCubicMeter)

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

func TestVolume_Arithmetic(t *testing.T) {
	factory := units.VolumeFactory{}
	a1, _ := factory.CreateVolume(30, units.VolumeCubicMeter)
	a2, _ := factory.CreateVolume(45, units.VolumeCubicMeter)

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