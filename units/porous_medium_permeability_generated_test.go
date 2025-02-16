// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPorousMediumPermeabilityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeter"}`
	
	factory := units.PorousMediumPermeabilityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PorousMediumPermeabilitySquareMeter {
		t.Errorf("expected unit %v, got %v", units.PorousMediumPermeabilitySquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPorousMediumPermeabilityDto_ToJSON(t *testing.T) {
	dto := units.PorousMediumPermeabilityDto{
		Value: 45,
		Unit:  units.PorousMediumPermeabilitySquareMeter,
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
	if result["unit"].(string) != string(units.PorousMediumPermeabilitySquareMeter) {
		t.Errorf("expected unit %s, got %v", units.PorousMediumPermeabilitySquareMeter, result["unit"])
	}
}

func TestNewPorousMediumPermeability_InvalidValue(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePorousMediumPermeability(math.NaN(), units.PorousMediumPermeabilitySquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePorousMediumPermeability(math.Inf(1), units.PorousMediumPermeabilitySquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPorousMediumPermeabilityConversions(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePorousMediumPermeability(180, units.PorousMediumPermeabilitySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Darcys.
		// No expected conversion value provided for Darcys, verifying result is not NaN.
		result := a.Darcys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Darcys returned NaN")
		}
	}
	{
		// Test conversion to SquareMeters.
		// No expected conversion value provided for SquareMeters, verifying result is not NaN.
		result := a.SquareMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeters returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeters.
		// No expected conversion value provided for SquareCentimeters, verifying result is not NaN.
		result := a.SquareCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to Microdarcys.
		// No expected conversion value provided for Microdarcys, verifying result is not NaN.
		result := a.Microdarcys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microdarcys returned NaN")
		}
	}
	{
		// Test conversion to Millidarcys.
		// No expected conversion value provided for Millidarcys, verifying result is not NaN.
		result := a.Millidarcys()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millidarcys returned NaN")
		}
	}
}

func TestPorousMediumPermeability_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	a, err := factory.CreatePorousMediumPermeability(90, units.PorousMediumPermeabilitySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PorousMediumPermeabilitySquareMeter {
		t.Errorf("expected default unit SquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PorousMediumPermeabilityDarcy
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PorousMediumPermeabilityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PorousMediumPermeabilitySquareMeter {
		t.Errorf("expected unit SquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPorousMediumPermeabilityFactory_FromDto(t *testing.T) {
    factory := units.PorousMediumPermeabilityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.PorousMediumPermeabilityDto{
        Value: 100,
        Unit:  units.PorousMediumPermeabilitySquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.PorousMediumPermeabilityDto{
        Value: math.NaN(),
        Unit:  units.PorousMediumPermeabilitySquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Darcy conversion
    darcysDto := units.PorousMediumPermeabilityDto{
        Value: 100,
        Unit:  units.PorousMediumPermeabilityDarcy,
    }
    
    var darcysResult *units.PorousMediumPermeability
    darcysResult, err = factory.FromDto(darcysDto)
    if err != nil {
        t.Errorf("FromDto() with Darcy returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = darcysResult.Convert(units.PorousMediumPermeabilityDarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Darcy = %v, want %v", converted, 100)
    }
    // Test SquareMeter conversion
    square_metersDto := units.PorousMediumPermeabilityDto{
        Value: 100,
        Unit:  units.PorousMediumPermeabilitySquareMeter,
    }
    
    var square_metersResult *units.PorousMediumPermeability
    square_metersResult, err = factory.FromDto(square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_metersResult.Convert(units.PorousMediumPermeabilitySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeter = %v, want %v", converted, 100)
    }
    // Test SquareCentimeter conversion
    square_centimetersDto := units.PorousMediumPermeabilityDto{
        Value: 100,
        Unit:  units.PorousMediumPermeabilitySquareCentimeter,
    }
    
    var square_centimetersResult *units.PorousMediumPermeability
    square_centimetersResult, err = factory.FromDto(square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with SquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimetersResult.Convert(units.PorousMediumPermeabilitySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeter = %v, want %v", converted, 100)
    }
    // Test Microdarcy conversion
    microdarcysDto := units.PorousMediumPermeabilityDto{
        Value: 100,
        Unit:  units.PorousMediumPermeabilityMicrodarcy,
    }
    
    var microdarcysResult *units.PorousMediumPermeability
    microdarcysResult, err = factory.FromDto(microdarcysDto)
    if err != nil {
        t.Errorf("FromDto() with Microdarcy returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microdarcysResult.Convert(units.PorousMediumPermeabilityMicrodarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microdarcy = %v, want %v", converted, 100)
    }
    // Test Millidarcy conversion
    millidarcysDto := units.PorousMediumPermeabilityDto{
        Value: 100,
        Unit:  units.PorousMediumPermeabilityMillidarcy,
    }
    
    var millidarcysResult *units.PorousMediumPermeability
    millidarcysResult, err = factory.FromDto(millidarcysDto)
    if err != nil {
        t.Errorf("FromDto() with Millidarcy returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidarcysResult.Convert(units.PorousMediumPermeabilityMillidarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millidarcy = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.PorousMediumPermeabilityDto{
        Value: 0,
        Unit:  units.PorousMediumPermeabilitySquareMeter,
    }
    
    var zeroResult *units.PorousMediumPermeability
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestPorousMediumPermeabilityFactory_FromDtoJSON(t *testing.T) {
    factory := units.PorousMediumPermeabilityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "SquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "SquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.PorousMediumPermeabilityDto{
        Value: nanValue,
        Unit:  units.PorousMediumPermeabilitySquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Darcy unit
    darcysJSON := []byte(`{"value": 100, "unit": "Darcy"}`)
    darcysResult, err := factory.FromDtoJSON(darcysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Darcy unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = darcysResult.Convert(units.PorousMediumPermeabilityDarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Darcy = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMeter unit
    square_metersJSON := []byte(`{"value": 100, "unit": "SquareMeter"}`)
    square_metersResult, err := factory.FromDtoJSON(square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_metersResult.Convert(units.PorousMediumPermeabilitySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with SquareCentimeter unit
    square_centimetersJSON := []byte(`{"value": 100, "unit": "SquareCentimeter"}`)
    square_centimetersResult, err := factory.FromDtoJSON(square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimetersResult.Convert(units.PorousMediumPermeabilitySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with Microdarcy unit
    microdarcysJSON := []byte(`{"value": 100, "unit": "Microdarcy"}`)
    microdarcysResult, err := factory.FromDtoJSON(microdarcysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microdarcy unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microdarcysResult.Convert(units.PorousMediumPermeabilityMicrodarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microdarcy = %v, want %v", converted, 100)
    }
    // Test JSON with Millidarcy unit
    millidarcysJSON := []byte(`{"value": 100, "unit": "Millidarcy"}`)
    millidarcysResult, err := factory.FromDtoJSON(millidarcysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millidarcy unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidarcysResult.Convert(units.PorousMediumPermeabilityMillidarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millidarcy = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "SquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDarcys function
func TestPorousMediumPermeabilityFactory_FromDarcys(t *testing.T) {
    factory := units.PorousMediumPermeabilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDarcys(100)
    if err != nil {
        t.Errorf("FromDarcys() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PorousMediumPermeabilityDarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDarcys() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDarcys(math.NaN())
    if err == nil {
        t.Error("FromDarcys() with NaN value should return error")
    }

    _, err = factory.FromDarcys(math.Inf(1))
    if err == nil {
        t.Error("FromDarcys() with +Inf value should return error")
    }

    _, err = factory.FromDarcys(math.Inf(-1))
    if err == nil {
        t.Error("FromDarcys() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDarcys(0)
    if err != nil {
        t.Errorf("FromDarcys() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PorousMediumPermeabilityDarcy)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDarcys() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMeters function
func TestPorousMediumPermeabilityFactory_FromSquareMeters(t *testing.T) {
    factory := units.PorousMediumPermeabilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMeters(100)
    if err != nil {
        t.Errorf("FromSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PorousMediumPermeabilitySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMeters(0)
    if err != nil {
        t.Errorf("FromSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PorousMediumPermeabilitySquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareCentimeters function
func TestPorousMediumPermeabilityFactory_FromSquareCentimeters(t *testing.T) {
    factory := units.PorousMediumPermeabilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PorousMediumPermeabilitySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PorousMediumPermeabilitySquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrodarcys function
func TestPorousMediumPermeabilityFactory_FromMicrodarcys(t *testing.T) {
    factory := units.PorousMediumPermeabilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrodarcys(100)
    if err != nil {
        t.Errorf("FromMicrodarcys() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PorousMediumPermeabilityMicrodarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrodarcys() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrodarcys(math.NaN())
    if err == nil {
        t.Error("FromMicrodarcys() with NaN value should return error")
    }

    _, err = factory.FromMicrodarcys(math.Inf(1))
    if err == nil {
        t.Error("FromMicrodarcys() with +Inf value should return error")
    }

    _, err = factory.FromMicrodarcys(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrodarcys() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrodarcys(0)
    if err != nil {
        t.Errorf("FromMicrodarcys() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PorousMediumPermeabilityMicrodarcy)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrodarcys() with zero value = %v, want 0", converted)
    }
}
// Test FromMillidarcys function
func TestPorousMediumPermeabilityFactory_FromMillidarcys(t *testing.T) {
    factory := units.PorousMediumPermeabilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillidarcys(100)
    if err != nil {
        t.Errorf("FromMillidarcys() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PorousMediumPermeabilityMillidarcy)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillidarcys() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillidarcys(math.NaN())
    if err == nil {
        t.Error("FromMillidarcys() with NaN value should return error")
    }

    _, err = factory.FromMillidarcys(math.Inf(1))
    if err == nil {
        t.Error("FromMillidarcys() with +Inf value should return error")
    }

    _, err = factory.FromMillidarcys(math.Inf(-1))
    if err == nil {
        t.Error("FromMillidarcys() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillidarcys(0)
    if err != nil {
        t.Errorf("FromMillidarcys() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PorousMediumPermeabilityMillidarcy)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillidarcys() with zero value = %v, want 0", converted)
    }
}

func TestPorousMediumPermeabilityToString(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	a, err := factory.CreatePorousMediumPermeability(45, units.PorousMediumPermeabilitySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PorousMediumPermeabilitySquareMeter, 2)
	expected := "45.00 " + units.GetPorousMediumPermeabilityAbbreviation(units.PorousMediumPermeabilitySquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PorousMediumPermeabilitySquareMeter, -1)
	expected = "45 " + units.GetPorousMediumPermeabilityAbbreviation(units.PorousMediumPermeabilitySquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPorousMediumPermeability_EqualityAndComparison(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	a1, _ := factory.CreatePorousMediumPermeability(60, units.PorousMediumPermeabilitySquareMeter)
	a2, _ := factory.CreatePorousMediumPermeability(60, units.PorousMediumPermeabilitySquareMeter)
	a3, _ := factory.CreatePorousMediumPermeability(120, units.PorousMediumPermeabilitySquareMeter)

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

func TestPorousMediumPermeability_Arithmetic(t *testing.T) {
	factory := units.PorousMediumPermeabilityFactory{}
	a1, _ := factory.CreatePorousMediumPermeability(30, units.PorousMediumPermeabilitySquareMeter)
	a2, _ := factory.CreatePorousMediumPermeability(45, units.PorousMediumPermeabilitySquareMeter)

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