// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

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
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CubicMetersPerMinute.
		// No expected conversion value provided for CubicMetersPerMinute, verifying result is not NaN.
		result := a.CubicMetersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CubicMetersPerHour.
		// No expected conversion value provided for CubicMetersPerHour, verifying result is not NaN.
		result := a.CubicMetersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to CubicMetersPerDay.
		// No expected conversion value provided for CubicMetersPerDay, verifying result is not NaN.
		result := a.CubicMetersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMetersPerDay returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerSecond.
		// No expected conversion value provided for CubicFeetPerSecond, verifying result is not NaN.
		result := a.CubicFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerMinute.
		// No expected conversion value provided for CubicFeetPerMinute, verifying result is not NaN.
		result := a.CubicFeetPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicFeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerHour.
		// No expected conversion value provided for CubicFeetPerHour, verifying result is not NaN.
		result := a.CubicFeetPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicFeetPerHour returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerSecond.
		// No expected conversion value provided for CubicYardsPerSecond, verifying result is not NaN.
		result := a.CubicYardsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicYardsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerMinute.
		// No expected conversion value provided for CubicYardsPerMinute, verifying result is not NaN.
		result := a.CubicYardsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicYardsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerHour.
		// No expected conversion value provided for CubicYardsPerHour, verifying result is not NaN.
		result := a.CubicYardsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicYardsPerHour returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerDay.
		// No expected conversion value provided for CubicYardsPerDay, verifying result is not NaN.
		result := a.CubicYardsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicYardsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MillionUsGallonsPerDay.
		// No expected conversion value provided for MillionUsGallonsPerDay, verifying result is not NaN.
		result := a.MillionUsGallonsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillionUsGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerDay.
		// No expected conversion value provided for UsGallonsPerDay, verifying result is not NaN.
		result := a.UsGallonsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to LitersPerSecond.
		// No expected conversion value provided for LitersPerSecond, verifying result is not NaN.
		result := a.LitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to LitersPerMinute.
		// No expected conversion value provided for LitersPerMinute, verifying result is not NaN.
		result := a.LitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to LitersPerHour.
		// No expected conversion value provided for LitersPerHour, verifying result is not NaN.
		result := a.LitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to LitersPerDay.
		// No expected conversion value provided for LitersPerDay, verifying result is not NaN.
		result := a.LitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerSecond.
		// No expected conversion value provided for UsGallonsPerSecond, verifying result is not NaN.
		result := a.UsGallonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsGallonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerMinute.
		// No expected conversion value provided for UsGallonsPerMinute, verifying result is not NaN.
		result := a.UsGallonsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsGallonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to UkGallonsPerDay.
		// No expected conversion value provided for UkGallonsPerDay, verifying result is not NaN.
		result := a.UkGallonsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to UkGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to UkGallonsPerHour.
		// No expected conversion value provided for UkGallonsPerHour, verifying result is not NaN.
		result := a.UkGallonsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to UkGallonsPerHour returned NaN")
		}
	}
	{
		// Test conversion to UkGallonsPerMinute.
		// No expected conversion value provided for UkGallonsPerMinute, verifying result is not NaN.
		result := a.UkGallonsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to UkGallonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to UkGallonsPerSecond.
		// No expected conversion value provided for UkGallonsPerSecond, verifying result is not NaN.
		result := a.UkGallonsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to UkGallonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilousGallonsPerMinute.
		// No expected conversion value provided for KilousGallonsPerMinute, verifying result is not NaN.
		result := a.KilousGallonsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilousGallonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerHour.
		// No expected conversion value provided for UsGallonsPerHour, verifying result is not NaN.
		result := a.UsGallonsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to UsGallonsPerHour returned NaN")
		}
	}
	{
		// Test conversion to CubicDecimetersPerMinute.
		// No expected conversion value provided for CubicDecimetersPerMinute, verifying result is not NaN.
		result := a.CubicDecimetersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicDecimetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerDay.
		// No expected conversion value provided for OilBarrelsPerDay, verifying result is not NaN.
		result := a.OilBarrelsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to OilBarrelsPerDay returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerMinute.
		// No expected conversion value provided for OilBarrelsPerMinute, verifying result is not NaN.
		result := a.OilBarrelsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to OilBarrelsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerHour.
		// No expected conversion value provided for OilBarrelsPerHour, verifying result is not NaN.
		result := a.OilBarrelsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to OilBarrelsPerHour returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerSecond.
		// No expected conversion value provided for OilBarrelsPerSecond, verifying result is not NaN.
		result := a.OilBarrelsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to OilBarrelsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CubicMillimetersPerSecond.
		// No expected conversion value provided for CubicMillimetersPerSecond, verifying result is not NaN.
		result := a.CubicMillimetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicMillimetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AcreFeetPerSecond.
		// No expected conversion value provided for AcreFeetPerSecond, verifying result is not NaN.
		result := a.AcreFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to AcreFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AcreFeetPerMinute.
		// No expected conversion value provided for AcreFeetPerMinute, verifying result is not NaN.
		result := a.AcreFeetPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to AcreFeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to AcreFeetPerHour.
		// No expected conversion value provided for AcreFeetPerHour, verifying result is not NaN.
		result := a.AcreFeetPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to AcreFeetPerHour returned NaN")
		}
	}
	{
		// Test conversion to AcreFeetPerDay.
		// No expected conversion value provided for AcreFeetPerDay, verifying result is not NaN.
		result := a.AcreFeetPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to AcreFeetPerDay returned NaN")
		}
	}
	{
		// Test conversion to CubicCentimetersPerMinute.
		// No expected conversion value provided for CubicCentimetersPerMinute, verifying result is not NaN.
		result := a.CubicCentimetersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to CubicCentimetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegausGallonsPerDay.
		// No expected conversion value provided for MegausGallonsPerDay, verifying result is not NaN.
		result := a.MegausGallonsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegausGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerSecond.
		// No expected conversion value provided for NanolitersPerSecond, verifying result is not NaN.
		result := a.NanolitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanolitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerSecond.
		// No expected conversion value provided for MicrolitersPerSecond, verifying result is not NaN.
		result := a.MicrolitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrolitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerSecond.
		// No expected conversion value provided for MillilitersPerSecond, verifying result is not NaN.
		result := a.MillilitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillilitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerSecond.
		// No expected conversion value provided for CentilitersPerSecond, verifying result is not NaN.
		result := a.CentilitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentilitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerSecond.
		// No expected conversion value provided for DecilitersPerSecond, verifying result is not NaN.
		result := a.DecilitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecilitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecalitersPerSecond.
		// No expected conversion value provided for DecalitersPerSecond, verifying result is not NaN.
		result := a.DecalitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecalitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to HectolitersPerSecond.
		// No expected conversion value provided for HectolitersPerSecond, verifying result is not NaN.
		result := a.HectolitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectolitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilolitersPerSecond.
		// No expected conversion value provided for KilolitersPerSecond, verifying result is not NaN.
		result := a.KilolitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilolitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegalitersPerSecond.
		// No expected conversion value provided for MegalitersPerSecond, verifying result is not NaN.
		result := a.MegalitersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegalitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerMinute.
		// No expected conversion value provided for NanolitersPerMinute, verifying result is not NaN.
		result := a.NanolitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanolitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerMinute.
		// No expected conversion value provided for MicrolitersPerMinute, verifying result is not NaN.
		result := a.MicrolitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrolitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerMinute.
		// No expected conversion value provided for MillilitersPerMinute, verifying result is not NaN.
		result := a.MillilitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillilitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerMinute.
		// No expected conversion value provided for CentilitersPerMinute, verifying result is not NaN.
		result := a.CentilitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentilitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerMinute.
		// No expected conversion value provided for DecilitersPerMinute, verifying result is not NaN.
		result := a.DecilitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecilitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DecalitersPerMinute.
		// No expected conversion value provided for DecalitersPerMinute, verifying result is not NaN.
		result := a.DecalitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecalitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to HectolitersPerMinute.
		// No expected conversion value provided for HectolitersPerMinute, verifying result is not NaN.
		result := a.HectolitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectolitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilolitersPerMinute.
		// No expected conversion value provided for KilolitersPerMinute, verifying result is not NaN.
		result := a.KilolitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilolitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegalitersPerMinute.
		// No expected conversion value provided for MegalitersPerMinute, verifying result is not NaN.
		result := a.MegalitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegalitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerHour.
		// No expected conversion value provided for NanolitersPerHour, verifying result is not NaN.
		result := a.NanolitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanolitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerHour.
		// No expected conversion value provided for MicrolitersPerHour, verifying result is not NaN.
		result := a.MicrolitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrolitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerHour.
		// No expected conversion value provided for MillilitersPerHour, verifying result is not NaN.
		result := a.MillilitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillilitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerHour.
		// No expected conversion value provided for CentilitersPerHour, verifying result is not NaN.
		result := a.CentilitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentilitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerHour.
		// No expected conversion value provided for DecilitersPerHour, verifying result is not NaN.
		result := a.DecilitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecilitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to DecalitersPerHour.
		// No expected conversion value provided for DecalitersPerHour, verifying result is not NaN.
		result := a.DecalitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecalitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to HectolitersPerHour.
		// No expected conversion value provided for HectolitersPerHour, verifying result is not NaN.
		result := a.HectolitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectolitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilolitersPerHour.
		// No expected conversion value provided for KilolitersPerHour, verifying result is not NaN.
		result := a.KilolitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilolitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to MegalitersPerHour.
		// No expected conversion value provided for MegalitersPerHour, verifying result is not NaN.
		result := a.MegalitersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegalitersPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerDay.
		// No expected conversion value provided for NanolitersPerDay, verifying result is not NaN.
		result := a.NanolitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanolitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerDay.
		// No expected conversion value provided for MicrolitersPerDay, verifying result is not NaN.
		result := a.MicrolitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrolitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerDay.
		// No expected conversion value provided for MillilitersPerDay, verifying result is not NaN.
		result := a.MillilitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillilitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerDay.
		// No expected conversion value provided for CentilitersPerDay, verifying result is not NaN.
		result := a.CentilitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentilitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerDay.
		// No expected conversion value provided for DecilitersPerDay, verifying result is not NaN.
		result := a.DecilitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecilitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to DecalitersPerDay.
		// No expected conversion value provided for DecalitersPerDay, verifying result is not NaN.
		result := a.DecalitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecalitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to HectolitersPerDay.
		// No expected conversion value provided for HectolitersPerDay, verifying result is not NaN.
		result := a.HectolitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectolitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to KilolitersPerDay.
		// No expected conversion value provided for KilolitersPerDay, verifying result is not NaN.
		result := a.KilolitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilolitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegalitersPerDay.
		// No expected conversion value provided for MegalitersPerDay, verifying result is not NaN.
		result := a.MegalitersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegalitersPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegaukGallonsPerDay.
		// No expected conversion value provided for MegaukGallonsPerDay, verifying result is not NaN.
		result := a.MegaukGallonsPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegaukGallonsPerDay returned NaN")
		}
	}
	{
		// Test conversion to MegaukGallonsPerSecond.
		// No expected conversion value provided for MegaukGallonsPerSecond, verifying result is not NaN.
		result := a.MegaukGallonsPerSecond()
		if math.IsNaN(result) {
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