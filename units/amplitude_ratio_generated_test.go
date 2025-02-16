// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAmplitudeRatioDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecibelVolt"}`
	
	factory := units.AmplitudeRatioDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AmplitudeRatioDecibelVolt {
		t.Errorf("expected unit %v, got %v", units.AmplitudeRatioDecibelVolt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecibelVolt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAmplitudeRatioDto_ToJSON(t *testing.T) {
	dto := units.AmplitudeRatioDto{
		Value: 45,
		Unit:  units.AmplitudeRatioDecibelVolt,
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
	if result["unit"].(string) != string(units.AmplitudeRatioDecibelVolt) {
		t.Errorf("expected unit %s, got %v", units.AmplitudeRatioDecibelVolt, result["unit"])
	}
}

func TestNewAmplitudeRatio_InvalidValue(t *testing.T) {
	factory := units.AmplitudeRatioFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAmplitudeRatio(math.NaN(), units.AmplitudeRatioDecibelVolt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAmplitudeRatio(math.Inf(1), units.AmplitudeRatioDecibelVolt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAmplitudeRatioConversions(t *testing.T) {
	factory := units.AmplitudeRatioFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAmplitudeRatio(180, units.AmplitudeRatioDecibelVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to DecibelVolts.
		// No expected conversion value provided for DecibelVolts, verifying result is not NaN.
		result := a.DecibelVolts()
		cacheResult := a.DecibelVolts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecibelVolts returned NaN")
		}
	}
	{
		// Test conversion to DecibelMicrovolts.
		// No expected conversion value provided for DecibelMicrovolts, verifying result is not NaN.
		result := a.DecibelMicrovolts()
		cacheResult := a.DecibelMicrovolts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecibelMicrovolts returned NaN")
		}
	}
	{
		// Test conversion to DecibelMillivolts.
		// No expected conversion value provided for DecibelMillivolts, verifying result is not NaN.
		result := a.DecibelMillivolts()
		cacheResult := a.DecibelMillivolts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecibelMillivolts returned NaN")
		}
	}
	{
		// Test conversion to DecibelsUnloaded.
		// No expected conversion value provided for DecibelsUnloaded, verifying result is not NaN.
		result := a.DecibelsUnloaded()
		cacheResult := a.DecibelsUnloaded()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecibelsUnloaded returned NaN")
		}
	}
}

func TestAmplitudeRatio_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AmplitudeRatioFactory{}
	a, err := factory.CreateAmplitudeRatio(90, units.AmplitudeRatioDecibelVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AmplitudeRatioDecibelVolt {
		t.Errorf("expected default unit DecibelVolt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AmplitudeRatioDecibelVolt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AmplitudeRatioDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AmplitudeRatioDecibelVolt {
		t.Errorf("expected unit DecibelVolt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAmplitudeRatioFactory_FromDto(t *testing.T) {
    factory := units.AmplitudeRatioFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AmplitudeRatioDto{
        Value: 100,
        Unit:  units.AmplitudeRatioDecibelVolt,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AmplitudeRatioDto{
        Value: math.NaN(),
        Unit:  units.AmplitudeRatioDecibelVolt,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test DecibelVolt conversion
    decibel_voltsDto := units.AmplitudeRatioDto{
        Value: 100,
        Unit:  units.AmplitudeRatioDecibelVolt,
    }
    
    var decibel_voltsResult *units.AmplitudeRatio
    decibel_voltsResult, err = factory.FromDto(decibel_voltsDto)
    if err != nil {
        t.Errorf("FromDto() with DecibelVolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_voltsResult.Convert(units.AmplitudeRatioDecibelVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelVolt = %v, want %v", converted, 100)
    }
    // Test DecibelMicrovolt conversion
    decibel_microvoltsDto := units.AmplitudeRatioDto{
        Value: 100,
        Unit:  units.AmplitudeRatioDecibelMicrovolt,
    }
    
    var decibel_microvoltsResult *units.AmplitudeRatio
    decibel_microvoltsResult, err = factory.FromDto(decibel_microvoltsDto)
    if err != nil {
        t.Errorf("FromDto() with DecibelMicrovolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_microvoltsResult.Convert(units.AmplitudeRatioDecibelMicrovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelMicrovolt = %v, want %v", converted, 100)
    }
    // Test DecibelMillivolt conversion
    decibel_millivoltsDto := units.AmplitudeRatioDto{
        Value: 100,
        Unit:  units.AmplitudeRatioDecibelMillivolt,
    }
    
    var decibel_millivoltsResult *units.AmplitudeRatio
    decibel_millivoltsResult, err = factory.FromDto(decibel_millivoltsDto)
    if err != nil {
        t.Errorf("FromDto() with DecibelMillivolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_millivoltsResult.Convert(units.AmplitudeRatioDecibelMillivolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelMillivolt = %v, want %v", converted, 100)
    }
    // Test DecibelUnloaded conversion
    decibels_unloadedDto := units.AmplitudeRatioDto{
        Value: 100,
        Unit:  units.AmplitudeRatioDecibelUnloaded,
    }
    
    var decibels_unloadedResult *units.AmplitudeRatio
    decibels_unloadedResult, err = factory.FromDto(decibels_unloadedDto)
    if err != nil {
        t.Errorf("FromDto() with DecibelUnloaded returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibels_unloadedResult.Convert(units.AmplitudeRatioDecibelUnloaded)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelUnloaded = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AmplitudeRatioDto{
        Value: 0,
        Unit:  units.AmplitudeRatioDecibelVolt,
    }
    
    var zeroResult *units.AmplitudeRatio
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAmplitudeRatioFactory_FromDtoJSON(t *testing.T) {
    factory := units.AmplitudeRatioFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "DecibelVolt"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "DecibelVolt"}`)
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
    nanJSON, _ := json.Marshal(units.AmplitudeRatioDto{
        Value: nanValue,
        Unit:  units.AmplitudeRatioDecibelVolt,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with DecibelVolt unit
    decibel_voltsJSON := []byte(`{"value": 100, "unit": "DecibelVolt"}`)
    decibel_voltsResult, err := factory.FromDtoJSON(decibel_voltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecibelVolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_voltsResult.Convert(units.AmplitudeRatioDecibelVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelVolt = %v, want %v", converted, 100)
    }
    // Test JSON with DecibelMicrovolt unit
    decibel_microvoltsJSON := []byte(`{"value": 100, "unit": "DecibelMicrovolt"}`)
    decibel_microvoltsResult, err := factory.FromDtoJSON(decibel_microvoltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecibelMicrovolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_microvoltsResult.Convert(units.AmplitudeRatioDecibelMicrovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelMicrovolt = %v, want %v", converted, 100)
    }
    // Test JSON with DecibelMillivolt unit
    decibel_millivoltsJSON := []byte(`{"value": 100, "unit": "DecibelMillivolt"}`)
    decibel_millivoltsResult, err := factory.FromDtoJSON(decibel_millivoltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecibelMillivolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_millivoltsResult.Convert(units.AmplitudeRatioDecibelMillivolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelMillivolt = %v, want %v", converted, 100)
    }
    // Test JSON with DecibelUnloaded unit
    decibels_unloadedJSON := []byte(`{"value": 100, "unit": "DecibelUnloaded"}`)
    decibels_unloadedResult, err := factory.FromDtoJSON(decibels_unloadedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecibelUnloaded unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibels_unloadedResult.Convert(units.AmplitudeRatioDecibelUnloaded)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelUnloaded = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "DecibelVolt"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDecibelVolts function
func TestAmplitudeRatioFactory_FromDecibelVolts(t *testing.T) {
    factory := units.AmplitudeRatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecibelVolts(100)
    if err != nil {
        t.Errorf("FromDecibelVolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmplitudeRatioDecibelVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecibelVolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecibelVolts(math.NaN())
    if err == nil {
        t.Error("FromDecibelVolts() with NaN value should return error")
    }

    _, err = factory.FromDecibelVolts(math.Inf(1))
    if err == nil {
        t.Error("FromDecibelVolts() with +Inf value should return error")
    }

    _, err = factory.FromDecibelVolts(math.Inf(-1))
    if err == nil {
        t.Error("FromDecibelVolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecibelVolts(0)
    if err != nil {
        t.Errorf("FromDecibelVolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmplitudeRatioDecibelVolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecibelVolts() with zero value = %v, want 0", converted)
    }
}
// Test FromDecibelMicrovolts function
func TestAmplitudeRatioFactory_FromDecibelMicrovolts(t *testing.T) {
    factory := units.AmplitudeRatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecibelMicrovolts(100)
    if err != nil {
        t.Errorf("FromDecibelMicrovolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmplitudeRatioDecibelMicrovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecibelMicrovolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecibelMicrovolts(math.NaN())
    if err == nil {
        t.Error("FromDecibelMicrovolts() with NaN value should return error")
    }

    _, err = factory.FromDecibelMicrovolts(math.Inf(1))
    if err == nil {
        t.Error("FromDecibelMicrovolts() with +Inf value should return error")
    }

    _, err = factory.FromDecibelMicrovolts(math.Inf(-1))
    if err == nil {
        t.Error("FromDecibelMicrovolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecibelMicrovolts(0)
    if err != nil {
        t.Errorf("FromDecibelMicrovolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmplitudeRatioDecibelMicrovolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecibelMicrovolts() with zero value = %v, want 0", converted)
    }
}
// Test FromDecibelMillivolts function
func TestAmplitudeRatioFactory_FromDecibelMillivolts(t *testing.T) {
    factory := units.AmplitudeRatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecibelMillivolts(100)
    if err != nil {
        t.Errorf("FromDecibelMillivolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmplitudeRatioDecibelMillivolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecibelMillivolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecibelMillivolts(math.NaN())
    if err == nil {
        t.Error("FromDecibelMillivolts() with NaN value should return error")
    }

    _, err = factory.FromDecibelMillivolts(math.Inf(1))
    if err == nil {
        t.Error("FromDecibelMillivolts() with +Inf value should return error")
    }

    _, err = factory.FromDecibelMillivolts(math.Inf(-1))
    if err == nil {
        t.Error("FromDecibelMillivolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecibelMillivolts(0)
    if err != nil {
        t.Errorf("FromDecibelMillivolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmplitudeRatioDecibelMillivolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecibelMillivolts() with zero value = %v, want 0", converted)
    }
}
// Test FromDecibelsUnloaded function
func TestAmplitudeRatioFactory_FromDecibelsUnloaded(t *testing.T) {
    factory := units.AmplitudeRatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecibelsUnloaded(100)
    if err != nil {
        t.Errorf("FromDecibelsUnloaded() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmplitudeRatioDecibelUnloaded)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecibelsUnloaded() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecibelsUnloaded(math.NaN())
    if err == nil {
        t.Error("FromDecibelsUnloaded() with NaN value should return error")
    }

    _, err = factory.FromDecibelsUnloaded(math.Inf(1))
    if err == nil {
        t.Error("FromDecibelsUnloaded() with +Inf value should return error")
    }

    _, err = factory.FromDecibelsUnloaded(math.Inf(-1))
    if err == nil {
        t.Error("FromDecibelsUnloaded() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecibelsUnloaded(0)
    if err != nil {
        t.Errorf("FromDecibelsUnloaded() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmplitudeRatioDecibelUnloaded)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecibelsUnloaded() with zero value = %v, want 0", converted)
    }
}

func TestAmplitudeRatioToString(t *testing.T) {
	factory := units.AmplitudeRatioFactory{}
	a, err := factory.CreateAmplitudeRatio(45, units.AmplitudeRatioDecibelVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AmplitudeRatioDecibelVolt, 2)
	expected := "45.00 " + units.GetAmplitudeRatioAbbreviation(units.AmplitudeRatioDecibelVolt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AmplitudeRatioDecibelVolt, -1)
	expected = "45 " + units.GetAmplitudeRatioAbbreviation(units.AmplitudeRatioDecibelVolt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAmplitudeRatio_EqualityAndComparison(t *testing.T) {
	factory := units.AmplitudeRatioFactory{}
	a1, _ := factory.CreateAmplitudeRatio(60, units.AmplitudeRatioDecibelVolt)
	a2, _ := factory.CreateAmplitudeRatio(60, units.AmplitudeRatioDecibelVolt)
	a3, _ := factory.CreateAmplitudeRatio(120, units.AmplitudeRatioDecibelVolt)

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

func TestAmplitudeRatio_Arithmetic(t *testing.T) {
	factory := units.AmplitudeRatioFactory{}
	a1, _ := factory.CreateAmplitudeRatio(30, units.AmplitudeRatioDecibelVolt)
	a2, _ := factory.CreateAmplitudeRatio(45, units.AmplitudeRatioDecibelVolt)

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


func TestGetAmplitudeRatioAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.AmplitudeRatioUnits
        want string
    }{
        {
            name: "DecibelVolt abbreviation",
            unit: units.AmplitudeRatioDecibelVolt,
            want: "dBV",
        },
        {
            name: "DecibelMicrovolt abbreviation",
            unit: units.AmplitudeRatioDecibelMicrovolt,
            want: "dBµV",
        },
        {
            name: "DecibelMillivolt abbreviation",
            unit: units.AmplitudeRatioDecibelMillivolt,
            want: "dBmV",
        },
        {
            name: "DecibelUnloaded abbreviation",
            unit: units.AmplitudeRatioDecibelUnloaded,
            want: "dBu",
        },
        {
            name: "invalid unit",
            unit: units.AmplitudeRatioUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetAmplitudeRatioAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetAmplitudeRatioAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestAmplitudeRatio_String(t *testing.T) {
    factory := units.AmplitudeRatioFactory{}
    
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
            unit, err := factory.CreateAmplitudeRatio(tt.value, units.AmplitudeRatioDecibelVolt)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("AmplitudeRatio.String() = %v, want %v", got, tt.want)
            }
        })
    }
}