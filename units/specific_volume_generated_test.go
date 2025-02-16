// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificVolumeDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CubicMeterPerKilogram"}`
	
	factory := units.SpecificVolumeDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificVolumeCubicMeterPerKilogram {
		t.Errorf("expected unit %v, got %v", units.SpecificVolumeCubicMeterPerKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CubicMeterPerKilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificVolumeDto_ToJSON(t *testing.T) {
	dto := units.SpecificVolumeDto{
		Value: 45,
		Unit:  units.SpecificVolumeCubicMeterPerKilogram,
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
	if result["unit"].(string) != string(units.SpecificVolumeCubicMeterPerKilogram) {
		t.Errorf("expected unit %s, got %v", units.SpecificVolumeCubicMeterPerKilogram, result["unit"])
	}
}

func TestNewSpecificVolume_InvalidValue(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificVolume(math.NaN(), units.SpecificVolumeCubicMeterPerKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificVolume(math.Inf(1), units.SpecificVolumeCubicMeterPerKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificVolumeConversions(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificVolume(180, units.SpecificVolumeCubicMeterPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CubicMetersPerKilogram.
		// No expected conversion value provided for CubicMetersPerKilogram, verifying result is not NaN.
		result := a.CubicMetersPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMetersPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerPound.
		// No expected conversion value provided for CubicFeetPerPound, verifying result is not NaN.
		result := a.CubicFeetPerPound()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicFeetPerPound returned NaN")
		}
	}
	{
		// Test conversion to MillicubicMetersPerKilogram.
		// No expected conversion value provided for MillicubicMetersPerKilogram, verifying result is not NaN.
		result := a.MillicubicMetersPerKilogram()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillicubicMetersPerKilogram returned NaN")
		}
	}
}

func TestSpecificVolume_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	a, err := factory.CreateSpecificVolume(90, units.SpecificVolumeCubicMeterPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificVolumeCubicMeterPerKilogram {
		t.Errorf("expected default unit CubicMeterPerKilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificVolumeCubicMeterPerKilogram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificVolumeDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificVolumeCubicMeterPerKilogram {
		t.Errorf("expected unit CubicMeterPerKilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificVolumeFactory_FromDto(t *testing.T) {
    factory := units.SpecificVolumeFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.SpecificVolumeDto{
        Value: 100,
        Unit:  units.SpecificVolumeCubicMeterPerKilogram,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.SpecificVolumeDto{
        Value: math.NaN(),
        Unit:  units.SpecificVolumeCubicMeterPerKilogram,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test CubicMeterPerKilogram conversion
    cubic_meters_per_kilogramDto := units.SpecificVolumeDto{
        Value: 100,
        Unit:  units.SpecificVolumeCubicMeterPerKilogram,
    }
    
    var cubic_meters_per_kilogramResult *units.SpecificVolume
    cubic_meters_per_kilogramResult, err = factory.FromDto(cubic_meters_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMeterPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_kilogramResult.Convert(units.SpecificVolumeCubicMeterPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerKilogram = %v, want %v", converted, 100)
    }
    // Test CubicFootPerPound conversion
    cubic_feet_per_poundDto := units.SpecificVolumeDto{
        Value: 100,
        Unit:  units.SpecificVolumeCubicFootPerPound,
    }
    
    var cubic_feet_per_poundResult *units.SpecificVolume
    cubic_feet_per_poundResult, err = factory.FromDto(cubic_feet_per_poundDto)
    if err != nil {
        t.Errorf("FromDto() with CubicFootPerPound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_poundResult.Convert(units.SpecificVolumeCubicFootPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerPound = %v, want %v", converted, 100)
    }
    // Test MillicubicMeterPerKilogram conversion
    millicubic_meters_per_kilogramDto := units.SpecificVolumeDto{
        Value: 100,
        Unit:  units.SpecificVolumeMillicubicMeterPerKilogram,
    }
    
    var millicubic_meters_per_kilogramResult *units.SpecificVolume
    millicubic_meters_per_kilogramResult, err = factory.FromDto(millicubic_meters_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MillicubicMeterPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicubic_meters_per_kilogramResult.Convert(units.SpecificVolumeMillicubicMeterPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillicubicMeterPerKilogram = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.SpecificVolumeDto{
        Value: 0,
        Unit:  units.SpecificVolumeCubicMeterPerKilogram,
    }
    
    var zeroResult *units.SpecificVolume
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestSpecificVolumeFactory_FromDtoJSON(t *testing.T) {
    factory := units.SpecificVolumeFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CubicMeterPerKilogram"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CubicMeterPerKilogram"}`)
    _, err = factory.FromDtoJSON(invalidJSON)
    if err == nil {
        t.Error("FromDtoJSON() with invalid JSON should return error")
    }

    // Test malformed JSON
    malformedJSON := []byte(`{malformed json`)
    _, err = factory.FromDtoJSON(malformedJSON)
    if err == nil {
        t.Error("FromDtoJSON() with malformed JSON should return error")
    }

    // Test empty JSON
    emptyJSON := []byte(`{}`)
    _, err = factory.FromDtoJSON(emptyJSON)
    if err == nil {
        t.Error("FromDtoJSON() with empty JSON should return error")
    }

    // Test JSON with invalid value (NaN)
    nanValue := math.NaN()
    nanJSON, _ := json.Marshal(units.SpecificVolumeDto{
        Value: nanValue,
        Unit:  units.SpecificVolumeCubicMeterPerKilogram,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with CubicMeterPerKilogram unit
    cubic_meters_per_kilogramJSON := []byte(`{"value": 100, "unit": "CubicMeterPerKilogram"}`)
    cubic_meters_per_kilogramResult, err := factory.FromDtoJSON(cubic_meters_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMeterPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_kilogramResult.Convert(units.SpecificVolumeCubicMeterPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with CubicFootPerPound unit
    cubic_feet_per_poundJSON := []byte(`{"value": 100, "unit": "CubicFootPerPound"}`)
    cubic_feet_per_poundResult, err := factory.FromDtoJSON(cubic_feet_per_poundJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicFootPerPound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_poundResult.Convert(units.SpecificVolumeCubicFootPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerPound = %v, want %v", converted, 100)
    }
    // Test JSON with MillicubicMeterPerKilogram unit
    millicubic_meters_per_kilogramJSON := []byte(`{"value": 100, "unit": "MillicubicMeterPerKilogram"}`)
    millicubic_meters_per_kilogramResult, err := factory.FromDtoJSON(millicubic_meters_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillicubicMeterPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicubic_meters_per_kilogramResult.Convert(units.SpecificVolumeMillicubicMeterPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillicubicMeterPerKilogram = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CubicMeterPerKilogram"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCubicMetersPerKilogram function
func TestSpecificVolumeFactory_FromCubicMetersPerKilogram(t *testing.T) {
    factory := units.SpecificVolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMetersPerKilogram(100)
    if err != nil {
        t.Errorf("FromCubicMetersPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificVolumeCubicMeterPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMetersPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMetersPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromCubicMetersPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromCubicMetersPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMetersPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromCubicMetersPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMetersPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMetersPerKilogram(0)
    if err != nil {
        t.Errorf("FromCubicMetersPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificVolumeCubicMeterPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMetersPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicFeetPerPound function
func TestSpecificVolumeFactory_FromCubicFeetPerPound(t *testing.T) {
    factory := units.SpecificVolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicFeetPerPound(100)
    if err != nil {
        t.Errorf("FromCubicFeetPerPound() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificVolumeCubicFootPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicFeetPerPound() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicFeetPerPound(math.NaN())
    if err == nil {
        t.Error("FromCubicFeetPerPound() with NaN value should return error")
    }

    _, err = factory.FromCubicFeetPerPound(math.Inf(1))
    if err == nil {
        t.Error("FromCubicFeetPerPound() with +Inf value should return error")
    }

    _, err = factory.FromCubicFeetPerPound(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicFeetPerPound() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicFeetPerPound(0)
    if err != nil {
        t.Errorf("FromCubicFeetPerPound() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificVolumeCubicFootPerPound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicFeetPerPound() with zero value = %v, want 0", converted)
    }
}
// Test FromMillicubicMetersPerKilogram function
func TestSpecificVolumeFactory_FromMillicubicMetersPerKilogram(t *testing.T) {
    factory := units.SpecificVolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillicubicMetersPerKilogram(100)
    if err != nil {
        t.Errorf("FromMillicubicMetersPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificVolumeMillicubicMeterPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillicubicMetersPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillicubicMetersPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMillicubicMetersPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMillicubicMetersPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMillicubicMetersPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMillicubicMetersPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMillicubicMetersPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillicubicMetersPerKilogram(0)
    if err != nil {
        t.Errorf("FromMillicubicMetersPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificVolumeMillicubicMeterPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillicubicMetersPerKilogram() with zero value = %v, want 0", converted)
    }
}

func TestSpecificVolumeToString(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	a, err := factory.CreateSpecificVolume(45, units.SpecificVolumeCubicMeterPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificVolumeCubicMeterPerKilogram, 2)
	expected := "45.00 " + units.GetSpecificVolumeAbbreviation(units.SpecificVolumeCubicMeterPerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificVolumeCubicMeterPerKilogram, -1)
	expected = "45 " + units.GetSpecificVolumeAbbreviation(units.SpecificVolumeCubicMeterPerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificVolume_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	a1, _ := factory.CreateSpecificVolume(60, units.SpecificVolumeCubicMeterPerKilogram)
	a2, _ := factory.CreateSpecificVolume(60, units.SpecificVolumeCubicMeterPerKilogram)
	a3, _ := factory.CreateSpecificVolume(120, units.SpecificVolumeCubicMeterPerKilogram)

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

func TestSpecificVolume_Arithmetic(t *testing.T) {
	factory := units.SpecificVolumeFactory{}
	a1, _ := factory.CreateSpecificVolume(30, units.SpecificVolumeCubicMeterPerKilogram)
	a2, _ := factory.CreateSpecificVolume(45, units.SpecificVolumeCubicMeterPerKilogram)

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