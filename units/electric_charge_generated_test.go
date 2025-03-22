// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricChargeDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Coulomb"}`
	
	factory := units.ElectricChargeDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricChargeCoulomb {
		t.Errorf("expected unit %v, got %v", units.ElectricChargeCoulomb, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Coulomb"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricChargeDto_ToJSON(t *testing.T) {
	dto := units.ElectricChargeDto{
		Value: 45,
		Unit:  units.ElectricChargeCoulomb,
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
	if result["unit"].(string) != string(units.ElectricChargeCoulomb) {
		t.Errorf("expected unit %s, got %v", units.ElectricChargeCoulomb, result["unit"])
	}
}

func TestNewElectricCharge_InvalidValue(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricCharge(math.NaN(), units.ElectricChargeCoulomb)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricCharge(math.Inf(1), units.ElectricChargeCoulomb)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricChargeConversions(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricCharge(180, units.ElectricChargeCoulomb)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Coulombs.
		// No expected conversion value provided for Coulombs, verifying result is not NaN.
		result := a.Coulombs()
		cacheResult := a.Coulombs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Coulombs returned NaN")
		}
	}
	{
		// Test conversion to AmpereHours.
		// No expected conversion value provided for AmpereHours, verifying result is not NaN.
		result := a.AmpereHours()
		cacheResult := a.AmpereHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmpereHours returned NaN")
		}
	}
	{
		// Test conversion to Picocoulombs.
		// No expected conversion value provided for Picocoulombs, verifying result is not NaN.
		result := a.Picocoulombs()
		cacheResult := a.Picocoulombs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picocoulombs returned NaN")
		}
	}
	{
		// Test conversion to Nanocoulombs.
		// No expected conversion value provided for Nanocoulombs, verifying result is not NaN.
		result := a.Nanocoulombs()
		cacheResult := a.Nanocoulombs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanocoulombs returned NaN")
		}
	}
	{
		// Test conversion to Microcoulombs.
		// No expected conversion value provided for Microcoulombs, verifying result is not NaN.
		result := a.Microcoulombs()
		cacheResult := a.Microcoulombs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microcoulombs returned NaN")
		}
	}
	{
		// Test conversion to Millicoulombs.
		// No expected conversion value provided for Millicoulombs, verifying result is not NaN.
		result := a.Millicoulombs()
		cacheResult := a.Millicoulombs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millicoulombs returned NaN")
		}
	}
	{
		// Test conversion to Kilocoulombs.
		// No expected conversion value provided for Kilocoulombs, verifying result is not NaN.
		result := a.Kilocoulombs()
		cacheResult := a.Kilocoulombs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilocoulombs returned NaN")
		}
	}
	{
		// Test conversion to Megacoulombs.
		// No expected conversion value provided for Megacoulombs, verifying result is not NaN.
		result := a.Megacoulombs()
		cacheResult := a.Megacoulombs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megacoulombs returned NaN")
		}
	}
	{
		// Test conversion to MilliampereHours.
		// No expected conversion value provided for MilliampereHours, verifying result is not NaN.
		result := a.MilliampereHours()
		cacheResult := a.MilliampereHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliampereHours returned NaN")
		}
	}
	{
		// Test conversion to KiloampereHours.
		// No expected conversion value provided for KiloampereHours, verifying result is not NaN.
		result := a.KiloampereHours()
		cacheResult := a.KiloampereHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KiloampereHours returned NaN")
		}
	}
	{
		// Test conversion to MegaampereHours.
		// No expected conversion value provided for MegaampereHours, verifying result is not NaN.
		result := a.MegaampereHours()
		cacheResult := a.MegaampereHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegaampereHours returned NaN")
		}
	}
}

func TestElectricCharge_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	a, err := factory.CreateElectricCharge(90, units.ElectricChargeCoulomb)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricChargeCoulomb {
		t.Errorf("expected default unit Coulomb, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricChargeCoulomb
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricChargeDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricChargeCoulomb {
		t.Errorf("expected unit Coulomb, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricChargeFactory_FromDto(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeCoulomb,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricChargeDto{
        Value: math.NaN(),
        Unit:  units.ElectricChargeCoulomb,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Coulomb conversion
    coulombsDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeCoulomb,
    }
    
    var coulombsResult *units.ElectricCharge
    coulombsResult, err = factory.FromDto(coulombsDto)
    if err != nil {
        t.Errorf("FromDto() with Coulomb returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombsResult.Convert(units.ElectricChargeCoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Coulomb = %v, want %v", converted, 100)
    }
    // Test AmpereHour conversion
    ampere_hoursDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeAmpereHour,
    }
    
    var ampere_hoursResult *units.ElectricCharge
    ampere_hoursResult, err = factory.FromDto(ampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with AmpereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ampere_hoursResult.Convert(units.ElectricChargeAmpereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmpereHour = %v, want %v", converted, 100)
    }
    // Test Picocoulomb conversion
    picocoulombsDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargePicocoulomb,
    }
    
    var picocoulombsResult *units.ElectricCharge
    picocoulombsResult, err = factory.FromDto(picocoulombsDto)
    if err != nil {
        t.Errorf("FromDto() with Picocoulomb returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picocoulombsResult.Convert(units.ElectricChargePicocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picocoulomb = %v, want %v", converted, 100)
    }
    // Test Nanocoulomb conversion
    nanocoulombsDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeNanocoulomb,
    }
    
    var nanocoulombsResult *units.ElectricCharge
    nanocoulombsResult, err = factory.FromDto(nanocoulombsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanocoulomb returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanocoulombsResult.Convert(units.ElectricChargeNanocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanocoulomb = %v, want %v", converted, 100)
    }
    // Test Microcoulomb conversion
    microcoulombsDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeMicrocoulomb,
    }
    
    var microcoulombsResult *units.ElectricCharge
    microcoulombsResult, err = factory.FromDto(microcoulombsDto)
    if err != nil {
        t.Errorf("FromDto() with Microcoulomb returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microcoulombsResult.Convert(units.ElectricChargeMicrocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microcoulomb = %v, want %v", converted, 100)
    }
    // Test Millicoulomb conversion
    millicoulombsDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeMillicoulomb,
    }
    
    var millicoulombsResult *units.ElectricCharge
    millicoulombsResult, err = factory.FromDto(millicoulombsDto)
    if err != nil {
        t.Errorf("FromDto() with Millicoulomb returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicoulombsResult.Convert(units.ElectricChargeMillicoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millicoulomb = %v, want %v", converted, 100)
    }
    // Test Kilocoulomb conversion
    kilocoulombsDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeKilocoulomb,
    }
    
    var kilocoulombsResult *units.ElectricCharge
    kilocoulombsResult, err = factory.FromDto(kilocoulombsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilocoulomb returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocoulombsResult.Convert(units.ElectricChargeKilocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilocoulomb = %v, want %v", converted, 100)
    }
    // Test Megacoulomb conversion
    megacoulombsDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeMegacoulomb,
    }
    
    var megacoulombsResult *units.ElectricCharge
    megacoulombsResult, err = factory.FromDto(megacoulombsDto)
    if err != nil {
        t.Errorf("FromDto() with Megacoulomb returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megacoulombsResult.Convert(units.ElectricChargeMegacoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megacoulomb = %v, want %v", converted, 100)
    }
    // Test MilliampereHour conversion
    milliampere_hoursDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeMilliampereHour,
    }
    
    var milliampere_hoursResult *units.ElectricCharge
    milliampere_hoursResult, err = factory.FromDto(milliampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MilliampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliampere_hoursResult.Convert(units.ElectricChargeMilliampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliampereHour = %v, want %v", converted, 100)
    }
    // Test KiloampereHour conversion
    kiloampere_hoursDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeKiloampereHour,
    }
    
    var kiloampere_hoursResult *units.ElectricCharge
    kiloampere_hoursResult, err = factory.FromDto(kiloampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with KiloampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloampere_hoursResult.Convert(units.ElectricChargeKiloampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloampereHour = %v, want %v", converted, 100)
    }
    // Test MegaampereHour conversion
    megaampere_hoursDto := units.ElectricChargeDto{
        Value: 100,
        Unit:  units.ElectricChargeMegaampereHour,
    }
    
    var megaampere_hoursResult *units.ElectricCharge
    megaampere_hoursResult, err = factory.FromDto(megaampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MegaampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaampere_hoursResult.Convert(units.ElectricChargeMegaampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaampereHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricChargeDto{
        Value: 0,
        Unit:  units.ElectricChargeCoulomb,
    }
    
    var zeroResult *units.ElectricCharge
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricChargeFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Coulomb"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Coulomb"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricChargeDto{
        Value: nanValue,
        Unit:  units.ElectricChargeCoulomb,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Coulomb unit
    coulombsJSON := []byte(`{"value": 100, "unit": "Coulomb"}`)
    coulombsResult, err := factory.FromDtoJSON(coulombsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Coulomb unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombsResult.Convert(units.ElectricChargeCoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Coulomb = %v, want %v", converted, 100)
    }
    // Test JSON with AmpereHour unit
    ampere_hoursJSON := []byte(`{"value": 100, "unit": "AmpereHour"}`)
    ampere_hoursResult, err := factory.FromDtoJSON(ampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmpereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ampere_hoursResult.Convert(units.ElectricChargeAmpereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmpereHour = %v, want %v", converted, 100)
    }
    // Test JSON with Picocoulomb unit
    picocoulombsJSON := []byte(`{"value": 100, "unit": "Picocoulomb"}`)
    picocoulombsResult, err := factory.FromDtoJSON(picocoulombsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picocoulomb unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picocoulombsResult.Convert(units.ElectricChargePicocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picocoulomb = %v, want %v", converted, 100)
    }
    // Test JSON with Nanocoulomb unit
    nanocoulombsJSON := []byte(`{"value": 100, "unit": "Nanocoulomb"}`)
    nanocoulombsResult, err := factory.FromDtoJSON(nanocoulombsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanocoulomb unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanocoulombsResult.Convert(units.ElectricChargeNanocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanocoulomb = %v, want %v", converted, 100)
    }
    // Test JSON with Microcoulomb unit
    microcoulombsJSON := []byte(`{"value": 100, "unit": "Microcoulomb"}`)
    microcoulombsResult, err := factory.FromDtoJSON(microcoulombsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microcoulomb unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microcoulombsResult.Convert(units.ElectricChargeMicrocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microcoulomb = %v, want %v", converted, 100)
    }
    // Test JSON with Millicoulomb unit
    millicoulombsJSON := []byte(`{"value": 100, "unit": "Millicoulomb"}`)
    millicoulombsResult, err := factory.FromDtoJSON(millicoulombsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millicoulomb unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicoulombsResult.Convert(units.ElectricChargeMillicoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millicoulomb = %v, want %v", converted, 100)
    }
    // Test JSON with Kilocoulomb unit
    kilocoulombsJSON := []byte(`{"value": 100, "unit": "Kilocoulomb"}`)
    kilocoulombsResult, err := factory.FromDtoJSON(kilocoulombsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilocoulomb unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocoulombsResult.Convert(units.ElectricChargeKilocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilocoulomb = %v, want %v", converted, 100)
    }
    // Test JSON with Megacoulomb unit
    megacoulombsJSON := []byte(`{"value": 100, "unit": "Megacoulomb"}`)
    megacoulombsResult, err := factory.FromDtoJSON(megacoulombsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megacoulomb unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megacoulombsResult.Convert(units.ElectricChargeMegacoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megacoulomb = %v, want %v", converted, 100)
    }
    // Test JSON with MilliampereHour unit
    milliampere_hoursJSON := []byte(`{"value": 100, "unit": "MilliampereHour"}`)
    milliampere_hoursResult, err := factory.FromDtoJSON(milliampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliampere_hoursResult.Convert(units.ElectricChargeMilliampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliampereHour = %v, want %v", converted, 100)
    }
    // Test JSON with KiloampereHour unit
    kiloampere_hoursJSON := []byte(`{"value": 100, "unit": "KiloampereHour"}`)
    kiloampere_hoursResult, err := factory.FromDtoJSON(kiloampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloampere_hoursResult.Convert(units.ElectricChargeKiloampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloampereHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegaampereHour unit
    megaampere_hoursJSON := []byte(`{"value": 100, "unit": "MegaampereHour"}`)
    megaampere_hoursResult, err := factory.FromDtoJSON(megaampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaampere_hoursResult.Convert(units.ElectricChargeMegaampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaampereHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Coulomb"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCoulombs function
func TestElectricChargeFactory_FromCoulombs(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCoulombs(100)
    if err != nil {
        t.Errorf("FromCoulombs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeCoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCoulombs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCoulombs(math.NaN())
    if err == nil {
        t.Error("FromCoulombs() with NaN value should return error")
    }

    _, err = factory.FromCoulombs(math.Inf(1))
    if err == nil {
        t.Error("FromCoulombs() with +Inf value should return error")
    }

    _, err = factory.FromCoulombs(math.Inf(-1))
    if err == nil {
        t.Error("FromCoulombs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCoulombs(0)
    if err != nil {
        t.Errorf("FromCoulombs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeCoulomb)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCoulombs() with zero value = %v, want 0", converted)
    }
}
// Test FromAmpereHours function
func TestElectricChargeFactory_FromAmpereHours(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmpereHours(100)
    if err != nil {
        t.Errorf("FromAmpereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeAmpereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmpereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmpereHours(math.NaN())
    if err == nil {
        t.Error("FromAmpereHours() with NaN value should return error")
    }

    _, err = factory.FromAmpereHours(math.Inf(1))
    if err == nil {
        t.Error("FromAmpereHours() with +Inf value should return error")
    }

    _, err = factory.FromAmpereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromAmpereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmpereHours(0)
    if err != nil {
        t.Errorf("FromAmpereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeAmpereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmpereHours() with zero value = %v, want 0", converted)
    }
}
// Test FromPicocoulombs function
func TestElectricChargeFactory_FromPicocoulombs(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicocoulombs(100)
    if err != nil {
        t.Errorf("FromPicocoulombs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargePicocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicocoulombs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicocoulombs(math.NaN())
    if err == nil {
        t.Error("FromPicocoulombs() with NaN value should return error")
    }

    _, err = factory.FromPicocoulombs(math.Inf(1))
    if err == nil {
        t.Error("FromPicocoulombs() with +Inf value should return error")
    }

    _, err = factory.FromPicocoulombs(math.Inf(-1))
    if err == nil {
        t.Error("FromPicocoulombs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicocoulombs(0)
    if err != nil {
        t.Errorf("FromPicocoulombs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargePicocoulomb)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicocoulombs() with zero value = %v, want 0", converted)
    }
}
// Test FromNanocoulombs function
func TestElectricChargeFactory_FromNanocoulombs(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanocoulombs(100)
    if err != nil {
        t.Errorf("FromNanocoulombs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeNanocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanocoulombs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanocoulombs(math.NaN())
    if err == nil {
        t.Error("FromNanocoulombs() with NaN value should return error")
    }

    _, err = factory.FromNanocoulombs(math.Inf(1))
    if err == nil {
        t.Error("FromNanocoulombs() with +Inf value should return error")
    }

    _, err = factory.FromNanocoulombs(math.Inf(-1))
    if err == nil {
        t.Error("FromNanocoulombs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanocoulombs(0)
    if err != nil {
        t.Errorf("FromNanocoulombs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeNanocoulomb)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanocoulombs() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrocoulombs function
func TestElectricChargeFactory_FromMicrocoulombs(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrocoulombs(100)
    if err != nil {
        t.Errorf("FromMicrocoulombs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeMicrocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrocoulombs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrocoulombs(math.NaN())
    if err == nil {
        t.Error("FromMicrocoulombs() with NaN value should return error")
    }

    _, err = factory.FromMicrocoulombs(math.Inf(1))
    if err == nil {
        t.Error("FromMicrocoulombs() with +Inf value should return error")
    }

    _, err = factory.FromMicrocoulombs(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrocoulombs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrocoulombs(0)
    if err != nil {
        t.Errorf("FromMicrocoulombs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeMicrocoulomb)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrocoulombs() with zero value = %v, want 0", converted)
    }
}
// Test FromMillicoulombs function
func TestElectricChargeFactory_FromMillicoulombs(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillicoulombs(100)
    if err != nil {
        t.Errorf("FromMillicoulombs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeMillicoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillicoulombs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillicoulombs(math.NaN())
    if err == nil {
        t.Error("FromMillicoulombs() with NaN value should return error")
    }

    _, err = factory.FromMillicoulombs(math.Inf(1))
    if err == nil {
        t.Error("FromMillicoulombs() with +Inf value should return error")
    }

    _, err = factory.FromMillicoulombs(math.Inf(-1))
    if err == nil {
        t.Error("FromMillicoulombs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillicoulombs(0)
    if err != nil {
        t.Errorf("FromMillicoulombs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeMillicoulomb)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillicoulombs() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocoulombs function
func TestElectricChargeFactory_FromKilocoulombs(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocoulombs(100)
    if err != nil {
        t.Errorf("FromKilocoulombs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeKilocoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocoulombs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocoulombs(math.NaN())
    if err == nil {
        t.Error("FromKilocoulombs() with NaN value should return error")
    }

    _, err = factory.FromKilocoulombs(math.Inf(1))
    if err == nil {
        t.Error("FromKilocoulombs() with +Inf value should return error")
    }

    _, err = factory.FromKilocoulombs(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocoulombs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocoulombs(0)
    if err != nil {
        t.Errorf("FromKilocoulombs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeKilocoulomb)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocoulombs() with zero value = %v, want 0", converted)
    }
}
// Test FromMegacoulombs function
func TestElectricChargeFactory_FromMegacoulombs(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegacoulombs(100)
    if err != nil {
        t.Errorf("FromMegacoulombs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeMegacoulomb)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegacoulombs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegacoulombs(math.NaN())
    if err == nil {
        t.Error("FromMegacoulombs() with NaN value should return error")
    }

    _, err = factory.FromMegacoulombs(math.Inf(1))
    if err == nil {
        t.Error("FromMegacoulombs() with +Inf value should return error")
    }

    _, err = factory.FromMegacoulombs(math.Inf(-1))
    if err == nil {
        t.Error("FromMegacoulombs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegacoulombs(0)
    if err != nil {
        t.Errorf("FromMegacoulombs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeMegacoulomb)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegacoulombs() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliampereHours function
func TestElectricChargeFactory_FromMilliampereHours(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliampereHours(100)
    if err != nil {
        t.Errorf("FromMilliampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeMilliampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliampereHours(math.NaN())
    if err == nil {
        t.Error("FromMilliampereHours() with NaN value should return error")
    }

    _, err = factory.FromMilliampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromMilliampereHours() with +Inf value should return error")
    }

    _, err = factory.FromMilliampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliampereHours(0)
    if err != nil {
        t.Errorf("FromMilliampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeMilliampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliampereHours() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloampereHours function
func TestElectricChargeFactory_FromKiloampereHours(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloampereHours(100)
    if err != nil {
        t.Errorf("FromKiloampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeKiloampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloampereHours(math.NaN())
    if err == nil {
        t.Error("FromKiloampereHours() with NaN value should return error")
    }

    _, err = factory.FromKiloampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromKiloampereHours() with +Inf value should return error")
    }

    _, err = factory.FromKiloampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloampereHours(0)
    if err != nil {
        t.Errorf("FromKiloampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeKiloampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloampereHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaampereHours function
func TestElectricChargeFactory_FromMegaampereHours(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaampereHours(100)
    if err != nil {
        t.Errorf("FromMegaampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeMegaampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaampereHours(math.NaN())
    if err == nil {
        t.Error("FromMegaampereHours() with NaN value should return error")
    }

    _, err = factory.FromMegaampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromMegaampereHours() with +Inf value should return error")
    }

    _, err = factory.FromMegaampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaampereHours(0)
    if err != nil {
        t.Errorf("FromMegaampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeMegaampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaampereHours() with zero value = %v, want 0", converted)
    }
}

func TestElectricChargeToString(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	a, err := factory.CreateElectricCharge(45, units.ElectricChargeCoulomb)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricChargeCoulomb, 2)
	expected := "45.00 " + units.GetElectricChargeAbbreviation(units.ElectricChargeCoulomb)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricChargeCoulomb, -1)
	expected = "45 " + units.GetElectricChargeAbbreviation(units.ElectricChargeCoulomb)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricCharge_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	a1, _ := factory.CreateElectricCharge(60, units.ElectricChargeCoulomb)
	a2, _ := factory.CreateElectricCharge(60, units.ElectricChargeCoulomb)
	a3, _ := factory.CreateElectricCharge(120, units.ElectricChargeCoulomb)

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

func TestElectricCharge_Arithmetic(t *testing.T) {
	factory := units.ElectricChargeFactory{}
	a1, _ := factory.CreateElectricCharge(30, units.ElectricChargeCoulomb)
	a2, _ := factory.CreateElectricCharge(45, units.ElectricChargeCoulomb)

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


func TestGetElectricChargeAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricChargeUnits
        want string
    }{
        {
            name: "Coulomb abbreviation",
            unit: units.ElectricChargeCoulomb,
            want: "C",
        },
        {
            name: "AmpereHour abbreviation",
            unit: units.ElectricChargeAmpereHour,
            want: "A-h",
        },
        {
            name: "Picocoulomb abbreviation",
            unit: units.ElectricChargePicocoulomb,
            want: "pC",
        },
        {
            name: "Nanocoulomb abbreviation",
            unit: units.ElectricChargeNanocoulomb,
            want: "nC",
        },
        {
            name: "Microcoulomb abbreviation",
            unit: units.ElectricChargeMicrocoulomb,
            want: "μC",
        },
        {
            name: "Millicoulomb abbreviation",
            unit: units.ElectricChargeMillicoulomb,
            want: "mC",
        },
        {
            name: "Kilocoulomb abbreviation",
            unit: units.ElectricChargeKilocoulomb,
            want: "kC",
        },
        {
            name: "Megacoulomb abbreviation",
            unit: units.ElectricChargeMegacoulomb,
            want: "MC",
        },
        {
            name: "MilliampereHour abbreviation",
            unit: units.ElectricChargeMilliampereHour,
            want: "mA-h",
        },
        {
            name: "KiloampereHour abbreviation",
            unit: units.ElectricChargeKiloampereHour,
            want: "kA-h",
        },
        {
            name: "MegaampereHour abbreviation",
            unit: units.ElectricChargeMegaampereHour,
            want: "MA-h",
        },
        {
            name: "invalid unit",
            unit: units.ElectricChargeUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricChargeAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricChargeAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricCharge_String(t *testing.T) {
    factory := units.ElectricChargeFactory{}
    
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
            unit, err := factory.CreateElectricCharge(tt.value, units.ElectricChargeCoulomb)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricCharge.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestElectricCharge_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ElectricChargeFactory{}

	_, err := uf.CreateElectricCharge(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}