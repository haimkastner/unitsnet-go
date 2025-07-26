// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumeFlowDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CubicMeterPerSecond"}`
	
	factory := units.VolumeFlowDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumeFlowCubicMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.VolumeFlowCubicMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CubicMeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumeFlowDto_ToJSON(t *testing.T) {
	dto := units.VolumeFlowDto{
		Value: 45,
		Unit:  units.VolumeFlowCubicMeterPerSecond,
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
	if result["unit"].(string) != string(units.VolumeFlowCubicMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.VolumeFlowCubicMeterPerSecond, result["unit"])
	}
}

func TestNewVolumeFlow_InvalidValue(t *testing.T) {
	factory := units.VolumeFlowFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolumeFlow(math.NaN(), units.VolumeFlowCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolumeFlow(math.Inf(1), units.VolumeFlowCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumeFlowConversions(t *testing.T) {
	factory := units.VolumeFlowFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolumeFlow(180, units.VolumeFlowCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CubicMetersPerSecond.
		// No expected conversion value provided for CubicMetersPerSecond, verifying result is not NaN.
		result := a.CubicMetersPerSecond()
		cacheResult := a.CubicMetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CubicMetersPerMinute.
		// No expected conversion value provided for CubicMetersPerMinute, verifying result is not NaN.
		result := a.CubicMetersPerMinute()
		cacheResult := a.CubicMetersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CubicMetersPerHour.
		// No expected conversion value provided for CubicMetersPerHour, verifying result is not NaN.
		result := a.CubicMetersPerHour()
		cacheResult := a.CubicMetersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to CubicMetersPerDay.
		// No expected conversion value provided for CubicMetersPerDay, verifying result is not NaN.
		result := a.CubicMetersPerDay()
		cacheResult := a.CubicMetersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMetersPerDay returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerSecond.
		// No expected conversion value provided for CubicFeetPerSecond, verifying result is not NaN.
		result := a.CubicFeetPerSecond()
		cacheResult := a.CubicFeetPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerMinute.
		// No expected conversion value provided for CubicFeetPerMinute, verifying result is not NaN.
		result := a.CubicFeetPerMinute()
		cacheResult := a.CubicFeetPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicFeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerHour.
		// No expected conversion value provided for CubicFeetPerHour, verifying result is not NaN.
		result := a.CubicFeetPerHour()
		cacheResult := a.CubicFeetPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicFeetPerHour returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerSecond.
		// No expected conversion value provided for CubicYardsPerSecond, verifying result is not NaN.
		result := a.CubicYardsPerSecond()
		cacheResult := a.CubicYardsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicYardsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerMinute.
		// No expected conversion value provided for CubicYardsPerMinute, verifying result is not NaN.
		result := a.CubicYardsPerMinute()
		cacheResult := a.CubicYardsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicYardsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerHour.
		// No expected conversion value provided for CubicYardsPerHour, verifying result is not NaN.
		result := a.CubicYardsPerHour()
		cacheResult := a.CubicYardsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicYardsPerHour returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerDay.
		// No expected conversion value provided for CubicYardsPerDay, verifying result is not NaN.
		result := a.CubicYardsPerDay()
		cacheResult := a.CubicYardsPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicYardsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MillionUsGallonsPerDay.
		// No expected conversion value provided for MillionUsGallonsPerDay, verifying result is not NaN.
		result := a.MillionUsGallonsPerDay()
		cacheResult := a.MillionUsGallonsPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillionUsGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerDay.
		// No expected conversion value provided for UsGallonsPerDay, verifying result is not NaN.
		result := a.UsGallonsPerDay()
		cacheResult := a.UsGallonsPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to LitersPerSecond.
		// No expected conversion value provided for LitersPerSecond, verifying result is not NaN.
		result := a.LitersPerSecond()
		cacheResult := a.LitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to LitersPerMinute.
		// No expected conversion value provided for LitersPerMinute, verifying result is not NaN.
		result := a.LitersPerMinute()
		cacheResult := a.LitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to LitersPerHour.
		// No expected conversion value provided for LitersPerHour, verifying result is not NaN.
		result := a.LitersPerHour()
		cacheResult := a.LitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to LitersPerDay.
		// No expected conversion value provided for LitersPerDay, verifying result is not NaN.
		result := a.LitersPerDay()
		cacheResult := a.LitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerSecond.
		// No expected conversion value provided for UsGallonsPerSecond, verifying result is not NaN.
		result := a.UsGallonsPerSecond()
		cacheResult := a.UsGallonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsGallonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerMinute.
		// No expected conversion value provided for UsGallonsPerMinute, verifying result is not NaN.
		result := a.UsGallonsPerMinute()
		cacheResult := a.UsGallonsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsGallonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to UkGallonsPerDay.
		// No expected conversion value provided for UkGallonsPerDay, verifying result is not NaN.
		result := a.UkGallonsPerDay()
		cacheResult := a.UkGallonsPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UkGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to UkGallonsPerHour.
		// No expected conversion value provided for UkGallonsPerHour, verifying result is not NaN.
		result := a.UkGallonsPerHour()
		cacheResult := a.UkGallonsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UkGallonsPerHour returned NaN")
		}
	}
	{
		// Test conversion to UkGallonsPerMinute.
		// No expected conversion value provided for UkGallonsPerMinute, verifying result is not NaN.
		result := a.UkGallonsPerMinute()
		cacheResult := a.UkGallonsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UkGallonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to UkGallonsPerSecond.
		// No expected conversion value provided for UkGallonsPerSecond, verifying result is not NaN.
		result := a.UkGallonsPerSecond()
		cacheResult := a.UkGallonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UkGallonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilousGallonsPerMinute.
		// No expected conversion value provided for KilousGallonsPerMinute, verifying result is not NaN.
		result := a.KilousGallonsPerMinute()
		cacheResult := a.KilousGallonsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilousGallonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerHour.
		// No expected conversion value provided for UsGallonsPerHour, verifying result is not NaN.
		result := a.UsGallonsPerHour()
		cacheResult := a.UsGallonsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsGallonsPerHour returned NaN")
		}
	}
	{
		// Test conversion to CubicDecimetersPerMinute.
		// No expected conversion value provided for CubicDecimetersPerMinute, verifying result is not NaN.
		result := a.CubicDecimetersPerMinute()
		cacheResult := a.CubicDecimetersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicDecimetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerDay.
		// No expected conversion value provided for OilBarrelsPerDay, verifying result is not NaN.
		result := a.OilBarrelsPerDay()
		cacheResult := a.OilBarrelsPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OilBarrelsPerDay returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerMinute.
		// No expected conversion value provided for OilBarrelsPerMinute, verifying result is not NaN.
		result := a.OilBarrelsPerMinute()
		cacheResult := a.OilBarrelsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OilBarrelsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerHour.
		// No expected conversion value provided for OilBarrelsPerHour, verifying result is not NaN.
		result := a.OilBarrelsPerHour()
		cacheResult := a.OilBarrelsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OilBarrelsPerHour returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerSecond.
		// No expected conversion value provided for OilBarrelsPerSecond, verifying result is not NaN.
		result := a.OilBarrelsPerSecond()
		cacheResult := a.OilBarrelsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OilBarrelsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CubicMillimetersPerSecond.
		// No expected conversion value provided for CubicMillimetersPerSecond, verifying result is not NaN.
		result := a.CubicMillimetersPerSecond()
		cacheResult := a.CubicMillimetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMillimetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AcreFeetPerSecond.
		// No expected conversion value provided for AcreFeetPerSecond, verifying result is not NaN.
		result := a.AcreFeetPerSecond()
		cacheResult := a.AcreFeetPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AcreFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AcreFeetPerMinute.
		// No expected conversion value provided for AcreFeetPerMinute, verifying result is not NaN.
		result := a.AcreFeetPerMinute()
		cacheResult := a.AcreFeetPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AcreFeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to AcreFeetPerHour.
		// No expected conversion value provided for AcreFeetPerHour, verifying result is not NaN.
		result := a.AcreFeetPerHour()
		cacheResult := a.AcreFeetPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AcreFeetPerHour returned NaN")
		}
	}
	{
		// Test conversion to AcreFeetPerDay.
		// No expected conversion value provided for AcreFeetPerDay, verifying result is not NaN.
		result := a.AcreFeetPerDay()
		cacheResult := a.AcreFeetPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AcreFeetPerDay returned NaN")
		}
	}
	{
		// Test conversion to CubicCentimetersPerMinute.
		// No expected conversion value provided for CubicCentimetersPerMinute, verifying result is not NaN.
		result := a.CubicCentimetersPerMinute()
		cacheResult := a.CubicCentimetersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicCentimetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegausGallonsPerDay.
		// No expected conversion value provided for MegausGallonsPerDay, verifying result is not NaN.
		result := a.MegausGallonsPerDay()
		cacheResult := a.MegausGallonsPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegausGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerSecond.
		// No expected conversion value provided for NanolitersPerSecond, verifying result is not NaN.
		result := a.NanolitersPerSecond()
		cacheResult := a.NanolitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanolitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerSecond.
		// No expected conversion value provided for MicrolitersPerSecond, verifying result is not NaN.
		result := a.MicrolitersPerSecond()
		cacheResult := a.MicrolitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrolitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerSecond.
		// No expected conversion value provided for MillilitersPerSecond, verifying result is not NaN.
		result := a.MillilitersPerSecond()
		cacheResult := a.MillilitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillilitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerSecond.
		// No expected conversion value provided for CentilitersPerSecond, verifying result is not NaN.
		result := a.CentilitersPerSecond()
		cacheResult := a.CentilitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentilitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerSecond.
		// No expected conversion value provided for DecilitersPerSecond, verifying result is not NaN.
		result := a.DecilitersPerSecond()
		cacheResult := a.DecilitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecilitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecalitersPerSecond.
		// No expected conversion value provided for DecalitersPerSecond, verifying result is not NaN.
		result := a.DecalitersPerSecond()
		cacheResult := a.DecalitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecalitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to HectolitersPerSecond.
		// No expected conversion value provided for HectolitersPerSecond, verifying result is not NaN.
		result := a.HectolitersPerSecond()
		cacheResult := a.HectolitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectolitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilolitersPerSecond.
		// No expected conversion value provided for KilolitersPerSecond, verifying result is not NaN.
		result := a.KilolitersPerSecond()
		cacheResult := a.KilolitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilolitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegalitersPerSecond.
		// No expected conversion value provided for MegalitersPerSecond, verifying result is not NaN.
		result := a.MegalitersPerSecond()
		cacheResult := a.MegalitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegalitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerMinute.
		// No expected conversion value provided for NanolitersPerMinute, verifying result is not NaN.
		result := a.NanolitersPerMinute()
		cacheResult := a.NanolitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanolitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerMinute.
		// No expected conversion value provided for MicrolitersPerMinute, verifying result is not NaN.
		result := a.MicrolitersPerMinute()
		cacheResult := a.MicrolitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrolitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerMinute.
		// No expected conversion value provided for MillilitersPerMinute, verifying result is not NaN.
		result := a.MillilitersPerMinute()
		cacheResult := a.MillilitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillilitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerMinute.
		// No expected conversion value provided for CentilitersPerMinute, verifying result is not NaN.
		result := a.CentilitersPerMinute()
		cacheResult := a.CentilitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentilitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerMinute.
		// No expected conversion value provided for DecilitersPerMinute, verifying result is not NaN.
		result := a.DecilitersPerMinute()
		cacheResult := a.DecilitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecilitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DecalitersPerMinute.
		// No expected conversion value provided for DecalitersPerMinute, verifying result is not NaN.
		result := a.DecalitersPerMinute()
		cacheResult := a.DecalitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecalitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to HectolitersPerMinute.
		// No expected conversion value provided for HectolitersPerMinute, verifying result is not NaN.
		result := a.HectolitersPerMinute()
		cacheResult := a.HectolitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectolitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilolitersPerMinute.
		// No expected conversion value provided for KilolitersPerMinute, verifying result is not NaN.
		result := a.KilolitersPerMinute()
		cacheResult := a.KilolitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilolitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegalitersPerMinute.
		// No expected conversion value provided for MegalitersPerMinute, verifying result is not NaN.
		result := a.MegalitersPerMinute()
		cacheResult := a.MegalitersPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegalitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerHour.
		// No expected conversion value provided for NanolitersPerHour, verifying result is not NaN.
		result := a.NanolitersPerHour()
		cacheResult := a.NanolitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanolitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerHour.
		// No expected conversion value provided for MicrolitersPerHour, verifying result is not NaN.
		result := a.MicrolitersPerHour()
		cacheResult := a.MicrolitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrolitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerHour.
		// No expected conversion value provided for MillilitersPerHour, verifying result is not NaN.
		result := a.MillilitersPerHour()
		cacheResult := a.MillilitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillilitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerHour.
		// No expected conversion value provided for CentilitersPerHour, verifying result is not NaN.
		result := a.CentilitersPerHour()
		cacheResult := a.CentilitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentilitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerHour.
		// No expected conversion value provided for DecilitersPerHour, verifying result is not NaN.
		result := a.DecilitersPerHour()
		cacheResult := a.DecilitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecilitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to DecalitersPerHour.
		// No expected conversion value provided for DecalitersPerHour, verifying result is not NaN.
		result := a.DecalitersPerHour()
		cacheResult := a.DecalitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecalitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to HectolitersPerHour.
		// No expected conversion value provided for HectolitersPerHour, verifying result is not NaN.
		result := a.HectolitersPerHour()
		cacheResult := a.HectolitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectolitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilolitersPerHour.
		// No expected conversion value provided for KilolitersPerHour, verifying result is not NaN.
		result := a.KilolitersPerHour()
		cacheResult := a.KilolitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilolitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to MegalitersPerHour.
		// No expected conversion value provided for MegalitersPerHour, verifying result is not NaN.
		result := a.MegalitersPerHour()
		cacheResult := a.MegalitersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegalitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerDay.
		// No expected conversion value provided for NanolitersPerDay, verifying result is not NaN.
		result := a.NanolitersPerDay()
		cacheResult := a.NanolitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanolitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerDay.
		// No expected conversion value provided for MicrolitersPerDay, verifying result is not NaN.
		result := a.MicrolitersPerDay()
		cacheResult := a.MicrolitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrolitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerDay.
		// No expected conversion value provided for MillilitersPerDay, verifying result is not NaN.
		result := a.MillilitersPerDay()
		cacheResult := a.MillilitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillilitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerDay.
		// No expected conversion value provided for CentilitersPerDay, verifying result is not NaN.
		result := a.CentilitersPerDay()
		cacheResult := a.CentilitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentilitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerDay.
		// No expected conversion value provided for DecilitersPerDay, verifying result is not NaN.
		result := a.DecilitersPerDay()
		cacheResult := a.DecilitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecilitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to DecalitersPerDay.
		// No expected conversion value provided for DecalitersPerDay, verifying result is not NaN.
		result := a.DecalitersPerDay()
		cacheResult := a.DecalitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecalitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to HectolitersPerDay.
		// No expected conversion value provided for HectolitersPerDay, verifying result is not NaN.
		result := a.HectolitersPerDay()
		cacheResult := a.HectolitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectolitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to KilolitersPerDay.
		// No expected conversion value provided for KilolitersPerDay, verifying result is not NaN.
		result := a.KilolitersPerDay()
		cacheResult := a.KilolitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilolitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegalitersPerDay.
		// No expected conversion value provided for MegalitersPerDay, verifying result is not NaN.
		result := a.MegalitersPerDay()
		cacheResult := a.MegalitersPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegalitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegaukGallonsPerDay.
		// No expected conversion value provided for MegaukGallonsPerDay, verifying result is not NaN.
		result := a.MegaukGallonsPerDay()
		cacheResult := a.MegaukGallonsPerDay()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegaukGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegaukGallonsPerSecond.
		// No expected conversion value provided for MegaukGallonsPerSecond, verifying result is not NaN.
		result := a.MegaukGallonsPerSecond()
		cacheResult := a.MegaukGallonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegaukGallonsPerSecond returned NaN")
		}
	}
}

func TestVolumeFlow_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumeFlowFactory{}
	a, err := factory.CreateVolumeFlow(90, units.VolumeFlowCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumeFlowCubicMeterPerSecond {
		t.Errorf("expected default unit CubicMeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumeFlowCubicMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumeFlowDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumeFlowCubicMeterPerSecond {
		t.Errorf("expected unit CubicMeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumeFlowFactory_FromDto(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicMeterPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.VolumeFlowDto{
        Value: math.NaN(),
        Unit:  units.VolumeFlowCubicMeterPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test CubicMeterPerSecond conversion
    cubic_meters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicMeterPerSecond,
    }
    
    var cubic_meters_per_secondResult *units.VolumeFlow
    cubic_meters_per_secondResult, err = factory.FromDto(cubic_meters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_secondResult.Convert(units.VolumeFlowCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test CubicMeterPerMinute conversion
    cubic_meters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicMeterPerMinute,
    }
    
    var cubic_meters_per_minuteResult *units.VolumeFlow
    cubic_meters_per_minuteResult, err = factory.FromDto(cubic_meters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_minuteResult.Convert(units.VolumeFlowCubicMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerMinute = %v, want %v", converted, 100)
    }
    // Test CubicMeterPerHour conversion
    cubic_meters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicMeterPerHour,
    }
    
    var cubic_meters_per_hourResult *units.VolumeFlow
    cubic_meters_per_hourResult, err = factory.FromDto(cubic_meters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMeterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_hourResult.Convert(units.VolumeFlowCubicMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerHour = %v, want %v", converted, 100)
    }
    // Test CubicMeterPerDay conversion
    cubic_meters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicMeterPerDay,
    }
    
    var cubic_meters_per_dayResult *units.VolumeFlow
    cubic_meters_per_dayResult, err = factory.FromDto(cubic_meters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMeterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_dayResult.Convert(units.VolumeFlowCubicMeterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerDay = %v, want %v", converted, 100)
    }
    // Test CubicFootPerSecond conversion
    cubic_feet_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicFootPerSecond,
    }
    
    var cubic_feet_per_secondResult *units.VolumeFlow
    cubic_feet_per_secondResult, err = factory.FromDto(cubic_feet_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CubicFootPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_secondResult.Convert(units.VolumeFlowCubicFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerSecond = %v, want %v", converted, 100)
    }
    // Test CubicFootPerMinute conversion
    cubic_feet_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicFootPerMinute,
    }
    
    var cubic_feet_per_minuteResult *units.VolumeFlow
    cubic_feet_per_minuteResult, err = factory.FromDto(cubic_feet_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with CubicFootPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_minuteResult.Convert(units.VolumeFlowCubicFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerMinute = %v, want %v", converted, 100)
    }
    // Test CubicFootPerHour conversion
    cubic_feet_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicFootPerHour,
    }
    
    var cubic_feet_per_hourResult *units.VolumeFlow
    cubic_feet_per_hourResult, err = factory.FromDto(cubic_feet_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with CubicFootPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_hourResult.Convert(units.VolumeFlowCubicFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerHour = %v, want %v", converted, 100)
    }
    // Test CubicYardPerSecond conversion
    cubic_yards_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicYardPerSecond,
    }
    
    var cubic_yards_per_secondResult *units.VolumeFlow
    cubic_yards_per_secondResult, err = factory.FromDto(cubic_yards_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CubicYardPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_secondResult.Convert(units.VolumeFlowCubicYardPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerSecond = %v, want %v", converted, 100)
    }
    // Test CubicYardPerMinute conversion
    cubic_yards_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicYardPerMinute,
    }
    
    var cubic_yards_per_minuteResult *units.VolumeFlow
    cubic_yards_per_minuteResult, err = factory.FromDto(cubic_yards_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with CubicYardPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_minuteResult.Convert(units.VolumeFlowCubicYardPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerMinute = %v, want %v", converted, 100)
    }
    // Test CubicYardPerHour conversion
    cubic_yards_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicYardPerHour,
    }
    
    var cubic_yards_per_hourResult *units.VolumeFlow
    cubic_yards_per_hourResult, err = factory.FromDto(cubic_yards_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with CubicYardPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_hourResult.Convert(units.VolumeFlowCubicYardPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerHour = %v, want %v", converted, 100)
    }
    // Test CubicYardPerDay conversion
    cubic_yards_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicYardPerDay,
    }
    
    var cubic_yards_per_dayResult *units.VolumeFlow
    cubic_yards_per_dayResult, err = factory.FromDto(cubic_yards_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with CubicYardPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_dayResult.Convert(units.VolumeFlowCubicYardPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerDay = %v, want %v", converted, 100)
    }
    // Test MillionUsGallonPerDay conversion
    million_us_gallons_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMillionUsGallonPerDay,
    }
    
    var million_us_gallons_per_dayResult *units.VolumeFlow
    million_us_gallons_per_dayResult, err = factory.FromDto(million_us_gallons_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MillionUsGallonPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = million_us_gallons_per_dayResult.Convert(units.VolumeFlowMillionUsGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillionUsGallonPerDay = %v, want %v", converted, 100)
    }
    // Test UsGallonPerDay conversion
    us_gallons_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowUsGallonPerDay,
    }
    
    var us_gallons_per_dayResult *units.VolumeFlow
    us_gallons_per_dayResult, err = factory.FromDto(us_gallons_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with UsGallonPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_dayResult.Convert(units.VolumeFlowUsGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerDay = %v, want %v", converted, 100)
    }
    // Test LiterPerSecond conversion
    liters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowLiterPerSecond,
    }
    
    var liters_per_secondResult *units.VolumeFlow
    liters_per_secondResult, err = factory.FromDto(liters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_secondResult.Convert(units.VolumeFlowLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerSecond = %v, want %v", converted, 100)
    }
    // Test LiterPerMinute conversion
    liters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowLiterPerMinute,
    }
    
    var liters_per_minuteResult *units.VolumeFlow
    liters_per_minuteResult, err = factory.FromDto(liters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_minuteResult.Convert(units.VolumeFlowLiterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerMinute = %v, want %v", converted, 100)
    }
    // Test LiterPerHour conversion
    liters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowLiterPerHour,
    }
    
    var liters_per_hourResult *units.VolumeFlow
    liters_per_hourResult, err = factory.FromDto(liters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_hourResult.Convert(units.VolumeFlowLiterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerHour = %v, want %v", converted, 100)
    }
    // Test LiterPerDay conversion
    liters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowLiterPerDay,
    }
    
    var liters_per_dayResult *units.VolumeFlow
    liters_per_dayResult, err = factory.FromDto(liters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_dayResult.Convert(units.VolumeFlowLiterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerDay = %v, want %v", converted, 100)
    }
    // Test UsGallonPerSecond conversion
    us_gallons_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowUsGallonPerSecond,
    }
    
    var us_gallons_per_secondResult *units.VolumeFlow
    us_gallons_per_secondResult, err = factory.FromDto(us_gallons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with UsGallonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_secondResult.Convert(units.VolumeFlowUsGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerSecond = %v, want %v", converted, 100)
    }
    // Test UsGallonPerMinute conversion
    us_gallons_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowUsGallonPerMinute,
    }
    
    var us_gallons_per_minuteResult *units.VolumeFlow
    us_gallons_per_minuteResult, err = factory.FromDto(us_gallons_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with UsGallonPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_minuteResult.Convert(units.VolumeFlowUsGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerMinute = %v, want %v", converted, 100)
    }
    // Test UkGallonPerDay conversion
    uk_gallons_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowUkGallonPerDay,
    }
    
    var uk_gallons_per_dayResult *units.VolumeFlow
    uk_gallons_per_dayResult, err = factory.FromDto(uk_gallons_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with UkGallonPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_gallons_per_dayResult.Convert(units.VolumeFlowUkGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkGallonPerDay = %v, want %v", converted, 100)
    }
    // Test UkGallonPerHour conversion
    uk_gallons_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowUkGallonPerHour,
    }
    
    var uk_gallons_per_hourResult *units.VolumeFlow
    uk_gallons_per_hourResult, err = factory.FromDto(uk_gallons_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with UkGallonPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_gallons_per_hourResult.Convert(units.VolumeFlowUkGallonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkGallonPerHour = %v, want %v", converted, 100)
    }
    // Test UkGallonPerMinute conversion
    uk_gallons_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowUkGallonPerMinute,
    }
    
    var uk_gallons_per_minuteResult *units.VolumeFlow
    uk_gallons_per_minuteResult, err = factory.FromDto(uk_gallons_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with UkGallonPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_gallons_per_minuteResult.Convert(units.VolumeFlowUkGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkGallonPerMinute = %v, want %v", converted, 100)
    }
    // Test UkGallonPerSecond conversion
    uk_gallons_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowUkGallonPerSecond,
    }
    
    var uk_gallons_per_secondResult *units.VolumeFlow
    uk_gallons_per_secondResult, err = factory.FromDto(uk_gallons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with UkGallonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_gallons_per_secondResult.Convert(units.VolumeFlowUkGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkGallonPerSecond = %v, want %v", converted, 100)
    }
    // Test KilousGallonPerMinute conversion
    kilous_gallons_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowKilousGallonPerMinute,
    }
    
    var kilous_gallons_per_minuteResult *units.VolumeFlow
    kilous_gallons_per_minuteResult, err = factory.FromDto(kilous_gallons_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KilousGallonPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilous_gallons_per_minuteResult.Convert(units.VolumeFlowKilousGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilousGallonPerMinute = %v, want %v", converted, 100)
    }
    // Test UsGallonPerHour conversion
    us_gallons_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowUsGallonPerHour,
    }
    
    var us_gallons_per_hourResult *units.VolumeFlow
    us_gallons_per_hourResult, err = factory.FromDto(us_gallons_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with UsGallonPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_hourResult.Convert(units.VolumeFlowUsGallonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerHour = %v, want %v", converted, 100)
    }
    // Test CubicDecimeterPerMinute conversion
    cubic_decimeters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicDecimeterPerMinute,
    }
    
    var cubic_decimeters_per_minuteResult *units.VolumeFlow
    cubic_decimeters_per_minuteResult, err = factory.FromDto(cubic_decimeters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with CubicDecimeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_decimeters_per_minuteResult.Convert(units.VolumeFlowCubicDecimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicDecimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test OilBarrelPerDay conversion
    oil_barrels_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowOilBarrelPerDay,
    }
    
    var oil_barrels_per_dayResult *units.VolumeFlow
    oil_barrels_per_dayResult, err = factory.FromDto(oil_barrels_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with OilBarrelPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_dayResult.Convert(units.VolumeFlowOilBarrelPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerDay = %v, want %v", converted, 100)
    }
    // Test OilBarrelPerMinute conversion
    oil_barrels_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowOilBarrelPerMinute,
    }
    
    var oil_barrels_per_minuteResult *units.VolumeFlow
    oil_barrels_per_minuteResult, err = factory.FromDto(oil_barrels_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with OilBarrelPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_minuteResult.Convert(units.VolumeFlowOilBarrelPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerMinute = %v, want %v", converted, 100)
    }
    // Test OilBarrelPerHour conversion
    oil_barrels_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowOilBarrelPerHour,
    }
    
    var oil_barrels_per_hourResult *units.VolumeFlow
    oil_barrels_per_hourResult, err = factory.FromDto(oil_barrels_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with OilBarrelPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_hourResult.Convert(units.VolumeFlowOilBarrelPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerHour = %v, want %v", converted, 100)
    }
    // Test OilBarrelPerSecond conversion
    oil_barrels_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowOilBarrelPerSecond,
    }
    
    var oil_barrels_per_secondResult *units.VolumeFlow
    oil_barrels_per_secondResult, err = factory.FromDto(oil_barrels_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with OilBarrelPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_secondResult.Convert(units.VolumeFlowOilBarrelPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerSecond = %v, want %v", converted, 100)
    }
    // Test CubicMillimeterPerSecond conversion
    cubic_millimeters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicMillimeterPerSecond,
    }
    
    var cubic_millimeters_per_secondResult *units.VolumeFlow
    cubic_millimeters_per_secondResult, err = factory.FromDto(cubic_millimeters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMillimeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_millimeters_per_secondResult.Convert(units.VolumeFlowCubicMillimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMillimeterPerSecond = %v, want %v", converted, 100)
    }
    // Test AcreFootPerSecond conversion
    acre_feet_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowAcreFootPerSecond,
    }
    
    var acre_feet_per_secondResult *units.VolumeFlow
    acre_feet_per_secondResult, err = factory.FromDto(acre_feet_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with AcreFootPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feet_per_secondResult.Convert(units.VolumeFlowAcreFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFootPerSecond = %v, want %v", converted, 100)
    }
    // Test AcreFootPerMinute conversion
    acre_feet_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowAcreFootPerMinute,
    }
    
    var acre_feet_per_minuteResult *units.VolumeFlow
    acre_feet_per_minuteResult, err = factory.FromDto(acre_feet_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with AcreFootPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feet_per_minuteResult.Convert(units.VolumeFlowAcreFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFootPerMinute = %v, want %v", converted, 100)
    }
    // Test AcreFootPerHour conversion
    acre_feet_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowAcreFootPerHour,
    }
    
    var acre_feet_per_hourResult *units.VolumeFlow
    acre_feet_per_hourResult, err = factory.FromDto(acre_feet_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with AcreFootPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feet_per_hourResult.Convert(units.VolumeFlowAcreFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFootPerHour = %v, want %v", converted, 100)
    }
    // Test AcreFootPerDay conversion
    acre_feet_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowAcreFootPerDay,
    }
    
    var acre_feet_per_dayResult *units.VolumeFlow
    acre_feet_per_dayResult, err = factory.FromDto(acre_feet_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with AcreFootPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feet_per_dayResult.Convert(units.VolumeFlowAcreFootPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFootPerDay = %v, want %v", converted, 100)
    }
    // Test CubicCentimeterPerMinute conversion
    cubic_centimeters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCubicCentimeterPerMinute,
    }
    
    var cubic_centimeters_per_minuteResult *units.VolumeFlow
    cubic_centimeters_per_minuteResult, err = factory.FromDto(cubic_centimeters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with CubicCentimeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_centimeters_per_minuteResult.Convert(units.VolumeFlowCubicCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicCentimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test MegausGallonPerDay conversion
    megaus_gallons_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMegausGallonPerDay,
    }
    
    var megaus_gallons_per_dayResult *units.VolumeFlow
    megaus_gallons_per_dayResult, err = factory.FromDto(megaus_gallons_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MegausGallonPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaus_gallons_per_dayResult.Convert(units.VolumeFlowMegausGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegausGallonPerDay = %v, want %v", converted, 100)
    }
    // Test NanoliterPerSecond conversion
    nanoliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowNanoliterPerSecond,
    }
    
    var nanoliters_per_secondResult *units.VolumeFlow
    nanoliters_per_secondResult, err = factory.FromDto(nanoliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NanoliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_secondResult.Convert(units.VolumeFlowNanoliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerSecond = %v, want %v", converted, 100)
    }
    // Test MicroliterPerSecond conversion
    microliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMicroliterPerSecond,
    }
    
    var microliters_per_secondResult *units.VolumeFlow
    microliters_per_secondResult, err = factory.FromDto(microliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicroliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_secondResult.Convert(units.VolumeFlowMicroliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerSecond = %v, want %v", converted, 100)
    }
    // Test MilliliterPerSecond conversion
    milliliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMilliliterPerSecond,
    }
    
    var milliliters_per_secondResult *units.VolumeFlow
    milliliters_per_secondResult, err = factory.FromDto(milliliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MilliliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_secondResult.Convert(units.VolumeFlowMilliliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerSecond = %v, want %v", converted, 100)
    }
    // Test CentiliterPerSecond conversion
    centiliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCentiliterPerSecond,
    }
    
    var centiliters_per_secondResult *units.VolumeFlow
    centiliters_per_secondResult, err = factory.FromDto(centiliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CentiliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_secondResult.Convert(units.VolumeFlowCentiliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerSecond = %v, want %v", converted, 100)
    }
    // Test DeciliterPerSecond conversion
    deciliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowDeciliterPerSecond,
    }
    
    var deciliters_per_secondResult *units.VolumeFlow
    deciliters_per_secondResult, err = factory.FromDto(deciliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DeciliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_secondResult.Convert(units.VolumeFlowDeciliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerSecond = %v, want %v", converted, 100)
    }
    // Test DecaliterPerSecond conversion
    decaliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowDecaliterPerSecond,
    }
    
    var decaliters_per_secondResult *units.VolumeFlow
    decaliters_per_secondResult, err = factory.FromDto(decaliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecaliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaliters_per_secondResult.Convert(units.VolumeFlowDecaliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecaliterPerSecond = %v, want %v", converted, 100)
    }
    // Test HectoliterPerSecond conversion
    hectoliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowHectoliterPerSecond,
    }
    
    var hectoliters_per_secondResult *units.VolumeFlow
    hectoliters_per_secondResult, err = factory.FromDto(hectoliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with HectoliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectoliters_per_secondResult.Convert(units.VolumeFlowHectoliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectoliterPerSecond = %v, want %v", converted, 100)
    }
    // Test KiloliterPerSecond conversion
    kiloliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowKiloliterPerSecond,
    }
    
    var kiloliters_per_secondResult *units.VolumeFlow
    kiloliters_per_secondResult, err = factory.FromDto(kiloliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KiloliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloliters_per_secondResult.Convert(units.VolumeFlowKiloliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloliterPerSecond = %v, want %v", converted, 100)
    }
    // Test MegaliterPerSecond conversion
    megaliters_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMegaliterPerSecond,
    }
    
    var megaliters_per_secondResult *units.VolumeFlow
    megaliters_per_secondResult, err = factory.FromDto(megaliters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegaliterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaliters_per_secondResult.Convert(units.VolumeFlowMegaliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaliterPerSecond = %v, want %v", converted, 100)
    }
    // Test NanoliterPerMinute conversion
    nanoliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowNanoliterPerMinute,
    }
    
    var nanoliters_per_minuteResult *units.VolumeFlow
    nanoliters_per_minuteResult, err = factory.FromDto(nanoliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with NanoliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_minuteResult.Convert(units.VolumeFlowNanoliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerMinute = %v, want %v", converted, 100)
    }
    // Test MicroliterPerMinute conversion
    microliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMicroliterPerMinute,
    }
    
    var microliters_per_minuteResult *units.VolumeFlow
    microliters_per_minuteResult, err = factory.FromDto(microliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MicroliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_minuteResult.Convert(units.VolumeFlowMicroliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerMinute = %v, want %v", converted, 100)
    }
    // Test MilliliterPerMinute conversion
    milliliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMilliliterPerMinute,
    }
    
    var milliliters_per_minuteResult *units.VolumeFlow
    milliliters_per_minuteResult, err = factory.FromDto(milliliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MilliliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_minuteResult.Convert(units.VolumeFlowMilliliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerMinute = %v, want %v", converted, 100)
    }
    // Test CentiliterPerMinute conversion
    centiliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCentiliterPerMinute,
    }
    
    var centiliters_per_minuteResult *units.VolumeFlow
    centiliters_per_minuteResult, err = factory.FromDto(centiliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with CentiliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_minuteResult.Convert(units.VolumeFlowCentiliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerMinute = %v, want %v", converted, 100)
    }
    // Test DeciliterPerMinute conversion
    deciliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowDeciliterPerMinute,
    }
    
    var deciliters_per_minuteResult *units.VolumeFlow
    deciliters_per_minuteResult, err = factory.FromDto(deciliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with DeciliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_minuteResult.Convert(units.VolumeFlowDeciliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerMinute = %v, want %v", converted, 100)
    }
    // Test DecaliterPerMinute conversion
    decaliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowDecaliterPerMinute,
    }
    
    var decaliters_per_minuteResult *units.VolumeFlow
    decaliters_per_minuteResult, err = factory.FromDto(decaliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with DecaliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaliters_per_minuteResult.Convert(units.VolumeFlowDecaliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecaliterPerMinute = %v, want %v", converted, 100)
    }
    // Test HectoliterPerMinute conversion
    hectoliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowHectoliterPerMinute,
    }
    
    var hectoliters_per_minuteResult *units.VolumeFlow
    hectoliters_per_minuteResult, err = factory.FromDto(hectoliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with HectoliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectoliters_per_minuteResult.Convert(units.VolumeFlowHectoliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectoliterPerMinute = %v, want %v", converted, 100)
    }
    // Test KiloliterPerMinute conversion
    kiloliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowKiloliterPerMinute,
    }
    
    var kiloliters_per_minuteResult *units.VolumeFlow
    kiloliters_per_minuteResult, err = factory.FromDto(kiloliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KiloliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloliters_per_minuteResult.Convert(units.VolumeFlowKiloliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloliterPerMinute = %v, want %v", converted, 100)
    }
    // Test MegaliterPerMinute conversion
    megaliters_per_minuteDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMegaliterPerMinute,
    }
    
    var megaliters_per_minuteResult *units.VolumeFlow
    megaliters_per_minuteResult, err = factory.FromDto(megaliters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MegaliterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaliters_per_minuteResult.Convert(units.VolumeFlowMegaliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaliterPerMinute = %v, want %v", converted, 100)
    }
    // Test NanoliterPerHour conversion
    nanoliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowNanoliterPerHour,
    }
    
    var nanoliters_per_hourResult *units.VolumeFlow
    nanoliters_per_hourResult, err = factory.FromDto(nanoliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with NanoliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_hourResult.Convert(units.VolumeFlowNanoliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerHour = %v, want %v", converted, 100)
    }
    // Test MicroliterPerHour conversion
    microliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMicroliterPerHour,
    }
    
    var microliters_per_hourResult *units.VolumeFlow
    microliters_per_hourResult, err = factory.FromDto(microliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MicroliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_hourResult.Convert(units.VolumeFlowMicroliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerHour = %v, want %v", converted, 100)
    }
    // Test MilliliterPerHour conversion
    milliliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMilliliterPerHour,
    }
    
    var milliliters_per_hourResult *units.VolumeFlow
    milliliters_per_hourResult, err = factory.FromDto(milliliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MilliliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_hourResult.Convert(units.VolumeFlowMilliliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerHour = %v, want %v", converted, 100)
    }
    // Test CentiliterPerHour conversion
    centiliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCentiliterPerHour,
    }
    
    var centiliters_per_hourResult *units.VolumeFlow
    centiliters_per_hourResult, err = factory.FromDto(centiliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with CentiliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_hourResult.Convert(units.VolumeFlowCentiliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerHour = %v, want %v", converted, 100)
    }
    // Test DeciliterPerHour conversion
    deciliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowDeciliterPerHour,
    }
    
    var deciliters_per_hourResult *units.VolumeFlow
    deciliters_per_hourResult, err = factory.FromDto(deciliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with DeciliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_hourResult.Convert(units.VolumeFlowDeciliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerHour = %v, want %v", converted, 100)
    }
    // Test DecaliterPerHour conversion
    decaliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowDecaliterPerHour,
    }
    
    var decaliters_per_hourResult *units.VolumeFlow
    decaliters_per_hourResult, err = factory.FromDto(decaliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with DecaliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaliters_per_hourResult.Convert(units.VolumeFlowDecaliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecaliterPerHour = %v, want %v", converted, 100)
    }
    // Test HectoliterPerHour conversion
    hectoliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowHectoliterPerHour,
    }
    
    var hectoliters_per_hourResult *units.VolumeFlow
    hectoliters_per_hourResult, err = factory.FromDto(hectoliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with HectoliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectoliters_per_hourResult.Convert(units.VolumeFlowHectoliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectoliterPerHour = %v, want %v", converted, 100)
    }
    // Test KiloliterPerHour conversion
    kiloliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowKiloliterPerHour,
    }
    
    var kiloliters_per_hourResult *units.VolumeFlow
    kiloliters_per_hourResult, err = factory.FromDto(kiloliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KiloliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloliters_per_hourResult.Convert(units.VolumeFlowKiloliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloliterPerHour = %v, want %v", converted, 100)
    }
    // Test MegaliterPerHour conversion
    megaliters_per_hourDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMegaliterPerHour,
    }
    
    var megaliters_per_hourResult *units.VolumeFlow
    megaliters_per_hourResult, err = factory.FromDto(megaliters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MegaliterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaliters_per_hourResult.Convert(units.VolumeFlowMegaliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaliterPerHour = %v, want %v", converted, 100)
    }
    // Test NanoliterPerDay conversion
    nanoliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowNanoliterPerDay,
    }
    
    var nanoliters_per_dayResult *units.VolumeFlow
    nanoliters_per_dayResult, err = factory.FromDto(nanoliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with NanoliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_dayResult.Convert(units.VolumeFlowNanoliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerDay = %v, want %v", converted, 100)
    }
    // Test MicroliterPerDay conversion
    microliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMicroliterPerDay,
    }
    
    var microliters_per_dayResult *units.VolumeFlow
    microliters_per_dayResult, err = factory.FromDto(microliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MicroliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_dayResult.Convert(units.VolumeFlowMicroliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerDay = %v, want %v", converted, 100)
    }
    // Test MilliliterPerDay conversion
    milliliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMilliliterPerDay,
    }
    
    var milliliters_per_dayResult *units.VolumeFlow
    milliliters_per_dayResult, err = factory.FromDto(milliliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MilliliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_dayResult.Convert(units.VolumeFlowMilliliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerDay = %v, want %v", converted, 100)
    }
    // Test CentiliterPerDay conversion
    centiliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowCentiliterPerDay,
    }
    
    var centiliters_per_dayResult *units.VolumeFlow
    centiliters_per_dayResult, err = factory.FromDto(centiliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with CentiliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_dayResult.Convert(units.VolumeFlowCentiliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerDay = %v, want %v", converted, 100)
    }
    // Test DeciliterPerDay conversion
    deciliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowDeciliterPerDay,
    }
    
    var deciliters_per_dayResult *units.VolumeFlow
    deciliters_per_dayResult, err = factory.FromDto(deciliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with DeciliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_dayResult.Convert(units.VolumeFlowDeciliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerDay = %v, want %v", converted, 100)
    }
    // Test DecaliterPerDay conversion
    decaliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowDecaliterPerDay,
    }
    
    var decaliters_per_dayResult *units.VolumeFlow
    decaliters_per_dayResult, err = factory.FromDto(decaliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with DecaliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaliters_per_dayResult.Convert(units.VolumeFlowDecaliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecaliterPerDay = %v, want %v", converted, 100)
    }
    // Test HectoliterPerDay conversion
    hectoliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowHectoliterPerDay,
    }
    
    var hectoliters_per_dayResult *units.VolumeFlow
    hectoliters_per_dayResult, err = factory.FromDto(hectoliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with HectoliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectoliters_per_dayResult.Convert(units.VolumeFlowHectoliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectoliterPerDay = %v, want %v", converted, 100)
    }
    // Test KiloliterPerDay conversion
    kiloliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowKiloliterPerDay,
    }
    
    var kiloliters_per_dayResult *units.VolumeFlow
    kiloliters_per_dayResult, err = factory.FromDto(kiloliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with KiloliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloliters_per_dayResult.Convert(units.VolumeFlowKiloliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloliterPerDay = %v, want %v", converted, 100)
    }
    // Test MegaliterPerDay conversion
    megaliters_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMegaliterPerDay,
    }
    
    var megaliters_per_dayResult *units.VolumeFlow
    megaliters_per_dayResult, err = factory.FromDto(megaliters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MegaliterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaliters_per_dayResult.Convert(units.VolumeFlowMegaliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaliterPerDay = %v, want %v", converted, 100)
    }
    // Test MegaukGallonPerDay conversion
    megauk_gallons_per_dayDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMegaukGallonPerDay,
    }
    
    var megauk_gallons_per_dayResult *units.VolumeFlow
    megauk_gallons_per_dayResult, err = factory.FromDto(megauk_gallons_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with MegaukGallonPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megauk_gallons_per_dayResult.Convert(units.VolumeFlowMegaukGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaukGallonPerDay = %v, want %v", converted, 100)
    }
    // Test MegaukGallonPerSecond conversion
    megauk_gallons_per_secondDto := units.VolumeFlowDto{
        Value: 100,
        Unit:  units.VolumeFlowMegaukGallonPerSecond,
    }
    
    var megauk_gallons_per_secondResult *units.VolumeFlow
    megauk_gallons_per_secondResult, err = factory.FromDto(megauk_gallons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegaukGallonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megauk_gallons_per_secondResult.Convert(units.VolumeFlowMegaukGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaukGallonPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.VolumeFlowDto{
        Value: 0,
        Unit:  units.VolumeFlowCubicMeterPerSecond,
    }
    
    var zeroResult *units.VolumeFlow
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestVolumeFlowFactory_FromDtoJSON(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CubicMeterPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CubicMeterPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.VolumeFlowDto{
        Value: nanValue,
        Unit:  units.VolumeFlowCubicMeterPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with CubicMeterPerSecond unit
    cubic_meters_per_secondJSON := []byte(`{"value": 100, "unit": "CubicMeterPerSecond"}`)
    cubic_meters_per_secondResult, err := factory.FromDtoJSON(cubic_meters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_secondResult.Convert(units.VolumeFlowCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CubicMeterPerMinute unit
    cubic_meters_per_minuteJSON := []byte(`{"value": 100, "unit": "CubicMeterPerMinute"}`)
    cubic_meters_per_minuteResult, err := factory.FromDtoJSON(cubic_meters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_minuteResult.Convert(units.VolumeFlowCubicMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with CubicMeterPerHour unit
    cubic_meters_per_hourJSON := []byte(`{"value": 100, "unit": "CubicMeterPerHour"}`)
    cubic_meters_per_hourResult, err := factory.FromDtoJSON(cubic_meters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMeterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_hourResult.Convert(units.VolumeFlowCubicMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with CubicMeterPerDay unit
    cubic_meters_per_dayJSON := []byte(`{"value": 100, "unit": "CubicMeterPerDay"}`)
    cubic_meters_per_dayResult, err := factory.FromDtoJSON(cubic_meters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMeterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_dayResult.Convert(units.VolumeFlowCubicMeterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with CubicFootPerSecond unit
    cubic_feet_per_secondJSON := []byte(`{"value": 100, "unit": "CubicFootPerSecond"}`)
    cubic_feet_per_secondResult, err := factory.FromDtoJSON(cubic_feet_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicFootPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_secondResult.Convert(units.VolumeFlowCubicFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CubicFootPerMinute unit
    cubic_feet_per_minuteJSON := []byte(`{"value": 100, "unit": "CubicFootPerMinute"}`)
    cubic_feet_per_minuteResult, err := factory.FromDtoJSON(cubic_feet_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicFootPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_minuteResult.Convert(units.VolumeFlowCubicFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with CubicFootPerHour unit
    cubic_feet_per_hourJSON := []byte(`{"value": 100, "unit": "CubicFootPerHour"}`)
    cubic_feet_per_hourResult, err := factory.FromDtoJSON(cubic_feet_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicFootPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_hourResult.Convert(units.VolumeFlowCubicFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with CubicYardPerSecond unit
    cubic_yards_per_secondJSON := []byte(`{"value": 100, "unit": "CubicYardPerSecond"}`)
    cubic_yards_per_secondResult, err := factory.FromDtoJSON(cubic_yards_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicYardPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_secondResult.Convert(units.VolumeFlowCubicYardPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CubicYardPerMinute unit
    cubic_yards_per_minuteJSON := []byte(`{"value": 100, "unit": "CubicYardPerMinute"}`)
    cubic_yards_per_minuteResult, err := factory.FromDtoJSON(cubic_yards_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicYardPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_minuteResult.Convert(units.VolumeFlowCubicYardPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with CubicYardPerHour unit
    cubic_yards_per_hourJSON := []byte(`{"value": 100, "unit": "CubicYardPerHour"}`)
    cubic_yards_per_hourResult, err := factory.FromDtoJSON(cubic_yards_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicYardPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_hourResult.Convert(units.VolumeFlowCubicYardPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with CubicYardPerDay unit
    cubic_yards_per_dayJSON := []byte(`{"value": 100, "unit": "CubicYardPerDay"}`)
    cubic_yards_per_dayResult, err := factory.FromDtoJSON(cubic_yards_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicYardPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_dayResult.Convert(units.VolumeFlowCubicYardPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MillionUsGallonPerDay unit
    million_us_gallons_per_dayJSON := []byte(`{"value": 100, "unit": "MillionUsGallonPerDay"}`)
    million_us_gallons_per_dayResult, err := factory.FromDtoJSON(million_us_gallons_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillionUsGallonPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = million_us_gallons_per_dayResult.Convert(units.VolumeFlowMillionUsGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillionUsGallonPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with UsGallonPerDay unit
    us_gallons_per_dayJSON := []byte(`{"value": 100, "unit": "UsGallonPerDay"}`)
    us_gallons_per_dayResult, err := factory.FromDtoJSON(us_gallons_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsGallonPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_dayResult.Convert(units.VolumeFlowUsGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerSecond unit
    liters_per_secondJSON := []byte(`{"value": 100, "unit": "LiterPerSecond"}`)
    liters_per_secondResult, err := factory.FromDtoJSON(liters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_secondResult.Convert(units.VolumeFlowLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerMinute unit
    liters_per_minuteJSON := []byte(`{"value": 100, "unit": "LiterPerMinute"}`)
    liters_per_minuteResult, err := factory.FromDtoJSON(liters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_minuteResult.Convert(units.VolumeFlowLiterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerHour unit
    liters_per_hourJSON := []byte(`{"value": 100, "unit": "LiterPerHour"}`)
    liters_per_hourResult, err := factory.FromDtoJSON(liters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_hourResult.Convert(units.VolumeFlowLiterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerDay unit
    liters_per_dayJSON := []byte(`{"value": 100, "unit": "LiterPerDay"}`)
    liters_per_dayResult, err := factory.FromDtoJSON(liters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_dayResult.Convert(units.VolumeFlowLiterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with UsGallonPerSecond unit
    us_gallons_per_secondJSON := []byte(`{"value": 100, "unit": "UsGallonPerSecond"}`)
    us_gallons_per_secondResult, err := factory.FromDtoJSON(us_gallons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsGallonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_secondResult.Convert(units.VolumeFlowUsGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with UsGallonPerMinute unit
    us_gallons_per_minuteJSON := []byte(`{"value": 100, "unit": "UsGallonPerMinute"}`)
    us_gallons_per_minuteResult, err := factory.FromDtoJSON(us_gallons_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsGallonPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_minuteResult.Convert(units.VolumeFlowUsGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with UkGallonPerDay unit
    uk_gallons_per_dayJSON := []byte(`{"value": 100, "unit": "UkGallonPerDay"}`)
    uk_gallons_per_dayResult, err := factory.FromDtoJSON(uk_gallons_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UkGallonPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_gallons_per_dayResult.Convert(units.VolumeFlowUkGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkGallonPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with UkGallonPerHour unit
    uk_gallons_per_hourJSON := []byte(`{"value": 100, "unit": "UkGallonPerHour"}`)
    uk_gallons_per_hourResult, err := factory.FromDtoJSON(uk_gallons_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UkGallonPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_gallons_per_hourResult.Convert(units.VolumeFlowUkGallonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkGallonPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with UkGallonPerMinute unit
    uk_gallons_per_minuteJSON := []byte(`{"value": 100, "unit": "UkGallonPerMinute"}`)
    uk_gallons_per_minuteResult, err := factory.FromDtoJSON(uk_gallons_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UkGallonPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_gallons_per_minuteResult.Convert(units.VolumeFlowUkGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkGallonPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with UkGallonPerSecond unit
    uk_gallons_per_secondJSON := []byte(`{"value": 100, "unit": "UkGallonPerSecond"}`)
    uk_gallons_per_secondResult, err := factory.FromDtoJSON(uk_gallons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UkGallonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_gallons_per_secondResult.Convert(units.VolumeFlowUkGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkGallonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilousGallonPerMinute unit
    kilous_gallons_per_minuteJSON := []byte(`{"value": 100, "unit": "KilousGallonPerMinute"}`)
    kilous_gallons_per_minuteResult, err := factory.FromDtoJSON(kilous_gallons_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilousGallonPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilous_gallons_per_minuteResult.Convert(units.VolumeFlowKilousGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilousGallonPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with UsGallonPerHour unit
    us_gallons_per_hourJSON := []byte(`{"value": 100, "unit": "UsGallonPerHour"}`)
    us_gallons_per_hourResult, err := factory.FromDtoJSON(us_gallons_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsGallonPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_hourResult.Convert(units.VolumeFlowUsGallonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with CubicDecimeterPerMinute unit
    cubic_decimeters_per_minuteJSON := []byte(`{"value": 100, "unit": "CubicDecimeterPerMinute"}`)
    cubic_decimeters_per_minuteResult, err := factory.FromDtoJSON(cubic_decimeters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicDecimeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_decimeters_per_minuteResult.Convert(units.VolumeFlowCubicDecimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicDecimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with OilBarrelPerDay unit
    oil_barrels_per_dayJSON := []byte(`{"value": 100, "unit": "OilBarrelPerDay"}`)
    oil_barrels_per_dayResult, err := factory.FromDtoJSON(oil_barrels_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OilBarrelPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_dayResult.Convert(units.VolumeFlowOilBarrelPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with OilBarrelPerMinute unit
    oil_barrels_per_minuteJSON := []byte(`{"value": 100, "unit": "OilBarrelPerMinute"}`)
    oil_barrels_per_minuteResult, err := factory.FromDtoJSON(oil_barrels_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OilBarrelPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_minuteResult.Convert(units.VolumeFlowOilBarrelPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with OilBarrelPerHour unit
    oil_barrels_per_hourJSON := []byte(`{"value": 100, "unit": "OilBarrelPerHour"}`)
    oil_barrels_per_hourResult, err := factory.FromDtoJSON(oil_barrels_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OilBarrelPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_hourResult.Convert(units.VolumeFlowOilBarrelPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with OilBarrelPerSecond unit
    oil_barrels_per_secondJSON := []byte(`{"value": 100, "unit": "OilBarrelPerSecond"}`)
    oil_barrels_per_secondResult, err := factory.FromDtoJSON(oil_barrels_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OilBarrelPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_secondResult.Convert(units.VolumeFlowOilBarrelPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CubicMillimeterPerSecond unit
    cubic_millimeters_per_secondJSON := []byte(`{"value": 100, "unit": "CubicMillimeterPerSecond"}`)
    cubic_millimeters_per_secondResult, err := factory.FromDtoJSON(cubic_millimeters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMillimeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_millimeters_per_secondResult.Convert(units.VolumeFlowCubicMillimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMillimeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with AcreFootPerSecond unit
    acre_feet_per_secondJSON := []byte(`{"value": 100, "unit": "AcreFootPerSecond"}`)
    acre_feet_per_secondResult, err := factory.FromDtoJSON(acre_feet_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AcreFootPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feet_per_secondResult.Convert(units.VolumeFlowAcreFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFootPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with AcreFootPerMinute unit
    acre_feet_per_minuteJSON := []byte(`{"value": 100, "unit": "AcreFootPerMinute"}`)
    acre_feet_per_minuteResult, err := factory.FromDtoJSON(acre_feet_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AcreFootPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feet_per_minuteResult.Convert(units.VolumeFlowAcreFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFootPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with AcreFootPerHour unit
    acre_feet_per_hourJSON := []byte(`{"value": 100, "unit": "AcreFootPerHour"}`)
    acre_feet_per_hourResult, err := factory.FromDtoJSON(acre_feet_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AcreFootPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feet_per_hourResult.Convert(units.VolumeFlowAcreFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFootPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with AcreFootPerDay unit
    acre_feet_per_dayJSON := []byte(`{"value": 100, "unit": "AcreFootPerDay"}`)
    acre_feet_per_dayResult, err := factory.FromDtoJSON(acre_feet_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AcreFootPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feet_per_dayResult.Convert(units.VolumeFlowAcreFootPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFootPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with CubicCentimeterPerMinute unit
    cubic_centimeters_per_minuteJSON := []byte(`{"value": 100, "unit": "CubicCentimeterPerMinute"}`)
    cubic_centimeters_per_minuteResult, err := factory.FromDtoJSON(cubic_centimeters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicCentimeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_centimeters_per_minuteResult.Convert(units.VolumeFlowCubicCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicCentimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MegausGallonPerDay unit
    megaus_gallons_per_dayJSON := []byte(`{"value": 100, "unit": "MegausGallonPerDay"}`)
    megaus_gallons_per_dayResult, err := factory.FromDtoJSON(megaus_gallons_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegausGallonPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaus_gallons_per_dayResult.Convert(units.VolumeFlowMegausGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegausGallonPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with NanoliterPerSecond unit
    nanoliters_per_secondJSON := []byte(`{"value": 100, "unit": "NanoliterPerSecond"}`)
    nanoliters_per_secondResult, err := factory.FromDtoJSON(nanoliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_secondResult.Convert(units.VolumeFlowNanoliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicroliterPerSecond unit
    microliters_per_secondJSON := []byte(`{"value": 100, "unit": "MicroliterPerSecond"}`)
    microliters_per_secondResult, err := factory.FromDtoJSON(microliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_secondResult.Convert(units.VolumeFlowMicroliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MilliliterPerSecond unit
    milliliters_per_secondJSON := []byte(`{"value": 100, "unit": "MilliliterPerSecond"}`)
    milliliters_per_secondResult, err := factory.FromDtoJSON(milliliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_secondResult.Convert(units.VolumeFlowMilliliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CentiliterPerSecond unit
    centiliters_per_secondJSON := []byte(`{"value": 100, "unit": "CentiliterPerSecond"}`)
    centiliters_per_secondResult, err := factory.FromDtoJSON(centiliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentiliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_secondResult.Convert(units.VolumeFlowCentiliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DeciliterPerSecond unit
    deciliters_per_secondJSON := []byte(`{"value": 100, "unit": "DeciliterPerSecond"}`)
    deciliters_per_secondResult, err := factory.FromDtoJSON(deciliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_secondResult.Convert(units.VolumeFlowDeciliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecaliterPerSecond unit
    decaliters_per_secondJSON := []byte(`{"value": 100, "unit": "DecaliterPerSecond"}`)
    decaliters_per_secondResult, err := factory.FromDtoJSON(decaliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecaliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaliters_per_secondResult.Convert(units.VolumeFlowDecaliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecaliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with HectoliterPerSecond unit
    hectoliters_per_secondJSON := []byte(`{"value": 100, "unit": "HectoliterPerSecond"}`)
    hectoliters_per_secondResult, err := factory.FromDtoJSON(hectoliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectoliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectoliters_per_secondResult.Convert(units.VolumeFlowHectoliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectoliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KiloliterPerSecond unit
    kiloliters_per_secondJSON := []byte(`{"value": 100, "unit": "KiloliterPerSecond"}`)
    kiloliters_per_secondResult, err := factory.FromDtoJSON(kiloliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloliters_per_secondResult.Convert(units.VolumeFlowKiloliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegaliterPerSecond unit
    megaliters_per_secondJSON := []byte(`{"value": 100, "unit": "MegaliterPerSecond"}`)
    megaliters_per_secondResult, err := factory.FromDtoJSON(megaliters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaliterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaliters_per_secondResult.Convert(units.VolumeFlowMegaliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaliterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with NanoliterPerMinute unit
    nanoliters_per_minuteJSON := []byte(`{"value": 100, "unit": "NanoliterPerMinute"}`)
    nanoliters_per_minuteResult, err := factory.FromDtoJSON(nanoliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_minuteResult.Convert(units.VolumeFlowNanoliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MicroliterPerMinute unit
    microliters_per_minuteJSON := []byte(`{"value": 100, "unit": "MicroliterPerMinute"}`)
    microliters_per_minuteResult, err := factory.FromDtoJSON(microliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_minuteResult.Convert(units.VolumeFlowMicroliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MilliliterPerMinute unit
    milliliters_per_minuteJSON := []byte(`{"value": 100, "unit": "MilliliterPerMinute"}`)
    milliliters_per_minuteResult, err := factory.FromDtoJSON(milliliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_minuteResult.Convert(units.VolumeFlowMilliliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with CentiliterPerMinute unit
    centiliters_per_minuteJSON := []byte(`{"value": 100, "unit": "CentiliterPerMinute"}`)
    centiliters_per_minuteResult, err := factory.FromDtoJSON(centiliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentiliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_minuteResult.Convert(units.VolumeFlowCentiliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with DeciliterPerMinute unit
    deciliters_per_minuteJSON := []byte(`{"value": 100, "unit": "DeciliterPerMinute"}`)
    deciliters_per_minuteResult, err := factory.FromDtoJSON(deciliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_minuteResult.Convert(units.VolumeFlowDeciliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with DecaliterPerMinute unit
    decaliters_per_minuteJSON := []byte(`{"value": 100, "unit": "DecaliterPerMinute"}`)
    decaliters_per_minuteResult, err := factory.FromDtoJSON(decaliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecaliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaliters_per_minuteResult.Convert(units.VolumeFlowDecaliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecaliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with HectoliterPerMinute unit
    hectoliters_per_minuteJSON := []byte(`{"value": 100, "unit": "HectoliterPerMinute"}`)
    hectoliters_per_minuteResult, err := factory.FromDtoJSON(hectoliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectoliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectoliters_per_minuteResult.Convert(units.VolumeFlowHectoliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectoliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KiloliterPerMinute unit
    kiloliters_per_minuteJSON := []byte(`{"value": 100, "unit": "KiloliterPerMinute"}`)
    kiloliters_per_minuteResult, err := factory.FromDtoJSON(kiloliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloliters_per_minuteResult.Convert(units.VolumeFlowKiloliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MegaliterPerMinute unit
    megaliters_per_minuteJSON := []byte(`{"value": 100, "unit": "MegaliterPerMinute"}`)
    megaliters_per_minuteResult, err := factory.FromDtoJSON(megaliters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaliterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaliters_per_minuteResult.Convert(units.VolumeFlowMegaliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaliterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with NanoliterPerHour unit
    nanoliters_per_hourJSON := []byte(`{"value": 100, "unit": "NanoliterPerHour"}`)
    nanoliters_per_hourResult, err := factory.FromDtoJSON(nanoliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_hourResult.Convert(units.VolumeFlowNanoliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MicroliterPerHour unit
    microliters_per_hourJSON := []byte(`{"value": 100, "unit": "MicroliterPerHour"}`)
    microliters_per_hourResult, err := factory.FromDtoJSON(microliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_hourResult.Convert(units.VolumeFlowMicroliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MilliliterPerHour unit
    milliliters_per_hourJSON := []byte(`{"value": 100, "unit": "MilliliterPerHour"}`)
    milliliters_per_hourResult, err := factory.FromDtoJSON(milliliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_hourResult.Convert(units.VolumeFlowMilliliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with CentiliterPerHour unit
    centiliters_per_hourJSON := []byte(`{"value": 100, "unit": "CentiliterPerHour"}`)
    centiliters_per_hourResult, err := factory.FromDtoJSON(centiliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentiliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_hourResult.Convert(units.VolumeFlowCentiliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with DeciliterPerHour unit
    deciliters_per_hourJSON := []byte(`{"value": 100, "unit": "DeciliterPerHour"}`)
    deciliters_per_hourResult, err := factory.FromDtoJSON(deciliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_hourResult.Convert(units.VolumeFlowDeciliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with DecaliterPerHour unit
    decaliters_per_hourJSON := []byte(`{"value": 100, "unit": "DecaliterPerHour"}`)
    decaliters_per_hourResult, err := factory.FromDtoJSON(decaliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecaliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaliters_per_hourResult.Convert(units.VolumeFlowDecaliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecaliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with HectoliterPerHour unit
    hectoliters_per_hourJSON := []byte(`{"value": 100, "unit": "HectoliterPerHour"}`)
    hectoliters_per_hourResult, err := factory.FromDtoJSON(hectoliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectoliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectoliters_per_hourResult.Convert(units.VolumeFlowHectoliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectoliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with KiloliterPerHour unit
    kiloliters_per_hourJSON := []byte(`{"value": 100, "unit": "KiloliterPerHour"}`)
    kiloliters_per_hourResult, err := factory.FromDtoJSON(kiloliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloliters_per_hourResult.Convert(units.VolumeFlowKiloliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegaliterPerHour unit
    megaliters_per_hourJSON := []byte(`{"value": 100, "unit": "MegaliterPerHour"}`)
    megaliters_per_hourResult, err := factory.FromDtoJSON(megaliters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaliterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaliters_per_hourResult.Convert(units.VolumeFlowMegaliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaliterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with NanoliterPerDay unit
    nanoliters_per_dayJSON := []byte(`{"value": 100, "unit": "NanoliterPerDay"}`)
    nanoliters_per_dayResult, err := factory.FromDtoJSON(nanoliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_dayResult.Convert(units.VolumeFlowNanoliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MicroliterPerDay unit
    microliters_per_dayJSON := []byte(`{"value": 100, "unit": "MicroliterPerDay"}`)
    microliters_per_dayResult, err := factory.FromDtoJSON(microliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_dayResult.Convert(units.VolumeFlowMicroliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MilliliterPerDay unit
    milliliters_per_dayJSON := []byte(`{"value": 100, "unit": "MilliliterPerDay"}`)
    milliliters_per_dayResult, err := factory.FromDtoJSON(milliliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_dayResult.Convert(units.VolumeFlowMilliliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with CentiliterPerDay unit
    centiliters_per_dayJSON := []byte(`{"value": 100, "unit": "CentiliterPerDay"}`)
    centiliters_per_dayResult, err := factory.FromDtoJSON(centiliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentiliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_dayResult.Convert(units.VolumeFlowCentiliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with DeciliterPerDay unit
    deciliters_per_dayJSON := []byte(`{"value": 100, "unit": "DeciliterPerDay"}`)
    deciliters_per_dayResult, err := factory.FromDtoJSON(deciliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_dayResult.Convert(units.VolumeFlowDeciliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with DecaliterPerDay unit
    decaliters_per_dayJSON := []byte(`{"value": 100, "unit": "DecaliterPerDay"}`)
    decaliters_per_dayResult, err := factory.FromDtoJSON(decaliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecaliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaliters_per_dayResult.Convert(units.VolumeFlowDecaliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecaliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with HectoliterPerDay unit
    hectoliters_per_dayJSON := []byte(`{"value": 100, "unit": "HectoliterPerDay"}`)
    hectoliters_per_dayResult, err := factory.FromDtoJSON(hectoliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectoliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectoliters_per_dayResult.Convert(units.VolumeFlowHectoliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectoliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with KiloliterPerDay unit
    kiloliters_per_dayJSON := []byte(`{"value": 100, "unit": "KiloliterPerDay"}`)
    kiloliters_per_dayResult, err := factory.FromDtoJSON(kiloliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloliters_per_dayResult.Convert(units.VolumeFlowKiloliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MegaliterPerDay unit
    megaliters_per_dayJSON := []byte(`{"value": 100, "unit": "MegaliterPerDay"}`)
    megaliters_per_dayResult, err := factory.FromDtoJSON(megaliters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaliterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaliters_per_dayResult.Convert(units.VolumeFlowMegaliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaliterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MegaukGallonPerDay unit
    megauk_gallons_per_dayJSON := []byte(`{"value": 100, "unit": "MegaukGallonPerDay"}`)
    megauk_gallons_per_dayResult, err := factory.FromDtoJSON(megauk_gallons_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaukGallonPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megauk_gallons_per_dayResult.Convert(units.VolumeFlowMegaukGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaukGallonPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with MegaukGallonPerSecond unit
    megauk_gallons_per_secondJSON := []byte(`{"value": 100, "unit": "MegaukGallonPerSecond"}`)
    megauk_gallons_per_secondResult, err := factory.FromDtoJSON(megauk_gallons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaukGallonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megauk_gallons_per_secondResult.Convert(units.VolumeFlowMegaukGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaukGallonPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CubicMeterPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCubicMetersPerSecond function
func TestVolumeFlowFactory_FromCubicMetersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMetersPerSecond(100)
    if err != nil {
        t.Errorf("FromCubicMetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCubicMetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCubicMetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCubicMetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMetersPerSecond(0)
    if err != nil {
        t.Errorf("FromCubicMetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicMeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicMetersPerMinute function
func TestVolumeFlowFactory_FromCubicMetersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMetersPerMinute(100)
    if err != nil {
        t.Errorf("FromCubicMetersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMetersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMetersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromCubicMetersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromCubicMetersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMetersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromCubicMetersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMetersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMetersPerMinute(0)
    if err != nil {
        t.Errorf("FromCubicMetersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicMeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMetersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicMetersPerHour function
func TestVolumeFlowFactory_FromCubicMetersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMetersPerHour(100)
    if err != nil {
        t.Errorf("FromCubicMetersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMetersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMetersPerHour(math.NaN())
    if err == nil {
        t.Error("FromCubicMetersPerHour() with NaN value should return error")
    }

    _, err = factory.FromCubicMetersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMetersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromCubicMetersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMetersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMetersPerHour(0)
    if err != nil {
        t.Errorf("FromCubicMetersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicMeterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMetersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicMetersPerDay function
func TestVolumeFlowFactory_FromCubicMetersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMetersPerDay(100)
    if err != nil {
        t.Errorf("FromCubicMetersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicMeterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMetersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMetersPerDay(math.NaN())
    if err == nil {
        t.Error("FromCubicMetersPerDay() with NaN value should return error")
    }

    _, err = factory.FromCubicMetersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMetersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromCubicMetersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMetersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMetersPerDay(0)
    if err != nil {
        t.Errorf("FromCubicMetersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicMeterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMetersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicFeetPerSecond function
func TestVolumeFlowFactory_FromCubicFeetPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicFeetPerSecond(100)
    if err != nil {
        t.Errorf("FromCubicFeetPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicFeetPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicFeetPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCubicFeetPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCubicFeetPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCubicFeetPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCubicFeetPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicFeetPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicFeetPerSecond(0)
    if err != nil {
        t.Errorf("FromCubicFeetPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicFootPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicFeetPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicFeetPerMinute function
func TestVolumeFlowFactory_FromCubicFeetPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicFeetPerMinute(100)
    if err != nil {
        t.Errorf("FromCubicFeetPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicFeetPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicFeetPerMinute(math.NaN())
    if err == nil {
        t.Error("FromCubicFeetPerMinute() with NaN value should return error")
    }

    _, err = factory.FromCubicFeetPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromCubicFeetPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromCubicFeetPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicFeetPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicFeetPerMinute(0)
    if err != nil {
        t.Errorf("FromCubicFeetPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicFootPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicFeetPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicFeetPerHour function
func TestVolumeFlowFactory_FromCubicFeetPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicFeetPerHour(100)
    if err != nil {
        t.Errorf("FromCubicFeetPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicFeetPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicFeetPerHour(math.NaN())
    if err == nil {
        t.Error("FromCubicFeetPerHour() with NaN value should return error")
    }

    _, err = factory.FromCubicFeetPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromCubicFeetPerHour() with +Inf value should return error")
    }

    _, err = factory.FromCubicFeetPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicFeetPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicFeetPerHour(0)
    if err != nil {
        t.Errorf("FromCubicFeetPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicFootPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicFeetPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicYardsPerSecond function
func TestVolumeFlowFactory_FromCubicYardsPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicYardsPerSecond(100)
    if err != nil {
        t.Errorf("FromCubicYardsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicYardPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicYardsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicYardsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCubicYardsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCubicYardsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCubicYardsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCubicYardsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicYardsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicYardsPerSecond(0)
    if err != nil {
        t.Errorf("FromCubicYardsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicYardPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicYardsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicYardsPerMinute function
func TestVolumeFlowFactory_FromCubicYardsPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicYardsPerMinute(100)
    if err != nil {
        t.Errorf("FromCubicYardsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicYardPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicYardsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicYardsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromCubicYardsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromCubicYardsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromCubicYardsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromCubicYardsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicYardsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicYardsPerMinute(0)
    if err != nil {
        t.Errorf("FromCubicYardsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicYardPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicYardsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicYardsPerHour function
func TestVolumeFlowFactory_FromCubicYardsPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicYardsPerHour(100)
    if err != nil {
        t.Errorf("FromCubicYardsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicYardPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicYardsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicYardsPerHour(math.NaN())
    if err == nil {
        t.Error("FromCubicYardsPerHour() with NaN value should return error")
    }

    _, err = factory.FromCubicYardsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromCubicYardsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromCubicYardsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicYardsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicYardsPerHour(0)
    if err != nil {
        t.Errorf("FromCubicYardsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicYardPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicYardsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicYardsPerDay function
func TestVolumeFlowFactory_FromCubicYardsPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicYardsPerDay(100)
    if err != nil {
        t.Errorf("FromCubicYardsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicYardPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicYardsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicYardsPerDay(math.NaN())
    if err == nil {
        t.Error("FromCubicYardsPerDay() with NaN value should return error")
    }

    _, err = factory.FromCubicYardsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromCubicYardsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromCubicYardsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicYardsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicYardsPerDay(0)
    if err != nil {
        t.Errorf("FromCubicYardsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicYardPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicYardsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMillionUsGallonsPerDay function
func TestVolumeFlowFactory_FromMillionUsGallonsPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillionUsGallonsPerDay(100)
    if err != nil {
        t.Errorf("FromMillionUsGallonsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMillionUsGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillionUsGallonsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillionUsGallonsPerDay(math.NaN())
    if err == nil {
        t.Error("FromMillionUsGallonsPerDay() with NaN value should return error")
    }

    _, err = factory.FromMillionUsGallonsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMillionUsGallonsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMillionUsGallonsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMillionUsGallonsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillionUsGallonsPerDay(0)
    if err != nil {
        t.Errorf("FromMillionUsGallonsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMillionUsGallonPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillionUsGallonsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromUsGallonsPerDay function
func TestVolumeFlowFactory_FromUsGallonsPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsGallonsPerDay(100)
    if err != nil {
        t.Errorf("FromUsGallonsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowUsGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsGallonsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsGallonsPerDay(math.NaN())
    if err == nil {
        t.Error("FromUsGallonsPerDay() with NaN value should return error")
    }

    _, err = factory.FromUsGallonsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromUsGallonsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromUsGallonsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromUsGallonsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsGallonsPerDay(0)
    if err != nil {
        t.Errorf("FromUsGallonsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowUsGallonPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsGallonsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerSecond function
func TestVolumeFlowFactory_FromLitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerSecond(100)
    if err != nil {
        t.Errorf("FromLitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromLitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromLitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerSecond(0)
    if err != nil {
        t.Errorf("FromLitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowLiterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerMinute function
func TestVolumeFlowFactory_FromLitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerMinute(100)
    if err != nil {
        t.Errorf("FromLitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowLiterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromLitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromLitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerMinute(0)
    if err != nil {
        t.Errorf("FromLitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowLiterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerHour function
func TestVolumeFlowFactory_FromLitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerHour(100)
    if err != nil {
        t.Errorf("FromLitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowLiterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromLitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromLitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerHour(0)
    if err != nil {
        t.Errorf("FromLitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowLiterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerDay function
func TestVolumeFlowFactory_FromLitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerDay(100)
    if err != nil {
        t.Errorf("FromLitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowLiterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromLitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromLitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerDay(0)
    if err != nil {
        t.Errorf("FromLitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowLiterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromUsGallonsPerSecond function
func TestVolumeFlowFactory_FromUsGallonsPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsGallonsPerSecond(100)
    if err != nil {
        t.Errorf("FromUsGallonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowUsGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsGallonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsGallonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromUsGallonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromUsGallonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromUsGallonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromUsGallonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromUsGallonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsGallonsPerSecond(0)
    if err != nil {
        t.Errorf("FromUsGallonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowUsGallonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsGallonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromUsGallonsPerMinute function
func TestVolumeFlowFactory_FromUsGallonsPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsGallonsPerMinute(100)
    if err != nil {
        t.Errorf("FromUsGallonsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowUsGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsGallonsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsGallonsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromUsGallonsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromUsGallonsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromUsGallonsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromUsGallonsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromUsGallonsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsGallonsPerMinute(0)
    if err != nil {
        t.Errorf("FromUsGallonsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowUsGallonPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsGallonsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromUkGallonsPerDay function
func TestVolumeFlowFactory_FromUkGallonsPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUkGallonsPerDay(100)
    if err != nil {
        t.Errorf("FromUkGallonsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowUkGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUkGallonsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUkGallonsPerDay(math.NaN())
    if err == nil {
        t.Error("FromUkGallonsPerDay() with NaN value should return error")
    }

    _, err = factory.FromUkGallonsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromUkGallonsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromUkGallonsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromUkGallonsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUkGallonsPerDay(0)
    if err != nil {
        t.Errorf("FromUkGallonsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowUkGallonPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUkGallonsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromUkGallonsPerHour function
func TestVolumeFlowFactory_FromUkGallonsPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUkGallonsPerHour(100)
    if err != nil {
        t.Errorf("FromUkGallonsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowUkGallonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUkGallonsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUkGallonsPerHour(math.NaN())
    if err == nil {
        t.Error("FromUkGallonsPerHour() with NaN value should return error")
    }

    _, err = factory.FromUkGallonsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromUkGallonsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromUkGallonsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromUkGallonsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUkGallonsPerHour(0)
    if err != nil {
        t.Errorf("FromUkGallonsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowUkGallonPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUkGallonsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromUkGallonsPerMinute function
func TestVolumeFlowFactory_FromUkGallonsPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUkGallonsPerMinute(100)
    if err != nil {
        t.Errorf("FromUkGallonsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowUkGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUkGallonsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUkGallonsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromUkGallonsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromUkGallonsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromUkGallonsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromUkGallonsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromUkGallonsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUkGallonsPerMinute(0)
    if err != nil {
        t.Errorf("FromUkGallonsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowUkGallonPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUkGallonsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromUkGallonsPerSecond function
func TestVolumeFlowFactory_FromUkGallonsPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUkGallonsPerSecond(100)
    if err != nil {
        t.Errorf("FromUkGallonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowUkGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUkGallonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUkGallonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromUkGallonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromUkGallonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromUkGallonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromUkGallonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromUkGallonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUkGallonsPerSecond(0)
    if err != nil {
        t.Errorf("FromUkGallonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowUkGallonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUkGallonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilousGallonsPerMinute function
func TestVolumeFlowFactory_FromKilousGallonsPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilousGallonsPerMinute(100)
    if err != nil {
        t.Errorf("FromKilousGallonsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowKilousGallonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilousGallonsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilousGallonsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilousGallonsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilousGallonsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilousGallonsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilousGallonsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilousGallonsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilousGallonsPerMinute(0)
    if err != nil {
        t.Errorf("FromKilousGallonsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowKilousGallonPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilousGallonsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromUsGallonsPerHour function
func TestVolumeFlowFactory_FromUsGallonsPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsGallonsPerHour(100)
    if err != nil {
        t.Errorf("FromUsGallonsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowUsGallonPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsGallonsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsGallonsPerHour(math.NaN())
    if err == nil {
        t.Error("FromUsGallonsPerHour() with NaN value should return error")
    }

    _, err = factory.FromUsGallonsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromUsGallonsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromUsGallonsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromUsGallonsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsGallonsPerHour(0)
    if err != nil {
        t.Errorf("FromUsGallonsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowUsGallonPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsGallonsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicDecimetersPerMinute function
func TestVolumeFlowFactory_FromCubicDecimetersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicDecimetersPerMinute(100)
    if err != nil {
        t.Errorf("FromCubicDecimetersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicDecimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicDecimetersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicDecimetersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromCubicDecimetersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromCubicDecimetersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromCubicDecimetersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromCubicDecimetersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicDecimetersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicDecimetersPerMinute(0)
    if err != nil {
        t.Errorf("FromCubicDecimetersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicDecimeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicDecimetersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromOilBarrelsPerDay function
func TestVolumeFlowFactory_FromOilBarrelsPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOilBarrelsPerDay(100)
    if err != nil {
        t.Errorf("FromOilBarrelsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowOilBarrelPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOilBarrelsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOilBarrelsPerDay(math.NaN())
    if err == nil {
        t.Error("FromOilBarrelsPerDay() with NaN value should return error")
    }

    _, err = factory.FromOilBarrelsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromOilBarrelsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromOilBarrelsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromOilBarrelsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOilBarrelsPerDay(0)
    if err != nil {
        t.Errorf("FromOilBarrelsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowOilBarrelPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOilBarrelsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromOilBarrelsPerMinute function
func TestVolumeFlowFactory_FromOilBarrelsPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOilBarrelsPerMinute(100)
    if err != nil {
        t.Errorf("FromOilBarrelsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowOilBarrelPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOilBarrelsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOilBarrelsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromOilBarrelsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromOilBarrelsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromOilBarrelsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromOilBarrelsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromOilBarrelsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOilBarrelsPerMinute(0)
    if err != nil {
        t.Errorf("FromOilBarrelsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowOilBarrelPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOilBarrelsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromOilBarrelsPerHour function
func TestVolumeFlowFactory_FromOilBarrelsPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOilBarrelsPerHour(100)
    if err != nil {
        t.Errorf("FromOilBarrelsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowOilBarrelPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOilBarrelsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOilBarrelsPerHour(math.NaN())
    if err == nil {
        t.Error("FromOilBarrelsPerHour() with NaN value should return error")
    }

    _, err = factory.FromOilBarrelsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromOilBarrelsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromOilBarrelsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromOilBarrelsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOilBarrelsPerHour(0)
    if err != nil {
        t.Errorf("FromOilBarrelsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowOilBarrelPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOilBarrelsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromOilBarrelsPerSecond function
func TestVolumeFlowFactory_FromOilBarrelsPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOilBarrelsPerSecond(100)
    if err != nil {
        t.Errorf("FromOilBarrelsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowOilBarrelPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOilBarrelsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOilBarrelsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromOilBarrelsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromOilBarrelsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromOilBarrelsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromOilBarrelsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromOilBarrelsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOilBarrelsPerSecond(0)
    if err != nil {
        t.Errorf("FromOilBarrelsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowOilBarrelPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOilBarrelsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicMillimetersPerSecond function
func TestVolumeFlowFactory_FromCubicMillimetersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMillimetersPerSecond(100)
    if err != nil {
        t.Errorf("FromCubicMillimetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicMillimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMillimetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMillimetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCubicMillimetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCubicMillimetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMillimetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCubicMillimetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMillimetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMillimetersPerSecond(0)
    if err != nil {
        t.Errorf("FromCubicMillimetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicMillimeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMillimetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromAcreFeetPerSecond function
func TestVolumeFlowFactory_FromAcreFeetPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAcreFeetPerSecond(100)
    if err != nil {
        t.Errorf("FromAcreFeetPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowAcreFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAcreFeetPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAcreFeetPerSecond(math.NaN())
    if err == nil {
        t.Error("FromAcreFeetPerSecond() with NaN value should return error")
    }

    _, err = factory.FromAcreFeetPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromAcreFeetPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromAcreFeetPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromAcreFeetPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAcreFeetPerSecond(0)
    if err != nil {
        t.Errorf("FromAcreFeetPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowAcreFootPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAcreFeetPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromAcreFeetPerMinute function
func TestVolumeFlowFactory_FromAcreFeetPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAcreFeetPerMinute(100)
    if err != nil {
        t.Errorf("FromAcreFeetPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowAcreFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAcreFeetPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAcreFeetPerMinute(math.NaN())
    if err == nil {
        t.Error("FromAcreFeetPerMinute() with NaN value should return error")
    }

    _, err = factory.FromAcreFeetPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromAcreFeetPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromAcreFeetPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromAcreFeetPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAcreFeetPerMinute(0)
    if err != nil {
        t.Errorf("FromAcreFeetPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowAcreFootPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAcreFeetPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromAcreFeetPerHour function
func TestVolumeFlowFactory_FromAcreFeetPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAcreFeetPerHour(100)
    if err != nil {
        t.Errorf("FromAcreFeetPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowAcreFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAcreFeetPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAcreFeetPerHour(math.NaN())
    if err == nil {
        t.Error("FromAcreFeetPerHour() with NaN value should return error")
    }

    _, err = factory.FromAcreFeetPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromAcreFeetPerHour() with +Inf value should return error")
    }

    _, err = factory.FromAcreFeetPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromAcreFeetPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAcreFeetPerHour(0)
    if err != nil {
        t.Errorf("FromAcreFeetPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowAcreFootPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAcreFeetPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromAcreFeetPerDay function
func TestVolumeFlowFactory_FromAcreFeetPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAcreFeetPerDay(100)
    if err != nil {
        t.Errorf("FromAcreFeetPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowAcreFootPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAcreFeetPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAcreFeetPerDay(math.NaN())
    if err == nil {
        t.Error("FromAcreFeetPerDay() with NaN value should return error")
    }

    _, err = factory.FromAcreFeetPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromAcreFeetPerDay() with +Inf value should return error")
    }

    _, err = factory.FromAcreFeetPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromAcreFeetPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAcreFeetPerDay(0)
    if err != nil {
        t.Errorf("FromAcreFeetPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowAcreFootPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAcreFeetPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicCentimetersPerMinute function
func TestVolumeFlowFactory_FromCubicCentimetersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicCentimetersPerMinute(100)
    if err != nil {
        t.Errorf("FromCubicCentimetersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCubicCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicCentimetersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicCentimetersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromCubicCentimetersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromCubicCentimetersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromCubicCentimetersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromCubicCentimetersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicCentimetersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicCentimetersPerMinute(0)
    if err != nil {
        t.Errorf("FromCubicCentimetersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCubicCentimeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicCentimetersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMegausGallonsPerDay function
func TestVolumeFlowFactory_FromMegausGallonsPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegausGallonsPerDay(100)
    if err != nil {
        t.Errorf("FromMegausGallonsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMegausGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegausGallonsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegausGallonsPerDay(math.NaN())
    if err == nil {
        t.Error("FromMegausGallonsPerDay() with NaN value should return error")
    }

    _, err = factory.FromMegausGallonsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMegausGallonsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMegausGallonsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMegausGallonsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegausGallonsPerDay(0)
    if err != nil {
        t.Errorf("FromMegausGallonsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMegausGallonPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegausGallonsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromNanolitersPerSecond function
func TestVolumeFlowFactory_FromNanolitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanolitersPerSecond(100)
    if err != nil {
        t.Errorf("FromNanolitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowNanoliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanolitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanolitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNanolitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNanolitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNanolitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNanolitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNanolitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanolitersPerSecond(0)
    if err != nil {
        t.Errorf("FromNanolitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowNanoliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanolitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrolitersPerSecond function
func TestVolumeFlowFactory_FromMicrolitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrolitersPerSecond(100)
    if err != nil {
        t.Errorf("FromMicrolitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMicroliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrolitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrolitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicrolitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicrolitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicrolitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicrolitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrolitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrolitersPerSecond(0)
    if err != nil {
        t.Errorf("FromMicrolitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMicroliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrolitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillilitersPerSecond function
func TestVolumeFlowFactory_FromMillilitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillilitersPerSecond(100)
    if err != nil {
        t.Errorf("FromMillilitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMilliliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillilitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillilitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillilitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillilitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillilitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillilitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillilitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillilitersPerSecond(0)
    if err != nil {
        t.Errorf("FromMillilitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMilliliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillilitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCentilitersPerSecond function
func TestVolumeFlowFactory_FromCentilitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentilitersPerSecond(100)
    if err != nil {
        t.Errorf("FromCentilitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCentiliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentilitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentilitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCentilitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCentilitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCentilitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCentilitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCentilitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentilitersPerSecond(0)
    if err != nil {
        t.Errorf("FromCentilitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCentiliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentilitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecilitersPerSecond function
func TestVolumeFlowFactory_FromDecilitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecilitersPerSecond(100)
    if err != nil {
        t.Errorf("FromDecilitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowDeciliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecilitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecilitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecilitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecilitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecilitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecilitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecilitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecilitersPerSecond(0)
    if err != nil {
        t.Errorf("FromDecilitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowDeciliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecilitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecalitersPerSecond function
func TestVolumeFlowFactory_FromDecalitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecalitersPerSecond(100)
    if err != nil {
        t.Errorf("FromDecalitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowDecaliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecalitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecalitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecalitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecalitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecalitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecalitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecalitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecalitersPerSecond(0)
    if err != nil {
        t.Errorf("FromDecalitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowDecaliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecalitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromHectolitersPerSecond function
func TestVolumeFlowFactory_FromHectolitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectolitersPerSecond(100)
    if err != nil {
        t.Errorf("FromHectolitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowHectoliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectolitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectolitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromHectolitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromHectolitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromHectolitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromHectolitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromHectolitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectolitersPerSecond(0)
    if err != nil {
        t.Errorf("FromHectolitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowHectoliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectolitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilolitersPerSecond function
func TestVolumeFlowFactory_FromKilolitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilolitersPerSecond(100)
    if err != nil {
        t.Errorf("FromKilolitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowKiloliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilolitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilolitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilolitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilolitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilolitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilolitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilolitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilolitersPerSecond(0)
    if err != nil {
        t.Errorf("FromKilolitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowKiloliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilolitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMegalitersPerSecond function
func TestVolumeFlowFactory_FromMegalitersPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegalitersPerSecond(100)
    if err != nil {
        t.Errorf("FromMegalitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMegaliterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegalitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegalitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegalitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegalitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegalitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegalitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegalitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegalitersPerSecond(0)
    if err != nil {
        t.Errorf("FromMegalitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMegaliterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegalitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromNanolitersPerMinute function
func TestVolumeFlowFactory_FromNanolitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanolitersPerMinute(100)
    if err != nil {
        t.Errorf("FromNanolitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowNanoliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanolitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanolitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromNanolitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromNanolitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromNanolitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromNanolitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromNanolitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanolitersPerMinute(0)
    if err != nil {
        t.Errorf("FromNanolitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowNanoliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanolitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrolitersPerMinute function
func TestVolumeFlowFactory_FromMicrolitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrolitersPerMinute(100)
    if err != nil {
        t.Errorf("FromMicrolitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMicroliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrolitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrolitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMicrolitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMicrolitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMicrolitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMicrolitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrolitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrolitersPerMinute(0)
    if err != nil {
        t.Errorf("FromMicrolitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMicroliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrolitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMillilitersPerMinute function
func TestVolumeFlowFactory_FromMillilitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillilitersPerMinute(100)
    if err != nil {
        t.Errorf("FromMillilitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMilliliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillilitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillilitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMillilitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMillilitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMillilitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMillilitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMillilitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillilitersPerMinute(0)
    if err != nil {
        t.Errorf("FromMillilitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMilliliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillilitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromCentilitersPerMinute function
func TestVolumeFlowFactory_FromCentilitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentilitersPerMinute(100)
    if err != nil {
        t.Errorf("FromCentilitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCentiliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentilitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentilitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromCentilitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromCentilitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromCentilitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromCentilitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromCentilitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentilitersPerMinute(0)
    if err != nil {
        t.Errorf("FromCentilitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCentiliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentilitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromDecilitersPerMinute function
func TestVolumeFlowFactory_FromDecilitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecilitersPerMinute(100)
    if err != nil {
        t.Errorf("FromDecilitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowDeciliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecilitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecilitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromDecilitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromDecilitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromDecilitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromDecilitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromDecilitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecilitersPerMinute(0)
    if err != nil {
        t.Errorf("FromDecilitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowDeciliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecilitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromDecalitersPerMinute function
func TestVolumeFlowFactory_FromDecalitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecalitersPerMinute(100)
    if err != nil {
        t.Errorf("FromDecalitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowDecaliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecalitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecalitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromDecalitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromDecalitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromDecalitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromDecalitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromDecalitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecalitersPerMinute(0)
    if err != nil {
        t.Errorf("FromDecalitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowDecaliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecalitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromHectolitersPerMinute function
func TestVolumeFlowFactory_FromHectolitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectolitersPerMinute(100)
    if err != nil {
        t.Errorf("FromHectolitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowHectoliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectolitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectolitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromHectolitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromHectolitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromHectolitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromHectolitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromHectolitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectolitersPerMinute(0)
    if err != nil {
        t.Errorf("FromHectolitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowHectoliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectolitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromKilolitersPerMinute function
func TestVolumeFlowFactory_FromKilolitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilolitersPerMinute(100)
    if err != nil {
        t.Errorf("FromKilolitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowKiloliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilolitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilolitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilolitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilolitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilolitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilolitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilolitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilolitersPerMinute(0)
    if err != nil {
        t.Errorf("FromKilolitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowKiloliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilolitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMegalitersPerMinute function
func TestVolumeFlowFactory_FromMegalitersPerMinute(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegalitersPerMinute(100)
    if err != nil {
        t.Errorf("FromMegalitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMegaliterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegalitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegalitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMegalitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMegalitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMegalitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMegalitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMegalitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegalitersPerMinute(0)
    if err != nil {
        t.Errorf("FromMegalitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMegaliterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegalitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromNanolitersPerHour function
func TestVolumeFlowFactory_FromNanolitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanolitersPerHour(100)
    if err != nil {
        t.Errorf("FromNanolitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowNanoliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanolitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanolitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromNanolitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromNanolitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromNanolitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromNanolitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromNanolitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanolitersPerHour(0)
    if err != nil {
        t.Errorf("FromNanolitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowNanoliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanolitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrolitersPerHour function
func TestVolumeFlowFactory_FromMicrolitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrolitersPerHour(100)
    if err != nil {
        t.Errorf("FromMicrolitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMicroliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrolitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrolitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromMicrolitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromMicrolitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMicrolitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMicrolitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrolitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrolitersPerHour(0)
    if err != nil {
        t.Errorf("FromMicrolitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMicroliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrolitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMillilitersPerHour function
func TestVolumeFlowFactory_FromMillilitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillilitersPerHour(100)
    if err != nil {
        t.Errorf("FromMillilitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMilliliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillilitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillilitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromMillilitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromMillilitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMillilitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMillilitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMillilitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillilitersPerHour(0)
    if err != nil {
        t.Errorf("FromMillilitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMilliliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillilitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromCentilitersPerHour function
func TestVolumeFlowFactory_FromCentilitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentilitersPerHour(100)
    if err != nil {
        t.Errorf("FromCentilitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCentiliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentilitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentilitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromCentilitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromCentilitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromCentilitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromCentilitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromCentilitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentilitersPerHour(0)
    if err != nil {
        t.Errorf("FromCentilitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCentiliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentilitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromDecilitersPerHour function
func TestVolumeFlowFactory_FromDecilitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecilitersPerHour(100)
    if err != nil {
        t.Errorf("FromDecilitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowDeciliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecilitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecilitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromDecilitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromDecilitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromDecilitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromDecilitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromDecilitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecilitersPerHour(0)
    if err != nil {
        t.Errorf("FromDecilitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowDeciliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecilitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromDecalitersPerHour function
func TestVolumeFlowFactory_FromDecalitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecalitersPerHour(100)
    if err != nil {
        t.Errorf("FromDecalitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowDecaliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecalitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecalitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromDecalitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromDecalitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromDecalitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromDecalitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromDecalitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecalitersPerHour(0)
    if err != nil {
        t.Errorf("FromDecalitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowDecaliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecalitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromHectolitersPerHour function
func TestVolumeFlowFactory_FromHectolitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectolitersPerHour(100)
    if err != nil {
        t.Errorf("FromHectolitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowHectoliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectolitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectolitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromHectolitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromHectolitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromHectolitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromHectolitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromHectolitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectolitersPerHour(0)
    if err != nil {
        t.Errorf("FromHectolitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowHectoliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectolitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilolitersPerHour function
func TestVolumeFlowFactory_FromKilolitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilolitersPerHour(100)
    if err != nil {
        t.Errorf("FromKilolitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowKiloliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilolitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilolitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromKilolitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromKilolitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromKilolitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromKilolitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKilolitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilolitersPerHour(0)
    if err != nil {
        t.Errorf("FromKilolitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowKiloliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilolitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMegalitersPerHour function
func TestVolumeFlowFactory_FromMegalitersPerHour(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegalitersPerHour(100)
    if err != nil {
        t.Errorf("FromMegalitersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMegaliterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegalitersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegalitersPerHour(math.NaN())
    if err == nil {
        t.Error("FromMegalitersPerHour() with NaN value should return error")
    }

    _, err = factory.FromMegalitersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMegalitersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMegalitersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMegalitersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegalitersPerHour(0)
    if err != nil {
        t.Errorf("FromMegalitersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMegaliterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegalitersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromNanolitersPerDay function
func TestVolumeFlowFactory_FromNanolitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanolitersPerDay(100)
    if err != nil {
        t.Errorf("FromNanolitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowNanoliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanolitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanolitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromNanolitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromNanolitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromNanolitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromNanolitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromNanolitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanolitersPerDay(0)
    if err != nil {
        t.Errorf("FromNanolitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowNanoliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanolitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrolitersPerDay function
func TestVolumeFlowFactory_FromMicrolitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrolitersPerDay(100)
    if err != nil {
        t.Errorf("FromMicrolitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMicroliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrolitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrolitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromMicrolitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromMicrolitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMicrolitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMicrolitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrolitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrolitersPerDay(0)
    if err != nil {
        t.Errorf("FromMicrolitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMicroliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrolitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMillilitersPerDay function
func TestVolumeFlowFactory_FromMillilitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillilitersPerDay(100)
    if err != nil {
        t.Errorf("FromMillilitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMilliliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillilitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillilitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromMillilitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromMillilitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMillilitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMillilitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMillilitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillilitersPerDay(0)
    if err != nil {
        t.Errorf("FromMillilitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMilliliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillilitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromCentilitersPerDay function
func TestVolumeFlowFactory_FromCentilitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentilitersPerDay(100)
    if err != nil {
        t.Errorf("FromCentilitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowCentiliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentilitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentilitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromCentilitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromCentilitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromCentilitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromCentilitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromCentilitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentilitersPerDay(0)
    if err != nil {
        t.Errorf("FromCentilitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowCentiliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentilitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromDecilitersPerDay function
func TestVolumeFlowFactory_FromDecilitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecilitersPerDay(100)
    if err != nil {
        t.Errorf("FromDecilitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowDeciliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecilitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecilitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromDecilitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromDecilitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromDecilitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromDecilitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromDecilitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecilitersPerDay(0)
    if err != nil {
        t.Errorf("FromDecilitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowDeciliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecilitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromDecalitersPerDay function
func TestVolumeFlowFactory_FromDecalitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecalitersPerDay(100)
    if err != nil {
        t.Errorf("FromDecalitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowDecaliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecalitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecalitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromDecalitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromDecalitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromDecalitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromDecalitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromDecalitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecalitersPerDay(0)
    if err != nil {
        t.Errorf("FromDecalitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowDecaliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecalitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromHectolitersPerDay function
func TestVolumeFlowFactory_FromHectolitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectolitersPerDay(100)
    if err != nil {
        t.Errorf("FromHectolitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowHectoliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectolitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectolitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromHectolitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromHectolitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromHectolitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromHectolitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromHectolitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectolitersPerDay(0)
    if err != nil {
        t.Errorf("FromHectolitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowHectoliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectolitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromKilolitersPerDay function
func TestVolumeFlowFactory_FromKilolitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilolitersPerDay(100)
    if err != nil {
        t.Errorf("FromKilolitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowKiloliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilolitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilolitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromKilolitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromKilolitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromKilolitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromKilolitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromKilolitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilolitersPerDay(0)
    if err != nil {
        t.Errorf("FromKilolitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowKiloliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilolitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMegalitersPerDay function
func TestVolumeFlowFactory_FromMegalitersPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegalitersPerDay(100)
    if err != nil {
        t.Errorf("FromMegalitersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMegaliterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegalitersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegalitersPerDay(math.NaN())
    if err == nil {
        t.Error("FromMegalitersPerDay() with NaN value should return error")
    }

    _, err = factory.FromMegalitersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMegalitersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMegalitersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMegalitersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegalitersPerDay(0)
    if err != nil {
        t.Errorf("FromMegalitersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMegaliterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegalitersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaukGallonsPerDay function
func TestVolumeFlowFactory_FromMegaukGallonsPerDay(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaukGallonsPerDay(100)
    if err != nil {
        t.Errorf("FromMegaukGallonsPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMegaukGallonPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaukGallonsPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaukGallonsPerDay(math.NaN())
    if err == nil {
        t.Error("FromMegaukGallonsPerDay() with NaN value should return error")
    }

    _, err = factory.FromMegaukGallonsPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromMegaukGallonsPerDay() with +Inf value should return error")
    }

    _, err = factory.FromMegaukGallonsPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaukGallonsPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaukGallonsPerDay(0)
    if err != nil {
        t.Errorf("FromMegaukGallonsPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMegaukGallonPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaukGallonsPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaukGallonsPerSecond function
func TestVolumeFlowFactory_FromMegaukGallonsPerSecond(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaukGallonsPerSecond(100)
    if err != nil {
        t.Errorf("FromMegaukGallonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowMegaukGallonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaukGallonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaukGallonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegaukGallonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegaukGallonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegaukGallonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegaukGallonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaukGallonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaukGallonsPerSecond(0)
    if err != nil {
        t.Errorf("FromMegaukGallonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowMegaukGallonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaukGallonsPerSecond() with zero value = %v, want 0", converted)
    }
}

func TestVolumeFlowToString(t *testing.T) {
	factory := units.VolumeFlowFactory{}
	a, err := factory.CreateVolumeFlow(45, units.VolumeFlowCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumeFlowCubicMeterPerSecond, 2)
	expected := "45.00 " + units.GetVolumeFlowAbbreviation(units.VolumeFlowCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumeFlowCubicMeterPerSecond, -1)
	expected = "45 " + units.GetVolumeFlowAbbreviation(units.VolumeFlowCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolumeFlow_EqualityAndComparison(t *testing.T) {
	factory := units.VolumeFlowFactory{}
	a1, _ := factory.CreateVolumeFlow(60, units.VolumeFlowCubicMeterPerSecond)
	a2, _ := factory.CreateVolumeFlow(60, units.VolumeFlowCubicMeterPerSecond)
	a3, _ := factory.CreateVolumeFlow(120, units.VolumeFlowCubicMeterPerSecond)

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

func TestVolumeFlow_Arithmetic(t *testing.T) {
	factory := units.VolumeFlowFactory{}
	a1, _ := factory.CreateVolumeFlow(30, units.VolumeFlowCubicMeterPerSecond)
	a2, _ := factory.CreateVolumeFlow(45, units.VolumeFlowCubicMeterPerSecond)

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


func TestGetVolumeFlowAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.VolumeFlowUnits
        want string
    }{
        {
            name: "CubicMeterPerSecond abbreviation",
            unit: units.VolumeFlowCubicMeterPerSecond,
            want: "m³/s",
        },
        {
            name: "CubicMeterPerMinute abbreviation",
            unit: units.VolumeFlowCubicMeterPerMinute,
            want: "m³/min",
        },
        {
            name: "CubicMeterPerHour abbreviation",
            unit: units.VolumeFlowCubicMeterPerHour,
            want: "m³/h",
        },
        {
            name: "CubicMeterPerDay abbreviation",
            unit: units.VolumeFlowCubicMeterPerDay,
            want: "m³/d",
        },
        {
            name: "CubicFootPerSecond abbreviation",
            unit: units.VolumeFlowCubicFootPerSecond,
            want: "ft³/s",
        },
        {
            name: "CubicFootPerMinute abbreviation",
            unit: units.VolumeFlowCubicFootPerMinute,
            want: "ft³/min",
        },
        {
            name: "CubicFootPerHour abbreviation",
            unit: units.VolumeFlowCubicFootPerHour,
            want: "ft³/h",
        },
        {
            name: "CubicYardPerSecond abbreviation",
            unit: units.VolumeFlowCubicYardPerSecond,
            want: "yd³/s",
        },
        {
            name: "CubicYardPerMinute abbreviation",
            unit: units.VolumeFlowCubicYardPerMinute,
            want: "yd³/min",
        },
        {
            name: "CubicYardPerHour abbreviation",
            unit: units.VolumeFlowCubicYardPerHour,
            want: "yd³/h",
        },
        {
            name: "CubicYardPerDay abbreviation",
            unit: units.VolumeFlowCubicYardPerDay,
            want: "cy/day",
        },
        {
            name: "MillionUsGallonPerDay abbreviation",
            unit: units.VolumeFlowMillionUsGallonPerDay,
            want: "MGD",
        },
        {
            name: "UsGallonPerDay abbreviation",
            unit: units.VolumeFlowUsGallonPerDay,
            want: "gpd",
        },
        {
            name: "LiterPerSecond abbreviation",
            unit: units.VolumeFlowLiterPerSecond,
            want: "l/s",
        },
        {
            name: "LiterPerMinute abbreviation",
            unit: units.VolumeFlowLiterPerMinute,
            want: "l/min",
        },
        {
            name: "LiterPerHour abbreviation",
            unit: units.VolumeFlowLiterPerHour,
            want: "l/h",
        },
        {
            name: "LiterPerDay abbreviation",
            unit: units.VolumeFlowLiterPerDay,
            want: "l/day",
        },
        {
            name: "UsGallonPerSecond abbreviation",
            unit: units.VolumeFlowUsGallonPerSecond,
            want: "gal (U.S.)/s",
        },
        {
            name: "UsGallonPerMinute abbreviation",
            unit: units.VolumeFlowUsGallonPerMinute,
            want: "gal (U.S.)/min",
        },
        {
            name: "UkGallonPerDay abbreviation",
            unit: units.VolumeFlowUkGallonPerDay,
            want: "gal (U. K.)/d",
        },
        {
            name: "UkGallonPerHour abbreviation",
            unit: units.VolumeFlowUkGallonPerHour,
            want: "gal (imp.)/h",
        },
        {
            name: "UkGallonPerMinute abbreviation",
            unit: units.VolumeFlowUkGallonPerMinute,
            want: "gal (imp.)/min",
        },
        {
            name: "UkGallonPerSecond abbreviation",
            unit: units.VolumeFlowUkGallonPerSecond,
            want: "gal (imp.)/s",
        },
        {
            name: "KilousGallonPerMinute abbreviation",
            unit: units.VolumeFlowKilousGallonPerMinute,
            want: "kgal (U.S.)/min",
        },
        {
            name: "UsGallonPerHour abbreviation",
            unit: units.VolumeFlowUsGallonPerHour,
            want: "gal (U.S.)/h",
        },
        {
            name: "CubicDecimeterPerMinute abbreviation",
            unit: units.VolumeFlowCubicDecimeterPerMinute,
            want: "dm³/min",
        },
        {
            name: "OilBarrelPerDay abbreviation",
            unit: units.VolumeFlowOilBarrelPerDay,
            want: "bbl/d",
        },
        {
            name: "OilBarrelPerMinute abbreviation",
            unit: units.VolumeFlowOilBarrelPerMinute,
            want: "bbl/min",
        },
        {
            name: "OilBarrelPerHour abbreviation",
            unit: units.VolumeFlowOilBarrelPerHour,
            want: "bbl/hr",
        },
        {
            name: "OilBarrelPerSecond abbreviation",
            unit: units.VolumeFlowOilBarrelPerSecond,
            want: "bbl/s",
        },
        {
            name: "CubicMillimeterPerSecond abbreviation",
            unit: units.VolumeFlowCubicMillimeterPerSecond,
            want: "mm³/s",
        },
        {
            name: "AcreFootPerSecond abbreviation",
            unit: units.VolumeFlowAcreFootPerSecond,
            want: "af/s",
        },
        {
            name: "AcreFootPerMinute abbreviation",
            unit: units.VolumeFlowAcreFootPerMinute,
            want: "af/m",
        },
        {
            name: "AcreFootPerHour abbreviation",
            unit: units.VolumeFlowAcreFootPerHour,
            want: "af/h",
        },
        {
            name: "AcreFootPerDay abbreviation",
            unit: units.VolumeFlowAcreFootPerDay,
            want: "af/d",
        },
        {
            name: "CubicCentimeterPerMinute abbreviation",
            unit: units.VolumeFlowCubicCentimeterPerMinute,
            want: "cm³/min",
        },
        {
            name: "MegausGallonPerDay abbreviation",
            unit: units.VolumeFlowMegausGallonPerDay,
            want: "Mgpd",
        },
        {
            name: "NanoliterPerSecond abbreviation",
            unit: units.VolumeFlowNanoliterPerSecond,
            want: "nl/s",
        },
        {
            name: "MicroliterPerSecond abbreviation",
            unit: units.VolumeFlowMicroliterPerSecond,
            want: "μl/s",
        },
        {
            name: "MilliliterPerSecond abbreviation",
            unit: units.VolumeFlowMilliliterPerSecond,
            want: "ml/s",
        },
        {
            name: "CentiliterPerSecond abbreviation",
            unit: units.VolumeFlowCentiliterPerSecond,
            want: "cl/s",
        },
        {
            name: "DeciliterPerSecond abbreviation",
            unit: units.VolumeFlowDeciliterPerSecond,
            want: "dl/s",
        },
        {
            name: "DecaliterPerSecond abbreviation",
            unit: units.VolumeFlowDecaliterPerSecond,
            want: "dal/s",
        },
        {
            name: "HectoliterPerSecond abbreviation",
            unit: units.VolumeFlowHectoliterPerSecond,
            want: "hl/s",
        },
        {
            name: "KiloliterPerSecond abbreviation",
            unit: units.VolumeFlowKiloliterPerSecond,
            want: "kl/s",
        },
        {
            name: "MegaliterPerSecond abbreviation",
            unit: units.VolumeFlowMegaliterPerSecond,
            want: "Ml/s",
        },
        {
            name: "NanoliterPerMinute abbreviation",
            unit: units.VolumeFlowNanoliterPerMinute,
            want: "nl/min",
        },
        {
            name: "MicroliterPerMinute abbreviation",
            unit: units.VolumeFlowMicroliterPerMinute,
            want: "μl/min",
        },
        {
            name: "MilliliterPerMinute abbreviation",
            unit: units.VolumeFlowMilliliterPerMinute,
            want: "ml/min",
        },
        {
            name: "CentiliterPerMinute abbreviation",
            unit: units.VolumeFlowCentiliterPerMinute,
            want: "cl/min",
        },
        {
            name: "DeciliterPerMinute abbreviation",
            unit: units.VolumeFlowDeciliterPerMinute,
            want: "dl/min",
        },
        {
            name: "DecaliterPerMinute abbreviation",
            unit: units.VolumeFlowDecaliterPerMinute,
            want: "dal/min",
        },
        {
            name: "HectoliterPerMinute abbreviation",
            unit: units.VolumeFlowHectoliterPerMinute,
            want: "hl/min",
        },
        {
            name: "KiloliterPerMinute abbreviation",
            unit: units.VolumeFlowKiloliterPerMinute,
            want: "kl/min",
        },
        {
            name: "MegaliterPerMinute abbreviation",
            unit: units.VolumeFlowMegaliterPerMinute,
            want: "Ml/min",
        },
        {
            name: "NanoliterPerHour abbreviation",
            unit: units.VolumeFlowNanoliterPerHour,
            want: "nl/h",
        },
        {
            name: "MicroliterPerHour abbreviation",
            unit: units.VolumeFlowMicroliterPerHour,
            want: "μl/h",
        },
        {
            name: "MilliliterPerHour abbreviation",
            unit: units.VolumeFlowMilliliterPerHour,
            want: "ml/h",
        },
        {
            name: "CentiliterPerHour abbreviation",
            unit: units.VolumeFlowCentiliterPerHour,
            want: "cl/h",
        },
        {
            name: "DeciliterPerHour abbreviation",
            unit: units.VolumeFlowDeciliterPerHour,
            want: "dl/h",
        },
        {
            name: "DecaliterPerHour abbreviation",
            unit: units.VolumeFlowDecaliterPerHour,
            want: "dal/h",
        },
        {
            name: "HectoliterPerHour abbreviation",
            unit: units.VolumeFlowHectoliterPerHour,
            want: "hl/h",
        },
        {
            name: "KiloliterPerHour abbreviation",
            unit: units.VolumeFlowKiloliterPerHour,
            want: "kl/h",
        },
        {
            name: "MegaliterPerHour abbreviation",
            unit: units.VolumeFlowMegaliterPerHour,
            want: "Ml/h",
        },
        {
            name: "NanoliterPerDay abbreviation",
            unit: units.VolumeFlowNanoliterPerDay,
            want: "nl/day",
        },
        {
            name: "MicroliterPerDay abbreviation",
            unit: units.VolumeFlowMicroliterPerDay,
            want: "μl/day",
        },
        {
            name: "MilliliterPerDay abbreviation",
            unit: units.VolumeFlowMilliliterPerDay,
            want: "ml/day",
        },
        {
            name: "CentiliterPerDay abbreviation",
            unit: units.VolumeFlowCentiliterPerDay,
            want: "cl/day",
        },
        {
            name: "DeciliterPerDay abbreviation",
            unit: units.VolumeFlowDeciliterPerDay,
            want: "dl/day",
        },
        {
            name: "DecaliterPerDay abbreviation",
            unit: units.VolumeFlowDecaliterPerDay,
            want: "dal/day",
        },
        {
            name: "HectoliterPerDay abbreviation",
            unit: units.VolumeFlowHectoliterPerDay,
            want: "hl/day",
        },
        {
            name: "KiloliterPerDay abbreviation",
            unit: units.VolumeFlowKiloliterPerDay,
            want: "kl/day",
        },
        {
            name: "MegaliterPerDay abbreviation",
            unit: units.VolumeFlowMegaliterPerDay,
            want: "Ml/day",
        },
        {
            name: "MegaukGallonPerDay abbreviation",
            unit: units.VolumeFlowMegaukGallonPerDay,
            want: "Mgal (U. K.)/d",
        },
        {
            name: "MegaukGallonPerSecond abbreviation",
            unit: units.VolumeFlowMegaukGallonPerSecond,
            want: "Mgal (imp.)/s",
        },
        {
            name: "invalid unit",
            unit: units.VolumeFlowUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetVolumeFlowAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetVolumeFlowAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestVolumeFlow_String(t *testing.T) {
    factory := units.VolumeFlowFactory{}
    
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
            unit, err := factory.CreateVolumeFlow(tt.value, units.VolumeFlowCubicMeterPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("VolumeFlow.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestVolumeFlow_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.VolumeFlowFactory{}

	_, err := uf.CreateVolumeFlow(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}