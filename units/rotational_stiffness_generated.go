// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RotationalStiffnessUnits defines various units of RotationalStiffness.
type RotationalStiffnessUnits string

const (
    
        // 
        RotationalStiffnessNewtonMeterPerRadian RotationalStiffnessUnits = "NewtonMeterPerRadian"
        // 
        RotationalStiffnessPoundForceFootPerDegrees RotationalStiffnessUnits = "PoundForceFootPerDegrees"
        // 
        RotationalStiffnessKilopoundForceFootPerDegrees RotationalStiffnessUnits = "KilopoundForceFootPerDegrees"
        // 
        RotationalStiffnessNewtonMillimeterPerDegree RotationalStiffnessUnits = "NewtonMillimeterPerDegree"
        // 
        RotationalStiffnessNewtonMeterPerDegree RotationalStiffnessUnits = "NewtonMeterPerDegree"
        // 
        RotationalStiffnessNewtonMillimeterPerRadian RotationalStiffnessUnits = "NewtonMillimeterPerRadian"
        // 
        RotationalStiffnessPoundForceFeetPerRadian RotationalStiffnessUnits = "PoundForceFeetPerRadian"
        // 
        RotationalStiffnessKilonewtonMeterPerRadian RotationalStiffnessUnits = "KilonewtonMeterPerRadian"
        // 
        RotationalStiffnessMeganewtonMeterPerRadian RotationalStiffnessUnits = "MeganewtonMeterPerRadian"
        // 
        RotationalStiffnessNanonewtonMillimeterPerDegree RotationalStiffnessUnits = "NanonewtonMillimeterPerDegree"
        // 
        RotationalStiffnessMicronewtonMillimeterPerDegree RotationalStiffnessUnits = "MicronewtonMillimeterPerDegree"
        // 
        RotationalStiffnessMillinewtonMillimeterPerDegree RotationalStiffnessUnits = "MillinewtonMillimeterPerDegree"
        // 
        RotationalStiffnessCentinewtonMillimeterPerDegree RotationalStiffnessUnits = "CentinewtonMillimeterPerDegree"
        // 
        RotationalStiffnessDecinewtonMillimeterPerDegree RotationalStiffnessUnits = "DecinewtonMillimeterPerDegree"
        // 
        RotationalStiffnessDecanewtonMillimeterPerDegree RotationalStiffnessUnits = "DecanewtonMillimeterPerDegree"
        // 
        RotationalStiffnessKilonewtonMillimeterPerDegree RotationalStiffnessUnits = "KilonewtonMillimeterPerDegree"
        // 
        RotationalStiffnessMeganewtonMillimeterPerDegree RotationalStiffnessUnits = "MeganewtonMillimeterPerDegree"
        // 
        RotationalStiffnessNanonewtonMeterPerDegree RotationalStiffnessUnits = "NanonewtonMeterPerDegree"
        // 
        RotationalStiffnessMicronewtonMeterPerDegree RotationalStiffnessUnits = "MicronewtonMeterPerDegree"
        // 
        RotationalStiffnessMillinewtonMeterPerDegree RotationalStiffnessUnits = "MillinewtonMeterPerDegree"
        // 
        RotationalStiffnessCentinewtonMeterPerDegree RotationalStiffnessUnits = "CentinewtonMeterPerDegree"
        // 
        RotationalStiffnessDecinewtonMeterPerDegree RotationalStiffnessUnits = "DecinewtonMeterPerDegree"
        // 
        RotationalStiffnessDecanewtonMeterPerDegree RotationalStiffnessUnits = "DecanewtonMeterPerDegree"
        // 
        RotationalStiffnessKilonewtonMeterPerDegree RotationalStiffnessUnits = "KilonewtonMeterPerDegree"
        // 
        RotationalStiffnessMeganewtonMeterPerDegree RotationalStiffnessUnits = "MeganewtonMeterPerDegree"
        // 
        RotationalStiffnessNanonewtonMillimeterPerRadian RotationalStiffnessUnits = "NanonewtonMillimeterPerRadian"
        // 
        RotationalStiffnessMicronewtonMillimeterPerRadian RotationalStiffnessUnits = "MicronewtonMillimeterPerRadian"
        // 
        RotationalStiffnessMillinewtonMillimeterPerRadian RotationalStiffnessUnits = "MillinewtonMillimeterPerRadian"
        // 
        RotationalStiffnessCentinewtonMillimeterPerRadian RotationalStiffnessUnits = "CentinewtonMillimeterPerRadian"
        // 
        RotationalStiffnessDecinewtonMillimeterPerRadian RotationalStiffnessUnits = "DecinewtonMillimeterPerRadian"
        // 
        RotationalStiffnessDecanewtonMillimeterPerRadian RotationalStiffnessUnits = "DecanewtonMillimeterPerRadian"
        // 
        RotationalStiffnessKilonewtonMillimeterPerRadian RotationalStiffnessUnits = "KilonewtonMillimeterPerRadian"
        // 
        RotationalStiffnessMeganewtonMillimeterPerRadian RotationalStiffnessUnits = "MeganewtonMillimeterPerRadian"
)

// RotationalStiffnessDto represents a RotationalStiffness measurement with a numerical value and its corresponding unit.
type RotationalStiffnessDto struct {
    // Value is the numerical representation of the RotationalStiffness.
	Value float64
    // Unit specifies the unit of measurement for the RotationalStiffness, as defined in the RotationalStiffnessUnits enumeration.
	Unit  RotationalStiffnessUnits
}

// RotationalStiffnessDtoFactory groups methods for creating and serializing RotationalStiffnessDto objects.
type RotationalStiffnessDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RotationalStiffnessDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RotationalStiffnessDtoFactory) FromJSON(data []byte) (*RotationalStiffnessDto, error) {
	a := RotationalStiffnessDto{}

    // Parse JSON into RotationalStiffnessDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a RotationalStiffnessDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RotationalStiffnessDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// RotationalStiffness represents a measurement in a of RotationalStiffness.
//
// https://en.wikipedia.org/wiki/Stiffness#Rotational_stiffness
type RotationalStiffness struct {
	// value is the base measurement stored internally.
	value       float64
    
    newton_meters_per_radianLazy *float64 
    pound_force_feet_per_degreesLazy *float64 
    kilopound_force_feet_per_degreesLazy *float64 
    newton_millimeters_per_degreeLazy *float64 
    newton_meters_per_degreeLazy *float64 
    newton_millimeters_per_radianLazy *float64 
    pound_force_feet_per_radianLazy *float64 
    kilonewton_meters_per_radianLazy *float64 
    meganewton_meters_per_radianLazy *float64 
    nanonewton_millimeters_per_degreeLazy *float64 
    micronewton_millimeters_per_degreeLazy *float64 
    millinewton_millimeters_per_degreeLazy *float64 
    centinewton_millimeters_per_degreeLazy *float64 
    decinewton_millimeters_per_degreeLazy *float64 
    decanewton_millimeters_per_degreeLazy *float64 
    kilonewton_millimeters_per_degreeLazy *float64 
    meganewton_millimeters_per_degreeLazy *float64 
    nanonewton_meters_per_degreeLazy *float64 
    micronewton_meters_per_degreeLazy *float64 
    millinewton_meters_per_degreeLazy *float64 
    centinewton_meters_per_degreeLazy *float64 
    decinewton_meters_per_degreeLazy *float64 
    decanewton_meters_per_degreeLazy *float64 
    kilonewton_meters_per_degreeLazy *float64 
    meganewton_meters_per_degreeLazy *float64 
    nanonewton_millimeters_per_radianLazy *float64 
    micronewton_millimeters_per_radianLazy *float64 
    millinewton_millimeters_per_radianLazy *float64 
    centinewton_millimeters_per_radianLazy *float64 
    decinewton_millimeters_per_radianLazy *float64 
    decanewton_millimeters_per_radianLazy *float64 
    kilonewton_millimeters_per_radianLazy *float64 
    meganewton_millimeters_per_radianLazy *float64 
}

// RotationalStiffnessFactory groups methods for creating RotationalStiffness instances.
type RotationalStiffnessFactory struct{}

// CreateRotationalStiffness creates a new RotationalStiffness instance from the given value and unit.
func (uf RotationalStiffnessFactory) CreateRotationalStiffness(value float64, unit RotationalStiffnessUnits) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, unit)
}

// FromDto converts a RotationalStiffnessDto to a RotationalStiffness instance.
func (uf RotationalStiffnessFactory) FromDto(dto RotationalStiffnessDto) (*RotationalStiffness, error) {
	return newRotationalStiffness(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RotationalStiffness instance.
func (uf RotationalStiffnessFactory) FromDtoJSON(data []byte) (*RotationalStiffness, error) {
	unitDto, err := RotationalStiffnessDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RotationalStiffnessDto from JSON: %w", err)
	}
	return RotationalStiffnessFactory{}.FromDto(*unitDto)
}


// FromNewtonMetersPerRadian creates a new RotationalStiffness instance from a value in NewtonMetersPerRadian.
func (uf RotationalStiffnessFactory) FromNewtonMetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNewtonMeterPerRadian)
}

// FromPoundForceFeetPerDegrees creates a new RotationalStiffness instance from a value in PoundForceFeetPerDegrees.
func (uf RotationalStiffnessFactory) FromPoundForceFeetPerDegrees(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessPoundForceFootPerDegrees)
}

// FromKilopoundForceFeetPerDegrees creates a new RotationalStiffness instance from a value in KilopoundForceFeetPerDegrees.
func (uf RotationalStiffnessFactory) FromKilopoundForceFeetPerDegrees(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilopoundForceFootPerDegrees)
}

// FromNewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in NewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromNewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNewtonMillimeterPerDegree)
}

// FromNewtonMetersPerDegree creates a new RotationalStiffness instance from a value in NewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromNewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNewtonMeterPerDegree)
}

// FromNewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in NewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromNewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNewtonMillimeterPerRadian)
}

// FromPoundForceFeetPerRadian creates a new RotationalStiffness instance from a value in PoundForceFeetPerRadian.
func (uf RotationalStiffnessFactory) FromPoundForceFeetPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessPoundForceFeetPerRadian)
}

// FromKilonewtonMetersPerRadian creates a new RotationalStiffness instance from a value in KilonewtonMetersPerRadian.
func (uf RotationalStiffnessFactory) FromKilonewtonMetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilonewtonMeterPerRadian)
}

// FromMeganewtonMetersPerRadian creates a new RotationalStiffness instance from a value in MeganewtonMetersPerRadian.
func (uf RotationalStiffnessFactory) FromMeganewtonMetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMeganewtonMeterPerRadian)
}

// FromNanonewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in NanonewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromNanonewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNanonewtonMillimeterPerDegree)
}

// FromMicronewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in MicronewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromMicronewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMicronewtonMillimeterPerDegree)
}

// FromMillinewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in MillinewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromMillinewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMillinewtonMillimeterPerDegree)
}

// FromCentinewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in CentinewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromCentinewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessCentinewtonMillimeterPerDegree)
}

// FromDecinewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in DecinewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromDecinewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecinewtonMillimeterPerDegree)
}

// FromDecanewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in DecanewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromDecanewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecanewtonMillimeterPerDegree)
}

// FromKilonewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in KilonewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromKilonewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilonewtonMillimeterPerDegree)
}

// FromMeganewtonMillimetersPerDegree creates a new RotationalStiffness instance from a value in MeganewtonMillimetersPerDegree.
func (uf RotationalStiffnessFactory) FromMeganewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMeganewtonMillimeterPerDegree)
}

// FromNanonewtonMetersPerDegree creates a new RotationalStiffness instance from a value in NanonewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromNanonewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNanonewtonMeterPerDegree)
}

// FromMicronewtonMetersPerDegree creates a new RotationalStiffness instance from a value in MicronewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromMicronewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMicronewtonMeterPerDegree)
}

// FromMillinewtonMetersPerDegree creates a new RotationalStiffness instance from a value in MillinewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromMillinewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMillinewtonMeterPerDegree)
}

// FromCentinewtonMetersPerDegree creates a new RotationalStiffness instance from a value in CentinewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromCentinewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessCentinewtonMeterPerDegree)
}

// FromDecinewtonMetersPerDegree creates a new RotationalStiffness instance from a value in DecinewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromDecinewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecinewtonMeterPerDegree)
}

// FromDecanewtonMetersPerDegree creates a new RotationalStiffness instance from a value in DecanewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromDecanewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecanewtonMeterPerDegree)
}

// FromKilonewtonMetersPerDegree creates a new RotationalStiffness instance from a value in KilonewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromKilonewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilonewtonMeterPerDegree)
}

// FromMeganewtonMetersPerDegree creates a new RotationalStiffness instance from a value in MeganewtonMetersPerDegree.
func (uf RotationalStiffnessFactory) FromMeganewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMeganewtonMeterPerDegree)
}

// FromNanonewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in NanonewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromNanonewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNanonewtonMillimeterPerRadian)
}

// FromMicronewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in MicronewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromMicronewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMicronewtonMillimeterPerRadian)
}

// FromMillinewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in MillinewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromMillinewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMillinewtonMillimeterPerRadian)
}

// FromCentinewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in CentinewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromCentinewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessCentinewtonMillimeterPerRadian)
}

// FromDecinewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in DecinewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromDecinewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecinewtonMillimeterPerRadian)
}

// FromDecanewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in DecanewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromDecanewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecanewtonMillimeterPerRadian)
}

// FromKilonewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in KilonewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromKilonewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilonewtonMillimeterPerRadian)
}

// FromMeganewtonMillimetersPerRadian creates a new RotationalStiffness instance from a value in MeganewtonMillimetersPerRadian.
func (uf RotationalStiffnessFactory) FromMeganewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMeganewtonMillimeterPerRadian)
}


// newRotationalStiffness creates a new RotationalStiffness.
func newRotationalStiffness(value float64, fromUnit RotationalStiffnessUnits) (*RotationalStiffness, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RotationalStiffness{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RotationalStiffness in NewtonMeterPerRadian unit (the base/default quantity).
func (a *RotationalStiffness) BaseValue() float64 {
	return a.value
}


// NewtonMetersPerRadian returns the RotationalStiffness value in NewtonMetersPerRadian.
//
// 
func (a *RotationalStiffness) NewtonMetersPerRadian() float64 {
	if a.newton_meters_per_radianLazy != nil {
		return *a.newton_meters_per_radianLazy
	}
	newton_meters_per_radian := a.convertFromBase(RotationalStiffnessNewtonMeterPerRadian)
	a.newton_meters_per_radianLazy = &newton_meters_per_radian
	return newton_meters_per_radian
}

// PoundForceFeetPerDegrees returns the RotationalStiffness value in PoundForceFeetPerDegrees.
//
// 
func (a *RotationalStiffness) PoundForceFeetPerDegrees() float64 {
	if a.pound_force_feet_per_degreesLazy != nil {
		return *a.pound_force_feet_per_degreesLazy
	}
	pound_force_feet_per_degrees := a.convertFromBase(RotationalStiffnessPoundForceFootPerDegrees)
	a.pound_force_feet_per_degreesLazy = &pound_force_feet_per_degrees
	return pound_force_feet_per_degrees
}

// KilopoundForceFeetPerDegrees returns the RotationalStiffness value in KilopoundForceFeetPerDegrees.
//
// 
func (a *RotationalStiffness) KilopoundForceFeetPerDegrees() float64 {
	if a.kilopound_force_feet_per_degreesLazy != nil {
		return *a.kilopound_force_feet_per_degreesLazy
	}
	kilopound_force_feet_per_degrees := a.convertFromBase(RotationalStiffnessKilopoundForceFootPerDegrees)
	a.kilopound_force_feet_per_degreesLazy = &kilopound_force_feet_per_degrees
	return kilopound_force_feet_per_degrees
}

// NewtonMillimetersPerDegree returns the RotationalStiffness value in NewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) NewtonMillimetersPerDegree() float64 {
	if a.newton_millimeters_per_degreeLazy != nil {
		return *a.newton_millimeters_per_degreeLazy
	}
	newton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessNewtonMillimeterPerDegree)
	a.newton_millimeters_per_degreeLazy = &newton_millimeters_per_degree
	return newton_millimeters_per_degree
}

// NewtonMetersPerDegree returns the RotationalStiffness value in NewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) NewtonMetersPerDegree() float64 {
	if a.newton_meters_per_degreeLazy != nil {
		return *a.newton_meters_per_degreeLazy
	}
	newton_meters_per_degree := a.convertFromBase(RotationalStiffnessNewtonMeterPerDegree)
	a.newton_meters_per_degreeLazy = &newton_meters_per_degree
	return newton_meters_per_degree
}

// NewtonMillimetersPerRadian returns the RotationalStiffness value in NewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) NewtonMillimetersPerRadian() float64 {
	if a.newton_millimeters_per_radianLazy != nil {
		return *a.newton_millimeters_per_radianLazy
	}
	newton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessNewtonMillimeterPerRadian)
	a.newton_millimeters_per_radianLazy = &newton_millimeters_per_radian
	return newton_millimeters_per_radian
}

// PoundForceFeetPerRadian returns the RotationalStiffness value in PoundForceFeetPerRadian.
//
// 
func (a *RotationalStiffness) PoundForceFeetPerRadian() float64 {
	if a.pound_force_feet_per_radianLazy != nil {
		return *a.pound_force_feet_per_radianLazy
	}
	pound_force_feet_per_radian := a.convertFromBase(RotationalStiffnessPoundForceFeetPerRadian)
	a.pound_force_feet_per_radianLazy = &pound_force_feet_per_radian
	return pound_force_feet_per_radian
}

// KilonewtonMetersPerRadian returns the RotationalStiffness value in KilonewtonMetersPerRadian.
//
// 
func (a *RotationalStiffness) KilonewtonMetersPerRadian() float64 {
	if a.kilonewton_meters_per_radianLazy != nil {
		return *a.kilonewton_meters_per_radianLazy
	}
	kilonewton_meters_per_radian := a.convertFromBase(RotationalStiffnessKilonewtonMeterPerRadian)
	a.kilonewton_meters_per_radianLazy = &kilonewton_meters_per_radian
	return kilonewton_meters_per_radian
}

// MeganewtonMetersPerRadian returns the RotationalStiffness value in MeganewtonMetersPerRadian.
//
// 
func (a *RotationalStiffness) MeganewtonMetersPerRadian() float64 {
	if a.meganewton_meters_per_radianLazy != nil {
		return *a.meganewton_meters_per_radianLazy
	}
	meganewton_meters_per_radian := a.convertFromBase(RotationalStiffnessMeganewtonMeterPerRadian)
	a.meganewton_meters_per_radianLazy = &meganewton_meters_per_radian
	return meganewton_meters_per_radian
}

// NanonewtonMillimetersPerDegree returns the RotationalStiffness value in NanonewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) NanonewtonMillimetersPerDegree() float64 {
	if a.nanonewton_millimeters_per_degreeLazy != nil {
		return *a.nanonewton_millimeters_per_degreeLazy
	}
	nanonewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessNanonewtonMillimeterPerDegree)
	a.nanonewton_millimeters_per_degreeLazy = &nanonewton_millimeters_per_degree
	return nanonewton_millimeters_per_degree
}

// MicronewtonMillimetersPerDegree returns the RotationalStiffness value in MicronewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) MicronewtonMillimetersPerDegree() float64 {
	if a.micronewton_millimeters_per_degreeLazy != nil {
		return *a.micronewton_millimeters_per_degreeLazy
	}
	micronewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessMicronewtonMillimeterPerDegree)
	a.micronewton_millimeters_per_degreeLazy = &micronewton_millimeters_per_degree
	return micronewton_millimeters_per_degree
}

// MillinewtonMillimetersPerDegree returns the RotationalStiffness value in MillinewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) MillinewtonMillimetersPerDegree() float64 {
	if a.millinewton_millimeters_per_degreeLazy != nil {
		return *a.millinewton_millimeters_per_degreeLazy
	}
	millinewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessMillinewtonMillimeterPerDegree)
	a.millinewton_millimeters_per_degreeLazy = &millinewton_millimeters_per_degree
	return millinewton_millimeters_per_degree
}

// CentinewtonMillimetersPerDegree returns the RotationalStiffness value in CentinewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) CentinewtonMillimetersPerDegree() float64 {
	if a.centinewton_millimeters_per_degreeLazy != nil {
		return *a.centinewton_millimeters_per_degreeLazy
	}
	centinewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessCentinewtonMillimeterPerDegree)
	a.centinewton_millimeters_per_degreeLazy = &centinewton_millimeters_per_degree
	return centinewton_millimeters_per_degree
}

// DecinewtonMillimetersPerDegree returns the RotationalStiffness value in DecinewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) DecinewtonMillimetersPerDegree() float64 {
	if a.decinewton_millimeters_per_degreeLazy != nil {
		return *a.decinewton_millimeters_per_degreeLazy
	}
	decinewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessDecinewtonMillimeterPerDegree)
	a.decinewton_millimeters_per_degreeLazy = &decinewton_millimeters_per_degree
	return decinewton_millimeters_per_degree
}

// DecanewtonMillimetersPerDegree returns the RotationalStiffness value in DecanewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) DecanewtonMillimetersPerDegree() float64 {
	if a.decanewton_millimeters_per_degreeLazy != nil {
		return *a.decanewton_millimeters_per_degreeLazy
	}
	decanewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessDecanewtonMillimeterPerDegree)
	a.decanewton_millimeters_per_degreeLazy = &decanewton_millimeters_per_degree
	return decanewton_millimeters_per_degree
}

// KilonewtonMillimetersPerDegree returns the RotationalStiffness value in KilonewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) KilonewtonMillimetersPerDegree() float64 {
	if a.kilonewton_millimeters_per_degreeLazy != nil {
		return *a.kilonewton_millimeters_per_degreeLazy
	}
	kilonewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessKilonewtonMillimeterPerDegree)
	a.kilonewton_millimeters_per_degreeLazy = &kilonewton_millimeters_per_degree
	return kilonewton_millimeters_per_degree
}

// MeganewtonMillimetersPerDegree returns the RotationalStiffness value in MeganewtonMillimetersPerDegree.
//
// 
func (a *RotationalStiffness) MeganewtonMillimetersPerDegree() float64 {
	if a.meganewton_millimeters_per_degreeLazy != nil {
		return *a.meganewton_millimeters_per_degreeLazy
	}
	meganewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessMeganewtonMillimeterPerDegree)
	a.meganewton_millimeters_per_degreeLazy = &meganewton_millimeters_per_degree
	return meganewton_millimeters_per_degree
}

// NanonewtonMetersPerDegree returns the RotationalStiffness value in NanonewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) NanonewtonMetersPerDegree() float64 {
	if a.nanonewton_meters_per_degreeLazy != nil {
		return *a.nanonewton_meters_per_degreeLazy
	}
	nanonewton_meters_per_degree := a.convertFromBase(RotationalStiffnessNanonewtonMeterPerDegree)
	a.nanonewton_meters_per_degreeLazy = &nanonewton_meters_per_degree
	return nanonewton_meters_per_degree
}

// MicronewtonMetersPerDegree returns the RotationalStiffness value in MicronewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) MicronewtonMetersPerDegree() float64 {
	if a.micronewton_meters_per_degreeLazy != nil {
		return *a.micronewton_meters_per_degreeLazy
	}
	micronewton_meters_per_degree := a.convertFromBase(RotationalStiffnessMicronewtonMeterPerDegree)
	a.micronewton_meters_per_degreeLazy = &micronewton_meters_per_degree
	return micronewton_meters_per_degree
}

// MillinewtonMetersPerDegree returns the RotationalStiffness value in MillinewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) MillinewtonMetersPerDegree() float64 {
	if a.millinewton_meters_per_degreeLazy != nil {
		return *a.millinewton_meters_per_degreeLazy
	}
	millinewton_meters_per_degree := a.convertFromBase(RotationalStiffnessMillinewtonMeterPerDegree)
	a.millinewton_meters_per_degreeLazy = &millinewton_meters_per_degree
	return millinewton_meters_per_degree
}

// CentinewtonMetersPerDegree returns the RotationalStiffness value in CentinewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) CentinewtonMetersPerDegree() float64 {
	if a.centinewton_meters_per_degreeLazy != nil {
		return *a.centinewton_meters_per_degreeLazy
	}
	centinewton_meters_per_degree := a.convertFromBase(RotationalStiffnessCentinewtonMeterPerDegree)
	a.centinewton_meters_per_degreeLazy = &centinewton_meters_per_degree
	return centinewton_meters_per_degree
}

// DecinewtonMetersPerDegree returns the RotationalStiffness value in DecinewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) DecinewtonMetersPerDegree() float64 {
	if a.decinewton_meters_per_degreeLazy != nil {
		return *a.decinewton_meters_per_degreeLazy
	}
	decinewton_meters_per_degree := a.convertFromBase(RotationalStiffnessDecinewtonMeterPerDegree)
	a.decinewton_meters_per_degreeLazy = &decinewton_meters_per_degree
	return decinewton_meters_per_degree
}

// DecanewtonMetersPerDegree returns the RotationalStiffness value in DecanewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) DecanewtonMetersPerDegree() float64 {
	if a.decanewton_meters_per_degreeLazy != nil {
		return *a.decanewton_meters_per_degreeLazy
	}
	decanewton_meters_per_degree := a.convertFromBase(RotationalStiffnessDecanewtonMeterPerDegree)
	a.decanewton_meters_per_degreeLazy = &decanewton_meters_per_degree
	return decanewton_meters_per_degree
}

// KilonewtonMetersPerDegree returns the RotationalStiffness value in KilonewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) KilonewtonMetersPerDegree() float64 {
	if a.kilonewton_meters_per_degreeLazy != nil {
		return *a.kilonewton_meters_per_degreeLazy
	}
	kilonewton_meters_per_degree := a.convertFromBase(RotationalStiffnessKilonewtonMeterPerDegree)
	a.kilonewton_meters_per_degreeLazy = &kilonewton_meters_per_degree
	return kilonewton_meters_per_degree
}

// MeganewtonMetersPerDegree returns the RotationalStiffness value in MeganewtonMetersPerDegree.
//
// 
func (a *RotationalStiffness) MeganewtonMetersPerDegree() float64 {
	if a.meganewton_meters_per_degreeLazy != nil {
		return *a.meganewton_meters_per_degreeLazy
	}
	meganewton_meters_per_degree := a.convertFromBase(RotationalStiffnessMeganewtonMeterPerDegree)
	a.meganewton_meters_per_degreeLazy = &meganewton_meters_per_degree
	return meganewton_meters_per_degree
}

// NanonewtonMillimetersPerRadian returns the RotationalStiffness value in NanonewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) NanonewtonMillimetersPerRadian() float64 {
	if a.nanonewton_millimeters_per_radianLazy != nil {
		return *a.nanonewton_millimeters_per_radianLazy
	}
	nanonewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessNanonewtonMillimeterPerRadian)
	a.nanonewton_millimeters_per_radianLazy = &nanonewton_millimeters_per_radian
	return nanonewton_millimeters_per_radian
}

// MicronewtonMillimetersPerRadian returns the RotationalStiffness value in MicronewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) MicronewtonMillimetersPerRadian() float64 {
	if a.micronewton_millimeters_per_radianLazy != nil {
		return *a.micronewton_millimeters_per_radianLazy
	}
	micronewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessMicronewtonMillimeterPerRadian)
	a.micronewton_millimeters_per_radianLazy = &micronewton_millimeters_per_radian
	return micronewton_millimeters_per_radian
}

// MillinewtonMillimetersPerRadian returns the RotationalStiffness value in MillinewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) MillinewtonMillimetersPerRadian() float64 {
	if a.millinewton_millimeters_per_radianLazy != nil {
		return *a.millinewton_millimeters_per_radianLazy
	}
	millinewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessMillinewtonMillimeterPerRadian)
	a.millinewton_millimeters_per_radianLazy = &millinewton_millimeters_per_radian
	return millinewton_millimeters_per_radian
}

// CentinewtonMillimetersPerRadian returns the RotationalStiffness value in CentinewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) CentinewtonMillimetersPerRadian() float64 {
	if a.centinewton_millimeters_per_radianLazy != nil {
		return *a.centinewton_millimeters_per_radianLazy
	}
	centinewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessCentinewtonMillimeterPerRadian)
	a.centinewton_millimeters_per_radianLazy = &centinewton_millimeters_per_radian
	return centinewton_millimeters_per_radian
}

// DecinewtonMillimetersPerRadian returns the RotationalStiffness value in DecinewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) DecinewtonMillimetersPerRadian() float64 {
	if a.decinewton_millimeters_per_radianLazy != nil {
		return *a.decinewton_millimeters_per_radianLazy
	}
	decinewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessDecinewtonMillimeterPerRadian)
	a.decinewton_millimeters_per_radianLazy = &decinewton_millimeters_per_radian
	return decinewton_millimeters_per_radian
}

// DecanewtonMillimetersPerRadian returns the RotationalStiffness value in DecanewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) DecanewtonMillimetersPerRadian() float64 {
	if a.decanewton_millimeters_per_radianLazy != nil {
		return *a.decanewton_millimeters_per_radianLazy
	}
	decanewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessDecanewtonMillimeterPerRadian)
	a.decanewton_millimeters_per_radianLazy = &decanewton_millimeters_per_radian
	return decanewton_millimeters_per_radian
}

// KilonewtonMillimetersPerRadian returns the RotationalStiffness value in KilonewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) KilonewtonMillimetersPerRadian() float64 {
	if a.kilonewton_millimeters_per_radianLazy != nil {
		return *a.kilonewton_millimeters_per_radianLazy
	}
	kilonewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessKilonewtonMillimeterPerRadian)
	a.kilonewton_millimeters_per_radianLazy = &kilonewton_millimeters_per_radian
	return kilonewton_millimeters_per_radian
}

// MeganewtonMillimetersPerRadian returns the RotationalStiffness value in MeganewtonMillimetersPerRadian.
//
// 
func (a *RotationalStiffness) MeganewtonMillimetersPerRadian() float64 {
	if a.meganewton_millimeters_per_radianLazy != nil {
		return *a.meganewton_millimeters_per_radianLazy
	}
	meganewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessMeganewtonMillimeterPerRadian)
	a.meganewton_millimeters_per_radianLazy = &meganewton_millimeters_per_radian
	return meganewton_millimeters_per_radian
}


// ToDto creates a RotationalStiffnessDto representation from the RotationalStiffness instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonMeterPerRadian by default.
func (a *RotationalStiffness) ToDto(holdInUnit *RotationalStiffnessUnits) RotationalStiffnessDto {
	if holdInUnit == nil {
		defaultUnit := RotationalStiffnessNewtonMeterPerRadian // Default value
		holdInUnit = &defaultUnit
	}

	return RotationalStiffnessDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the RotationalStiffness instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonMeterPerRadian by default.
func (a *RotationalStiffness) ToDtoJSON(holdInUnit *RotationalStiffnessUnits) ([]byte, error) {
	// Convert to RotationalStiffnessDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RotationalStiffness to a specific unit value.
// The function uses the provided unit type (RotationalStiffnessUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *RotationalStiffness) Convert(toUnit RotationalStiffnessUnits) float64 {
	switch toUnit { 
    case RotationalStiffnessNewtonMeterPerRadian:
		return a.NewtonMetersPerRadian()
    case RotationalStiffnessPoundForceFootPerDegrees:
		return a.PoundForceFeetPerDegrees()
    case RotationalStiffnessKilopoundForceFootPerDegrees:
		return a.KilopoundForceFeetPerDegrees()
    case RotationalStiffnessNewtonMillimeterPerDegree:
		return a.NewtonMillimetersPerDegree()
    case RotationalStiffnessNewtonMeterPerDegree:
		return a.NewtonMetersPerDegree()
    case RotationalStiffnessNewtonMillimeterPerRadian:
		return a.NewtonMillimetersPerRadian()
    case RotationalStiffnessPoundForceFeetPerRadian:
		return a.PoundForceFeetPerRadian()
    case RotationalStiffnessKilonewtonMeterPerRadian:
		return a.KilonewtonMetersPerRadian()
    case RotationalStiffnessMeganewtonMeterPerRadian:
		return a.MeganewtonMetersPerRadian()
    case RotationalStiffnessNanonewtonMillimeterPerDegree:
		return a.NanonewtonMillimetersPerDegree()
    case RotationalStiffnessMicronewtonMillimeterPerDegree:
		return a.MicronewtonMillimetersPerDegree()
    case RotationalStiffnessMillinewtonMillimeterPerDegree:
		return a.MillinewtonMillimetersPerDegree()
    case RotationalStiffnessCentinewtonMillimeterPerDegree:
		return a.CentinewtonMillimetersPerDegree()
    case RotationalStiffnessDecinewtonMillimeterPerDegree:
		return a.DecinewtonMillimetersPerDegree()
    case RotationalStiffnessDecanewtonMillimeterPerDegree:
		return a.DecanewtonMillimetersPerDegree()
    case RotationalStiffnessKilonewtonMillimeterPerDegree:
		return a.KilonewtonMillimetersPerDegree()
    case RotationalStiffnessMeganewtonMillimeterPerDegree:
		return a.MeganewtonMillimetersPerDegree()
    case RotationalStiffnessNanonewtonMeterPerDegree:
		return a.NanonewtonMetersPerDegree()
    case RotationalStiffnessMicronewtonMeterPerDegree:
		return a.MicronewtonMetersPerDegree()
    case RotationalStiffnessMillinewtonMeterPerDegree:
		return a.MillinewtonMetersPerDegree()
    case RotationalStiffnessCentinewtonMeterPerDegree:
		return a.CentinewtonMetersPerDegree()
    case RotationalStiffnessDecinewtonMeterPerDegree:
		return a.DecinewtonMetersPerDegree()
    case RotationalStiffnessDecanewtonMeterPerDegree:
		return a.DecanewtonMetersPerDegree()
    case RotationalStiffnessKilonewtonMeterPerDegree:
		return a.KilonewtonMetersPerDegree()
    case RotationalStiffnessMeganewtonMeterPerDegree:
		return a.MeganewtonMetersPerDegree()
    case RotationalStiffnessNanonewtonMillimeterPerRadian:
		return a.NanonewtonMillimetersPerRadian()
    case RotationalStiffnessMicronewtonMillimeterPerRadian:
		return a.MicronewtonMillimetersPerRadian()
    case RotationalStiffnessMillinewtonMillimeterPerRadian:
		return a.MillinewtonMillimetersPerRadian()
    case RotationalStiffnessCentinewtonMillimeterPerRadian:
		return a.CentinewtonMillimetersPerRadian()
    case RotationalStiffnessDecinewtonMillimeterPerRadian:
		return a.DecinewtonMillimetersPerRadian()
    case RotationalStiffnessDecanewtonMillimeterPerRadian:
		return a.DecanewtonMillimetersPerRadian()
    case RotationalStiffnessKilonewtonMillimeterPerRadian:
		return a.KilonewtonMillimetersPerRadian()
    case RotationalStiffnessMeganewtonMillimeterPerRadian:
		return a.MeganewtonMillimetersPerRadian()
	default:
		return math.NaN()
	}
}

func (a *RotationalStiffness) convertFromBase(toUnit RotationalStiffnessUnits) float64 {
    value := a.value
	switch toUnit { 
	case RotationalStiffnessNewtonMeterPerRadian:
		return (value) 
	case RotationalStiffnessPoundForceFootPerDegrees:
		return (value / 77.6826) 
	case RotationalStiffnessKilopoundForceFootPerDegrees:
		return (value / 77682.6) 
	case RotationalStiffnessNewtonMillimeterPerDegree:
		return (value / 180 * math.Pi * 1000) 
	case RotationalStiffnessNewtonMeterPerDegree:
		return (value / (180 / math.Pi)) 
	case RotationalStiffnessNewtonMillimeterPerRadian:
		return (value * 1000) 
	case RotationalStiffnessPoundForceFeetPerRadian:
		return (value / 1.3558179483314) 
	case RotationalStiffnessKilonewtonMeterPerRadian:
		return ((value) / 1000.0) 
	case RotationalStiffnessMeganewtonMeterPerRadian:
		return ((value) / 1000000.0) 
	case RotationalStiffnessNanonewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 1e-09) 
	case RotationalStiffnessMicronewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 1e-06) 
	case RotationalStiffnessMillinewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 0.001) 
	case RotationalStiffnessCentinewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 0.01) 
	case RotationalStiffnessDecinewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 0.1) 
	case RotationalStiffnessDecanewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 10.0) 
	case RotationalStiffnessKilonewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 1000.0) 
	case RotationalStiffnessMeganewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 1000000.0) 
	case RotationalStiffnessNanonewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 1e-09) 
	case RotationalStiffnessMicronewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 1e-06) 
	case RotationalStiffnessMillinewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 0.001) 
	case RotationalStiffnessCentinewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 0.01) 
	case RotationalStiffnessDecinewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 0.1) 
	case RotationalStiffnessDecanewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 10.0) 
	case RotationalStiffnessKilonewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 1000.0) 
	case RotationalStiffnessMeganewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 1000000.0) 
	case RotationalStiffnessNanonewtonMillimeterPerRadian:
		return ((value * 1000) / 1e-09) 
	case RotationalStiffnessMicronewtonMillimeterPerRadian:
		return ((value * 1000) / 1e-06) 
	case RotationalStiffnessMillinewtonMillimeterPerRadian:
		return ((value * 1000) / 0.001) 
	case RotationalStiffnessCentinewtonMillimeterPerRadian:
		return ((value * 1000) / 0.01) 
	case RotationalStiffnessDecinewtonMillimeterPerRadian:
		return ((value * 1000) / 0.1) 
	case RotationalStiffnessDecanewtonMillimeterPerRadian:
		return ((value * 1000) / 10.0) 
	case RotationalStiffnessKilonewtonMillimeterPerRadian:
		return ((value * 1000) / 1000.0) 
	case RotationalStiffnessMeganewtonMillimeterPerRadian:
		return ((value * 1000) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *RotationalStiffness) convertToBase(value float64, fromUnit RotationalStiffnessUnits) float64 {
	switch fromUnit { 
	case RotationalStiffnessNewtonMeterPerRadian:
		return (value) 
	case RotationalStiffnessPoundForceFootPerDegrees:
		return (value * 77.6826) 
	case RotationalStiffnessKilopoundForceFootPerDegrees:
		return (value * 77682.6) 
	case RotationalStiffnessNewtonMillimeterPerDegree:
		return (value * 180 / math.Pi * 0.001) 
	case RotationalStiffnessNewtonMeterPerDegree:
		return (value * (180 / math.Pi)) 
	case RotationalStiffnessNewtonMillimeterPerRadian:
		return (value * 0.001) 
	case RotationalStiffnessPoundForceFeetPerRadian:
		return (value * 1.3558179483314) 
	case RotationalStiffnessKilonewtonMeterPerRadian:
		return ((value) * 1000.0) 
	case RotationalStiffnessMeganewtonMeterPerRadian:
		return ((value) * 1000000.0) 
	case RotationalStiffnessNanonewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 1e-09) 
	case RotationalStiffnessMicronewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 1e-06) 
	case RotationalStiffnessMillinewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 0.001) 
	case RotationalStiffnessCentinewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 0.01) 
	case RotationalStiffnessDecinewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 0.1) 
	case RotationalStiffnessDecanewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 10.0) 
	case RotationalStiffnessKilonewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 1000.0) 
	case RotationalStiffnessMeganewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 1000000.0) 
	case RotationalStiffnessNanonewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 1e-09) 
	case RotationalStiffnessMicronewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 1e-06) 
	case RotationalStiffnessMillinewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 0.001) 
	case RotationalStiffnessCentinewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 0.01) 
	case RotationalStiffnessDecinewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 0.1) 
	case RotationalStiffnessDecanewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 10.0) 
	case RotationalStiffnessKilonewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 1000.0) 
	case RotationalStiffnessMeganewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 1000000.0) 
	case RotationalStiffnessNanonewtonMillimeterPerRadian:
		return ((value * 0.001) * 1e-09) 
	case RotationalStiffnessMicronewtonMillimeterPerRadian:
		return ((value * 0.001) * 1e-06) 
	case RotationalStiffnessMillinewtonMillimeterPerRadian:
		return ((value * 0.001) * 0.001) 
	case RotationalStiffnessCentinewtonMillimeterPerRadian:
		return ((value * 0.001) * 0.01) 
	case RotationalStiffnessDecinewtonMillimeterPerRadian:
		return ((value * 0.001) * 0.1) 
	case RotationalStiffnessDecanewtonMillimeterPerRadian:
		return ((value * 0.001) * 10.0) 
	case RotationalStiffnessKilonewtonMillimeterPerRadian:
		return ((value * 0.001) * 1000.0) 
	case RotationalStiffnessMeganewtonMillimeterPerRadian:
		return ((value * 0.001) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the RotationalStiffness in the default unit (NewtonMeterPerRadian),
// formatted to two decimal places.
func (a RotationalStiffness) String() string {
	return a.ToString(RotationalStiffnessNewtonMeterPerRadian, 2)
}

// ToString formats the RotationalStiffness value as a string with the specified unit and fractional digits.
// It converts the RotationalStiffness to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RotationalStiffness value will be converted (e.g., NewtonMeterPerRadian))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RotationalStiffness with the unit abbreviation.
func (a *RotationalStiffness) ToString(unit RotationalStiffnessUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RotationalStiffness) getUnitAbbreviation(unit RotationalStiffnessUnits) string {
	switch unit { 
	case RotationalStiffnessNewtonMeterPerRadian:
		return "N·m/rad" 
	case RotationalStiffnessPoundForceFootPerDegrees:
		return "lbf·ft/deg" 
	case RotationalStiffnessKilopoundForceFootPerDegrees:
		return "kipf·ft/°" 
	case RotationalStiffnessNewtonMillimeterPerDegree:
		return "N·mm/deg" 
	case RotationalStiffnessNewtonMeterPerDegree:
		return "N·m/deg" 
	case RotationalStiffnessNewtonMillimeterPerRadian:
		return "N·mm/rad" 
	case RotationalStiffnessPoundForceFeetPerRadian:
		return "lbf·ft/rad" 
	case RotationalStiffnessKilonewtonMeterPerRadian:
		return "kN·m/rad" 
	case RotationalStiffnessMeganewtonMeterPerRadian:
		return "MN·m/rad" 
	case RotationalStiffnessNanonewtonMillimeterPerDegree:
		return "nN·mm/deg" 
	case RotationalStiffnessMicronewtonMillimeterPerDegree:
		return "μN·mm/deg" 
	case RotationalStiffnessMillinewtonMillimeterPerDegree:
		return "mN·mm/deg" 
	case RotationalStiffnessCentinewtonMillimeterPerDegree:
		return "cN·mm/deg" 
	case RotationalStiffnessDecinewtonMillimeterPerDegree:
		return "dN·mm/deg" 
	case RotationalStiffnessDecanewtonMillimeterPerDegree:
		return "daN·mm/deg" 
	case RotationalStiffnessKilonewtonMillimeterPerDegree:
		return "kN·mm/deg" 
	case RotationalStiffnessMeganewtonMillimeterPerDegree:
		return "MN·mm/deg" 
	case RotationalStiffnessNanonewtonMeterPerDegree:
		return "nN·m/deg" 
	case RotationalStiffnessMicronewtonMeterPerDegree:
		return "μN·m/deg" 
	case RotationalStiffnessMillinewtonMeterPerDegree:
		return "mN·m/deg" 
	case RotationalStiffnessCentinewtonMeterPerDegree:
		return "cN·m/deg" 
	case RotationalStiffnessDecinewtonMeterPerDegree:
		return "dN·m/deg" 
	case RotationalStiffnessDecanewtonMeterPerDegree:
		return "daN·m/deg" 
	case RotationalStiffnessKilonewtonMeterPerDegree:
		return "kN·m/deg" 
	case RotationalStiffnessMeganewtonMeterPerDegree:
		return "MN·m/deg" 
	case RotationalStiffnessNanonewtonMillimeterPerRadian:
		return "nN·mm/rad" 
	case RotationalStiffnessMicronewtonMillimeterPerRadian:
		return "μN·mm/rad" 
	case RotationalStiffnessMillinewtonMillimeterPerRadian:
		return "mN·mm/rad" 
	case RotationalStiffnessCentinewtonMillimeterPerRadian:
		return "cN·mm/rad" 
	case RotationalStiffnessDecinewtonMillimeterPerRadian:
		return "dN·mm/rad" 
	case RotationalStiffnessDecanewtonMillimeterPerRadian:
		return "daN·mm/rad" 
	case RotationalStiffnessKilonewtonMillimeterPerRadian:
		return "kN·mm/rad" 
	case RotationalStiffnessMeganewtonMillimeterPerRadian:
		return "MN·mm/rad" 
	default:
		return ""
	}
}

// Equals checks if the given RotationalStiffness is equal to the current RotationalStiffness.
//
// Parameters:
//    other: The RotationalStiffness to compare against.
//
// Returns:
//    bool: Returns true if both RotationalStiffness are equal, false otherwise.
func (a *RotationalStiffness) Equals(other *RotationalStiffness) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RotationalStiffness with another RotationalStiffness.
// It returns -1 if the current RotationalStiffness is less than the other RotationalStiffness, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RotationalStiffness to compare against.
//
// Returns:
//    int: -1 if the current RotationalStiffness is less, 1 if greater, and 0 if equal.
func (a *RotationalStiffness) CompareTo(other *RotationalStiffness) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given RotationalStiffness to the current RotationalStiffness and returns the result.
// The result is a new RotationalStiffness instance with the sum of the values.
//
// Parameters:
//    other: The RotationalStiffness to add to the current RotationalStiffness.
//
// Returns:
//    *RotationalStiffness: A new RotationalStiffness instance representing the sum of both RotationalStiffness.
func (a *RotationalStiffness) Add(other *RotationalStiffness) *RotationalStiffness {
	return &RotationalStiffness{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RotationalStiffness from the current RotationalStiffness and returns the result.
// The result is a new RotationalStiffness instance with the difference of the values.
//
// Parameters:
//    other: The RotationalStiffness to subtract from the current RotationalStiffness.
//
// Returns:
//    *RotationalStiffness: A new RotationalStiffness instance representing the difference of both RotationalStiffness.
func (a *RotationalStiffness) Subtract(other *RotationalStiffness) *RotationalStiffness {
	return &RotationalStiffness{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RotationalStiffness by the given RotationalStiffness and returns the result.
// The result is a new RotationalStiffness instance with the product of the values.
//
// Parameters:
//    other: The RotationalStiffness to multiply with the current RotationalStiffness.
//
// Returns:
//    *RotationalStiffness: A new RotationalStiffness instance representing the product of both RotationalStiffness.
func (a *RotationalStiffness) Multiply(other *RotationalStiffness) *RotationalStiffness {
	return &RotationalStiffness{value: a.value * other.BaseValue()}
}

// Divide divides the current RotationalStiffness by the given RotationalStiffness and returns the result.
// The result is a new RotationalStiffness instance with the quotient of the values.
//
// Parameters:
//    other: The RotationalStiffness to divide the current RotationalStiffness by.
//
// Returns:
//    *RotationalStiffness: A new RotationalStiffness instance representing the quotient of both RotationalStiffness.
func (a *RotationalStiffness) Divide(other *RotationalStiffness) *RotationalStiffness {
	return &RotationalStiffness{value: a.value / other.BaseValue()}
}