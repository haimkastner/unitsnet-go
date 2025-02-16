// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRotationalAccelerationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "RadianPerSecondSquared"}`
	
	factory := units.RotationalAccelerationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RotationalAccelerationRadianPerSecondSquared {
		t.Errorf("expected unit %v, got %v", units.RotationalAccelerationRadianPerSecondSquared, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "RadianPerSecondSquared"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRotationalAccelerationDto_ToJSON(t *testing.T) {
	dto := units.RotationalAccelerationDto{
		Value: 45,
		Unit:  units.RotationalAccelerationRadianPerSecondSquared,
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
	if result["unit"].(string) != string(units.RotationalAccelerationRadianPerSecondSquared) {
		t.Errorf("expected unit %s, got %v", units.RotationalAccelerationRadianPerSecondSquared, result["unit"])
	}
}

func TestNewRotationalAcceleration_InvalidValue(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRotationalAcceleration(math.NaN(), units.RotationalAccelerationRadianPerSecondSquared)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRotationalAcceleration(math.Inf(1), units.RotationalAccelerationRadianPerSecondSquared)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRotationalAccelerationConversions(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRotationalAcceleration(180, units.RotationalAccelerationRadianPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to RadiansPerSecondSquared.
		// No expected conversion value provided for RadiansPerSecondSquared, verifying result is not NaN.
		result := a.RadiansPerSecondSquared()
		cacheResult := a.RadiansPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to RadiansPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to DegreesPerSecondSquared.
		// No expected conversion value provided for DegreesPerSecondSquared, verifying result is not NaN.
		result := a.DegreesPerSecondSquared()
		cacheResult := a.DegreesPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to RevolutionsPerMinutePerSecond.
		// No expected conversion value provided for RevolutionsPerMinutePerSecond, verifying result is not NaN.
		result := a.RevolutionsPerMinutePerSecond()
		cacheResult := a.RevolutionsPerMinutePerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to RevolutionsPerMinutePerSecond returned NaN")
		}
	}
	{
		// Test conversion to RevolutionsPerSecondSquared.
		// No expected conversion value provided for RevolutionsPerSecondSquared, verifying result is not NaN.
		result := a.RevolutionsPerSecondSquared()
		cacheResult := a.RevolutionsPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to RevolutionsPerSecondSquared returned NaN")
		}
	}
}

func TestRotationalAcceleration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	a, err := factory.CreateRotationalAcceleration(90, units.RotationalAccelerationRadianPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RotationalAccelerationRadianPerSecondSquared {
		t.Errorf("expected default unit RadianPerSecondSquared, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RotationalAccelerationRadianPerSecondSquared
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RotationalAccelerationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RotationalAccelerationRadianPerSecondSquared {
		t.Errorf("expected unit RadianPerSecondSquared, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRotationalAccelerationFactory_FromDto(t *testing.T) {
    factory := units.RotationalAccelerationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RotationalAccelerationDto{
        Value: 100,
        Unit:  units.RotationalAccelerationRadianPerSecondSquared,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RotationalAccelerationDto{
        Value: math.NaN(),
        Unit:  units.RotationalAccelerationRadianPerSecondSquared,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test RadianPerSecondSquared conversion
    radians_per_second_squaredDto := units.RotationalAccelerationDto{
        Value: 100,
        Unit:  units.RotationalAccelerationRadianPerSecondSquared,
    }
    
    var radians_per_second_squaredResult *units.RotationalAcceleration
    radians_per_second_squaredResult, err = factory.FromDto(radians_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with RadianPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radians_per_second_squaredResult.Convert(units.RotationalAccelerationRadianPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RadianPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test DegreePerSecondSquared conversion
    degrees_per_second_squaredDto := units.RotationalAccelerationDto{
        Value: 100,
        Unit:  units.RotationalAccelerationDegreePerSecondSquared,
    }
    
    var degrees_per_second_squaredResult *units.RotationalAcceleration
    degrees_per_second_squaredResult, err = factory.FromDto(degrees_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with DegreePerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_per_second_squaredResult.Convert(units.RotationalAccelerationDegreePerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreePerSecondSquared = %v, want %v", converted, 100)
    }
    // Test RevolutionPerMinutePerSecond conversion
    revolutions_per_minute_per_secondDto := units.RotationalAccelerationDto{
        Value: 100,
        Unit:  units.RotationalAccelerationRevolutionPerMinutePerSecond,
    }
    
    var revolutions_per_minute_per_secondResult *units.RotationalAcceleration
    revolutions_per_minute_per_secondResult, err = factory.FromDto(revolutions_per_minute_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with RevolutionPerMinutePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutions_per_minute_per_secondResult.Convert(units.RotationalAccelerationRevolutionPerMinutePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RevolutionPerMinutePerSecond = %v, want %v", converted, 100)
    }
    // Test RevolutionPerSecondSquared conversion
    revolutions_per_second_squaredDto := units.RotationalAccelerationDto{
        Value: 100,
        Unit:  units.RotationalAccelerationRevolutionPerSecondSquared,
    }
    
    var revolutions_per_second_squaredResult *units.RotationalAcceleration
    revolutions_per_second_squaredResult, err = factory.FromDto(revolutions_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with RevolutionPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutions_per_second_squaredResult.Convert(units.RotationalAccelerationRevolutionPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RevolutionPerSecondSquared = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RotationalAccelerationDto{
        Value: 0,
        Unit:  units.RotationalAccelerationRadianPerSecondSquared,
    }
    
    var zeroResult *units.RotationalAcceleration
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRotationalAccelerationFactory_FromDtoJSON(t *testing.T) {
    factory := units.RotationalAccelerationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "RadianPerSecondSquared"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "RadianPerSecondSquared"}`)
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
    nanJSON, _ := json.Marshal(units.RotationalAccelerationDto{
        Value: nanValue,
        Unit:  units.RotationalAccelerationRadianPerSecondSquared,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with RadianPerSecondSquared unit
    radians_per_second_squaredJSON := []byte(`{"value": 100, "unit": "RadianPerSecondSquared"}`)
    radians_per_second_squaredResult, err := factory.FromDtoJSON(radians_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RadianPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radians_per_second_squaredResult.Convert(units.RotationalAccelerationRadianPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RadianPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with DegreePerSecondSquared unit
    degrees_per_second_squaredJSON := []byte(`{"value": 100, "unit": "DegreePerSecondSquared"}`)
    degrees_per_second_squaredResult, err := factory.FromDtoJSON(degrees_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreePerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_per_second_squaredResult.Convert(units.RotationalAccelerationDegreePerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreePerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with RevolutionPerMinutePerSecond unit
    revolutions_per_minute_per_secondJSON := []byte(`{"value": 100, "unit": "RevolutionPerMinutePerSecond"}`)
    revolutions_per_minute_per_secondResult, err := factory.FromDtoJSON(revolutions_per_minute_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RevolutionPerMinutePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutions_per_minute_per_secondResult.Convert(units.RotationalAccelerationRevolutionPerMinutePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RevolutionPerMinutePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with RevolutionPerSecondSquared unit
    revolutions_per_second_squaredJSON := []byte(`{"value": 100, "unit": "RevolutionPerSecondSquared"}`)
    revolutions_per_second_squaredResult, err := factory.FromDtoJSON(revolutions_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RevolutionPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutions_per_second_squaredResult.Convert(units.RotationalAccelerationRevolutionPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RevolutionPerSecondSquared = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "RadianPerSecondSquared"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromRadiansPerSecondSquared function
func TestRotationalAccelerationFactory_FromRadiansPerSecondSquared(t *testing.T) {
    factory := units.RotationalAccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRadiansPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromRadiansPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalAccelerationRadianPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRadiansPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRadiansPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromRadiansPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromRadiansPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromRadiansPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromRadiansPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromRadiansPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRadiansPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromRadiansPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalAccelerationRadianPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRadiansPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesPerSecondSquared function
func TestRotationalAccelerationFactory_FromDegreesPerSecondSquared(t *testing.T) {
    factory := units.RotationalAccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromDegreesPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalAccelerationDegreePerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromDegreesPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromDegreesPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromDegreesPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromDegreesPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalAccelerationDegreePerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromRevolutionsPerMinutePerSecond function
func TestRotationalAccelerationFactory_FromRevolutionsPerMinutePerSecond(t *testing.T) {
    factory := units.RotationalAccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRevolutionsPerMinutePerSecond(100)
    if err != nil {
        t.Errorf("FromRevolutionsPerMinutePerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalAccelerationRevolutionPerMinutePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRevolutionsPerMinutePerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRevolutionsPerMinutePerSecond(math.NaN())
    if err == nil {
        t.Error("FromRevolutionsPerMinutePerSecond() with NaN value should return error")
    }

    _, err = factory.FromRevolutionsPerMinutePerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromRevolutionsPerMinutePerSecond() with +Inf value should return error")
    }

    _, err = factory.FromRevolutionsPerMinutePerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromRevolutionsPerMinutePerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRevolutionsPerMinutePerSecond(0)
    if err != nil {
        t.Errorf("FromRevolutionsPerMinutePerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalAccelerationRevolutionPerMinutePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRevolutionsPerMinutePerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromRevolutionsPerSecondSquared function
func TestRotationalAccelerationFactory_FromRevolutionsPerSecondSquared(t *testing.T) {
    factory := units.RotationalAccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRevolutionsPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromRevolutionsPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalAccelerationRevolutionPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRevolutionsPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRevolutionsPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromRevolutionsPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromRevolutionsPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromRevolutionsPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromRevolutionsPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromRevolutionsPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRevolutionsPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromRevolutionsPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalAccelerationRevolutionPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRevolutionsPerSecondSquared() with zero value = %v, want 0", converted)
    }
}

func TestRotationalAccelerationToString(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	a, err := factory.CreateRotationalAcceleration(45, units.RotationalAccelerationRadianPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RotationalAccelerationRadianPerSecondSquared, 2)
	expected := "45.00 " + units.GetRotationalAccelerationAbbreviation(units.RotationalAccelerationRadianPerSecondSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RotationalAccelerationRadianPerSecondSquared, -1)
	expected = "45 " + units.GetRotationalAccelerationAbbreviation(units.RotationalAccelerationRadianPerSecondSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRotationalAcceleration_EqualityAndComparison(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	a1, _ := factory.CreateRotationalAcceleration(60, units.RotationalAccelerationRadianPerSecondSquared)
	a2, _ := factory.CreateRotationalAcceleration(60, units.RotationalAccelerationRadianPerSecondSquared)
	a3, _ := factory.CreateRotationalAcceleration(120, units.RotationalAccelerationRadianPerSecondSquared)

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

func TestRotationalAcceleration_Arithmetic(t *testing.T) {
	factory := units.RotationalAccelerationFactory{}
	a1, _ := factory.CreateRotationalAcceleration(30, units.RotationalAccelerationRadianPerSecondSquared)
	a2, _ := factory.CreateRotationalAcceleration(45, units.RotationalAccelerationRadianPerSecondSquared)

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


func TestGetRotationalAccelerationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.RotationalAccelerationUnits
        want string
    }{
        {
            name: "RadianPerSecondSquared abbreviation",
            unit: units.RotationalAccelerationRadianPerSecondSquared,
            want: "rad/s²",
        },
        {
            name: "DegreePerSecondSquared abbreviation",
            unit: units.RotationalAccelerationDegreePerSecondSquared,
            want: "°/s²",
        },
        {
            name: "RevolutionPerMinutePerSecond abbreviation",
            unit: units.RotationalAccelerationRevolutionPerMinutePerSecond,
            want: "rpm/s",
        },
        {
            name: "RevolutionPerSecondSquared abbreviation",
            unit: units.RotationalAccelerationRevolutionPerSecondSquared,
            want: "r/s²",
        },
        {
            name: "invalid unit",
            unit: units.RotationalAccelerationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetRotationalAccelerationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetRotationalAccelerationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestRotationalAcceleration_String(t *testing.T) {
    factory := units.RotationalAccelerationFactory{}
    
    tests := []struct {
        name  string
        value float64
        want  string
    }{
        {
            name:  "positive integer",
            value: 100,
            want:  "100.00",
        },
        {
            name:  "negative integer",
            value: -100,
            want:  "-100.00",
        },
        {
            name:  "zero",
            value: 0,
            want:  "0.00",
        },
        {
            name:  "positive decimal",
            value: 123.456,
            want:  "123.46",
        },
        {
            name:  "negative decimal",
            value: -123.456,
            want:  "-123.46",
        },
        {
            name:  "small decimal",
            value: 0.123,
            want:  "0.12",
        },
        {
            name:  "large number",
            value: 1000000,
            want:  "1000000.00",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            unit, err := factory.CreateRotationalAcceleration(tt.value, units.RotationalAccelerationRadianPerSecondSquared)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("RotationalAcceleration.String() = %v, want %v", got, tt.want)
            }
        })
    }
}