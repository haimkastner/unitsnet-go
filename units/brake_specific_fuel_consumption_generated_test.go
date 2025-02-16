// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestBrakeSpecificFuelConsumptionDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerJoule"}`
	
	factory := units.BrakeSpecificFuelConsumptionDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.BrakeSpecificFuelConsumptionKilogramPerJoule {
		t.Errorf("expected unit %v, got %v", units.BrakeSpecificFuelConsumptionKilogramPerJoule, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerJoule"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestBrakeSpecificFuelConsumptionDto_ToJSON(t *testing.T) {
	dto := units.BrakeSpecificFuelConsumptionDto{
		Value: 45,
		Unit:  units.BrakeSpecificFuelConsumptionKilogramPerJoule,
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
	if result["unit"].(string) != string(units.BrakeSpecificFuelConsumptionKilogramPerJoule) {
		t.Errorf("expected unit %s, got %v", units.BrakeSpecificFuelConsumptionKilogramPerJoule, result["unit"])
	}
}

func TestNewBrakeSpecificFuelConsumption_InvalidValue(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	// NaN value should return an error.
	_, err := factory.CreateBrakeSpecificFuelConsumption(math.NaN(), units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateBrakeSpecificFuelConsumption(math.Inf(1), units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestBrakeSpecificFuelConsumptionConversions(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateBrakeSpecificFuelConsumption(180, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerKiloWattHour.
		// No expected conversion value provided for GramsPerKiloWattHour, verifying result is not NaN.
		result := a.GramsPerKiloWattHour()
		cacheResult := a.GramsPerKiloWattHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerKiloWattHour returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerJoule.
		// No expected conversion value provided for KilogramsPerJoule, verifying result is not NaN.
		result := a.KilogramsPerJoule()
		cacheResult := a.KilogramsPerJoule()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerJoule returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerMechanicalHorsepowerHour.
		// No expected conversion value provided for PoundsPerMechanicalHorsepowerHour, verifying result is not NaN.
		result := a.PoundsPerMechanicalHorsepowerHour()
		cacheResult := a.PoundsPerMechanicalHorsepowerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsPerMechanicalHorsepowerHour returned NaN")
		}
	}
}

func TestBrakeSpecificFuelConsumption_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	a, err := factory.CreateBrakeSpecificFuelConsumption(90, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.BrakeSpecificFuelConsumptionKilogramPerJoule {
		t.Errorf("expected default unit KilogramPerJoule, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.BrakeSpecificFuelConsumptionGramPerKiloWattHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.BrakeSpecificFuelConsumptionDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.BrakeSpecificFuelConsumptionKilogramPerJoule {
		t.Errorf("expected unit KilogramPerJoule, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestBrakeSpecificFuelConsumptionFactory_FromDto(t *testing.T) {
    factory := units.BrakeSpecificFuelConsumptionFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.BrakeSpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.BrakeSpecificFuelConsumptionKilogramPerJoule,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.BrakeSpecificFuelConsumptionDto{
        Value: math.NaN(),
        Unit:  units.BrakeSpecificFuelConsumptionKilogramPerJoule,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GramPerKiloWattHour conversion
    grams_per_kilo_watt_hourDto := units.BrakeSpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.BrakeSpecificFuelConsumptionGramPerKiloWattHour,
    }
    
    var grams_per_kilo_watt_hourResult *units.BrakeSpecificFuelConsumption
    grams_per_kilo_watt_hourResult, err = factory.FromDto(grams_per_kilo_watt_hourDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerKiloWattHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_kilo_watt_hourResult.Convert(units.BrakeSpecificFuelConsumptionGramPerKiloWattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerKiloWattHour = %v, want %v", converted, 100)
    }
    // Test KilogramPerJoule conversion
    kilograms_per_jouleDto := units.BrakeSpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.BrakeSpecificFuelConsumptionKilogramPerJoule,
    }
    
    var kilograms_per_jouleResult *units.BrakeSpecificFuelConsumption
    kilograms_per_jouleResult, err = factory.FromDto(kilograms_per_jouleDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerJoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_jouleResult.Convert(units.BrakeSpecificFuelConsumptionKilogramPerJoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerJoule = %v, want %v", converted, 100)
    }
    // Test PoundPerMechanicalHorsepowerHour conversion
    pounds_per_mechanical_horsepower_hourDto := units.BrakeSpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour,
    }
    
    var pounds_per_mechanical_horsepower_hourResult *units.BrakeSpecificFuelConsumption
    pounds_per_mechanical_horsepower_hourResult, err = factory.FromDto(pounds_per_mechanical_horsepower_hourDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerMechanicalHorsepowerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_mechanical_horsepower_hourResult.Convert(units.BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerMechanicalHorsepowerHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.BrakeSpecificFuelConsumptionDto{
        Value: 0,
        Unit:  units.BrakeSpecificFuelConsumptionKilogramPerJoule,
    }
    
    var zeroResult *units.BrakeSpecificFuelConsumption
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestBrakeSpecificFuelConsumptionFactory_FromDtoJSON(t *testing.T) {
    factory := units.BrakeSpecificFuelConsumptionFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KilogramPerJoule"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KilogramPerJoule"}`)
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
    nanJSON, _ := json.Marshal(units.BrakeSpecificFuelConsumptionDto{
        Value: nanValue,
        Unit:  units.BrakeSpecificFuelConsumptionKilogramPerJoule,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with GramPerKiloWattHour unit
    grams_per_kilo_watt_hourJSON := []byte(`{"value": 100, "unit": "GramPerKiloWattHour"}`)
    grams_per_kilo_watt_hourResult, err := factory.FromDtoJSON(grams_per_kilo_watt_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerKiloWattHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_kilo_watt_hourResult.Convert(units.BrakeSpecificFuelConsumptionGramPerKiloWattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerKiloWattHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerJoule unit
    kilograms_per_jouleJSON := []byte(`{"value": 100, "unit": "KilogramPerJoule"}`)
    kilograms_per_jouleResult, err := factory.FromDtoJSON(kilograms_per_jouleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerJoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_jouleResult.Convert(units.BrakeSpecificFuelConsumptionKilogramPerJoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerJoule = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerMechanicalHorsepowerHour unit
    pounds_per_mechanical_horsepower_hourJSON := []byte(`{"value": 100, "unit": "PoundPerMechanicalHorsepowerHour"}`)
    pounds_per_mechanical_horsepower_hourResult, err := factory.FromDtoJSON(pounds_per_mechanical_horsepower_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerMechanicalHorsepowerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_mechanical_horsepower_hourResult.Convert(units.BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerMechanicalHorsepowerHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KilogramPerJoule"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGramsPerKiloWattHour function
func TestBrakeSpecificFuelConsumptionFactory_FromGramsPerKiloWattHour(t *testing.T) {
    factory := units.BrakeSpecificFuelConsumptionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerKiloWattHour(100)
    if err != nil {
        t.Errorf("FromGramsPerKiloWattHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BrakeSpecificFuelConsumptionGramPerKiloWattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerKiloWattHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerKiloWattHour(math.NaN())
    if err == nil {
        t.Error("FromGramsPerKiloWattHour() with NaN value should return error")
    }

    _, err = factory.FromGramsPerKiloWattHour(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerKiloWattHour() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerKiloWattHour(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerKiloWattHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerKiloWattHour(0)
    if err != nil {
        t.Errorf("FromGramsPerKiloWattHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BrakeSpecificFuelConsumptionGramPerKiloWattHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerKiloWattHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerJoule function
func TestBrakeSpecificFuelConsumptionFactory_FromKilogramsPerJoule(t *testing.T) {
    factory := units.BrakeSpecificFuelConsumptionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerJoule(100)
    if err != nil {
        t.Errorf("FromKilogramsPerJoule() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BrakeSpecificFuelConsumptionKilogramPerJoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerJoule() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerJoule(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerJoule() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerJoule(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerJoule() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerJoule(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerJoule() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerJoule(0)
    if err != nil {
        t.Errorf("FromKilogramsPerJoule() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BrakeSpecificFuelConsumptionKilogramPerJoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerJoule() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerMechanicalHorsepowerHour function
func TestBrakeSpecificFuelConsumptionFactory_FromPoundsPerMechanicalHorsepowerHour(t *testing.T) {
    factory := units.BrakeSpecificFuelConsumptionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerMechanicalHorsepowerHour(100)
    if err != nil {
        t.Errorf("FromPoundsPerMechanicalHorsepowerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerMechanicalHorsepowerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerMechanicalHorsepowerHour(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerMechanicalHorsepowerHour() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerMechanicalHorsepowerHour(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerMechanicalHorsepowerHour() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerMechanicalHorsepowerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerMechanicalHorsepowerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerMechanicalHorsepowerHour(0)
    if err != nil {
        t.Errorf("FromPoundsPerMechanicalHorsepowerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerMechanicalHorsepowerHour() with zero value = %v, want 0", converted)
    }
}

func TestBrakeSpecificFuelConsumptionToString(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	a, err := factory.CreateBrakeSpecificFuelConsumption(45, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.BrakeSpecificFuelConsumptionKilogramPerJoule, 2)
	expected := "45.00 " + units.GetBrakeSpecificFuelConsumptionAbbreviation(units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.BrakeSpecificFuelConsumptionKilogramPerJoule, -1)
	expected = "45 " + units.GetBrakeSpecificFuelConsumptionAbbreviation(units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestBrakeSpecificFuelConsumption_EqualityAndComparison(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	a1, _ := factory.CreateBrakeSpecificFuelConsumption(60, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	a2, _ := factory.CreateBrakeSpecificFuelConsumption(60, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	a3, _ := factory.CreateBrakeSpecificFuelConsumption(120, units.BrakeSpecificFuelConsumptionKilogramPerJoule)

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

func TestBrakeSpecificFuelConsumption_Arithmetic(t *testing.T) {
	factory := units.BrakeSpecificFuelConsumptionFactory{}
	a1, _ := factory.CreateBrakeSpecificFuelConsumption(30, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
	a2, _ := factory.CreateBrakeSpecificFuelConsumption(45, units.BrakeSpecificFuelConsumptionKilogramPerJoule)

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


func TestGetBrakeSpecificFuelConsumptionAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.BrakeSpecificFuelConsumptionUnits
        want string
    }{
        {
            name: "GramPerKiloWattHour abbreviation",
            unit: units.BrakeSpecificFuelConsumptionGramPerKiloWattHour,
            want: "g/kWh",
        },
        {
            name: "KilogramPerJoule abbreviation",
            unit: units.BrakeSpecificFuelConsumptionKilogramPerJoule,
            want: "kg/J",
        },
        {
            name: "PoundPerMechanicalHorsepowerHour abbreviation",
            unit: units.BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour,
            want: "lb/hph",
        },
        {
            name: "invalid unit",
            unit: units.BrakeSpecificFuelConsumptionUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetBrakeSpecificFuelConsumptionAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetBrakeSpecificFuelConsumptionAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestBrakeSpecificFuelConsumption_String(t *testing.T) {
    factory := units.BrakeSpecificFuelConsumptionFactory{}
    
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
            unit, err := factory.CreateBrakeSpecificFuelConsumption(tt.value, units.BrakeSpecificFuelConsumptionKilogramPerJoule)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("BrakeSpecificFuelConsumption.String() = %v, want %v", got, tt.want)
            }
        })
    }
}