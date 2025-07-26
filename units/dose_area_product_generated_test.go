// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestDoseAreaProductDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "GraySquareMeter"}`
	
	factory := units.DoseAreaProductDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.DoseAreaProductGraySquareMeter {
		t.Errorf("expected unit %v, got %v", units.DoseAreaProductGraySquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "GraySquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestDoseAreaProductDto_ToJSON(t *testing.T) {
	dto := units.DoseAreaProductDto{
		Value: 45,
		Unit:  units.DoseAreaProductGraySquareMeter,
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
	if result["unit"].(string) != string(units.DoseAreaProductGraySquareMeter) {
		t.Errorf("expected unit %s, got %v", units.DoseAreaProductGraySquareMeter, result["unit"])
	}
}

func TestNewDoseAreaProduct_InvalidValue(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	// NaN value should return an error.
	_, err := factory.CreateDoseAreaProduct(math.NaN(), units.DoseAreaProductGraySquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateDoseAreaProduct(math.Inf(1), units.DoseAreaProductGraySquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestDoseAreaProductConversions(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateDoseAreaProduct(180, units.DoseAreaProductGraySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GraySquareMeters.
		// No expected conversion value provided for GraySquareMeters, verifying result is not NaN.
		result := a.GraySquareMeters()
		cacheResult := a.GraySquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GraySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to GraySquareDecimeters.
		// No expected conversion value provided for GraySquareDecimeters, verifying result is not NaN.
		result := a.GraySquareDecimeters()
		cacheResult := a.GraySquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GraySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to GraySquareCentimeters.
		// No expected conversion value provided for GraySquareCentimeters, verifying result is not NaN.
		result := a.GraySquareCentimeters()
		cacheResult := a.GraySquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GraySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to GraySquareMillimeters.
		// No expected conversion value provided for GraySquareMillimeters, verifying result is not NaN.
		result := a.GraySquareMillimeters()
		cacheResult := a.GraySquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GraySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to GraySquareMicrometers.
		// No expected conversion value provided for GraySquareMicrometers, verifying result is not NaN.
		result := a.GraySquareMicrometers()
		cacheResult := a.GraySquareMicrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GraySquareMicrometers returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareMeters.
		// No expected conversion value provided for MicrograySquareMeters, verifying result is not NaN.
		result := a.MicrograySquareMeters()
		cacheResult := a.MicrograySquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrograySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareMeters.
		// No expected conversion value provided for MilligraySquareMeters, verifying result is not NaN.
		result := a.MilligraySquareMeters()
		cacheResult := a.MilligraySquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligraySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareMeters.
		// No expected conversion value provided for CentigraySquareMeters, verifying result is not NaN.
		result := a.CentigraySquareMeters()
		cacheResult := a.CentigraySquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigraySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareMeters.
		// No expected conversion value provided for DecigraySquareMeters, verifying result is not NaN.
		result := a.DecigraySquareMeters()
		cacheResult := a.DecigraySquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigraySquareMeters returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareDecimeters.
		// No expected conversion value provided for MicrograySquareDecimeters, verifying result is not NaN.
		result := a.MicrograySquareDecimeters()
		cacheResult := a.MicrograySquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrograySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareDecimeters.
		// No expected conversion value provided for MilligraySquareDecimeters, verifying result is not NaN.
		result := a.MilligraySquareDecimeters()
		cacheResult := a.MilligraySquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligraySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareDecimeters.
		// No expected conversion value provided for CentigraySquareDecimeters, verifying result is not NaN.
		result := a.CentigraySquareDecimeters()
		cacheResult := a.CentigraySquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigraySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareDecimeters.
		// No expected conversion value provided for DecigraySquareDecimeters, verifying result is not NaN.
		result := a.DecigraySquareDecimeters()
		cacheResult := a.DecigraySquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigraySquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareCentimeters.
		// No expected conversion value provided for MicrograySquareCentimeters, verifying result is not NaN.
		result := a.MicrograySquareCentimeters()
		cacheResult := a.MicrograySquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrograySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareCentimeters.
		// No expected conversion value provided for MilligraySquareCentimeters, verifying result is not NaN.
		result := a.MilligraySquareCentimeters()
		cacheResult := a.MilligraySquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligraySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareCentimeters.
		// No expected conversion value provided for CentigraySquareCentimeters, verifying result is not NaN.
		result := a.CentigraySquareCentimeters()
		cacheResult := a.CentigraySquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigraySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareCentimeters.
		// No expected conversion value provided for DecigraySquareCentimeters, verifying result is not NaN.
		result := a.DecigraySquareCentimeters()
		cacheResult := a.DecigraySquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigraySquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareMillimeters.
		// No expected conversion value provided for MicrograySquareMillimeters, verifying result is not NaN.
		result := a.MicrograySquareMillimeters()
		cacheResult := a.MicrograySquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrograySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareMillimeters.
		// No expected conversion value provided for MilligraySquareMillimeters, verifying result is not NaN.
		result := a.MilligraySquareMillimeters()
		cacheResult := a.MilligraySquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligraySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareMillimeters.
		// No expected conversion value provided for CentigraySquareMillimeters, verifying result is not NaN.
		result := a.CentigraySquareMillimeters()
		cacheResult := a.CentigraySquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigraySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareMillimeters.
		// No expected conversion value provided for DecigraySquareMillimeters, verifying result is not NaN.
		result := a.DecigraySquareMillimeters()
		cacheResult := a.DecigraySquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigraySquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to MicrograySquareMicrometers.
		// No expected conversion value provided for MicrograySquareMicrometers, verifying result is not NaN.
		result := a.MicrograySquareMicrometers()
		cacheResult := a.MicrograySquareMicrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrograySquareMicrometers returned NaN")
		}
	}
	{
		// Test conversion to MilligraySquareMicrometers.
		// No expected conversion value provided for MilligraySquareMicrometers, verifying result is not NaN.
		result := a.MilligraySquareMicrometers()
		cacheResult := a.MilligraySquareMicrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligraySquareMicrometers returned NaN")
		}
	}
	{
		// Test conversion to CentigraySquareMicrometers.
		// No expected conversion value provided for CentigraySquareMicrometers, verifying result is not NaN.
		result := a.CentigraySquareMicrometers()
		cacheResult := a.CentigraySquareMicrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigraySquareMicrometers returned NaN")
		}
	}
	{
		// Test conversion to DecigraySquareMicrometers.
		// No expected conversion value provided for DecigraySquareMicrometers, verifying result is not NaN.
		result := a.DecigraySquareMicrometers()
		cacheResult := a.DecigraySquareMicrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigraySquareMicrometers returned NaN")
		}
	}
}

func TestDoseAreaProduct_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	a, err := factory.CreateDoseAreaProduct(90, units.DoseAreaProductGraySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.DoseAreaProductGraySquareMeter {
		t.Errorf("expected default unit GraySquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.DoseAreaProductGraySquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.DoseAreaProductDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.DoseAreaProductGraySquareMeter {
		t.Errorf("expected unit GraySquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestDoseAreaProductFactory_FromDto(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductGraySquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.DoseAreaProductDto{
        Value: math.NaN(),
        Unit:  units.DoseAreaProductGraySquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GraySquareMeter conversion
    gray_square_metersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductGraySquareMeter,
    }
    
    var gray_square_metersResult *units.DoseAreaProduct
    gray_square_metersResult, err = factory.FromDto(gray_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with GraySquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_metersResult.Convert(units.DoseAreaProductGraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareMeter = %v, want %v", converted, 100)
    }
    // Test GraySquareDecimeter conversion
    gray_square_decimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductGraySquareDecimeter,
    }
    
    var gray_square_decimetersResult *units.DoseAreaProduct
    gray_square_decimetersResult, err = factory.FromDto(gray_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with GraySquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_decimetersResult.Convert(units.DoseAreaProductGraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test GraySquareCentimeter conversion
    gray_square_centimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductGraySquareCentimeter,
    }
    
    var gray_square_centimetersResult *units.DoseAreaProduct
    gray_square_centimetersResult, err = factory.FromDto(gray_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with GraySquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_centimetersResult.Convert(units.DoseAreaProductGraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test GraySquareMillimeter conversion
    gray_square_millimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductGraySquareMillimeter,
    }
    
    var gray_square_millimetersResult *units.DoseAreaProduct
    gray_square_millimetersResult, err = factory.FromDto(gray_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with GraySquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_millimetersResult.Convert(units.DoseAreaProductGraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test GraySquareMicrometer conversion
    gray_square_micrometersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductGraySquareMicrometer,
    }
    
    var gray_square_micrometersResult *units.DoseAreaProduct
    gray_square_micrometersResult, err = factory.FromDto(gray_square_micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with GraySquareMicrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_micrometersResult.Convert(units.DoseAreaProductGraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareMicrometer = %v, want %v", converted, 100)
    }
    // Test MicrograySquareMeter conversion
    microgray_square_metersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMicrograySquareMeter,
    }
    
    var microgray_square_metersResult *units.DoseAreaProduct
    microgray_square_metersResult, err = factory.FromDto(microgray_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with MicrograySquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_metersResult.Convert(units.DoseAreaProductMicrograySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareMeter = %v, want %v", converted, 100)
    }
    // Test MilligraySquareMeter conversion
    milligray_square_metersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMilligraySquareMeter,
    }
    
    var milligray_square_metersResult *units.DoseAreaProduct
    milligray_square_metersResult, err = factory.FromDto(milligray_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligraySquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_metersResult.Convert(units.DoseAreaProductMilligraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareMeter = %v, want %v", converted, 100)
    }
    // Test CentigraySquareMeter conversion
    centigray_square_metersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductCentigraySquareMeter,
    }
    
    var centigray_square_metersResult *units.DoseAreaProduct
    centigray_square_metersResult, err = factory.FromDto(centigray_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with CentigraySquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_metersResult.Convert(units.DoseAreaProductCentigraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareMeter = %v, want %v", converted, 100)
    }
    // Test DecigraySquareMeter conversion
    decigray_square_metersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductDecigraySquareMeter,
    }
    
    var decigray_square_metersResult *units.DoseAreaProduct
    decigray_square_metersResult, err = factory.FromDto(decigray_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with DecigraySquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_metersResult.Convert(units.DoseAreaProductDecigraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareMeter = %v, want %v", converted, 100)
    }
    // Test MicrograySquareDecimeter conversion
    microgray_square_decimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMicrograySquareDecimeter,
    }
    
    var microgray_square_decimetersResult *units.DoseAreaProduct
    microgray_square_decimetersResult, err = factory.FromDto(microgray_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MicrograySquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_decimetersResult.Convert(units.DoseAreaProductMicrograySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test MilligraySquareDecimeter conversion
    milligray_square_decimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMilligraySquareDecimeter,
    }
    
    var milligray_square_decimetersResult *units.DoseAreaProduct
    milligray_square_decimetersResult, err = factory.FromDto(milligray_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligraySquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_decimetersResult.Convert(units.DoseAreaProductMilligraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test CentigraySquareDecimeter conversion
    centigray_square_decimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductCentigraySquareDecimeter,
    }
    
    var centigray_square_decimetersResult *units.DoseAreaProduct
    centigray_square_decimetersResult, err = factory.FromDto(centigray_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with CentigraySquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_decimetersResult.Convert(units.DoseAreaProductCentigraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test DecigraySquareDecimeter conversion
    decigray_square_decimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductDecigraySquareDecimeter,
    }
    
    var decigray_square_decimetersResult *units.DoseAreaProduct
    decigray_square_decimetersResult, err = factory.FromDto(decigray_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with DecigraySquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_decimetersResult.Convert(units.DoseAreaProductDecigraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test MicrograySquareCentimeter conversion
    microgray_square_centimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMicrograySquareCentimeter,
    }
    
    var microgray_square_centimetersResult *units.DoseAreaProduct
    microgray_square_centimetersResult, err = factory.FromDto(microgray_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MicrograySquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_centimetersResult.Convert(units.DoseAreaProductMicrograySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test MilligraySquareCentimeter conversion
    milligray_square_centimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMilligraySquareCentimeter,
    }
    
    var milligray_square_centimetersResult *units.DoseAreaProduct
    milligray_square_centimetersResult, err = factory.FromDto(milligray_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligraySquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_centimetersResult.Convert(units.DoseAreaProductMilligraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test CentigraySquareCentimeter conversion
    centigray_square_centimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductCentigraySquareCentimeter,
    }
    
    var centigray_square_centimetersResult *units.DoseAreaProduct
    centigray_square_centimetersResult, err = factory.FromDto(centigray_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with CentigraySquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_centimetersResult.Convert(units.DoseAreaProductCentigraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test DecigraySquareCentimeter conversion
    decigray_square_centimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductDecigraySquareCentimeter,
    }
    
    var decigray_square_centimetersResult *units.DoseAreaProduct
    decigray_square_centimetersResult, err = factory.FromDto(decigray_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with DecigraySquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_centimetersResult.Convert(units.DoseAreaProductDecigraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test MicrograySquareMillimeter conversion
    microgray_square_millimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMicrograySquareMillimeter,
    }
    
    var microgray_square_millimetersResult *units.DoseAreaProduct
    microgray_square_millimetersResult, err = factory.FromDto(microgray_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MicrograySquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_millimetersResult.Convert(units.DoseAreaProductMicrograySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test MilligraySquareMillimeter conversion
    milligray_square_millimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMilligraySquareMillimeter,
    }
    
    var milligray_square_millimetersResult *units.DoseAreaProduct
    milligray_square_millimetersResult, err = factory.FromDto(milligray_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligraySquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_millimetersResult.Convert(units.DoseAreaProductMilligraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test CentigraySquareMillimeter conversion
    centigray_square_millimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductCentigraySquareMillimeter,
    }
    
    var centigray_square_millimetersResult *units.DoseAreaProduct
    centigray_square_millimetersResult, err = factory.FromDto(centigray_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with CentigraySquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_millimetersResult.Convert(units.DoseAreaProductCentigraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test DecigraySquareMillimeter conversion
    decigray_square_millimetersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductDecigraySquareMillimeter,
    }
    
    var decigray_square_millimetersResult *units.DoseAreaProduct
    decigray_square_millimetersResult, err = factory.FromDto(decigray_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with DecigraySquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_millimetersResult.Convert(units.DoseAreaProductDecigraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test MicrograySquareMicrometer conversion
    microgray_square_micrometersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMicrograySquareMicrometer,
    }
    
    var microgray_square_micrometersResult *units.DoseAreaProduct
    microgray_square_micrometersResult, err = factory.FromDto(microgray_square_micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with MicrograySquareMicrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_micrometersResult.Convert(units.DoseAreaProductMicrograySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareMicrometer = %v, want %v", converted, 100)
    }
    // Test MilligraySquareMicrometer conversion
    milligray_square_micrometersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductMilligraySquareMicrometer,
    }
    
    var milligray_square_micrometersResult *units.DoseAreaProduct
    milligray_square_micrometersResult, err = factory.FromDto(milligray_square_micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with MilligraySquareMicrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_micrometersResult.Convert(units.DoseAreaProductMilligraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareMicrometer = %v, want %v", converted, 100)
    }
    // Test CentigraySquareMicrometer conversion
    centigray_square_micrometersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductCentigraySquareMicrometer,
    }
    
    var centigray_square_micrometersResult *units.DoseAreaProduct
    centigray_square_micrometersResult, err = factory.FromDto(centigray_square_micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with CentigraySquareMicrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_micrometersResult.Convert(units.DoseAreaProductCentigraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareMicrometer = %v, want %v", converted, 100)
    }
    // Test DecigraySquareMicrometer conversion
    decigray_square_micrometersDto := units.DoseAreaProductDto{
        Value: 100,
        Unit:  units.DoseAreaProductDecigraySquareMicrometer,
    }
    
    var decigray_square_micrometersResult *units.DoseAreaProduct
    decigray_square_micrometersResult, err = factory.FromDto(decigray_square_micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with DecigraySquareMicrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_micrometersResult.Convert(units.DoseAreaProductDecigraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareMicrometer = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.DoseAreaProductDto{
        Value: 0,
        Unit:  units.DoseAreaProductGraySquareMeter,
    }
    
    var zeroResult *units.DoseAreaProduct
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestDoseAreaProductFactory_FromDtoJSON(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "GraySquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "GraySquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.DoseAreaProductDto{
        Value: nanValue,
        Unit:  units.DoseAreaProductGraySquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with GraySquareMeter unit
    gray_square_metersJSON := []byte(`{"value": 100, "unit": "GraySquareMeter"}`)
    gray_square_metersResult, err := factory.FromDtoJSON(gray_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GraySquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_metersResult.Convert(units.DoseAreaProductGraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GraySquareDecimeter unit
    gray_square_decimetersJSON := []byte(`{"value": 100, "unit": "GraySquareDecimeter"}`)
    gray_square_decimetersResult, err := factory.FromDtoJSON(gray_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GraySquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_decimetersResult.Convert(units.DoseAreaProductGraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GraySquareCentimeter unit
    gray_square_centimetersJSON := []byte(`{"value": 100, "unit": "GraySquareCentimeter"}`)
    gray_square_centimetersResult, err := factory.FromDtoJSON(gray_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GraySquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_centimetersResult.Convert(units.DoseAreaProductGraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GraySquareMillimeter unit
    gray_square_millimetersJSON := []byte(`{"value": 100, "unit": "GraySquareMillimeter"}`)
    gray_square_millimetersResult, err := factory.FromDtoJSON(gray_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GraySquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_millimetersResult.Convert(units.DoseAreaProductGraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GraySquareMicrometer unit
    gray_square_micrometersJSON := []byte(`{"value": 100, "unit": "GraySquareMicrometer"}`)
    gray_square_micrometersResult, err := factory.FromDtoJSON(gray_square_micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GraySquareMicrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gray_square_micrometersResult.Convert(units.DoseAreaProductGraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GraySquareMicrometer = %v, want %v", converted, 100)
    }
    // Test JSON with MicrograySquareMeter unit
    microgray_square_metersJSON := []byte(`{"value": 100, "unit": "MicrograySquareMeter"}`)
    microgray_square_metersResult, err := factory.FromDtoJSON(microgray_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrograySquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_metersResult.Convert(units.DoseAreaProductMicrograySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligraySquareMeter unit
    milligray_square_metersJSON := []byte(`{"value": 100, "unit": "MilligraySquareMeter"}`)
    milligray_square_metersResult, err := factory.FromDtoJSON(milligray_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligraySquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_metersResult.Convert(units.DoseAreaProductMilligraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigraySquareMeter unit
    centigray_square_metersJSON := []byte(`{"value": 100, "unit": "CentigraySquareMeter"}`)
    centigray_square_metersResult, err := factory.FromDtoJSON(centigray_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigraySquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_metersResult.Convert(units.DoseAreaProductCentigraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigraySquareMeter unit
    decigray_square_metersJSON := []byte(`{"value": 100, "unit": "DecigraySquareMeter"}`)
    decigray_square_metersResult, err := factory.FromDtoJSON(decigray_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigraySquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_metersResult.Convert(units.DoseAreaProductDecigraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrograySquareDecimeter unit
    microgray_square_decimetersJSON := []byte(`{"value": 100, "unit": "MicrograySquareDecimeter"}`)
    microgray_square_decimetersResult, err := factory.FromDtoJSON(microgray_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrograySquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_decimetersResult.Convert(units.DoseAreaProductMicrograySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligraySquareDecimeter unit
    milligray_square_decimetersJSON := []byte(`{"value": 100, "unit": "MilligraySquareDecimeter"}`)
    milligray_square_decimetersResult, err := factory.FromDtoJSON(milligray_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligraySquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_decimetersResult.Convert(units.DoseAreaProductMilligraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigraySquareDecimeter unit
    centigray_square_decimetersJSON := []byte(`{"value": 100, "unit": "CentigraySquareDecimeter"}`)
    centigray_square_decimetersResult, err := factory.FromDtoJSON(centigray_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigraySquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_decimetersResult.Convert(units.DoseAreaProductCentigraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigraySquareDecimeter unit
    decigray_square_decimetersJSON := []byte(`{"value": 100, "unit": "DecigraySquareDecimeter"}`)
    decigray_square_decimetersResult, err := factory.FromDtoJSON(decigray_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigraySquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_decimetersResult.Convert(units.DoseAreaProductDecigraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrograySquareCentimeter unit
    microgray_square_centimetersJSON := []byte(`{"value": 100, "unit": "MicrograySquareCentimeter"}`)
    microgray_square_centimetersResult, err := factory.FromDtoJSON(microgray_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrograySquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_centimetersResult.Convert(units.DoseAreaProductMicrograySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligraySquareCentimeter unit
    milligray_square_centimetersJSON := []byte(`{"value": 100, "unit": "MilligraySquareCentimeter"}`)
    milligray_square_centimetersResult, err := factory.FromDtoJSON(milligray_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligraySquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_centimetersResult.Convert(units.DoseAreaProductMilligraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigraySquareCentimeter unit
    centigray_square_centimetersJSON := []byte(`{"value": 100, "unit": "CentigraySquareCentimeter"}`)
    centigray_square_centimetersResult, err := factory.FromDtoJSON(centigray_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigraySquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_centimetersResult.Convert(units.DoseAreaProductCentigraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigraySquareCentimeter unit
    decigray_square_centimetersJSON := []byte(`{"value": 100, "unit": "DecigraySquareCentimeter"}`)
    decigray_square_centimetersResult, err := factory.FromDtoJSON(decigray_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigraySquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_centimetersResult.Convert(units.DoseAreaProductDecigraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrograySquareMillimeter unit
    microgray_square_millimetersJSON := []byte(`{"value": 100, "unit": "MicrograySquareMillimeter"}`)
    microgray_square_millimetersResult, err := factory.FromDtoJSON(microgray_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrograySquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_millimetersResult.Convert(units.DoseAreaProductMicrograySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligraySquareMillimeter unit
    milligray_square_millimetersJSON := []byte(`{"value": 100, "unit": "MilligraySquareMillimeter"}`)
    milligray_square_millimetersResult, err := factory.FromDtoJSON(milligray_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligraySquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_millimetersResult.Convert(units.DoseAreaProductMilligraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigraySquareMillimeter unit
    centigray_square_millimetersJSON := []byte(`{"value": 100, "unit": "CentigraySquareMillimeter"}`)
    centigray_square_millimetersResult, err := factory.FromDtoJSON(centigray_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigraySquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_millimetersResult.Convert(units.DoseAreaProductCentigraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigraySquareMillimeter unit
    decigray_square_millimetersJSON := []byte(`{"value": 100, "unit": "DecigraySquareMillimeter"}`)
    decigray_square_millimetersResult, err := factory.FromDtoJSON(decigray_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigraySquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_millimetersResult.Convert(units.DoseAreaProductDecigraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrograySquareMicrometer unit
    microgray_square_micrometersJSON := []byte(`{"value": 100, "unit": "MicrograySquareMicrometer"}`)
    microgray_square_micrometersResult, err := factory.FromDtoJSON(microgray_square_micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrograySquareMicrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microgray_square_micrometersResult.Convert(units.DoseAreaProductMicrograySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrograySquareMicrometer = %v, want %v", converted, 100)
    }
    // Test JSON with MilligraySquareMicrometer unit
    milligray_square_micrometersJSON := []byte(`{"value": 100, "unit": "MilligraySquareMicrometer"}`)
    milligray_square_micrometersResult, err := factory.FromDtoJSON(milligray_square_micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligraySquareMicrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligray_square_micrometersResult.Convert(units.DoseAreaProductMilligraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligraySquareMicrometer = %v, want %v", converted, 100)
    }
    // Test JSON with CentigraySquareMicrometer unit
    centigray_square_micrometersJSON := []byte(`{"value": 100, "unit": "CentigraySquareMicrometer"}`)
    centigray_square_micrometersResult, err := factory.FromDtoJSON(centigray_square_micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigraySquareMicrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigray_square_micrometersResult.Convert(units.DoseAreaProductCentigraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigraySquareMicrometer = %v, want %v", converted, 100)
    }
    // Test JSON with DecigraySquareMicrometer unit
    decigray_square_micrometersJSON := []byte(`{"value": 100, "unit": "DecigraySquareMicrometer"}`)
    decigray_square_micrometersResult, err := factory.FromDtoJSON(decigray_square_micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigraySquareMicrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigray_square_micrometersResult.Convert(units.DoseAreaProductDecigraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigraySquareMicrometer = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "GraySquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGraySquareMeters function
func TestDoseAreaProductFactory_FromGraySquareMeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGraySquareMeters(100)
    if err != nil {
        t.Errorf("FromGraySquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductGraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGraySquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGraySquareMeters(math.NaN())
    if err == nil {
        t.Error("FromGraySquareMeters() with NaN value should return error")
    }

    _, err = factory.FromGraySquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromGraySquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromGraySquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGraySquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGraySquareMeters(0)
    if err != nil {
        t.Errorf("FromGraySquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductGraySquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGraySquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGraySquareDecimeters function
func TestDoseAreaProductFactory_FromGraySquareDecimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGraySquareDecimeters(100)
    if err != nil {
        t.Errorf("FromGraySquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductGraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGraySquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGraySquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromGraySquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromGraySquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromGraySquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromGraySquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGraySquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGraySquareDecimeters(0)
    if err != nil {
        t.Errorf("FromGraySquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductGraySquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGraySquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGraySquareCentimeters function
func TestDoseAreaProductFactory_FromGraySquareCentimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGraySquareCentimeters(100)
    if err != nil {
        t.Errorf("FromGraySquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductGraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGraySquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGraySquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromGraySquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromGraySquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromGraySquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromGraySquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGraySquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGraySquareCentimeters(0)
    if err != nil {
        t.Errorf("FromGraySquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductGraySquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGraySquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGraySquareMillimeters function
func TestDoseAreaProductFactory_FromGraySquareMillimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGraySquareMillimeters(100)
    if err != nil {
        t.Errorf("FromGraySquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductGraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGraySquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGraySquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromGraySquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromGraySquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromGraySquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromGraySquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGraySquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGraySquareMillimeters(0)
    if err != nil {
        t.Errorf("FromGraySquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductGraySquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGraySquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGraySquareMicrometers function
func TestDoseAreaProductFactory_FromGraySquareMicrometers(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGraySquareMicrometers(100)
    if err != nil {
        t.Errorf("FromGraySquareMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductGraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGraySquareMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGraySquareMicrometers(math.NaN())
    if err == nil {
        t.Error("FromGraySquareMicrometers() with NaN value should return error")
    }

    _, err = factory.FromGraySquareMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromGraySquareMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromGraySquareMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromGraySquareMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGraySquareMicrometers(0)
    if err != nil {
        t.Errorf("FromGraySquareMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductGraySquareMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGraySquareMicrometers() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrograySquareMeters function
func TestDoseAreaProductFactory_FromMicrograySquareMeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrograySquareMeters(100)
    if err != nil {
        t.Errorf("FromMicrograySquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMicrograySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrograySquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrograySquareMeters(math.NaN())
    if err == nil {
        t.Error("FromMicrograySquareMeters() with NaN value should return error")
    }

    _, err = factory.FromMicrograySquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMicrograySquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromMicrograySquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrograySquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrograySquareMeters(0)
    if err != nil {
        t.Errorf("FromMicrograySquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMicrograySquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrograySquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligraySquareMeters function
func TestDoseAreaProductFactory_FromMilligraySquareMeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligraySquareMeters(100)
    if err != nil {
        t.Errorf("FromMilligraySquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMilligraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligraySquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligraySquareMeters(math.NaN())
    if err == nil {
        t.Error("FromMilligraySquareMeters() with NaN value should return error")
    }

    _, err = factory.FromMilligraySquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilligraySquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromMilligraySquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligraySquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligraySquareMeters(0)
    if err != nil {
        t.Errorf("FromMilligraySquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMilligraySquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligraySquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigraySquareMeters function
func TestDoseAreaProductFactory_FromCentigraySquareMeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigraySquareMeters(100)
    if err != nil {
        t.Errorf("FromCentigraySquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductCentigraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigraySquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigraySquareMeters(math.NaN())
    if err == nil {
        t.Error("FromCentigraySquareMeters() with NaN value should return error")
    }

    _, err = factory.FromCentigraySquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromCentigraySquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromCentigraySquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigraySquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigraySquareMeters(0)
    if err != nil {
        t.Errorf("FromCentigraySquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductCentigraySquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigraySquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigraySquareMeters function
func TestDoseAreaProductFactory_FromDecigraySquareMeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigraySquareMeters(100)
    if err != nil {
        t.Errorf("FromDecigraySquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductDecigraySquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigraySquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigraySquareMeters(math.NaN())
    if err == nil {
        t.Error("FromDecigraySquareMeters() with NaN value should return error")
    }

    _, err = factory.FromDecigraySquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromDecigraySquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromDecigraySquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigraySquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigraySquareMeters(0)
    if err != nil {
        t.Errorf("FromDecigraySquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductDecigraySquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigraySquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrograySquareDecimeters function
func TestDoseAreaProductFactory_FromMicrograySquareDecimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrograySquareDecimeters(100)
    if err != nil {
        t.Errorf("FromMicrograySquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMicrograySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrograySquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrograySquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromMicrograySquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromMicrograySquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMicrograySquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromMicrograySquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrograySquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrograySquareDecimeters(0)
    if err != nil {
        t.Errorf("FromMicrograySquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMicrograySquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrograySquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligraySquareDecimeters function
func TestDoseAreaProductFactory_FromMilligraySquareDecimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligraySquareDecimeters(100)
    if err != nil {
        t.Errorf("FromMilligraySquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMilligraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligraySquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligraySquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromMilligraySquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromMilligraySquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilligraySquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromMilligraySquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligraySquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligraySquareDecimeters(0)
    if err != nil {
        t.Errorf("FromMilligraySquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMilligraySquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligraySquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigraySquareDecimeters function
func TestDoseAreaProductFactory_FromCentigraySquareDecimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigraySquareDecimeters(100)
    if err != nil {
        t.Errorf("FromCentigraySquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductCentigraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigraySquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigraySquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromCentigraySquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromCentigraySquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromCentigraySquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromCentigraySquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigraySquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigraySquareDecimeters(0)
    if err != nil {
        t.Errorf("FromCentigraySquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductCentigraySquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigraySquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigraySquareDecimeters function
func TestDoseAreaProductFactory_FromDecigraySquareDecimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigraySquareDecimeters(100)
    if err != nil {
        t.Errorf("FromDecigraySquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductDecigraySquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigraySquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigraySquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromDecigraySquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromDecigraySquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromDecigraySquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromDecigraySquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigraySquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigraySquareDecimeters(0)
    if err != nil {
        t.Errorf("FromDecigraySquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductDecigraySquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigraySquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrograySquareCentimeters function
func TestDoseAreaProductFactory_FromMicrograySquareCentimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrograySquareCentimeters(100)
    if err != nil {
        t.Errorf("FromMicrograySquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMicrograySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrograySquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrograySquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromMicrograySquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromMicrograySquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMicrograySquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromMicrograySquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrograySquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrograySquareCentimeters(0)
    if err != nil {
        t.Errorf("FromMicrograySquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMicrograySquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrograySquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligraySquareCentimeters function
func TestDoseAreaProductFactory_FromMilligraySquareCentimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligraySquareCentimeters(100)
    if err != nil {
        t.Errorf("FromMilligraySquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMilligraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligraySquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligraySquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromMilligraySquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromMilligraySquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilligraySquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromMilligraySquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligraySquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligraySquareCentimeters(0)
    if err != nil {
        t.Errorf("FromMilligraySquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMilligraySquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligraySquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigraySquareCentimeters function
func TestDoseAreaProductFactory_FromCentigraySquareCentimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigraySquareCentimeters(100)
    if err != nil {
        t.Errorf("FromCentigraySquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductCentigraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigraySquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigraySquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromCentigraySquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromCentigraySquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromCentigraySquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromCentigraySquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigraySquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigraySquareCentimeters(0)
    if err != nil {
        t.Errorf("FromCentigraySquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductCentigraySquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigraySquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigraySquareCentimeters function
func TestDoseAreaProductFactory_FromDecigraySquareCentimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigraySquareCentimeters(100)
    if err != nil {
        t.Errorf("FromDecigraySquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductDecigraySquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigraySquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigraySquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromDecigraySquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromDecigraySquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromDecigraySquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromDecigraySquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigraySquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigraySquareCentimeters(0)
    if err != nil {
        t.Errorf("FromDecigraySquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductDecigraySquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigraySquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrograySquareMillimeters function
func TestDoseAreaProductFactory_FromMicrograySquareMillimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrograySquareMillimeters(100)
    if err != nil {
        t.Errorf("FromMicrograySquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMicrograySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrograySquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrograySquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromMicrograySquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromMicrograySquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMicrograySquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromMicrograySquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrograySquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrograySquareMillimeters(0)
    if err != nil {
        t.Errorf("FromMicrograySquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMicrograySquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrograySquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligraySquareMillimeters function
func TestDoseAreaProductFactory_FromMilligraySquareMillimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligraySquareMillimeters(100)
    if err != nil {
        t.Errorf("FromMilligraySquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMilligraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligraySquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligraySquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromMilligraySquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromMilligraySquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMilligraySquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromMilligraySquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligraySquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligraySquareMillimeters(0)
    if err != nil {
        t.Errorf("FromMilligraySquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMilligraySquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligraySquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigraySquareMillimeters function
func TestDoseAreaProductFactory_FromCentigraySquareMillimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigraySquareMillimeters(100)
    if err != nil {
        t.Errorf("FromCentigraySquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductCentigraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigraySquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigraySquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromCentigraySquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromCentigraySquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromCentigraySquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromCentigraySquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigraySquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigraySquareMillimeters(0)
    if err != nil {
        t.Errorf("FromCentigraySquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductCentigraySquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigraySquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigraySquareMillimeters function
func TestDoseAreaProductFactory_FromDecigraySquareMillimeters(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigraySquareMillimeters(100)
    if err != nil {
        t.Errorf("FromDecigraySquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductDecigraySquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigraySquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigraySquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromDecigraySquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromDecigraySquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromDecigraySquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromDecigraySquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigraySquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigraySquareMillimeters(0)
    if err != nil {
        t.Errorf("FromDecigraySquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductDecigraySquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigraySquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrograySquareMicrometers function
func TestDoseAreaProductFactory_FromMicrograySquareMicrometers(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrograySquareMicrometers(100)
    if err != nil {
        t.Errorf("FromMicrograySquareMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMicrograySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrograySquareMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrograySquareMicrometers(math.NaN())
    if err == nil {
        t.Error("FromMicrograySquareMicrometers() with NaN value should return error")
    }

    _, err = factory.FromMicrograySquareMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromMicrograySquareMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromMicrograySquareMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrograySquareMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrograySquareMicrometers(0)
    if err != nil {
        t.Errorf("FromMicrograySquareMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMicrograySquareMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrograySquareMicrometers() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligraySquareMicrometers function
func TestDoseAreaProductFactory_FromMilligraySquareMicrometers(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligraySquareMicrometers(100)
    if err != nil {
        t.Errorf("FromMilligraySquareMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductMilligraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligraySquareMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligraySquareMicrometers(math.NaN())
    if err == nil {
        t.Error("FromMilligraySquareMicrometers() with NaN value should return error")
    }

    _, err = factory.FromMilligraySquareMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromMilligraySquareMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromMilligraySquareMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligraySquareMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligraySquareMicrometers(0)
    if err != nil {
        t.Errorf("FromMilligraySquareMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductMilligraySquareMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligraySquareMicrometers() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigraySquareMicrometers function
func TestDoseAreaProductFactory_FromCentigraySquareMicrometers(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigraySquareMicrometers(100)
    if err != nil {
        t.Errorf("FromCentigraySquareMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductCentigraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigraySquareMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigraySquareMicrometers(math.NaN())
    if err == nil {
        t.Error("FromCentigraySquareMicrometers() with NaN value should return error")
    }

    _, err = factory.FromCentigraySquareMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromCentigraySquareMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromCentigraySquareMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigraySquareMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigraySquareMicrometers(0)
    if err != nil {
        t.Errorf("FromCentigraySquareMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductCentigraySquareMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigraySquareMicrometers() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigraySquareMicrometers function
func TestDoseAreaProductFactory_FromDecigraySquareMicrometers(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigraySquareMicrometers(100)
    if err != nil {
        t.Errorf("FromDecigraySquareMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DoseAreaProductDecigraySquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigraySquareMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigraySquareMicrometers(math.NaN())
    if err == nil {
        t.Error("FromDecigraySquareMicrometers() with NaN value should return error")
    }

    _, err = factory.FromDecigraySquareMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromDecigraySquareMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromDecigraySquareMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigraySquareMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigraySquareMicrometers(0)
    if err != nil {
        t.Errorf("FromDecigraySquareMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DoseAreaProductDecigraySquareMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigraySquareMicrometers() with zero value = %v, want 0", converted)
    }
}

func TestDoseAreaProductToString(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	a, err := factory.CreateDoseAreaProduct(45, units.DoseAreaProductGraySquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.DoseAreaProductGraySquareMeter, 2)
	expected := "45.00 " + units.GetDoseAreaProductAbbreviation(units.DoseAreaProductGraySquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.DoseAreaProductGraySquareMeter, -1)
	expected = "45 " + units.GetDoseAreaProductAbbreviation(units.DoseAreaProductGraySquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestDoseAreaProduct_EqualityAndComparison(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	a1, _ := factory.CreateDoseAreaProduct(60, units.DoseAreaProductGraySquareMeter)
	a2, _ := factory.CreateDoseAreaProduct(60, units.DoseAreaProductGraySquareMeter)
	a3, _ := factory.CreateDoseAreaProduct(120, units.DoseAreaProductGraySquareMeter)

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

func TestDoseAreaProduct_Arithmetic(t *testing.T) {
	factory := units.DoseAreaProductFactory{}
	a1, _ := factory.CreateDoseAreaProduct(30, units.DoseAreaProductGraySquareMeter)
	a2, _ := factory.CreateDoseAreaProduct(45, units.DoseAreaProductGraySquareMeter)

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


func TestGetDoseAreaProductAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.DoseAreaProductUnits
        want string
    }{
        {
            name: "GraySquareMeter abbreviation",
            unit: units.DoseAreaProductGraySquareMeter,
            want: "Gy·m²",
        },
        {
            name: "GraySquareDecimeter abbreviation",
            unit: units.DoseAreaProductGraySquareDecimeter,
            want: "Gy·dm²",
        },
        {
            name: "GraySquareCentimeter abbreviation",
            unit: units.DoseAreaProductGraySquareCentimeter,
            want: "Gy·cm²",
        },
        {
            name: "GraySquareMillimeter abbreviation",
            unit: units.DoseAreaProductGraySquareMillimeter,
            want: "Gy·mm²",
        },
        {
            name: "GraySquareMicrometer abbreviation",
            unit: units.DoseAreaProductGraySquareMicrometer,
            want: "Gy·μm²",
        },
        {
            name: "MicrograySquareMeter abbreviation",
            unit: units.DoseAreaProductMicrograySquareMeter,
            want: "μGy·m²",
        },
        {
            name: "MilligraySquareMeter abbreviation",
            unit: units.DoseAreaProductMilligraySquareMeter,
            want: "mGy·m²",
        },
        {
            name: "CentigraySquareMeter abbreviation",
            unit: units.DoseAreaProductCentigraySquareMeter,
            want: "cGy·m²",
        },
        {
            name: "DecigraySquareMeter abbreviation",
            unit: units.DoseAreaProductDecigraySquareMeter,
            want: "dGy·m²",
        },
        {
            name: "MicrograySquareDecimeter abbreviation",
            unit: units.DoseAreaProductMicrograySquareDecimeter,
            want: "μGy·dm²",
        },
        {
            name: "MilligraySquareDecimeter abbreviation",
            unit: units.DoseAreaProductMilligraySquareDecimeter,
            want: "mGy·dm²",
        },
        {
            name: "CentigraySquareDecimeter abbreviation",
            unit: units.DoseAreaProductCentigraySquareDecimeter,
            want: "cGy·dm²",
        },
        {
            name: "DecigraySquareDecimeter abbreviation",
            unit: units.DoseAreaProductDecigraySquareDecimeter,
            want: "dGy·dm²",
        },
        {
            name: "MicrograySquareCentimeter abbreviation",
            unit: units.DoseAreaProductMicrograySquareCentimeter,
            want: "μGy·cm²",
        },
        {
            name: "MilligraySquareCentimeter abbreviation",
            unit: units.DoseAreaProductMilligraySquareCentimeter,
            want: "mGy·cm²",
        },
        {
            name: "CentigraySquareCentimeter abbreviation",
            unit: units.DoseAreaProductCentigraySquareCentimeter,
            want: "cGy·cm²",
        },
        {
            name: "DecigraySquareCentimeter abbreviation",
            unit: units.DoseAreaProductDecigraySquareCentimeter,
            want: "dGy·cm²",
        },
        {
            name: "MicrograySquareMillimeter abbreviation",
            unit: units.DoseAreaProductMicrograySquareMillimeter,
            want: "μGy·mm²",
        },
        {
            name: "MilligraySquareMillimeter abbreviation",
            unit: units.DoseAreaProductMilligraySquareMillimeter,
            want: "mGy·mm²",
        },
        {
            name: "CentigraySquareMillimeter abbreviation",
            unit: units.DoseAreaProductCentigraySquareMillimeter,
            want: "cGy·mm²",
        },
        {
            name: "DecigraySquareMillimeter abbreviation",
            unit: units.DoseAreaProductDecigraySquareMillimeter,
            want: "dGy·mm²",
        },
        {
            name: "MicrograySquareMicrometer abbreviation",
            unit: units.DoseAreaProductMicrograySquareMicrometer,
            want: "μGy·μm²",
        },
        {
            name: "MilligraySquareMicrometer abbreviation",
            unit: units.DoseAreaProductMilligraySquareMicrometer,
            want: "mGy·μm²",
        },
        {
            name: "CentigraySquareMicrometer abbreviation",
            unit: units.DoseAreaProductCentigraySquareMicrometer,
            want: "cGy·μm²",
        },
        {
            name: "DecigraySquareMicrometer abbreviation",
            unit: units.DoseAreaProductDecigraySquareMicrometer,
            want: "dGy·μm²",
        },
        {
            name: "invalid unit",
            unit: units.DoseAreaProductUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetDoseAreaProductAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetDoseAreaProductAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestDoseAreaProduct_String(t *testing.T) {
    factory := units.DoseAreaProductFactory{}
    
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
            unit, err := factory.CreateDoseAreaProduct(tt.value, units.DoseAreaProductGraySquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("DoseAreaProduct.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestDoseAreaProduct_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.DoseAreaProductFactory{}

	_, err := uf.CreateDoseAreaProduct(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}