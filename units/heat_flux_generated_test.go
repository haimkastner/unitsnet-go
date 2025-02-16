// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestHeatFluxDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerSquareMeter"}`
	
	factory := units.HeatFluxDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.HeatFluxWattPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.HeatFluxWattPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestHeatFluxDto_ToJSON(t *testing.T) {
	dto := units.HeatFluxDto{
		Value: 45,
		Unit:  units.HeatFluxWattPerSquareMeter,
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
	if result["unit"].(string) != string(units.HeatFluxWattPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.HeatFluxWattPerSquareMeter, result["unit"])
	}
}

func TestNewHeatFlux_InvalidValue(t *testing.T) {
	factory := units.HeatFluxFactory{}
	// NaN value should return an error.
	_, err := factory.CreateHeatFlux(math.NaN(), units.HeatFluxWattPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateHeatFlux(math.Inf(1), units.HeatFluxWattPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestHeatFluxConversions(t *testing.T) {
	factory := units.HeatFluxFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateHeatFlux(180, units.HeatFluxWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerSquareMeter.
		// No expected conversion value provided for WattsPerSquareMeter, verifying result is not NaN.
		result := a.WattsPerSquareMeter()
		cacheResult := a.WattsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerSquareInch.
		// No expected conversion value provided for WattsPerSquareInch, verifying result is not NaN.
		result := a.WattsPerSquareInch()
		cacheResult := a.WattsPerSquareInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to WattsPerSquareFoot.
		// No expected conversion value provided for WattsPerSquareFoot, verifying result is not NaN.
		result := a.WattsPerSquareFoot()
		cacheResult := a.WattsPerSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to BtusPerSecondSquareInch.
		// No expected conversion value provided for BtusPerSecondSquareInch, verifying result is not NaN.
		result := a.BtusPerSecondSquareInch()
		cacheResult := a.BtusPerSecondSquareInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BtusPerSecondSquareInch returned NaN")
		}
	}
	{
		// Test conversion to BtusPerSecondSquareFoot.
		// No expected conversion value provided for BtusPerSecondSquareFoot, verifying result is not NaN.
		result := a.BtusPerSecondSquareFoot()
		cacheResult := a.BtusPerSecondSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BtusPerSecondSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to BtusPerMinuteSquareFoot.
		// No expected conversion value provided for BtusPerMinuteSquareFoot, verifying result is not NaN.
		result := a.BtusPerMinuteSquareFoot()
		cacheResult := a.BtusPerMinuteSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BtusPerMinuteSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to BtusPerHourSquareFoot.
		// No expected conversion value provided for BtusPerHourSquareFoot, verifying result is not NaN.
		result := a.BtusPerHourSquareFoot()
		cacheResult := a.BtusPerHourSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BtusPerHourSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerSecondSquareCentimeter.
		// No expected conversion value provided for CaloriesPerSecondSquareCentimeter, verifying result is not NaN.
		result := a.CaloriesPerSecondSquareCentimeter()
		cacheResult := a.CaloriesPerSecondSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CaloriesPerSecondSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerHourSquareMeter.
		// No expected conversion value provided for KilocaloriesPerHourSquareMeter, verifying result is not NaN.
		result := a.KilocaloriesPerHourSquareMeter()
		cacheResult := a.KilocaloriesPerHourSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocaloriesPerHourSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerFootSecond.
		// No expected conversion value provided for PoundsForcePerFootSecond, verifying result is not NaN.
		result := a.PoundsForcePerFootSecond()
		cacheResult := a.PoundsForcePerFootSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerFootSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerSecondCubed.
		// No expected conversion value provided for PoundsPerSecondCubed, verifying result is not NaN.
		result := a.PoundsPerSecondCubed()
		cacheResult := a.PoundsPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerSquareMeter.
		// No expected conversion value provided for NanowattsPerSquareMeter, verifying result is not NaN.
		result := a.NanowattsPerSquareMeter()
		cacheResult := a.NanowattsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerSquareMeter.
		// No expected conversion value provided for MicrowattsPerSquareMeter, verifying result is not NaN.
		result := a.MicrowattsPerSquareMeter()
		cacheResult := a.MicrowattsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerSquareMeter.
		// No expected conversion value provided for MilliwattsPerSquareMeter, verifying result is not NaN.
		result := a.MilliwattsPerSquareMeter()
		cacheResult := a.MilliwattsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to CentiwattsPerSquareMeter.
		// No expected conversion value provided for CentiwattsPerSquareMeter, verifying result is not NaN.
		result := a.CentiwattsPerSquareMeter()
		cacheResult := a.CentiwattsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentiwattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerSquareMeter.
		// No expected conversion value provided for DeciwattsPerSquareMeter, verifying result is not NaN.
		result := a.DeciwattsPerSquareMeter()
		cacheResult := a.DeciwattsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DeciwattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerSquareMeter.
		// No expected conversion value provided for KilowattsPerSquareMeter, verifying result is not NaN.
		result := a.KilowattsPerSquareMeter()
		cacheResult := a.KilowattsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerSecondSquareCentimeter.
		// No expected conversion value provided for KilocaloriesPerSecondSquareCentimeter, verifying result is not NaN.
		result := a.KilocaloriesPerSecondSquareCentimeter()
		cacheResult := a.KilocaloriesPerSecondSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocaloriesPerSecondSquareCentimeter returned NaN")
		}
	}
}

func TestHeatFlux_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.HeatFluxFactory{}
	a, err := factory.CreateHeatFlux(90, units.HeatFluxWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.HeatFluxWattPerSquareMeter {
		t.Errorf("expected default unit WattPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.HeatFluxWattPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.HeatFluxDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.HeatFluxWattPerSquareMeter {
		t.Errorf("expected unit WattPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestHeatFluxFactory_FromDto(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxWattPerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.HeatFluxDto{
        Value: math.NaN(),
        Unit:  units.HeatFluxWattPerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test WattPerSquareMeter conversion
    watts_per_square_meterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxWattPerSquareMeter,
    }
    
    var watts_per_square_meterResult *units.HeatFlux
    watts_per_square_meterResult, err = factory.FromDto(watts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_meterResult.Convert(units.HeatFluxWattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test WattPerSquareInch conversion
    watts_per_square_inchDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxWattPerSquareInch,
    }
    
    var watts_per_square_inchResult *units.HeatFlux
    watts_per_square_inchResult, err = factory.FromDto(watts_per_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_inchResult.Convert(units.HeatFluxWattPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareInch = %v, want %v", converted, 100)
    }
    // Test WattPerSquareFoot conversion
    watts_per_square_footDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxWattPerSquareFoot,
    }
    
    var watts_per_square_footResult *units.HeatFlux
    watts_per_square_footResult, err = factory.FromDto(watts_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_footResult.Convert(units.HeatFluxWattPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareFoot = %v, want %v", converted, 100)
    }
    // Test BtuPerSecondSquareInch conversion
    btus_per_second_square_inchDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxBtuPerSecondSquareInch,
    }
    
    var btus_per_second_square_inchResult *units.HeatFlux
    btus_per_second_square_inchResult, err = factory.FromDto(btus_per_second_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerSecondSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_second_square_inchResult.Convert(units.HeatFluxBtuPerSecondSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerSecondSquareInch = %v, want %v", converted, 100)
    }
    // Test BtuPerSecondSquareFoot conversion
    btus_per_second_square_footDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxBtuPerSecondSquareFoot,
    }
    
    var btus_per_second_square_footResult *units.HeatFlux
    btus_per_second_square_footResult, err = factory.FromDto(btus_per_second_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerSecondSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_second_square_footResult.Convert(units.HeatFluxBtuPerSecondSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerSecondSquareFoot = %v, want %v", converted, 100)
    }
    // Test BtuPerMinuteSquareFoot conversion
    btus_per_minute_square_footDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxBtuPerMinuteSquareFoot,
    }
    
    var btus_per_minute_square_footResult *units.HeatFlux
    btus_per_minute_square_footResult, err = factory.FromDto(btus_per_minute_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerMinuteSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_minute_square_footResult.Convert(units.HeatFluxBtuPerMinuteSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerMinuteSquareFoot = %v, want %v", converted, 100)
    }
    // Test BtuPerHourSquareFoot conversion
    btus_per_hour_square_footDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxBtuPerHourSquareFoot,
    }
    
    var btus_per_hour_square_footResult *units.HeatFlux
    btus_per_hour_square_footResult, err = factory.FromDto(btus_per_hour_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerHourSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_hour_square_footResult.Convert(units.HeatFluxBtuPerHourSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerHourSquareFoot = %v, want %v", converted, 100)
    }
    // Test CaloriePerSecondSquareCentimeter conversion
    calories_per_second_square_centimeterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxCaloriePerSecondSquareCentimeter,
    }
    
    var calories_per_second_square_centimeterResult *units.HeatFlux
    calories_per_second_square_centimeterResult, err = factory.FromDto(calories_per_second_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with CaloriePerSecondSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_second_square_centimeterResult.Convert(units.HeatFluxCaloriePerSecondSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerSecondSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilocaloriePerHourSquareMeter conversion
    kilocalories_per_hour_square_meterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxKilocaloriePerHourSquareMeter,
    }
    
    var kilocalories_per_hour_square_meterResult *units.HeatFlux
    kilocalories_per_hour_square_meterResult, err = factory.FromDto(kilocalories_per_hour_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilocaloriePerHourSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_hour_square_meterResult.Convert(units.HeatFluxKilocaloriePerHourSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerHourSquareMeter = %v, want %v", converted, 100)
    }
    // Test PoundForcePerFootSecond conversion
    pounds_force_per_foot_secondDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxPoundForcePerFootSecond,
    }
    
    var pounds_force_per_foot_secondResult *units.HeatFlux
    pounds_force_per_foot_secondResult, err = factory.FromDto(pounds_force_per_foot_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerFootSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_foot_secondResult.Convert(units.HeatFluxPoundForcePerFootSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerFootSecond = %v, want %v", converted, 100)
    }
    // Test PoundPerSecondCubed conversion
    pounds_per_second_cubedDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxPoundPerSecondCubed,
    }
    
    var pounds_per_second_cubedResult *units.HeatFlux
    pounds_per_second_cubedResult, err = factory.FromDto(pounds_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_second_cubedResult.Convert(units.HeatFluxPoundPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test NanowattPerSquareMeter conversion
    nanowatts_per_square_meterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxNanowattPerSquareMeter,
    }
    
    var nanowatts_per_square_meterResult *units.HeatFlux
    nanowatts_per_square_meterResult, err = factory.FromDto(nanowatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NanowattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_square_meterResult.Convert(units.HeatFluxNanowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MicrowattPerSquareMeter conversion
    microwatts_per_square_meterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxMicrowattPerSquareMeter,
    }
    
    var microwatts_per_square_meterResult *units.HeatFlux
    microwatts_per_square_meterResult, err = factory.FromDto(microwatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrowattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_square_meterResult.Convert(units.HeatFluxMicrowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MilliwattPerSquareMeter conversion
    milliwatts_per_square_meterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxMilliwattPerSquareMeter,
    }
    
    var milliwatts_per_square_meterResult *units.HeatFlux
    milliwatts_per_square_meterResult, err = factory.FromDto(milliwatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_square_meterResult.Convert(units.HeatFluxMilliwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test CentiwattPerSquareMeter conversion
    centiwatts_per_square_meterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxCentiwattPerSquareMeter,
    }
    
    var centiwatts_per_square_meterResult *units.HeatFlux
    centiwatts_per_square_meterResult, err = factory.FromDto(centiwatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with CentiwattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiwatts_per_square_meterResult.Convert(units.HeatFluxCentiwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiwattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test DeciwattPerSquareMeter conversion
    deciwatts_per_square_meterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxDeciwattPerSquareMeter,
    }
    
    var deciwatts_per_square_meterResult *units.HeatFlux
    deciwatts_per_square_meterResult, err = factory.FromDto(deciwatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with DeciwattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_square_meterResult.Convert(units.HeatFluxDeciwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilowattPerSquareMeter conversion
    kilowatts_per_square_meterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxKilowattPerSquareMeter,
    }
    
    var kilowatts_per_square_meterResult *units.HeatFlux
    kilowatts_per_square_meterResult, err = factory.FromDto(kilowatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_square_meterResult.Convert(units.HeatFluxKilowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilocaloriePerSecondSquareCentimeter conversion
    kilocalories_per_second_square_centimeterDto := units.HeatFluxDto{
        Value: 100,
        Unit:  units.HeatFluxKilocaloriePerSecondSquareCentimeter,
    }
    
    var kilocalories_per_second_square_centimeterResult *units.HeatFlux
    kilocalories_per_second_square_centimeterResult, err = factory.FromDto(kilocalories_per_second_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilocaloriePerSecondSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_second_square_centimeterResult.Convert(units.HeatFluxKilocaloriePerSecondSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerSecondSquareCentimeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.HeatFluxDto{
        Value: 0,
        Unit:  units.HeatFluxWattPerSquareMeter,
    }
    
    var zeroResult *units.HeatFlux
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestHeatFluxFactory_FromDtoJSON(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "WattPerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "WattPerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.HeatFluxDto{
        Value: nanValue,
        Unit:  units.HeatFluxWattPerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with WattPerSquareMeter unit
    watts_per_square_meterJSON := []byte(`{"value": 100, "unit": "WattPerSquareMeter"}`)
    watts_per_square_meterResult, err := factory.FromDtoJSON(watts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_meterResult.Convert(units.HeatFluxWattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerSquareInch unit
    watts_per_square_inchJSON := []byte(`{"value": 100, "unit": "WattPerSquareInch"}`)
    watts_per_square_inchResult, err := factory.FromDtoJSON(watts_per_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_inchResult.Convert(units.HeatFluxWattPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerSquareFoot unit
    watts_per_square_footJSON := []byte(`{"value": 100, "unit": "WattPerSquareFoot"}`)
    watts_per_square_footResult, err := factory.FromDtoJSON(watts_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_footResult.Convert(units.HeatFluxWattPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerSecondSquareInch unit
    btus_per_second_square_inchJSON := []byte(`{"value": 100, "unit": "BtuPerSecondSquareInch"}`)
    btus_per_second_square_inchResult, err := factory.FromDtoJSON(btus_per_second_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerSecondSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_second_square_inchResult.Convert(units.HeatFluxBtuPerSecondSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerSecondSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerSecondSquareFoot unit
    btus_per_second_square_footJSON := []byte(`{"value": 100, "unit": "BtuPerSecondSquareFoot"}`)
    btus_per_second_square_footResult, err := factory.FromDtoJSON(btus_per_second_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerSecondSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_second_square_footResult.Convert(units.HeatFluxBtuPerSecondSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerSecondSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerMinuteSquareFoot unit
    btus_per_minute_square_footJSON := []byte(`{"value": 100, "unit": "BtuPerMinuteSquareFoot"}`)
    btus_per_minute_square_footResult, err := factory.FromDtoJSON(btus_per_minute_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerMinuteSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_minute_square_footResult.Convert(units.HeatFluxBtuPerMinuteSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerMinuteSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerHourSquareFoot unit
    btus_per_hour_square_footJSON := []byte(`{"value": 100, "unit": "BtuPerHourSquareFoot"}`)
    btus_per_hour_square_footResult, err := factory.FromDtoJSON(btus_per_hour_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerHourSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_hour_square_footResult.Convert(units.HeatFluxBtuPerHourSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerHourSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with CaloriePerSecondSquareCentimeter unit
    calories_per_second_square_centimeterJSON := []byte(`{"value": 100, "unit": "CaloriePerSecondSquareCentimeter"}`)
    calories_per_second_square_centimeterResult, err := factory.FromDtoJSON(calories_per_second_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CaloriePerSecondSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_second_square_centimeterResult.Convert(units.HeatFluxCaloriePerSecondSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerSecondSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilocaloriePerHourSquareMeter unit
    kilocalories_per_hour_square_meterJSON := []byte(`{"value": 100, "unit": "KilocaloriePerHourSquareMeter"}`)
    kilocalories_per_hour_square_meterResult, err := factory.FromDtoJSON(kilocalories_per_hour_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocaloriePerHourSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_hour_square_meterResult.Convert(units.HeatFluxKilocaloriePerHourSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerHourSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerFootSecond unit
    pounds_force_per_foot_secondJSON := []byte(`{"value": 100, "unit": "PoundForcePerFootSecond"}`)
    pounds_force_per_foot_secondResult, err := factory.FromDtoJSON(pounds_force_per_foot_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerFootSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_foot_secondResult.Convert(units.HeatFluxPoundForcePerFootSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerFootSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerSecondCubed unit
    pounds_per_second_cubedJSON := []byte(`{"value": 100, "unit": "PoundPerSecondCubed"}`)
    pounds_per_second_cubedResult, err := factory.FromDtoJSON(pounds_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_second_cubedResult.Convert(units.HeatFluxPoundPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with NanowattPerSquareMeter unit
    nanowatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "NanowattPerSquareMeter"}`)
    nanowatts_per_square_meterResult, err := factory.FromDtoJSON(nanowatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanowattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_square_meterResult.Convert(units.HeatFluxNanowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrowattPerSquareMeter unit
    microwatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "MicrowattPerSquareMeter"}`)
    microwatts_per_square_meterResult, err := factory.FromDtoJSON(microwatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrowattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_square_meterResult.Convert(units.HeatFluxMicrowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerSquareMeter unit
    milliwatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "MilliwattPerSquareMeter"}`)
    milliwatts_per_square_meterResult, err := factory.FromDtoJSON(milliwatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_square_meterResult.Convert(units.HeatFluxMilliwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with CentiwattPerSquareMeter unit
    centiwatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "CentiwattPerSquareMeter"}`)
    centiwatts_per_square_meterResult, err := factory.FromDtoJSON(centiwatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentiwattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiwatts_per_square_meterResult.Convert(units.HeatFluxCentiwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiwattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DeciwattPerSquareMeter unit
    deciwatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "DeciwattPerSquareMeter"}`)
    deciwatts_per_square_meterResult, err := factory.FromDtoJSON(deciwatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciwattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_square_meterResult.Convert(units.HeatFluxDeciwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerSquareMeter unit
    kilowatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilowattPerSquareMeter"}`)
    kilowatts_per_square_meterResult, err := factory.FromDtoJSON(kilowatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_square_meterResult.Convert(units.HeatFluxKilowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilocaloriePerSecondSquareCentimeter unit
    kilocalories_per_second_square_centimeterJSON := []byte(`{"value": 100, "unit": "KilocaloriePerSecondSquareCentimeter"}`)
    kilocalories_per_second_square_centimeterResult, err := factory.FromDtoJSON(kilocalories_per_second_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocaloriePerSecondSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_second_square_centimeterResult.Convert(units.HeatFluxKilocaloriePerSecondSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerSecondSquareCentimeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "WattPerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromWattsPerSquareMeter function
func TestHeatFluxFactory_FromWattsPerSquareMeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromWattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxWattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromWattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromWattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromWattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxWattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerSquareInch function
func TestHeatFluxFactory_FromWattsPerSquareInch(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerSquareInch(100)
    if err != nil {
        t.Errorf("FromWattsPerSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxWattPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerSquareInch(math.NaN())
    if err == nil {
        t.Error("FromWattsPerSquareInch() with NaN value should return error")
    }

    _, err = factory.FromWattsPerSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerSquareInch(0)
    if err != nil {
        t.Errorf("FromWattsPerSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxWattPerSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerSquareInch() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerSquareFoot function
func TestHeatFluxFactory_FromWattsPerSquareFoot(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerSquareFoot(100)
    if err != nil {
        t.Errorf("FromWattsPerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxWattPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromWattsPerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromWattsPerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerSquareFoot(0)
    if err != nil {
        t.Errorf("FromWattsPerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxWattPerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerSecondSquareInch function
func TestHeatFluxFactory_FromBtusPerSecondSquareInch(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerSecondSquareInch(100)
    if err != nil {
        t.Errorf("FromBtusPerSecondSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxBtuPerSecondSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerSecondSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerSecondSquareInch(math.NaN())
    if err == nil {
        t.Error("FromBtusPerSecondSquareInch() with NaN value should return error")
    }

    _, err = factory.FromBtusPerSecondSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerSecondSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerSecondSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerSecondSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerSecondSquareInch(0)
    if err != nil {
        t.Errorf("FromBtusPerSecondSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxBtuPerSecondSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerSecondSquareInch() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerSecondSquareFoot function
func TestHeatFluxFactory_FromBtusPerSecondSquareFoot(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerSecondSquareFoot(100)
    if err != nil {
        t.Errorf("FromBtusPerSecondSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxBtuPerSecondSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerSecondSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerSecondSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromBtusPerSecondSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromBtusPerSecondSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerSecondSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerSecondSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerSecondSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerSecondSquareFoot(0)
    if err != nil {
        t.Errorf("FromBtusPerSecondSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxBtuPerSecondSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerSecondSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerMinuteSquareFoot function
func TestHeatFluxFactory_FromBtusPerMinuteSquareFoot(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerMinuteSquareFoot(100)
    if err != nil {
        t.Errorf("FromBtusPerMinuteSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxBtuPerMinuteSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerMinuteSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerMinuteSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromBtusPerMinuteSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromBtusPerMinuteSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerMinuteSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerMinuteSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerMinuteSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerMinuteSquareFoot(0)
    if err != nil {
        t.Errorf("FromBtusPerMinuteSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxBtuPerMinuteSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerMinuteSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerHourSquareFoot function
func TestHeatFluxFactory_FromBtusPerHourSquareFoot(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerHourSquareFoot(100)
    if err != nil {
        t.Errorf("FromBtusPerHourSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxBtuPerHourSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerHourSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerHourSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromBtusPerHourSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromBtusPerHourSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerHourSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerHourSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerHourSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerHourSquareFoot(0)
    if err != nil {
        t.Errorf("FromBtusPerHourSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxBtuPerHourSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerHourSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromCaloriesPerSecondSquareCentimeter function
func TestHeatFluxFactory_FromCaloriesPerSecondSquareCentimeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCaloriesPerSecondSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromCaloriesPerSecondSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxCaloriePerSecondSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCaloriesPerSecondSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCaloriesPerSecondSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromCaloriesPerSecondSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromCaloriesPerSecondSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromCaloriesPerSecondSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromCaloriesPerSecondSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCaloriesPerSecondSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCaloriesPerSecondSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromCaloriesPerSecondSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxCaloriePerSecondSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCaloriesPerSecondSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocaloriesPerHourSquareMeter function
func TestHeatFluxFactory_FromKilocaloriesPerHourSquareMeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocaloriesPerHourSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilocaloriesPerHourSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxKilocaloriePerHourSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocaloriesPerHourSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocaloriesPerHourSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilocaloriesPerHourSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilocaloriesPerHourSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilocaloriesPerHourSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilocaloriesPerHourSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocaloriesPerHourSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocaloriesPerHourSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilocaloriesPerHourSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxKilocaloriePerHourSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocaloriesPerHourSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerFootSecond function
func TestHeatFluxFactory_FromPoundsForcePerFootSecond(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerFootSecond(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerFootSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxPoundForcePerFootSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerFootSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerFootSecond(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerFootSecond() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerFootSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerFootSecond() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerFootSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerFootSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerFootSecond(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerFootSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxPoundForcePerFootSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerFootSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerSecondCubed function
func TestHeatFluxFactory_FromPoundsPerSecondCubed(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromPoundsPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxPoundPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromPoundsPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxPoundPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowattsPerSquareMeter function
func TestHeatFluxFactory_FromNanowattsPerSquareMeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromNanowattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxNanowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromNanowattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromNanowattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanowattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromNanowattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromNanowattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxNanowattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowattsPerSquareMeter function
func TestHeatFluxFactory_FromMicrowattsPerSquareMeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMicrowattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxMicrowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMicrowattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMicrowattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrowattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMicrowattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxMicrowattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerSquareMeter function
func TestHeatFluxFactory_FromMilliwattsPerSquareMeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxMilliwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxMilliwattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentiwattsPerSquareMeter function
func TestHeatFluxFactory_FromCentiwattsPerSquareMeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentiwattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromCentiwattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxCentiwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentiwattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentiwattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromCentiwattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromCentiwattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromCentiwattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromCentiwattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentiwattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentiwattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromCentiwattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxCentiwattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentiwattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciwattsPerSquareMeter function
func TestHeatFluxFactory_FromDeciwattsPerSquareMeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciwattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromDeciwattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxDeciwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciwattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciwattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromDeciwattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromDeciwattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromDeciwattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromDeciwattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciwattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciwattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromDeciwattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxDeciwattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciwattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerSquareMeter function
func TestHeatFluxFactory_FromKilowattsPerSquareMeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilowattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxKilowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilowattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxKilowattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocaloriesPerSecondSquareCentimeter function
func TestHeatFluxFactory_FromKilocaloriesPerSecondSquareCentimeter(t *testing.T) {
    factory := units.HeatFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocaloriesPerSecondSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromKilocaloriesPerSecondSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatFluxKilocaloriePerSecondSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocaloriesPerSecondSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocaloriesPerSecondSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilocaloriesPerSecondSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilocaloriesPerSecondSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilocaloriesPerSecondSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilocaloriesPerSecondSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocaloriesPerSecondSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocaloriesPerSecondSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromKilocaloriesPerSecondSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatFluxKilocaloriePerSecondSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocaloriesPerSecondSquareCentimeter() with zero value = %v, want 0", converted)
    }
}

func TestHeatFluxToString(t *testing.T) {
	factory := units.HeatFluxFactory{}
	a, err := factory.CreateHeatFlux(45, units.HeatFluxWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.HeatFluxWattPerSquareMeter, 2)
	expected := "45.00 " + units.GetHeatFluxAbbreviation(units.HeatFluxWattPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.HeatFluxWattPerSquareMeter, -1)
	expected = "45 " + units.GetHeatFluxAbbreviation(units.HeatFluxWattPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestHeatFlux_EqualityAndComparison(t *testing.T) {
	factory := units.HeatFluxFactory{}
	a1, _ := factory.CreateHeatFlux(60, units.HeatFluxWattPerSquareMeter)
	a2, _ := factory.CreateHeatFlux(60, units.HeatFluxWattPerSquareMeter)
	a3, _ := factory.CreateHeatFlux(120, units.HeatFluxWattPerSquareMeter)

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

func TestHeatFlux_Arithmetic(t *testing.T) {
	factory := units.HeatFluxFactory{}
	a1, _ := factory.CreateHeatFlux(30, units.HeatFluxWattPerSquareMeter)
	a2, _ := factory.CreateHeatFlux(45, units.HeatFluxWattPerSquareMeter)

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


func TestGetHeatFluxAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.HeatFluxUnits
        want string
    }{
        {
            name: "WattPerSquareMeter abbreviation",
            unit: units.HeatFluxWattPerSquareMeter,
            want: "W/m²",
        },
        {
            name: "WattPerSquareInch abbreviation",
            unit: units.HeatFluxWattPerSquareInch,
            want: "W/in²",
        },
        {
            name: "WattPerSquareFoot abbreviation",
            unit: units.HeatFluxWattPerSquareFoot,
            want: "W/ft²",
        },
        {
            name: "BtuPerSecondSquareInch abbreviation",
            unit: units.HeatFluxBtuPerSecondSquareInch,
            want: "BTU/s·in²",
        },
        {
            name: "BtuPerSecondSquareFoot abbreviation",
            unit: units.HeatFluxBtuPerSecondSquareFoot,
            want: "BTU/s·ft²",
        },
        {
            name: "BtuPerMinuteSquareFoot abbreviation",
            unit: units.HeatFluxBtuPerMinuteSquareFoot,
            want: "BTU/min·ft²",
        },
        {
            name: "BtuPerHourSquareFoot abbreviation",
            unit: units.HeatFluxBtuPerHourSquareFoot,
            want: "BTU/h·ft²",
        },
        {
            name: "CaloriePerSecondSquareCentimeter abbreviation",
            unit: units.HeatFluxCaloriePerSecondSquareCentimeter,
            want: "cal/s·cm²",
        },
        {
            name: "KilocaloriePerHourSquareMeter abbreviation",
            unit: units.HeatFluxKilocaloriePerHourSquareMeter,
            want: "kcal/h·m²",
        },
        {
            name: "PoundForcePerFootSecond abbreviation",
            unit: units.HeatFluxPoundForcePerFootSecond,
            want: "lbf/(ft·s)",
        },
        {
            name: "PoundPerSecondCubed abbreviation",
            unit: units.HeatFluxPoundPerSecondCubed,
            want: "lb/s³",
        },
        {
            name: "NanowattPerSquareMeter abbreviation",
            unit: units.HeatFluxNanowattPerSquareMeter,
            want: "nW/m²",
        },
        {
            name: "MicrowattPerSquareMeter abbreviation",
            unit: units.HeatFluxMicrowattPerSquareMeter,
            want: "μW/m²",
        },
        {
            name: "MilliwattPerSquareMeter abbreviation",
            unit: units.HeatFluxMilliwattPerSquareMeter,
            want: "mW/m²",
        },
        {
            name: "CentiwattPerSquareMeter abbreviation",
            unit: units.HeatFluxCentiwattPerSquareMeter,
            want: "cW/m²",
        },
        {
            name: "DeciwattPerSquareMeter abbreviation",
            unit: units.HeatFluxDeciwattPerSquareMeter,
            want: "dW/m²",
        },
        {
            name: "KilowattPerSquareMeter abbreviation",
            unit: units.HeatFluxKilowattPerSquareMeter,
            want: "kW/m²",
        },
        {
            name: "KilocaloriePerSecondSquareCentimeter abbreviation",
            unit: units.HeatFluxKilocaloriePerSecondSquareCentimeter,
            want: "kcal/s·cm²",
        },
        {
            name: "invalid unit",
            unit: units.HeatFluxUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetHeatFluxAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetHeatFluxAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestHeatFlux_String(t *testing.T) {
    factory := units.HeatFluxFactory{}
    
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
            unit, err := factory.CreateHeatFlux(tt.value, units.HeatFluxWattPerSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("HeatFlux.String() = %v, want %v", got, tt.want)
            }
        })
    }
}