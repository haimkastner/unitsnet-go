// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestFuelEfficiencyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilometerPerLiter"}`
	
	factory := units.FuelEfficiencyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.FuelEfficiencyKilometerPerLiter {
		t.Errorf("expected unit %v, got %v", units.FuelEfficiencyKilometerPerLiter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilometerPerLiter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestFuelEfficiencyDto_ToJSON(t *testing.T) {
	dto := units.FuelEfficiencyDto{
		Value: 45,
		Unit:  units.FuelEfficiencyKilometerPerLiter,
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
	if result["unit"].(string) != string(units.FuelEfficiencyKilometerPerLiter) {
		t.Errorf("expected unit %s, got %v", units.FuelEfficiencyKilometerPerLiter, result["unit"])
	}
}

func TestNewFuelEfficiency_InvalidValue(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateFuelEfficiency(math.NaN(), units.FuelEfficiencyKilometerPerLiter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateFuelEfficiency(math.Inf(1), units.FuelEfficiencyKilometerPerLiter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestFuelEfficiencyConversions(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateFuelEfficiency(180, units.FuelEfficiencyKilometerPerLiter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to LitersPer100Kilometers.
		// No expected conversion value provided for LitersPer100Kilometers, verifying result is not NaN.
		result := a.LitersPer100Kilometers()
		cacheResult := a.LitersPer100Kilometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPer100Kilometers returned NaN")
		}
	}
	{
		// Test conversion to MilesPerUsGallon.
		// No expected conversion value provided for MilesPerUsGallon, verifying result is not NaN.
		result := a.MilesPerUsGallon()
		cacheResult := a.MilesPerUsGallon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilesPerUsGallon returned NaN")
		}
	}
	{
		// Test conversion to MilesPerUkGallon.
		// No expected conversion value provided for MilesPerUkGallon, verifying result is not NaN.
		result := a.MilesPerUkGallon()
		cacheResult := a.MilesPerUkGallon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilesPerUkGallon returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerLiter.
		// No expected conversion value provided for KilometersPerLiter, verifying result is not NaN.
		result := a.KilometersPerLiter()
		cacheResult := a.KilometersPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilometersPerLiter returned NaN")
		}
	}
}

func TestFuelEfficiency_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	a, err := factory.CreateFuelEfficiency(90, units.FuelEfficiencyKilometerPerLiter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.FuelEfficiencyKilometerPerLiter {
		t.Errorf("expected default unit KilometerPerLiter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.FuelEfficiencyLiterPer100Kilometers
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.FuelEfficiencyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.FuelEfficiencyKilometerPerLiter {
		t.Errorf("expected unit KilometerPerLiter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestFuelEfficiencyFactory_FromDto(t *testing.T) {
    factory := units.FuelEfficiencyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.FuelEfficiencyDto{
        Value: 100,
        Unit:  units.FuelEfficiencyKilometerPerLiter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.FuelEfficiencyDto{
        Value: math.NaN(),
        Unit:  units.FuelEfficiencyKilometerPerLiter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test LiterPer100Kilometers conversion
    liters_per100_kilometersDto := units.FuelEfficiencyDto{
        Value: 100,
        Unit:  units.FuelEfficiencyLiterPer100Kilometers,
    }
    
    var liters_per100_kilometersResult *units.FuelEfficiency
    liters_per100_kilometersResult, err = factory.FromDto(liters_per100_kilometersDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPer100Kilometers returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per100_kilometersResult.Convert(units.FuelEfficiencyLiterPer100Kilometers)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPer100Kilometers = %v, want %v", converted, 100)
    }
    // Test MilePerUsGallon conversion
    miles_per_us_gallonDto := units.FuelEfficiencyDto{
        Value: 100,
        Unit:  units.FuelEfficiencyMilePerUsGallon,
    }
    
    var miles_per_us_gallonResult *units.FuelEfficiency
    miles_per_us_gallonResult, err = factory.FromDto(miles_per_us_gallonDto)
    if err != nil {
        t.Errorf("FromDto() with MilePerUsGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = miles_per_us_gallonResult.Convert(units.FuelEfficiencyMilePerUsGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilePerUsGallon = %v, want %v", converted, 100)
    }
    // Test MilePerUkGallon conversion
    miles_per_uk_gallonDto := units.FuelEfficiencyDto{
        Value: 100,
        Unit:  units.FuelEfficiencyMilePerUkGallon,
    }
    
    var miles_per_uk_gallonResult *units.FuelEfficiency
    miles_per_uk_gallonResult, err = factory.FromDto(miles_per_uk_gallonDto)
    if err != nil {
        t.Errorf("FromDto() with MilePerUkGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = miles_per_uk_gallonResult.Convert(units.FuelEfficiencyMilePerUkGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilePerUkGallon = %v, want %v", converted, 100)
    }
    // Test KilometerPerLiter conversion
    kilometers_per_literDto := units.FuelEfficiencyDto{
        Value: 100,
        Unit:  units.FuelEfficiencyKilometerPerLiter,
    }
    
    var kilometers_per_literResult *units.FuelEfficiency
    kilometers_per_literResult, err = factory.FromDto(kilometers_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with KilometerPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_literResult.Convert(units.FuelEfficiencyKilometerPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerLiter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.FuelEfficiencyDto{
        Value: 0,
        Unit:  units.FuelEfficiencyKilometerPerLiter,
    }
    
    var zeroResult *units.FuelEfficiency
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestFuelEfficiencyFactory_FromDtoJSON(t *testing.T) {
    factory := units.FuelEfficiencyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KilometerPerLiter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KilometerPerLiter"}`)
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
    nanJSON, _ := json.Marshal(units.FuelEfficiencyDto{
        Value: nanValue,
        Unit:  units.FuelEfficiencyKilometerPerLiter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with LiterPer100Kilometers unit
    liters_per100_kilometersJSON := []byte(`{"value": 100, "unit": "LiterPer100Kilometers"}`)
    liters_per100_kilometersResult, err := factory.FromDtoJSON(liters_per100_kilometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPer100Kilometers unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per100_kilometersResult.Convert(units.FuelEfficiencyLiterPer100Kilometers)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPer100Kilometers = %v, want %v", converted, 100)
    }
    // Test JSON with MilePerUsGallon unit
    miles_per_us_gallonJSON := []byte(`{"value": 100, "unit": "MilePerUsGallon"}`)
    miles_per_us_gallonResult, err := factory.FromDtoJSON(miles_per_us_gallonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilePerUsGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = miles_per_us_gallonResult.Convert(units.FuelEfficiencyMilePerUsGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilePerUsGallon = %v, want %v", converted, 100)
    }
    // Test JSON with MilePerUkGallon unit
    miles_per_uk_gallonJSON := []byte(`{"value": 100, "unit": "MilePerUkGallon"}`)
    miles_per_uk_gallonResult, err := factory.FromDtoJSON(miles_per_uk_gallonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilePerUkGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = miles_per_uk_gallonResult.Convert(units.FuelEfficiencyMilePerUkGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilePerUkGallon = %v, want %v", converted, 100)
    }
    // Test JSON with KilometerPerLiter unit
    kilometers_per_literJSON := []byte(`{"value": 100, "unit": "KilometerPerLiter"}`)
    kilometers_per_literResult, err := factory.FromDtoJSON(kilometers_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilometerPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_literResult.Convert(units.FuelEfficiencyKilometerPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerLiter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KilometerPerLiter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromLitersPer100Kilometers function
func TestFuelEfficiencyFactory_FromLitersPer100Kilometers(t *testing.T) {
    factory := units.FuelEfficiencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPer100Kilometers(100)
    if err != nil {
        t.Errorf("FromLitersPer100Kilometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FuelEfficiencyLiterPer100Kilometers)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPer100Kilometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPer100Kilometers(math.NaN())
    if err == nil {
        t.Error("FromLitersPer100Kilometers() with NaN value should return error")
    }

    _, err = factory.FromLitersPer100Kilometers(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPer100Kilometers() with +Inf value should return error")
    }

    _, err = factory.FromLitersPer100Kilometers(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPer100Kilometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPer100Kilometers(0)
    if err != nil {
        t.Errorf("FromLitersPer100Kilometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FuelEfficiencyLiterPer100Kilometers)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPer100Kilometers() with zero value = %v, want 0", converted)
    }
}
// Test FromMilesPerUsGallon function
func TestFuelEfficiencyFactory_FromMilesPerUsGallon(t *testing.T) {
    factory := units.FuelEfficiencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilesPerUsGallon(100)
    if err != nil {
        t.Errorf("FromMilesPerUsGallon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FuelEfficiencyMilePerUsGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilesPerUsGallon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilesPerUsGallon(math.NaN())
    if err == nil {
        t.Error("FromMilesPerUsGallon() with NaN value should return error")
    }

    _, err = factory.FromMilesPerUsGallon(math.Inf(1))
    if err == nil {
        t.Error("FromMilesPerUsGallon() with +Inf value should return error")
    }

    _, err = factory.FromMilesPerUsGallon(math.Inf(-1))
    if err == nil {
        t.Error("FromMilesPerUsGallon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilesPerUsGallon(0)
    if err != nil {
        t.Errorf("FromMilesPerUsGallon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FuelEfficiencyMilePerUsGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilesPerUsGallon() with zero value = %v, want 0", converted)
    }
}
// Test FromMilesPerUkGallon function
func TestFuelEfficiencyFactory_FromMilesPerUkGallon(t *testing.T) {
    factory := units.FuelEfficiencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilesPerUkGallon(100)
    if err != nil {
        t.Errorf("FromMilesPerUkGallon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FuelEfficiencyMilePerUkGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilesPerUkGallon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilesPerUkGallon(math.NaN())
    if err == nil {
        t.Error("FromMilesPerUkGallon() with NaN value should return error")
    }

    _, err = factory.FromMilesPerUkGallon(math.Inf(1))
    if err == nil {
        t.Error("FromMilesPerUkGallon() with +Inf value should return error")
    }

    _, err = factory.FromMilesPerUkGallon(math.Inf(-1))
    if err == nil {
        t.Error("FromMilesPerUkGallon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilesPerUkGallon(0)
    if err != nil {
        t.Errorf("FromMilesPerUkGallon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FuelEfficiencyMilePerUkGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilesPerUkGallon() with zero value = %v, want 0", converted)
    }
}
// Test FromKilometersPerLiter function
func TestFuelEfficiencyFactory_FromKilometersPerLiter(t *testing.T) {
    factory := units.FuelEfficiencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilometersPerLiter(100)
    if err != nil {
        t.Errorf("FromKilometersPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FuelEfficiencyKilometerPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilometersPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilometersPerLiter(math.NaN())
    if err == nil {
        t.Error("FromKilometersPerLiter() with NaN value should return error")
    }

    _, err = factory.FromKilometersPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromKilometersPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromKilometersPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilometersPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilometersPerLiter(0)
    if err != nil {
        t.Errorf("FromKilometersPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FuelEfficiencyKilometerPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilometersPerLiter() with zero value = %v, want 0", converted)
    }
}

func TestFuelEfficiencyToString(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	a, err := factory.CreateFuelEfficiency(45, units.FuelEfficiencyKilometerPerLiter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.FuelEfficiencyKilometerPerLiter, 2)
	expected := "45.00 " + units.GetFuelEfficiencyAbbreviation(units.FuelEfficiencyKilometerPerLiter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.FuelEfficiencyKilometerPerLiter, -1)
	expected = "45 " + units.GetFuelEfficiencyAbbreviation(units.FuelEfficiencyKilometerPerLiter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestFuelEfficiency_EqualityAndComparison(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	a1, _ := factory.CreateFuelEfficiency(60, units.FuelEfficiencyKilometerPerLiter)
	a2, _ := factory.CreateFuelEfficiency(60, units.FuelEfficiencyKilometerPerLiter)
	a3, _ := factory.CreateFuelEfficiency(120, units.FuelEfficiencyKilometerPerLiter)

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

func TestFuelEfficiency_Arithmetic(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	a1, _ := factory.CreateFuelEfficiency(30, units.FuelEfficiencyKilometerPerLiter)
	a2, _ := factory.CreateFuelEfficiency(45, units.FuelEfficiencyKilometerPerLiter)

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


func TestGetFuelEfficiencyAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.FuelEfficiencyUnits
        want string
    }{
        {
            name: "LiterPer100Kilometers abbreviation",
            unit: units.FuelEfficiencyLiterPer100Kilometers,
            want: "l/100km",
        },
        {
            name: "MilePerUsGallon abbreviation",
            unit: units.FuelEfficiencyMilePerUsGallon,
            want: "mpg (U.S.)",
        },
        {
            name: "MilePerUkGallon abbreviation",
            unit: units.FuelEfficiencyMilePerUkGallon,
            want: "mpg (imp.)",
        },
        {
            name: "KilometerPerLiter abbreviation",
            unit: units.FuelEfficiencyKilometerPerLiter,
            want: "km/l",
        },
        {
            name: "invalid unit",
            unit: units.FuelEfficiencyUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetFuelEfficiencyAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetFuelEfficiencyAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestFuelEfficiency_String(t *testing.T) {
    factory := units.FuelEfficiencyFactory{}
    
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
            unit, err := factory.CreateFuelEfficiency(tt.value, units.FuelEfficiencyKilometerPerLiter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("FuelEfficiency.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestFuelEfficiency_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.FuelEfficiencyFactory{}

	_, err := uf.CreateFuelEfficiency(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}