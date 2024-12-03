package unitsnet_go_test

import (
	"testing"

	"github.com/haimkastner/unitsnet-go/unitsnet_go"
)

func TestUnitDto(t *testing.T) {
	lf := unitsnet_go.LengthFactory{}

	// Create a length instance from meters
	length1, _ := lf.FromMeters(100.01)

	// Test: Create JSON from default unit (Meter)
	t.Run("Test Create JSON From Default Unit", func(t *testing.T) {
		dtoJson, _ := length1.ToDtoJSON(nil)
		expected := `{"value":100.01,"unit":"Meter"}`
		if string(dtoJson) != expected {
			t.Errorf("Expected %s, got %s", expected, dtoJson)
		}
	})

	// Test: Create JSON from a specific unit (Centimeter)
	t.Run("Test Create JSON From Specific Unit", func(t *testing.T) {
		lengthInCm := unitsnet_go.LengthCentimeter
		dtoJson, _ := length1.ToDtoJSON(&lengthInCm)
		expected := `{"value":10001,"unit":"Centimeter"}`
		if string(dtoJson) != expected {
			t.Errorf("Expected %s, got %s", expected, dtoJson)
		}
	})

	// Test: Directly create JSON from default unit
	t.Run("Test Directly Create JSON From Default Unit", func(t *testing.T) {
		dtoJson, _ := length1.ToDtoJSON(nil)
		expected := `{"value":100.01,"unit":"Meter"}`
		if string(dtoJson) != expected {
			t.Errorf("Expected %s, got %s", expected, dtoJson)
		}
	})

	// Test: Directly create JSON from a specific unit
	t.Run("Test Directly Create JSON From Specific Unit", func(t *testing.T) {
		lengthInCm := unitsnet_go.LengthCentimeter
		dtoJson, _ := length1.ToDtoJSON(&lengthInCm)
		expected := `{"value":10001,"unit":"Centimeter"}`
		if string(dtoJson) != expected {
			t.Errorf("Expected %s, got %s", expected, dtoJson)
		}
	})

	// Test: Create DTO from default unit
	t.Run("Test Create DTO From Default Unit", func(t *testing.T) {
		dto := length1.ToDto(nil)
		if dto.Value != 100.01 || dto.Unit != "Meter" {
			t.Errorf("Expected {100.01, Meter}, got {%f, %s}", dto.Value, dto.Unit)
		}
	})

	// Test: Create DTO from a specific unit
	t.Run("Test Create DTO From Specific Unit", func(t *testing.T) {
		lengthInCm := unitsnet_go.LengthCentimeter
		dto := length1.ToDto(&lengthInCm)
		if dto.Value != 10001 || dto.Unit != "Centimeter" {
			t.Errorf("Expected {10001, Centimeter}, got {%f, %s}", dto.Value, dto.Unit)
		}
	})

	// Test: Load from default unit JSON
	t.Run("Test Load From Default Unit JSON", func(t *testing.T) {
		dtoJson, _ := length1.ToDtoJSON(nil)
		lengthFromJson, _ := lf.FromDtoJSON(dtoJson)
		if lengthFromJson.Meters() != 100.01 {
			t.Errorf("Expected 100.01 meters, got %f", lengthFromJson.Meters())
		}
	})

	// Test: Load from specific unit JSON
	t.Run("Test Load From Specific Unit JSON", func(t *testing.T) {
		lengthInCm := unitsnet_go.LengthCentimeter
		dtoJson, _ := length1.ToDtoJSON(&lengthInCm)
		lengthFromJson, _ := lf.FromDtoJSON(dtoJson)
		if lengthFromJson.Decimeters() != 1000.1 {
			t.Errorf("Expected 1000.1 decimeters, got %f", lengthFromJson.Decimeters())
		}
	})

	// Test: Load from default unit DTO
	t.Run("Test Load From Default Unit DTO", func(t *testing.T) {
		dto := length1.ToDto(nil)
		lengthFromDto, _ := lf.FromDto(dto)
		if lengthFromDto.Meters() != 100.01 {
			t.Errorf("Expected 100.01 meters, got %f", lengthFromDto.Meters())
		}
	})

	// Test: Load from specific unit DTO
	t.Run("Test Load From Specific Unit DTO", func(t *testing.T) {
		lengthInCm := unitsnet_go.LengthCentimeter
		dto := length1.ToDto(&lengthInCm)
		lengthFromDto, _ := lf.FromDto(dto)
		if lengthFromDto.Decimeters() != 1000.1 {
			t.Errorf("Expected 1000.1 decimeters, got %f", lengthFromDto.Decimeters())
		}
	})

	// Test: Load from JSON
	t.Run("Test Load From JSON", func(t *testing.T) {
		json := `{"value":100.01,"unit":"Meter"}`
		lengthFromJson, _ := lf.FromDtoJSON([]byte(json))
		if lengthFromJson.Meters() != 100.01 {
			t.Errorf("Expected 100.01 meters, got %f", lengthFromJson.Meters())
		}
	})

	// Test: Load directly from JSON
	t.Run("Test Load Directly From JSON", func(t *testing.T) {
		json := `{"value":100.01,"unit":"Meter"}`
		lengthFromJson, _ := lf.FromDtoJSON([]byte(json))
		if lengthFromJson.Meters() != 100.01 {
			t.Errorf("Expected 100.01 meters, got %f", lengthFromJson.Meters())
		}
	})

	// Test: Load directly from specific unit JSON
	t.Run("Test Load Directly From Specific Unit JSON", func(t *testing.T) {
		json := `{"value":10001,"unit":"Centimeter"}`
		lengthFromJson, _ := lf.FromDtoJSON([]byte(json))
		if lengthFromJson.Meters() != 100.01 {
			t.Errorf("Expected 100.01 meters, got %f", lengthFromJson.Meters())
		}
	})

	// Test: Should be similar values from two DTO representations
	t.Run("Test Should Be Similar Values From Two DTO Representations", func(t *testing.T) {
		length, _ := unitsnet_go.LengthFactory{}.FromMeters(100.01)

		// Obtain DTO object as JSON, represented by default unit (Meter)
		lengthDtoJson, _ := length.ToDtoJSON(nil)

		// Obtain same value represented in kilometers
		lengthKilometer := unitsnet_go.LengthKilometer
		lengthDtoRepresentsInKmJson, _ := length.ToDtoJSON(&lengthKilometer)

		// Load JSON to DTO, and load
		lengthFromMetersDto, _ := lf.FromDtoJSON(lengthDtoJson)
		lengthFromKmDto, _ := lf.FromDtoJSON(lengthDtoRepresentsInKmJson)

		if lengthFromMetersDto.Meters() != lengthFromKmDto.Meters() {
			t.Errorf("Expected similar values, got %f and %f", lengthFromMetersDto.Meters(), lengthFromKmDto.Meters())
		}
	})
}
