// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRadiationEquivalentDoseDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Sievert"}`
	
	factory := units.RadiationEquivalentDoseDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RadiationEquivalentDoseSievert {
		t.Errorf("expected unit %v, got %v", units.RadiationEquivalentDoseSievert, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Sievert"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRadiationEquivalentDoseDto_ToJSON(t *testing.T) {
	dto := units.RadiationEquivalentDoseDto{
		Value: 45,
		Unit:  units.RadiationEquivalentDoseSievert,
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
	if result["unit"].(string) != string(units.RadiationEquivalentDoseSievert) {
		t.Errorf("expected unit %s, got %v", units.RadiationEquivalentDoseSievert, result["unit"])
	}
}

func TestNewRadiationEquivalentDose_InvalidValue(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRadiationEquivalentDose(math.NaN(), units.RadiationEquivalentDoseSievert)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRadiationEquivalentDose(math.Inf(1), units.RadiationEquivalentDoseSievert)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRadiationEquivalentDoseConversions(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRadiationEquivalentDose(180, units.RadiationEquivalentDoseSievert)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Sieverts.
		// No expected conversion value provided for Sieverts, verifying result is not NaN.
		result := a.Sieverts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Sieverts returned NaN")
		}
	}
	{
		// Test conversion to RoentgensEquivalentMan.
		// No expected conversion value provided for RoentgensEquivalentMan, verifying result is not NaN.
		result := a.RoentgensEquivalentMan()
		if math.IsNaN(result) {
			t.Errorf("conversion to RoentgensEquivalentMan returned NaN")
		}
	}
	{
		// Test conversion to Nanosieverts.
		// No expected conversion value provided for Nanosieverts, verifying result is not NaN.
		result := a.Nanosieverts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanosieverts returned NaN")
		}
	}
	{
		// Test conversion to Microsieverts.
		// No expected conversion value provided for Microsieverts, verifying result is not NaN.
		result := a.Microsieverts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microsieverts returned NaN")
		}
	}
	{
		// Test conversion to Millisieverts.
		// No expected conversion value provided for Millisieverts, verifying result is not NaN.
		result := a.Millisieverts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millisieverts returned NaN")
		}
	}
	{
		// Test conversion to MilliroentgensEquivalentMan.
		// No expected conversion value provided for MilliroentgensEquivalentMan, verifying result is not NaN.
		result := a.MilliroentgensEquivalentMan()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliroentgensEquivalentMan returned NaN")
		}
	}
}

func TestRadiationEquivalentDose_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	a, err := factory.CreateRadiationEquivalentDose(90, units.RadiationEquivalentDoseSievert)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RadiationEquivalentDoseSievert {
		t.Errorf("expected default unit Sievert, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RadiationEquivalentDoseSievert
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RadiationEquivalentDoseDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RadiationEquivalentDoseSievert {
		t.Errorf("expected unit Sievert, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRadiationEquivalentDoseFactory_FromDto(t *testing.T) {
    factory := units.RadiationEquivalentDoseFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RadiationEquivalentDoseDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseSievert,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RadiationEquivalentDoseDto{
        Value: math.NaN(),
        Unit:  units.RadiationEquivalentDoseSievert,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Sievert conversion
    sievertsDto := units.RadiationEquivalentDoseDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseSievert,
    }
    
    var sievertsResult *units.RadiationEquivalentDose
    sievertsResult, err = factory.FromDto(sievertsDto)
    if err != nil {
        t.Errorf("FromDto() with Sievert returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = sievertsResult.Convert(units.RadiationEquivalentDoseSievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Sievert = %v, want %v", converted, 100)
    }
    // Test RoentgenEquivalentMan conversion
    roentgens_equivalent_manDto := units.RadiationEquivalentDoseDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRoentgenEquivalentMan,
    }
    
    var roentgens_equivalent_manResult *units.RadiationEquivalentDose
    roentgens_equivalent_manResult, err = factory.FromDto(roentgens_equivalent_manDto)
    if err != nil {
        t.Errorf("FromDto() with RoentgenEquivalentMan returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = roentgens_equivalent_manResult.Convert(units.RadiationEquivalentDoseRoentgenEquivalentMan)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RoentgenEquivalentMan = %v, want %v", converted, 100)
    }
    // Test Nanosievert conversion
    nanosievertsDto := units.RadiationEquivalentDoseDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseNanosievert,
    }
    
    var nanosievertsResult *units.RadiationEquivalentDose
    nanosievertsResult, err = factory.FromDto(nanosievertsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanosievert returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosievertsResult.Convert(units.RadiationEquivalentDoseNanosievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanosievert = %v, want %v", converted, 100)
    }
    // Test Microsievert conversion
    microsievertsDto := units.RadiationEquivalentDoseDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseMicrosievert,
    }
    
    var microsievertsResult *units.RadiationEquivalentDose
    microsievertsResult, err = factory.FromDto(microsievertsDto)
    if err != nil {
        t.Errorf("FromDto() with Microsievert returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsievertsResult.Convert(units.RadiationEquivalentDoseMicrosievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microsievert = %v, want %v", converted, 100)
    }
    // Test Millisievert conversion
    millisievertsDto := units.RadiationEquivalentDoseDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseMillisievert,
    }
    
    var millisievertsResult *units.RadiationEquivalentDose
    millisievertsResult, err = factory.FromDto(millisievertsDto)
    if err != nil {
        t.Errorf("FromDto() with Millisievert returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisievertsResult.Convert(units.RadiationEquivalentDoseMillisievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millisievert = %v, want %v", converted, 100)
    }
    // Test MilliroentgenEquivalentMan conversion
    milliroentgens_equivalent_manDto := units.RadiationEquivalentDoseDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseMilliroentgenEquivalentMan,
    }
    
    var milliroentgens_equivalent_manResult *units.RadiationEquivalentDose
    milliroentgens_equivalent_manResult, err = factory.FromDto(milliroentgens_equivalent_manDto)
    if err != nil {
        t.Errorf("FromDto() with MilliroentgenEquivalentMan returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliroentgens_equivalent_manResult.Convert(units.RadiationEquivalentDoseMilliroentgenEquivalentMan)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliroentgenEquivalentMan = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RadiationEquivalentDoseDto{
        Value: 0,
        Unit:  units.RadiationEquivalentDoseSievert,
    }
    
    var zeroResult *units.RadiationEquivalentDose
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRadiationEquivalentDoseFactory_FromDtoJSON(t *testing.T) {
    factory := units.RadiationEquivalentDoseFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Sievert"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Sievert"}`)
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
    nanJSON, _ := json.Marshal(units.RadiationEquivalentDoseDto{
        Value: nanValue,
        Unit:  units.RadiationEquivalentDoseSievert,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Sievert unit
    sievertsJSON := []byte(`{"value": 100, "unit": "Sievert"}`)
    sievertsResult, err := factory.FromDtoJSON(sievertsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Sievert unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = sievertsResult.Convert(units.RadiationEquivalentDoseSievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Sievert = %v, want %v", converted, 100)
    }
    // Test JSON with RoentgenEquivalentMan unit
    roentgens_equivalent_manJSON := []byte(`{"value": 100, "unit": "RoentgenEquivalentMan"}`)
    roentgens_equivalent_manResult, err := factory.FromDtoJSON(roentgens_equivalent_manJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RoentgenEquivalentMan unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = roentgens_equivalent_manResult.Convert(units.RadiationEquivalentDoseRoentgenEquivalentMan)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RoentgenEquivalentMan = %v, want %v", converted, 100)
    }
    // Test JSON with Nanosievert unit
    nanosievertsJSON := []byte(`{"value": 100, "unit": "Nanosievert"}`)
    nanosievertsResult, err := factory.FromDtoJSON(nanosievertsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanosievert unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosievertsResult.Convert(units.RadiationEquivalentDoseNanosievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanosievert = %v, want %v", converted, 100)
    }
    // Test JSON with Microsievert unit
    microsievertsJSON := []byte(`{"value": 100, "unit": "Microsievert"}`)
    microsievertsResult, err := factory.FromDtoJSON(microsievertsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microsievert unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsievertsResult.Convert(units.RadiationEquivalentDoseMicrosievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microsievert = %v, want %v", converted, 100)
    }
    // Test JSON with Millisievert unit
    millisievertsJSON := []byte(`{"value": 100, "unit": "Millisievert"}`)
    millisievertsResult, err := factory.FromDtoJSON(millisievertsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millisievert unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisievertsResult.Convert(units.RadiationEquivalentDoseMillisievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millisievert = %v, want %v", converted, 100)
    }
    // Test JSON with MilliroentgenEquivalentMan unit
    milliroentgens_equivalent_manJSON := []byte(`{"value": 100, "unit": "MilliroentgenEquivalentMan"}`)
    milliroentgens_equivalent_manResult, err := factory.FromDtoJSON(milliroentgens_equivalent_manJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliroentgenEquivalentMan unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliroentgens_equivalent_manResult.Convert(units.RadiationEquivalentDoseMilliroentgenEquivalentMan)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliroentgenEquivalentMan = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Sievert"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSieverts function
func TestRadiationEquivalentDoseFactory_FromSieverts(t *testing.T) {
    factory := units.RadiationEquivalentDoseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSieverts(100)
    if err != nil {
        t.Errorf("FromSieverts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseSievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSieverts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSieverts(math.NaN())
    if err == nil {
        t.Error("FromSieverts() with NaN value should return error")
    }

    _, err = factory.FromSieverts(math.Inf(1))
    if err == nil {
        t.Error("FromSieverts() with +Inf value should return error")
    }

    _, err = factory.FromSieverts(math.Inf(-1))
    if err == nil {
        t.Error("FromSieverts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSieverts(0)
    if err != nil {
        t.Errorf("FromSieverts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseSievert)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSieverts() with zero value = %v, want 0", converted)
    }
}
// Test FromRoentgensEquivalentMan function
func TestRadiationEquivalentDoseFactory_FromRoentgensEquivalentMan(t *testing.T) {
    factory := units.RadiationEquivalentDoseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRoentgensEquivalentMan(100)
    if err != nil {
        t.Errorf("FromRoentgensEquivalentMan() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRoentgenEquivalentMan)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRoentgensEquivalentMan() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRoentgensEquivalentMan(math.NaN())
    if err == nil {
        t.Error("FromRoentgensEquivalentMan() with NaN value should return error")
    }

    _, err = factory.FromRoentgensEquivalentMan(math.Inf(1))
    if err == nil {
        t.Error("FromRoentgensEquivalentMan() with +Inf value should return error")
    }

    _, err = factory.FromRoentgensEquivalentMan(math.Inf(-1))
    if err == nil {
        t.Error("FromRoentgensEquivalentMan() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRoentgensEquivalentMan(0)
    if err != nil {
        t.Errorf("FromRoentgensEquivalentMan() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRoentgenEquivalentMan)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRoentgensEquivalentMan() with zero value = %v, want 0", converted)
    }
}
// Test FromNanosieverts function
func TestRadiationEquivalentDoseFactory_FromNanosieverts(t *testing.T) {
    factory := units.RadiationEquivalentDoseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanosieverts(100)
    if err != nil {
        t.Errorf("FromNanosieverts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseNanosievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanosieverts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanosieverts(math.NaN())
    if err == nil {
        t.Error("FromNanosieverts() with NaN value should return error")
    }

    _, err = factory.FromNanosieverts(math.Inf(1))
    if err == nil {
        t.Error("FromNanosieverts() with +Inf value should return error")
    }

    _, err = factory.FromNanosieverts(math.Inf(-1))
    if err == nil {
        t.Error("FromNanosieverts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanosieverts(0)
    if err != nil {
        t.Errorf("FromNanosieverts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseNanosievert)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanosieverts() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrosieverts function
func TestRadiationEquivalentDoseFactory_FromMicrosieverts(t *testing.T) {
    factory := units.RadiationEquivalentDoseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrosieverts(100)
    if err != nil {
        t.Errorf("FromMicrosieverts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseMicrosievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrosieverts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrosieverts(math.NaN())
    if err == nil {
        t.Error("FromMicrosieverts() with NaN value should return error")
    }

    _, err = factory.FromMicrosieverts(math.Inf(1))
    if err == nil {
        t.Error("FromMicrosieverts() with +Inf value should return error")
    }

    _, err = factory.FromMicrosieverts(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrosieverts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrosieverts(0)
    if err != nil {
        t.Errorf("FromMicrosieverts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseMicrosievert)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrosieverts() with zero value = %v, want 0", converted)
    }
}
// Test FromMillisieverts function
func TestRadiationEquivalentDoseFactory_FromMillisieverts(t *testing.T) {
    factory := units.RadiationEquivalentDoseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillisieverts(100)
    if err != nil {
        t.Errorf("FromMillisieverts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseMillisievert)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillisieverts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillisieverts(math.NaN())
    if err == nil {
        t.Error("FromMillisieverts() with NaN value should return error")
    }

    _, err = factory.FromMillisieverts(math.Inf(1))
    if err == nil {
        t.Error("FromMillisieverts() with +Inf value should return error")
    }

    _, err = factory.FromMillisieverts(math.Inf(-1))
    if err == nil {
        t.Error("FromMillisieverts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillisieverts(0)
    if err != nil {
        t.Errorf("FromMillisieverts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseMillisievert)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillisieverts() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliroentgensEquivalentMan function
func TestRadiationEquivalentDoseFactory_FromMilliroentgensEquivalentMan(t *testing.T) {
    factory := units.RadiationEquivalentDoseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliroentgensEquivalentMan(100)
    if err != nil {
        t.Errorf("FromMilliroentgensEquivalentMan() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseMilliroentgenEquivalentMan)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliroentgensEquivalentMan() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliroentgensEquivalentMan(math.NaN())
    if err == nil {
        t.Error("FromMilliroentgensEquivalentMan() with NaN value should return error")
    }

    _, err = factory.FromMilliroentgensEquivalentMan(math.Inf(1))
    if err == nil {
        t.Error("FromMilliroentgensEquivalentMan() with +Inf value should return error")
    }

    _, err = factory.FromMilliroentgensEquivalentMan(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliroentgensEquivalentMan() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliroentgensEquivalentMan(0)
    if err != nil {
        t.Errorf("FromMilliroentgensEquivalentMan() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseMilliroentgenEquivalentMan)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliroentgensEquivalentMan() with zero value = %v, want 0", converted)
    }
}

func TestRadiationEquivalentDoseToString(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	a, err := factory.CreateRadiationEquivalentDose(45, units.RadiationEquivalentDoseSievert)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RadiationEquivalentDoseSievert, 2)
	expected := "45.00 " + units.GetRadiationEquivalentDoseAbbreviation(units.RadiationEquivalentDoseSievert)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RadiationEquivalentDoseSievert, -1)
	expected = "45 " + units.GetRadiationEquivalentDoseAbbreviation(units.RadiationEquivalentDoseSievert)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRadiationEquivalentDose_EqualityAndComparison(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	a1, _ := factory.CreateRadiationEquivalentDose(60, units.RadiationEquivalentDoseSievert)
	a2, _ := factory.CreateRadiationEquivalentDose(60, units.RadiationEquivalentDoseSievert)
	a3, _ := factory.CreateRadiationEquivalentDose(120, units.RadiationEquivalentDoseSievert)

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

func TestRadiationEquivalentDose_Arithmetic(t *testing.T) {
	factory := units.RadiationEquivalentDoseFactory{}
	a1, _ := factory.CreateRadiationEquivalentDose(30, units.RadiationEquivalentDoseSievert)
	a2, _ := factory.CreateRadiationEquivalentDose(45, units.RadiationEquivalentDoseSievert)

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