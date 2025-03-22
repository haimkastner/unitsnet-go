// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricInductanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Henry"}`
	
	factory := units.ElectricInductanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricInductanceHenry {
		t.Errorf("expected unit %v, got %v", units.ElectricInductanceHenry, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Henry"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricInductanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricInductanceDto{
		Value: 45,
		Unit:  units.ElectricInductanceHenry,
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
	if result["unit"].(string) != string(units.ElectricInductanceHenry) {
		t.Errorf("expected unit %s, got %v", units.ElectricInductanceHenry, result["unit"])
	}
}

func TestNewElectricInductance_InvalidValue(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricInductance(math.NaN(), units.ElectricInductanceHenry)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricInductance(math.Inf(1), units.ElectricInductanceHenry)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricInductanceConversions(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricInductance(180, units.ElectricInductanceHenry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Henries.
		// No expected conversion value provided for Henries, verifying result is not NaN.
		result := a.Henries()
		cacheResult := a.Henries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Henries returned NaN")
		}
	}
	{
		// Test conversion to Picohenries.
		// No expected conversion value provided for Picohenries, verifying result is not NaN.
		result := a.Picohenries()
		cacheResult := a.Picohenries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picohenries returned NaN")
		}
	}
	{
		// Test conversion to Nanohenries.
		// No expected conversion value provided for Nanohenries, verifying result is not NaN.
		result := a.Nanohenries()
		cacheResult := a.Nanohenries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanohenries returned NaN")
		}
	}
	{
		// Test conversion to Microhenries.
		// No expected conversion value provided for Microhenries, verifying result is not NaN.
		result := a.Microhenries()
		cacheResult := a.Microhenries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microhenries returned NaN")
		}
	}
	{
		// Test conversion to Millihenries.
		// No expected conversion value provided for Millihenries, verifying result is not NaN.
		result := a.Millihenries()
		cacheResult := a.Millihenries()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millihenries returned NaN")
		}
	}
}

func TestElectricInductance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	a, err := factory.CreateElectricInductance(90, units.ElectricInductanceHenry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricInductanceHenry {
		t.Errorf("expected default unit Henry, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricInductanceHenry
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricInductanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricInductanceHenry {
		t.Errorf("expected unit Henry, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricInductanceFactory_FromDto(t *testing.T) {
    factory := units.ElectricInductanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricInductanceDto{
        Value: 100,
        Unit:  units.ElectricInductanceHenry,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricInductanceDto{
        Value: math.NaN(),
        Unit:  units.ElectricInductanceHenry,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Henry conversion
    henriesDto := units.ElectricInductanceDto{
        Value: 100,
        Unit:  units.ElectricInductanceHenry,
    }
    
    var henriesResult *units.ElectricInductance
    henriesResult, err = factory.FromDto(henriesDto)
    if err != nil {
        t.Errorf("FromDto() with Henry returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = henriesResult.Convert(units.ElectricInductanceHenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Henry = %v, want %v", converted, 100)
    }
    // Test Picohenry conversion
    picohenriesDto := units.ElectricInductanceDto{
        Value: 100,
        Unit:  units.ElectricInductancePicohenry,
    }
    
    var picohenriesResult *units.ElectricInductance
    picohenriesResult, err = factory.FromDto(picohenriesDto)
    if err != nil {
        t.Errorf("FromDto() with Picohenry returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picohenriesResult.Convert(units.ElectricInductancePicohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picohenry = %v, want %v", converted, 100)
    }
    // Test Nanohenry conversion
    nanohenriesDto := units.ElectricInductanceDto{
        Value: 100,
        Unit:  units.ElectricInductanceNanohenry,
    }
    
    var nanohenriesResult *units.ElectricInductance
    nanohenriesResult, err = factory.FromDto(nanohenriesDto)
    if err != nil {
        t.Errorf("FromDto() with Nanohenry returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanohenriesResult.Convert(units.ElectricInductanceNanohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanohenry = %v, want %v", converted, 100)
    }
    // Test Microhenry conversion
    microhenriesDto := units.ElectricInductanceDto{
        Value: 100,
        Unit:  units.ElectricInductanceMicrohenry,
    }
    
    var microhenriesResult *units.ElectricInductance
    microhenriesResult, err = factory.FromDto(microhenriesDto)
    if err != nil {
        t.Errorf("FromDto() with Microhenry returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microhenriesResult.Convert(units.ElectricInductanceMicrohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microhenry = %v, want %v", converted, 100)
    }
    // Test Millihenry conversion
    millihenriesDto := units.ElectricInductanceDto{
        Value: 100,
        Unit:  units.ElectricInductanceMillihenry,
    }
    
    var millihenriesResult *units.ElectricInductance
    millihenriesResult, err = factory.FromDto(millihenriesDto)
    if err != nil {
        t.Errorf("FromDto() with Millihenry returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millihenriesResult.Convert(units.ElectricInductanceMillihenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millihenry = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricInductanceDto{
        Value: 0,
        Unit:  units.ElectricInductanceHenry,
    }
    
    var zeroResult *units.ElectricInductance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricInductanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricInductanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Henry"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Henry"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricInductanceDto{
        Value: nanValue,
        Unit:  units.ElectricInductanceHenry,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Henry unit
    henriesJSON := []byte(`{"value": 100, "unit": "Henry"}`)
    henriesResult, err := factory.FromDtoJSON(henriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Henry unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = henriesResult.Convert(units.ElectricInductanceHenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Henry = %v, want %v", converted, 100)
    }
    // Test JSON with Picohenry unit
    picohenriesJSON := []byte(`{"value": 100, "unit": "Picohenry"}`)
    picohenriesResult, err := factory.FromDtoJSON(picohenriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picohenry unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picohenriesResult.Convert(units.ElectricInductancePicohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picohenry = %v, want %v", converted, 100)
    }
    // Test JSON with Nanohenry unit
    nanohenriesJSON := []byte(`{"value": 100, "unit": "Nanohenry"}`)
    nanohenriesResult, err := factory.FromDtoJSON(nanohenriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanohenry unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanohenriesResult.Convert(units.ElectricInductanceNanohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanohenry = %v, want %v", converted, 100)
    }
    // Test JSON with Microhenry unit
    microhenriesJSON := []byte(`{"value": 100, "unit": "Microhenry"}`)
    microhenriesResult, err := factory.FromDtoJSON(microhenriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microhenry unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microhenriesResult.Convert(units.ElectricInductanceMicrohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microhenry = %v, want %v", converted, 100)
    }
    // Test JSON with Millihenry unit
    millihenriesJSON := []byte(`{"value": 100, "unit": "Millihenry"}`)
    millihenriesResult, err := factory.FromDtoJSON(millihenriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millihenry unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millihenriesResult.Convert(units.ElectricInductanceMillihenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millihenry = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Henry"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromHenries function
func TestElectricInductanceFactory_FromHenries(t *testing.T) {
    factory := units.ElectricInductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHenries(100)
    if err != nil {
        t.Errorf("FromHenries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricInductanceHenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHenries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHenries(math.NaN())
    if err == nil {
        t.Error("FromHenries() with NaN value should return error")
    }

    _, err = factory.FromHenries(math.Inf(1))
    if err == nil {
        t.Error("FromHenries() with +Inf value should return error")
    }

    _, err = factory.FromHenries(math.Inf(-1))
    if err == nil {
        t.Error("FromHenries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHenries(0)
    if err != nil {
        t.Errorf("FromHenries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricInductanceHenry)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHenries() with zero value = %v, want 0", converted)
    }
}
// Test FromPicohenries function
func TestElectricInductanceFactory_FromPicohenries(t *testing.T) {
    factory := units.ElectricInductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicohenries(100)
    if err != nil {
        t.Errorf("FromPicohenries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricInductancePicohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicohenries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicohenries(math.NaN())
    if err == nil {
        t.Error("FromPicohenries() with NaN value should return error")
    }

    _, err = factory.FromPicohenries(math.Inf(1))
    if err == nil {
        t.Error("FromPicohenries() with +Inf value should return error")
    }

    _, err = factory.FromPicohenries(math.Inf(-1))
    if err == nil {
        t.Error("FromPicohenries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicohenries(0)
    if err != nil {
        t.Errorf("FromPicohenries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricInductancePicohenry)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicohenries() with zero value = %v, want 0", converted)
    }
}
// Test FromNanohenries function
func TestElectricInductanceFactory_FromNanohenries(t *testing.T) {
    factory := units.ElectricInductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanohenries(100)
    if err != nil {
        t.Errorf("FromNanohenries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricInductanceNanohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanohenries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanohenries(math.NaN())
    if err == nil {
        t.Error("FromNanohenries() with NaN value should return error")
    }

    _, err = factory.FromNanohenries(math.Inf(1))
    if err == nil {
        t.Error("FromNanohenries() with +Inf value should return error")
    }

    _, err = factory.FromNanohenries(math.Inf(-1))
    if err == nil {
        t.Error("FromNanohenries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanohenries(0)
    if err != nil {
        t.Errorf("FromNanohenries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricInductanceNanohenry)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanohenries() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrohenries function
func TestElectricInductanceFactory_FromMicrohenries(t *testing.T) {
    factory := units.ElectricInductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrohenries(100)
    if err != nil {
        t.Errorf("FromMicrohenries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricInductanceMicrohenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrohenries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrohenries(math.NaN())
    if err == nil {
        t.Error("FromMicrohenries() with NaN value should return error")
    }

    _, err = factory.FromMicrohenries(math.Inf(1))
    if err == nil {
        t.Error("FromMicrohenries() with +Inf value should return error")
    }

    _, err = factory.FromMicrohenries(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrohenries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrohenries(0)
    if err != nil {
        t.Errorf("FromMicrohenries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricInductanceMicrohenry)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrohenries() with zero value = %v, want 0", converted)
    }
}
// Test FromMillihenries function
func TestElectricInductanceFactory_FromMillihenries(t *testing.T) {
    factory := units.ElectricInductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillihenries(100)
    if err != nil {
        t.Errorf("FromMillihenries() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricInductanceMillihenry)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillihenries() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillihenries(math.NaN())
    if err == nil {
        t.Error("FromMillihenries() with NaN value should return error")
    }

    _, err = factory.FromMillihenries(math.Inf(1))
    if err == nil {
        t.Error("FromMillihenries() with +Inf value should return error")
    }

    _, err = factory.FromMillihenries(math.Inf(-1))
    if err == nil {
        t.Error("FromMillihenries() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillihenries(0)
    if err != nil {
        t.Errorf("FromMillihenries() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricInductanceMillihenry)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillihenries() with zero value = %v, want 0", converted)
    }
}

func TestElectricInductanceToString(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	a, err := factory.CreateElectricInductance(45, units.ElectricInductanceHenry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricInductanceHenry, 2)
	expected := "45.00 " + units.GetElectricInductanceAbbreviation(units.ElectricInductanceHenry)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricInductanceHenry, -1)
	expected = "45 " + units.GetElectricInductanceAbbreviation(units.ElectricInductanceHenry)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricInductance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	a1, _ := factory.CreateElectricInductance(60, units.ElectricInductanceHenry)
	a2, _ := factory.CreateElectricInductance(60, units.ElectricInductanceHenry)
	a3, _ := factory.CreateElectricInductance(120, units.ElectricInductanceHenry)

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

func TestElectricInductance_Arithmetic(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	a1, _ := factory.CreateElectricInductance(30, units.ElectricInductanceHenry)
	a2, _ := factory.CreateElectricInductance(45, units.ElectricInductanceHenry)

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


func TestGetElectricInductanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricInductanceUnits
        want string
    }{
        {
            name: "Henry abbreviation",
            unit: units.ElectricInductanceHenry,
            want: "H",
        },
        {
            name: "Picohenry abbreviation",
            unit: units.ElectricInductancePicohenry,
            want: "pH",
        },
        {
            name: "Nanohenry abbreviation",
            unit: units.ElectricInductanceNanohenry,
            want: "nH",
        },
        {
            name: "Microhenry abbreviation",
            unit: units.ElectricInductanceMicrohenry,
            want: "μH",
        },
        {
            name: "Millihenry abbreviation",
            unit: units.ElectricInductanceMillihenry,
            want: "mH",
        },
        {
            name: "invalid unit",
            unit: units.ElectricInductanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricInductanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricInductanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricInductance_String(t *testing.T) {
    factory := units.ElectricInductanceFactory{}
    
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
            unit, err := factory.CreateElectricInductance(tt.value, units.ElectricInductanceHenry)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricInductance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestElectricInductance_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ElectricInductanceFactory{}

	_, err := uf.CreateElectricInductance(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}